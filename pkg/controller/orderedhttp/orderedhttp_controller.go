package orderedhttp

import (
	"context"
	"strconv"

	orderedhttpv1alpha1 "github.com/splicemaahs/orderedhttp-operator/pkg/apis/orderedhttp/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_orderedhttp")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new OrderedHttp Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileOrderedHttp{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("orderedhttp-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource OrderedHttp
	err = c.Watch(&source.Kind{Type: &orderedhttpv1alpha1.OrderedHttp{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner OrderedHttp
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &orderedhttpv1alpha1.OrderedHttp{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileOrderedHttp implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileOrderedHttp{}

// ReconcileOrderedHttp reconciles a OrderedHttp object
type ReconcileOrderedHttp struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a OrderedHttp object and makes changes based on the state read
// and what is in the OrderedHttp.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileOrderedHttp) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("1. Reconciling OrderedHttp")

	// Fetch the OrderedHttp instance
	orderedHttp := &orderedhttpv1alpha1.OrderedHttp{}
	err := r.client.Get(context.TODO(), request.NamespacedName, orderedHttp)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// List all pods owned by this OrderedHttp instance
	lbls := labels.Set{
		"app":     orderedHttp.Name,
		"version": "0.1",
	}
	var allPodsReady bool
	allPodsReady = true
	existingPods := &corev1.PodList{}
	err = r.client.List(context.TODO(),
		&client.ListOptions{
			Namespace:     request.Namespace,
			LabelSelector: labels.SelectorFromSet(lbls),
		},
		existingPods)
	if err != nil {
		reqLogger.Error(err, "failed to list existing pods in the orderedHttp pod")
		return reconcile.Result{}, err
	}
	existingPodNames := []string{}

	// Count the pods that are pending or running as available
	for _, pod := range existingPods.Items {
		reqLogger.Info("1.1 Loop Pods", "PodName: ", pod.Name)
		if pod.GetObjectMeta().GetDeletionTimestamp() != nil {
			continue
		}
		reqLogger.Info("1.2 Pod Phase", "Phase: ", pod.Status.Phase)
		if pod.Status.Phase == corev1.PodPending || pod.Status.Phase == corev1.PodRunning {
			existingPodNames = append(existingPodNames, pod.GetObjectMeta().GetName())
		}
		if pod.Status.Phase != corev1.PodRunning {
			allPodsReady = false
		}
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.Ready == false {
				allPodsReady = false
			}
			reqLogger.Info("1.2.1 Container Statuses: ", "Ready: ", strconv.FormatBool(containerStatus.Ready), "All Pods Ready: ", strconv.FormatBool(allPodsReady))
		}
	}
	reqLogger.Info("2. Checking orderedHttp", "expected size", orderedHttp.Spec.Replicas, "Pod.Names", existingPodNames)

	// // Update the status if necessary
	// status := orderedhttpv1alpha1.OrderedHttpStatus{
	// 	PodNames: existingPodNames,
	// }
	// // reqLogger.Info("Direct Updating orderedHttp status")
	// // orderedHttp.Status = status
	// if !reflect.DeepEqual(orderedHttp.Status, status) {
	// 	reqLogger.Info("Updating orderedHttp status")
	// 	orderedHttp.Status = status
	// 	err := r.client.Update(context.TODO(), orderedHttp)
	// 	if err != nil {
	// 		reqLogger.Error(err, "failed to update the orderedHttp pod")
	// 		return reconcile.Result{}, err
	// 	}
	// }

	// List the pods for this deployment
	// podList := &corev1.PodList{}
	// podNames := getPodNames(podList.Items)
	// orderedHttp.Status.PodNames = podNames
	// reqLogger.Info("0. Setting Pod Names in status", "Pod.Names", existingPodNames)
	orderedHttp.Status.PodNames = existingPodNames
	err = r.client.Status().Update(context.TODO(), orderedHttp)
	if err != nil {
		reqLogger.Error(err, "failed to update the orderedHttp pod")
		return reconcile.Result{}, err
	}
	// Update status.PodNames if needed
	// if !reflect.DeepEqual(podNames, orderedHttp.Status.PodNames) {
	// 	orderedHttp.Status.PodNames = podNames
	// 	err := r.client.Update(context.TODO(), orderedHttp)
	// 	if err != nil {
	// 		// log.Printf("failed to update node status: %v", err)
	// 		reqLogger.Error(err, "failed to update the orderedHttp pod")
	// 		return reconcile.Result{}, err
	// 	}
	// }

	// Scale Down Pods
	if int32(len(existingPodNames)) > orderedHttp.Spec.Replicas {
		// When scaling down, just delete one, and allow the process to continue, the next loop will determine additional removals
		reqLogger.Info("2.1 Deleting a pod in orderedHttp set", "expected size", orderedHttp.Spec.Replicas, "Pod.Names", existingPodNames)
		pod := existingPods.Items[0]
		err = r.client.Delete(context.TODO(), &pod)
		if err != nil {
			reqLogger.Error(err, "failed to delete a pod")
			return reconcile.Result{}, err
		}
	}

	// Scale Up Pods
	if int32(len(existingPodNames)) < orderedHttp.Spec.Replicas {
		// When scaling up, just add one, and allow the process to contiue, the next loop will add more pods if needed.
		reqLogger.Info("2.2 Pod Ready Check", "All Pods Ready: ", strconv.FormatBool(allPodsReady))
		if allPodsReady == true {
			reqLogger.Info("2.2.1 Adding a pod in orderedHttp set", "expected size", orderedHttp.Spec.Replicas, "Pod.Names", existingPodNames)
			pod := newPodForCR(orderedHttp)
			if err := controllerutil.SetControllerReference(orderedHttp, pod, r.scheme); err != nil {
				reqLogger.Error(err, "unable to set owner reference on new pod")
				return reconcile.Result{}, err
			}
			reqLogger.Info("2.2.2 Create Pod")
			err = r.client.Create(context.TODO(), pod)
			if err != nil {
				reqLogger.Error(err, "failed to create a pod")
				return reconcile.Result{}, err
			}
		}
	}
	return reconcile.Result{Requeue: true}, nil

}

// getPodNames returns the pod names of the array of pods passed in.
func getPodNames(pods []corev1.Pod) []string {
	var podNames []string
	for _, pod := range pods {
		podNames = append(podNames, pod.Name)
	}
	return podNames
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *orderedhttpv1alpha1.OrderedHttp) *corev1.Pod {
	labels := map[string]string{
		"app":     cr.Name,
		"version": "0.1",
	}

	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: cr.Name + "-",
			Namespace:    cr.Namespace,
			Labels:       labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx-delay",
					Image: "splicemaahs/nginx-delay:latest",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 80,
						},
					},
					ReadinessProbe: &corev1.Probe{
						Handler: corev1.Handler{
							TCPSocket: &corev1.TCPSocketAction{
								Port: intstr.IntOrString{
									Type:   intstr.Int,
									IntVal: 80,
								},
							},
						},
						InitialDelaySeconds: 5,
						PeriodSeconds:       10,
					},
				},
			},
		},
	}
}
