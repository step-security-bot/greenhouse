// SPDX-FileCopyrightText: 2024-2026 SAP SE or an SAP affiliate company and Greenhouse contributors
// SPDX-License-Identifier: Apache-2.0

package admission

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	greenhouseapis "github.com/cloudoperators/greenhouse/pkg/apis"
	"github.com/cloudoperators/greenhouse/pkg/clientutil"
)

// Webhook for the core Secret type resource.

func SetupSecretWebhookWithManager(mgr ctrl.Manager) error {
	return setupWebhook(mgr,
		&corev1.Secret{},
		webhookFuncs{
			defaultFunc:        DefaultSecret,
			validateCreateFunc: ValidateCreateSecret,
			validateUpdateFunc: ValidateUpdateSecret,
			validateDeleteFunc: ValidateDeleteSecret,
		},
	)
}

//+kubebuilder:webhook:path=/mutate--v1-secret,mutating=true,failurePolicy=ignore,sideEffects=None,groups="",matchPolicy=Exact,resources=secrets,verbs=create;update,versions=v1,name=msecret.kb.io,admissionReviewVersions=v1

func DefaultSecret(_ context.Context, _ client.Client, _ runtime.Object) error {
	return nil
}

//+kubebuilder:webhook:path=/validate--v1-secret,mutating=false,failurePolicy=ignore,sideEffects=None,groups="",matchPolicy=Exact,resources=secrets,verbs=create;update,versions=v1,name=vsecret.kb.io,admissionReviewVersions=v1

func ValidateCreateSecret(_ context.Context, _ client.Client, o runtime.Object) (admission.Warnings, error) {
	secret, ok := o.(*corev1.Secret)
	if !ok {
		return nil, nil
	}
	if err := validateSecretGreenHouseType(secret); err != nil {
		return nil, err
	}
	return nil, validateKubeconfigInSecret(secret)
}

func ValidateUpdateSecret(_ context.Context, _ client.Client, _, o runtime.Object) (admission.Warnings, error) {
	secret, ok := o.(*corev1.Secret)
	if !ok {
		return nil, nil
	}
	if err := validateSecretGreenHouseType(secret); err != nil {
		return nil, err
	}
	return nil, validateKubeconfigInSecret(secret)
}

func ValidateDeleteSecret(_ context.Context, _ client.Client, _ runtime.Object) (admission.Warnings, error) {
	return nil, nil
}

func validateSecretGreenHouseType(secret *corev1.Secret) error {
	// if not greenhouse kubeconfig secret, skip validation
	if secret.Type != greenhouseapis.SecretTypeKubeConfig {
		return nil
	}
	// Check if the secret contains kubeconfig provided by the client
	if !clientutil.IsSecretContainsKey(secret, greenhouseapis.KubeConfigKey) {
		return apierrors.NewInvalid(secret.GroupVersionKind().GroupKind(), secret.GetName(), field.ErrorList{
			field.Required(field.NewPath("data").Child(greenhouseapis.KubeConfigKey),
				"This type of secrets without Data.kubeconfig is invalid."),
		})
	}
	return nil
}

func validateKubeconfigInSecret(secret *corev1.Secret) error {
	switch {
	case clientutil.IsSecretContainsKey(secret, greenhouseapis.KubeConfigKey):
		if len(secret.Data[greenhouseapis.KubeConfigKey]) == 0 {
			return apierrors.NewInvalid(secret.GroupVersionKind().GroupKind(), secret.GetName(), field.ErrorList{
				field.Required(field.NewPath("data").Child(greenhouseapis.KubeConfigKey), "The kubeconfig could not be empty."),
			})
		}
		if err := validateKubeConfig(secret.Data[greenhouseapis.KubeConfigKey]); err != nil {
			return apierrors.NewInvalid(secret.GroupVersionKind().GroupKind(), secret.GetName(), field.ErrorList{
				field.Invalid(field.NewPath("data").Child(greenhouseapis.KubeConfigKey), string(secret.Data[greenhouseapis.KubeConfigKey]),
					"The provided kubeconfig is invalid or not usable."),
			})
		}
	case clientutil.IsSecretContainsKey(secret, greenhouseapis.GreenHouseKubeConfigKey):
		if len(secret.Data[greenhouseapis.GreenHouseKubeConfigKey]) == 0 {
			return apierrors.NewInvalid(secret.GroupVersionKind().GroupKind(), secret.GetName(), field.ErrorList{
				field.Required(field.NewPath("data").Child(greenhouseapis.GreenHouseKubeConfigKey), "The greenhousekubeconfig could not be empty."),
			})
		}
		if err := validateKubeConfig(secret.Data[greenhouseapis.GreenHouseKubeConfigKey]); err != nil {
			return apierrors.NewInvalid(secret.GroupVersionKind().GroupKind(), secret.GetName(), field.ErrorList{
				field.Invalid(field.NewPath("data").Child(greenhouseapis.GreenHouseKubeConfigKey), string(secret.Data[greenhouseapis.GreenHouseKubeConfigKey]),
					"The provided greenhousekubeconfig is invalid or not usable."),
			})
		}
	}
	return nil
}

func validateKubeConfig(kubeconfig []byte) error {
	apiConfig, err := clientcmd.Load(kubeconfig)
	if err != nil {
		return err
	}
	return clientcmd.ConfirmUsable(*apiConfig, apiConfig.CurrentContext)
}
