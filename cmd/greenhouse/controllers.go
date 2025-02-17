// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Greenhouse contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"sort"

	ctrl "sigs.k8s.io/controller-runtime"

	clustercontrollers "github.com/cloudoperators/greenhouse/pkg/controllers/cluster"
	organizationcontrollers "github.com/cloudoperators/greenhouse/pkg/controllers/organization"
	plugincontrollers "github.com/cloudoperators/greenhouse/pkg/controllers/plugin"
	teamcontrollers "github.com/cloudoperators/greenhouse/pkg/controllers/team"
	teammembershipcontrollers "github.com/cloudoperators/greenhouse/pkg/controllers/teammembership"
	teamrbaccontrollers "github.com/cloudoperators/greenhouse/pkg/controllers/teamrbac"
)

// knownControllers contains all controllers to be registered when starting the operator.
var knownControllers = map[string]func(controllerName string, mgr ctrl.Manager) error{
	// Organization controllers.
	"organizationController": startOrganizationReconciler,

	// Team controllers.
	"teamPropagation": (&teamcontrollers.TeamPropagationReconciler{}).SetupWithManager,

	// TeamMembership controllers.
	"teamMembershipUpdater":     (&teammembershipcontrollers.TeamMembershipUpdaterController{}).SetupWithManager,
	"teamMembershipPropagation": (&teammembershipcontrollers.TeamMembershipPropagationReconciler{}).SetupWithManager,

	// Team RBAC controllers.
	"teamRoleBindingController": (&teamrbaccontrollers.TeamRoleBindingReconciler{}).SetupWithManager,

	// Plugin controllers.
	"plugin": (&plugincontrollers.PluginReconciler{
		KubeRuntimeOpts: kubeClientOpts,
	}).SetupWithManager,
	"pluginPreset": (&plugincontrollers.PluginPresetReconciler{}).SetupWithManager,

	// Cluster controllers
	"bootStrap":         (&clustercontrollers.BootstrapReconciler{}).SetupWithManager,
	"clusterReconciler": startClusterReconciler,
	"kubeconfig":        (&clustercontrollers.KubeconfigReconciler{}).SetupWithManager,
}

// knownControllers lists the name of known controllers.
func knownControllersNames() []string {
	controllerStringSlice := make([]string, 0)
	for controllerName := range knownControllers {
		controllerStringSlice = append(controllerStringSlice, controllerName)
	}
	sort.Strings(controllerStringSlice)
	return controllerStringSlice
}

// isControllerEnabled checks whether the given controller or regex is enabled
func isControllerEnabled(controllerName string) bool {
	for _, c := range enabledControllers {
		if controllerName == "*" || controllerName == c {
			return true
		}
	}
	return false
}

func startOrganizationReconciler(name string, mgr ctrl.Manager) error {
	namespace := "greenhouse"
	if v, ok := os.LookupEnv("POD_NAMESPACE"); ok {
		namespace = v
	}
	return (&organizationcontrollers.OrganizationReconciler{
		Namespace: namespace,
	}).SetupWithManager(name, mgr)
}

func startClusterReconciler(name string, mgr ctrl.Manager) error {
	if renewRemoteClusterBearerTokenAfter > remoteClusterBearerTokenValidity {
		setupLog.Info("WARN: remoteClusterBearerTokenValidity is less than renewRemoteClusterBearerTokenAfter")
		setupLog.Info("Setting renewRemoteClusterBearerTokenAfter to half of remoteClusterBearerTokenValidity")
		renewRemoteClusterBearerTokenAfter = remoteClusterBearerTokenValidity / 2
	}
	return (&clustercontrollers.RemoteClusterReconciler{
		RemoteClusterBearerTokenValidity:   remoteClusterBearerTokenValidity,
		RenewRemoteClusterBearerTokenAfter: renewRemoteClusterBearerTokenAfter,
	}).SetupWithManager(name, mgr)
}
