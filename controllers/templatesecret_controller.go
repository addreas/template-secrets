/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	api "github.com/addreas/template-secrets/api/v1alpha1"
)

// TemplateSecretReconciler reconciles a TemplateSecret object
type TemplateSecretReconciler struct {
	client.Client
	Log         logr.Logger
	SchemeField *runtime.Scheme
}

// Scheme has to be a method for controllerutil to be happy about it
func (r *TemplateSecretReconciler) Scheme() *runtime.Scheme {
	return r.SchemeField
}

//+kubebuilder:rbac:groups=addem.se,resources=templatesecrets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=addem.se,resources=templatesecrets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=addem.se,resources=templatesecrets/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch

func (r *TemplateSecretReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("templatesecret", req.NamespacedName)

	// your logic here
	var tSecret api.TemplateSecret
	if err := r.Get(ctx, req.NamespacedName, &tSecret); err != nil {
		log.Error(err, "unable to fetch TemplateSecret")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var secret corev1.Secret
	secret.Name = tSecret.Spec.SecretName
	secret.Namespace = tSecret.Namespace

	if err := controllerutil.SetControllerReference(&tSecret, &secret, r.Scheme()); err != nil {
		return ctrl.Result{}, err
	}

	template, err := getTemplate(ctx, r, tSecret)
	if err != nil {
		return ctrl.Result{}, err
	}

	replacements, err := getReplacements(ctx, r, tSecret)
	if err != nil {
		return ctrl.Result{}, err
	}

	controllerutil.CreateOrUpdate(ctx, r, &secret, func() error {
		return modifySecret(&secret, template, replacements)
	})

	return ctrl.Result{}, nil
}

func modifySecret(secret *corev1.Secret, template map[string]string, replacements []Replacement) error {
	secret.StringData = make(map[string]string)
	for key, val := range template {
		secret.StringData[key] = applyReplacements(val, replacements)
	}
	return nil
}

// Replacement contains a resolved replacement match/replacement pair
type Replacement struct {
	Match       api.MatchSpec
	Replacement string
}

func getReplacements(ctx context.Context, r *TemplateSecretReconciler, tSecret api.TemplateSecret) ([]Replacement, error) {
	ret := make([]Replacement, len(tSecret.Spec.Replacements))
	for _, val := range tSecret.Spec.Replacements {
		repl, err := getReplacement(ctx, r, tSecret, val.Replacement)
		if err != nil {
			return nil, err
		}
		ret = append(ret, Replacement{
			Match:       val.Match,
			Replacement: repl,
		})
	}
	return ret, nil
}

func applyReplacements(template string, replacements []Replacement) string {
	ret := template
	for _, val := range replacements {
		ret = strings.ReplaceAll(ret, val.Match.Exact, val.Replacement)
	}
	return ret
}

func getReplacement(ctx context.Context, r *TemplateSecretReconciler, tSecret api.TemplateSecret, source api.ReplacementSource) (string, error) {
	switch {
	case source.SecretKeyRef.Name != "":
		var secret corev1.Secret

		nsName := types.NamespacedName{
			Name:      source.SecretKeyRef.Name,
			Namespace: tSecret.Namespace,
		}

		if err := r.Get(ctx, nsName, &secret); err != nil {
			return "", err
		}

		res, found := secret.Data[source.SecretKeyRef.Key]
		if !found {
			return "", fmt.Errorf("Missing key %s in secret %s", source.SecretKeyRef.Key, source.SecretKeyRef.Name)
		}

		return string(res), nil

	case source.ConfigMapKeyRef.Name != "":
		var configMap corev1.ConfigMap

		nsName := types.NamespacedName{
			Name:      source.ConfigMapKeyRef.Name,
			Namespace: tSecret.Namespace,
		}

		if err := r.Get(ctx, nsName, &configMap); err != nil {
			return "", err
		}

		res, err := configMap.Data[source.ConfigMapKeyRef.Key]
		if err {
			return "", fmt.Errorf("Missing key %s in ConfigMap %s", source.ConfigMapKeyRef.Key, source.ConfigMapKeyRef.Name)
		}

		return res, nil
	default:
		return "", fmt.Errorf("Invalid replacement config")
	}
}

func getTemplate(ctx context.Context, r *TemplateSecretReconciler, tSecret api.TemplateSecret) (map[string]string, error) {
	switch {
	case tSecret.Spec.Template.Inline != nil:
		return tSecret.Spec.Template.Inline, nil
	case tSecret.Spec.Template.Secret.Name != "":
		var secret corev1.Secret

		nsName := types.NamespacedName{
			Name:      tSecret.Spec.Template.Secret.Name,
			Namespace: tSecret.Namespace,
		}

		if err := r.Get(ctx, nsName, &secret); err != nil {
			return nil, err
		}

		return secret.StringData, nil
	case tSecret.Spec.Template.ConfigMap.Name != "":
		var configMap corev1.ConfigMap

		nsName := types.NamespacedName{
			Name:      tSecret.Spec.Template.ConfigMap.Name,
			Namespace: tSecret.Namespace,
		}

		if err := r.Get(ctx, nsName, &configMap); err != nil {
			return nil, err
		}

		return configMap.Data, nil
	default:
		return nil, fmt.Errorf("TemplateSecret had no source")
	}
}

func (r *TemplateSecretReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&api.TemplateSecret{}).
		Owns(&corev1.Secret{}).
		Complete(r)
}
