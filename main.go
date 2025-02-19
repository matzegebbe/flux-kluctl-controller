/*
Copyright 2022.

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

package main

import (
	"fmt"
	ssh_pool "github.com/kluctl/kluctl/v2/pkg/git/ssh-pool"
	"k8s.io/client-go/kubernetes"
	"os"

	helper "github.com/fluxcd/pkg/runtime/controller"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/kluctl/flux-kluctl-controller/controllers"

	flag "github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/fluxcd/pkg/runtime/acl"
	"github.com/fluxcd/pkg/runtime/client"
	"github.com/fluxcd/pkg/runtime/leaderelection"
	"github.com/fluxcd/pkg/runtime/logger"
	"github.com/fluxcd/pkg/runtime/pprof"
	"github.com/fluxcd/pkg/runtime/probes"
	sourcev1 "github.com/fluxcd/source-controller/api/v1beta2"

	fluxkluctliov1alpha1 "github.com/kluctl/flux-kluctl-controller/api/v1alpha1"
	kluctliov1alpha1 "github.com/kluctl/flux-kluctl-controller/api/v1alpha1"
	// +kubebuilder:scaffold:imports
)

const controllerName = "flux-kluctl-controller"

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(sourcev1.AddToScheme(scheme))
	utilruntime.Must(kluctliov1alpha1.AddToScheme(scheme))
	utilruntime.Must(kluctliov1alpha1.AddToScheme(scheme))
	utilruntime.Must(fluxkluctliov1alpha1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

func main() {
	var (
		metricsAddr           string
		healthAddr            string
		concurrent            int
		clientOptions         client.Options
		logOptions            logger.Options
		leaderElectionOptions leaderelection.Options
		aclOptions            acl.Options
		watchAllNamespaces    bool
		httpRetry             int
		defaultServiceAccount string
	)

	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&healthAddr, "health-addr", ":9440", "The address the health endpoint binds to.")
	flag.IntVar(&concurrent, "concurrent", 4, "The number of concurrent kluctl deployments.")
	flag.BoolVar(&watchAllNamespaces, "watch-all-namespaces", true,
		"Watch for custom resources in all namespaces, if set to false it will only watch the runtime namespace.")
	flag.IntVar(&httpRetry, "http-retry", 9, "The maximum number of retries when failing to fetch artifacts over HTTP.")
	flag.StringVar(&defaultServiceAccount, "default-service-account", "", "Default service account used for impersonation.")

	clientOptions.BindFlags(flag.CommandLine)
	logOptions.BindFlags(flag.CommandLine)
	leaderElectionOptions.BindFlags(flag.CommandLine)
	aclOptions.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(logger.NewLogger(logOptions))

	watchNamespace := ""
	if !watchAllNamespaces {
		watchNamespace = os.Getenv("RUNTIME_NAMESPACE")
	}

	restConfig := client.GetConfigOrDie(clientOptions)
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	mgr, err := ctrl.NewManager(restConfig, ctrl.Options{
		Scheme:                        scheme,
		MetricsBindAddress:            metricsAddr,
		HealthProbeBindAddress:        healthAddr,
		Port:                          9443,
		LeaderElection:                leaderElectionOptions.Enable,
		LeaderElectionReleaseOnCancel: leaderElectionOptions.ReleaseOnCancel,
		LeaseDuration:                 &leaderElectionOptions.LeaseDuration,
		RenewDeadline:                 &leaderElectionOptions.RenewDeadline,
		RetryPeriod:                   &leaderElectionOptions.RetryPeriod,
		LeaderElectionID:              fmt.Sprintf("%s-leader-election", controllerName),
		Namespace:                     watchNamespace,
		Logger:                        ctrl.Log,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	probes.SetupChecks(mgr, setupLog)
	pprof.SetupHandlers(mgr, setupLog)

	eventRecorder := mgr.GetEventRecorderFor(controllerName)

	metricsH := helper.MustMakeMetrics(mgr)
	sshPool := &ssh_pool.SshPool{}

	r := controllers.KluctlDeploymentReconciler{
		ControllerName:        controllerName,
		DefaultServiceAccount: defaultServiceAccount,
		RestConfig:            restConfig,
		Client:                mgr.GetClient(),
		ClientSet:             clientSet,
		Scheme:                mgr.GetScheme(),
		EventRecorder:         eventRecorder,
		MetricsRecorder:       metricsH.MetricsRecorder,
		NoCrossNamespaceRefs:  aclOptions.NoCrossNamespaceRefs,
		SshPool:               sshPool,
	}

	if err = r.SetupWithManager(mgr, controllers.KluctlDeploymentReconcilerOpts{
		MaxConcurrentReconciles: concurrent,
		HTTPRetry:               httpRetry,
	}); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", kluctliov1alpha1.KluctlDeploymentKind)
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
