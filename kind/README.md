
# 单节点

kind create cluster --image=kindest/node-arm64:v1.22.0 --name=dev

darren@darrendeMacBook-Pro GolangStudy % kubectl get node
NAME                STATUS   ROLES                  AGE   VERSION
dev-control-plane   Ready    control-plane,master   32s   v1.22.0

kubectl get pod -n kube-system

# 多节点搭建

删除
kind delete cluster --name=dev4



kind create cluster --config multi-node-config.yaml --image=kindest/node-arm64:v1.22.0  --name=dev4
 
kind cluster-info --context kind-dev4

kubectl get node 

## 三主三从

kind create cluster --config ha-config.yaml --image=kindest/node-arm64:v1.22.0  --name=dev6




