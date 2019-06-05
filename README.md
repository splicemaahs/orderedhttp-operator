# orderedhttp-operator

Example k8s operator using operator-sdk

## What the example looks to accomplish

The goal for this example is to demonstrate a method of launching PODs serially.  Waiting to launch
additional PODs until the currently running POD has reached a **live** state.

## Requirements

```bash
brew install go
brew install operator-sdk

go version
# go version go1.12.5 darwin/amd64
operator-sdk version
# operator-sdk version: v0.8.1, commit: 33b3bfe10176f8647f5354516fff29dea42b6342
```

These were the current brew versions as of June 1st, 2019.

## Clone and Run

If you have a kubernetes cluster available and you want to see this operator in action, follow these
steps.

### Docker Hub Repositories

Logon to your Docker Hub and create yourself two repositories:

- orderedhttp-operator
- nginx-delay

```bash
# change into a directory in your $GOPATH/src/
git clone https://github.com/splicemaahs/orderedhttp-operator.git
cd orderedhttp-operator/nginx-delay
vi Makefile
# change the `splicemaahs` reference to YOUR docker id, so it will push to your newly created
# repository
make build
make push
cd ..
export GO111MODULE=on
operator-sdk build YOURDOCKERID/orderedhttp-operator:latest
docker push YOURDOCKERID/orderedhttp-operator:latest
```

### Create Kubernetes Resources

```bash
kubectl create -f deploy/crds/orderedhttp_v1alpha1_orderedhttp_crd.yaml
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml
kubectl create -f deploy/operator.yaml
kubectl get pods
kubectl create -f deploy/crds/orderedhttp_v1alpha1_orderedhttp_cr.yaml
```

### Check on Resources

```bash
kubectl get pods
kubectl describe OrderedHttp
kubectl logs $(kubectl get pods | grep orderedhttp-operator | tr -s ' ' | cut -d' ' -f1)
```

### Delete Resources

```bash
kubectl delete -f deploy/crds/orderedhttp_v1alpha1_orderedhttp_cr.yaml
kubectl delete -f deploy/operator.yaml

kubectl create -f deploy/crds/orderedhttp_v1alpha1_orderedhttp_crd.yaml
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml

kubectl get pods
```

## Create Operator using operator-sdk

### Create Docker Hub Repositories, if not already existing

Logon to your Docker Hub and create yourself two repositories:

- orderedhttp-operator
- nginx-delay

### Build our custom nginx-delay docker image

```bash
mkdir -p nginx-delay
curl TODO: pull nginx-delay folder from github
cd nginx-delay
vi Makefile
# change the `splicemaahs` reference to YOUR docker id, so it will push to your newly created
# repository
make build
make push
```

### Create new operator

```bash
# change into a directory in your $GOPATH/src/
export GO111MODULE=on
operator-sdk new orderedhttp-operator
cd orderedhttp-operator
```

### Add an api

This output produces some warning lines, this appears to be normal as the resulting operator works without issue.

```bash
operator-sdk add api --api-version=orderedhttp.splicemachine.io/v1alpha1 --kind=OrderedHttp
```

### Add Properties to the api

Create this patch file, then apply to the current sources

```bash
mkdir -p patches
curl TODO: from repository
```

```bash
git apply patches/apicode.patch
# the 'generate k8s' process builds code based on the properties added to the Spec and Status
# sections of ./pkg/apis/orderedhttp/v1alpha1/orderedhttp_types.go'
operator-sdk generate k8s
```

### Add a controller

```bash
operator-sdk add controller --api-version=orderedhttp.splicemachine.io/v1alpha1 --kind=OrderedHttp
```

### Add reconciler code to the controller

```bash
mkdir -p patches
curl TODO: from repository
```

```bash
# ./pkg/controller/orderedhttp/orderedhttp_controller.go
git apply patches/controllercode.patch
```

### Update operator deploy for docker image name

```bash
vi deploy/operator.yaml
# change the 'splicemaahs' reference to your own docker ID
```

### Build the operator docker image

```bash
go mod vendor # <- you need only run this once, and can rebuild with the 'build' command
operator-sdk build YOURDOCKERID/orderedhttp-operator:latest
# this process will fail on go syntax errors as it builds the code as part of the docker image build.
docker push YOURDOCKERID/orderedhttp-operator:latest
```

### Create Kubernetes Resources (same as above)

```bash
# this installs/defines the Custom Resource Definition
kubectl create -f deploy/crds/orderedhttp_v1alpha1_orderedhttp_crd.yaml
# these create the ability for the operator to interact with the k8s controller
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml
# this deploys the operator pod itself
kubectl create -f deploy/operator.yaml
kubectl get pods
# this creates an instance of our custom 'Kind'
kubectl create -f deploy/crds/orderedhttp_v1alpha1_orderedhttp_cr.yaml
```

### Check on Resources (same as above)

```bash
kubectl get pods
kubectl describe OrderedHttp
kubectl logs $(kubectl get pods | grep orderedhttp-operator | tr -s ' ' | cut -d' ' -f1)
```

### Delete Resources (same as above)

```bash
kubectl delete -f deploy/crds/orderedhttp_v1alpha1_orderedhttp_cr.yaml
kubectl delete -f deploy/operator.yaml

kubectl create -f deploy/crds/orderedhttp_v1alpha1_orderedhttp_crd.yaml
kubectl create -f deploy/service_account.yaml
kubectl create -f deploy/role.yaml
kubectl create -f deploy/role_binding.yaml

kubectl get pods
```
