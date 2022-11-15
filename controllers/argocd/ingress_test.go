package argocd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/types"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/argoproj-labs/argocd-operator/api/v1alpha1"
	argoprojv1alpha1 "github.com/argoproj-labs/argocd-operator/api/v1alpha1"
	"github.com/argoproj-labs/argocd-operator/common"
)

func TestReconcileArgoCD_reconcile_ServerIngress_ingressClassName(t *testing.T) {
	logf.SetLogger(ZapLogger(true))

	nginx := "nginx"

	tests := []struct {
		name             string
		ingressClassName *string
	}{
		{
			name:             "undefined ingress class name",
			ingressClassName: nil,
		},
		{
			name:             "ingress class name specified",
			ingressClassName: &nginx,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			a := makeTestArgoCD(func(a *argoprojv1alpha1.ArgoCD) {
				a.Spec.Server.Ingress.Enabled = true
				a.Spec.Server.Ingress.IngressClassName = test.ingressClassName
			})
			r := makeTestReconciler(t, a)

			err := r.reconcileArgoServerIngress(a)
			assert.NoError(t, err)

			ingress := &networkingv1.Ingress{}
			err = r.Client.Get(context.TODO(), types.NamespacedName{
				Name:      "argocd-server",
				Namespace: testNamespace,
			}, ingress)
			assert.NoError(t, err)
			assert.Equal(t, test.ingressClassName, ingress.Spec.IngressClassName)
		})
	}
}

func TestReconcileArgoCD_reconcile_ServerGRPCIngress_ingressClassName(t *testing.T) {
	logf.SetLogger(ZapLogger(true))

	nginx := "nginx"

	tests := []struct {
		name             string
		ingressClassName *string
	}{
		{
			name:             "undefined ingress class name",
			ingressClassName: nil,
		},
		{
			name:             "ingress class name specified",
			ingressClassName: &nginx,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			a := makeTestArgoCD(func(a *argoprojv1alpha1.ArgoCD) {
				a.Spec.Server.GRPC.Ingress.Enabled = true
				a.Spec.Server.GRPC.Ingress.IngressClassName = test.ingressClassName
			})
			r := makeTestReconciler(t, a)

			err := r.reconcileArgoServerGRPCIngress(a)
			assert.NoError(t, err)

			ingress := &networkingv1.Ingress{}
			err = r.Client.Get(context.TODO(), types.NamespacedName{
				Name:      "argocd-grpc",
				Namespace: testNamespace,
			}, ingress)
			assert.NoError(t, err)
			assert.Equal(t, test.ingressClassName, ingress.Spec.IngressClassName)
		})
	}
}

func TestReconcileArgoCD_reconcile_GrafanaIngress_ingressClassName(t *testing.T) {
	logf.SetLogger(ZapLogger(true))

	nginx := "nginx"

	tests := []struct {
		name             string
		ingressClassName *string
	}{
		{
			name:             "undefined ingress class name",
			ingressClassName: nil,
		},
		{
			name:             "ingress class name specified",
			ingressClassName: &nginx,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			a := makeTestArgoCD(func(a *argoprojv1alpha1.ArgoCD) {
				a.Spec.Grafana.Enabled = true
				a.Spec.Grafana.Ingress.Enabled = true
				a.Spec.Grafana.Ingress.IngressClassName = test.ingressClassName
			})
			r := makeTestReconciler(t, a)

			err := r.reconcileGrafanaIngress(a)
			assert.NoError(t, err)

			ingress := &networkingv1.Ingress{}
			err = r.Client.Get(context.TODO(), types.NamespacedName{
				Name:      "argocd-grafana",
				Namespace: testNamespace,
			}, ingress)
			assert.NoError(t, err)
			assert.Equal(t, test.ingressClassName, ingress.Spec.IngressClassName)
		})
	}
}

func TestReconcileArgoCD_reconcile_PrometheusIngress_ingressClassName(t *testing.T) {
	logf.SetLogger(ZapLogger(true))

	nginx := "nginx"

	tests := []struct {
		name             string
		ingressClassName *string
	}{
		{
			name:             "undefined ingress class name",
			ingressClassName: nil,
		},
		{
			name:             "ingress class name specified",
			ingressClassName: &nginx,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			a := makeTestArgoCD(func(a *argoprojv1alpha1.ArgoCD) {
				a.Spec.Prometheus.Enabled = true
				a.Spec.Prometheus.Ingress.Enabled = true
				a.Spec.Prometheus.Ingress.IngressClassName = test.ingressClassName
			})
			r := makeTestReconciler(t, a)

			err := r.reconcilePrometheusIngress(a)
			assert.NoError(t, err)

			ingress := &networkingv1.Ingress{}
			err = r.Client.Get(context.TODO(), types.NamespacedName{
				Name:      "argocd-prometheus",
				Namespace: testNamespace,
			}, ingress)
			assert.NoError(t, err)
			assert.Equal(t, test.ingressClassName, ingress.Spec.IngressClassName)
		})
	}
}

func TestReconcileApplicationSetService_Ingress(t *testing.T) {
	logf.SetLogger(ZapLogger(true))
	a := makeTestArgoCD()
	obj := v1alpha1.ArgoCDApplicationSet{
		WebhookServer: v1alpha1.WebhookServerSpec{
			Ingress: v1alpha1.ArgoCDIngressSpec{
				Enabled: true,
			},
		},
	}
	a.Spec.ApplicationSet = &obj
	r := makeTestReconciler(t, a)
	ingress := newIngressWithSuffix(common.ApplicationSetServiceNameSuffix, a)
	assert.NoError(t, r.reconcileApplicationSetControllerIngress(a))
	assert.NoError(t, r.Client.Get(context.TODO(), types.NamespacedName{Namespace: ingress.Namespace, Name: ingress.Name}, ingress))
}
