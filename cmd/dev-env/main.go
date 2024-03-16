// SPDX-FileCopyrightText: 2024-2026 SAP SE or an SAP affiliate company and Greenhouse contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	cp "github.com/otiai10/copy"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd/api"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"github.com/cloudoperators/greenhouse/pkg/clientutil"
	"github.com/cloudoperators/greenhouse/pkg/rbac"
)

const (
	UpdateInterval  = 1 * time.Second
	UpdateTimeout   = 30 * time.Second
	healthCheckFile = "/tmp/.envtest-running"
)

var (
	devEnvDataDir     string
	certDir           string
	webhookHost       string
	webhookPort       int
	kubeBuilderAssets string
	kubeProxyHost     string
	kubeProxyPort     string
	graceFullShutDown bool
	userData          = map[string][]string{
		"test-org-admin":  {rbac.GetOrganizationRoleName("test-org"), rbac.GetAdminRoleNameForOrganization("test-org"), rbac.GetTeamRoleName("test-team-1")},
		"test-org-member": {rbac.GetOrganizationRoleName("test-org"), rbac.GetTeamRoleName("test-team-1")},
	}
)

func main() {
	logger := logrus.New()
	err := os.RemoveAll(healthCheckFile)
	if err != nil {
		logger.Infof("Failed to delete healthcheck file:: %s", err)
	}

	flag.StringVar(&devEnvDataDir, "data-dir", "/envtest", "The directory dev env data is written to")
	flag.StringVar(&certDir, "cert-dir", clientutil.GetEnvOrDefault("WEBHOOK_CERT_DIR", "/webhook-certs"), "directory the autogenerated certs for serving webhook server are persisted to - defaults to WEBHOOK_CERT_DIR env var and /webhook-certs if unset - leave empty if no copy needed")
	flag.StringVar(&webhookHost, "webhook-host", clientutil.GetEnvOrDefault("WEBHOOK_HOST", "127.0.0.1"), "host the webhooks are served on - defaults to WEBHOOK_HOST env var and \"127.0.0.1\" if unset")
	flag.IntVar(&webhookPort, "webhook-port", clientutil.GetIntEnvWithDefault("WEBHOOK_PORT", 9443), "port the webhook server is served on - defaults to WEBHOOK_PORT env var and 6884 if unset")
	flag.StringVar(&kubeBuilderAssets, "kubebuilder-assets", os.Getenv("KUBEBUILDER_ASSETS"), "directory containing testenv binaries - defaults to KUBEBUILDER_ASSETS env var")
	flag.StringVar(&kubeProxyHost, "kubeproxy-host", "127.0.0.1", "host proxied by kubectl proxy - defaults to 127.0.01")
	flag.StringVar(&kubeProxyPort, "kubeproxy-port", "8090", "port proxied by kubectl proxy -  defaults to 8090")
	flag.BoolVar(&graceFullShutDown, "graceful-shutdown", false, "Shutdown envtest on SIGTERM and SIGINT. Defaults to false.")
	flag.Parse()

	if kubeBuilderAssets == "" {
		logger.Fatal("KUBEBUILDER_ASSETS env needs to be set for dev-env to run")
	}

	absPathConfigBasePath, err := clientutil.FindDirUpwards(".", "config", 10)
	if err != nil {
		logger.Fatalf("failed finding config base path: %s", err)
	}
	crdPaths := []string{filepath.Join(absPathConfigBasePath, "crd", "bases"), filepath.Join(absPathConfigBasePath, "..", "charts", "idproxy", "crds")}
	webhookPaths := []string{filepath.Join(absPathConfigBasePath, "webhook")}

	envTest := &envtest.Environment{
		CRDDirectoryPaths:     crdPaths,
		ErrorIfCRDPathMissing: true,
		WebhookInstallOptions: envtest.WebhookInstallOptions{
			Paths:            webhookPaths,
			LocalServingPort: webhookPort,
			LocalServingHost: webhookHost,
		},
	}
	envTest.ControlPlane.GetAPIServer().SecureServing.ListenAddr.Address = "127.0.0.1"
	envTest.ControlPlane.GetAPIServer().SecureServing.ListenAddr.Port = "6884"
	envTest.ControlPlane.GetAPIServer().Configure().Append("cors-allowed-origins", ".*")
	envTest.ControlPlane.GetAPIServer().Configure().Append("enable-admission-plugins", "MutatingAdmissionWebhook", "ValidatingAdmissionWebhook")

	//starting dev env
	logger.Info("Starting apiserver & etcd")
	cfg, err := envTest.Start()
	if err != nil {
		logger.Fatalf("failed starting envTest: %s", err)
	}
	logger.Info("apiserver running - host: ", cfg.Host)

	createInternalKubeConfigFile(logger, envTest, cfg)

	createAdditionalKubeConfigFiles(logger)

	if certDir != "" {
		copyWebhookCertsToAccessibleDir(logger, envTest)
	}

	logger.Infof("Expecting webhook server at %s:%v", envTest.WebhookInstallOptions.LocalServingHost, envTest.WebhookInstallOptions.LocalServingPort)
	logger.Info("dev-env running")
	_, err = os.OpenFile("/tmp/.envtest-running", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		logger.Infof("Failed to write healthcheck file:: %s", err)
	}

	if graceFullShutDown {
		ctx := ctrl.SetupSignalHandler()

		for {
			select {
			case <-ctx.Done():
				logger.Info("Stopping apiserver & etcd")
				err := envTest.Stop()
				if err != nil {
					logger.Fatalf("failed stopping envTest: %s", err)
				}
				break
			default:
				continue
			}
			break
		}
	}
}

func createInternalKubeConfigFile(logger *logrus.Logger, envTest *envtest.Environment, cfg *rest.Config) {
	internalKubeConfig := KubeConfig{
		dataDir: devEnvDataDir,
		logger:  logger,
		config: api.Config{
			Kind:       "Config",
			APIVersion: "v1",
			Clusters: map[string]*api.Cluster{
				"default": {
					Server:                   cfg.Host,
					CertificateAuthorityData: cfg.CAData,
				}},
			Contexts:  make(map[string]*api.Context),
			AuthInfos: make(map[string]*api.AuthInfo),
		}}

	internalKubeConfig.addUser("cluster-admin", cfg, "")

	//create users for test-org and add to kubeConfig
	for name, groups := range userData {
		user, err := envTest.ControlPlane.AddUser(envtest.User{
			Name:   name,
			Groups: groups,
		}, nil)
		if err != nil {
			logger.Fatalf("Failed adding user: %s", err)
		}
		userCfg := user.Config()
		internalKubeConfig.addUser(name, userCfg, "test-org")
	}

	internalKubeConfig.config.CurrentContext = "cluster-admin"

	err := internalKubeConfig.writeFile("internal.kubeconfig")
	if err != nil {
		logger.Fatalf("Failed writing kubeconfig: %s", err)
	}

	err = internalKubeConfig.writeCertDataToFiles()
	if err != nil {
		logger.Fatalf("Failed writing cert data to files: %s", err)
	}
}

func createAdditionalKubeConfigFiles(logger *logrus.Logger) {
	proxyKubeConfig := KubeConfig{

		dataDir: devEnvDataDir,
		logger:  logger,
		config: api.Config{
			Kind:       "Config",
			APIVersion: "v1",
			Clusters: map[string]*api.Cluster{
				"default": {
					Server: "http://" + kubeProxyHost + ":" + kubeProxyPort,
				}},
			Contexts: map[string]*api.Context{
				"default": {
					Cluster:  "default",
					AuthInfo: "default",
				},
			},
			AuthInfos: map[string]*api.AuthInfo{
				"default": {},
			},
			CurrentContext: "default",
		}}

	err := proxyKubeConfig.writeFile("kubeconfig")
	if err != nil {
		logger.Fatalf("Failed writing kubeconfig: %s", err)
	}
}

func copyWebhookCertsToAccessibleDir(logger *logrus.Logger, envTest *envtest.Environment) {
	err := os.MkdirAll(certDir, os.ModePerm)
	if err != nil {
		logger.Fatalf("Failed to create target dir: %s", err)
	}
	if err := cp.Copy(envTest.WebhookInstallOptions.LocalServingCertDir, certDir, cp.Options{AddPermission: 0007}); err != nil {
		logger.Fatalf("Failed to copy webhook cert dir %s: %s", envTest.WebhookInstallOptions.LocalServingCertDir, err)
	}
	logger.Infof("Successfully copied generated webhook server certs from %s to %s", envTest.WebhookInstallOptions.LocalServingCertDir, certDir)
}
