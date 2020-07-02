/*


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
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	alibabacloudv1alpha1 "github.com/zzxwill/aliyun-ecs-k8s-operator/api/v1alpha1"
)

// ECSInstanceReconciler reconciles a ECSInstance object
type ECSInstanceReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=alibabacloud.zhouzhengxi.com,resources=ecsinstances,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=alibabacloud.zhouzhengxi.com,resources=ecsinstances/status,verbs=get;update;patch

func (r *ECSInstanceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("ecsinstance", req.NamespacedName)

	// your logic here
	ecsInstance := &alibabacloudv1alpha1.ECSInstance{}

	if err := r.Get(ctx, req.NamespacedName, ecsInstance); err != nil {
		if err := client.IgnoreNotFound(err); err != nil {
			log.Info("AAAAAAAAAAAA")
			return ctrl.Result{}, nil
		} else {
			log.Info("BBBBBBBBBBBB")
			return ctrl.Result{}, err
		}
	}

	accessKeyID := "xxx"
	accessKeySecret := "xxxxx"

	regionID := ecsInstance.Spec.RegionID

	client, err := ecs.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)

	request := ecs.CreateCreateInstanceRequest()
	request.Scheme = "https"

	request.InstanceType = ecsInstance.Spec.InstanceType
	request.ImageId = ecsInstance.Spec.ImageID
	request.InstanceName = ecsInstance.Spec.InstanceName

	response, err := client.CreateInstance(request)

	if err != nil {
		log.Error(err, "Failed to create an instance")
	} else {
		log.Info(response.String())
		ecsInstance.Status.InstanceID = response.InstanceId
	}

	return ctrl.Result{}, nil
}

func (r *ECSInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&alibabacloudv1alpha1.ECSInstance{}).
		Complete(r)
}
