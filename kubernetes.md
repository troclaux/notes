
# kubernetes

[kubernetes documentation](https://kubernetes.io/docs/)
[kubernetes tutorials by katacoda](https://www.katacoda.com/courses/kubernetes)
[docker](/docker.md)

- kubernetes (K8s) is an open-source platform for managing containerized applications at scale
- automates tasks like deployment, scaling, and maintaining containerized applications

- self-healing: if a pod fails, kubernetes automatically replaces it

## Key Components

- workload: application(s) running on the cluster
- pods: smallest deployable unit in kubernetes
  - can contain 1+ containers
  - shared storage, network, and configuration for containers
  - e.g. pod that contains a web server container and a logging container
- node: single physical/visual machine that runs components required to execute/manage containers within pods
  - each node can host multiple pods
  - node types:
    - worker node
      - higher workload
      - higher compute
    - master node
      - more important
      - should have at least 3 replicas
- cluster: group of nodes
- object: cluster's desired state

![single kubernetes node](./images/kubernetes_node.png)

- kubernetes cluster consists of:
  - master node (control plane): manages the cluster and schedules workloads
    - API server: provides the interaction for management tools (e.g. kubectl or kubernetes dashboard)
    - etcd: key-value storate that holds the current state of the kubernetes cluster
      - basically a cluster backup
    - scheduler: allocates unscheduled pods on different nodes based on the workload
    - controller-manager: set of containers that manages the cluster with actions (e.g. replicating pods)
      - checks if something needs to be repaired or if a container died and needs to be restarted
  - worker nodes: run the application workloads
    - kubelet
      - communicates with control plane to receive and execute workload instructions
      - ensures containers are running and healthy
      - reports node status and pod health back to control plane
    - container runtime: manages lifecycle of container
      - examples: docker, containerd
    - kube-proxy: manages networking rules so that pods and services communicate effectively
    - pods: smallest deployable unit in kubernetes

TODO: explain workload resources and Services/Deployments/ConfigMaps/Secrets/Ingress
TODO: organize explanation from this point below

master node/control plane ensures that your applications (pods) are placed on the appropriate worker nodes

- workload resources:
  - deployment
  - ReplicaSet (RS)
  - StatefulSet
  - DaemonSet
  - jobs
  - automatic cleanup for finished jobs
  - CronJob
  - Replication Controller (RC)

- Services
  - Expose pods to the network
  - Example: A ClusterIP service routes traffic to a set of web server pods inside the cluster
- Deployments
  - Define how pods should run and how many replicas you want
  - Example: A Deployment might specify 3 replicas of your web server
- ConfigMaps & Secrets
  - Manage configuration and sensitive data
  - Example: Store database credentials securely using Secrets
- Ingress
  - Manages external access to your services
  - Example: Configure a domain name and SSL for your web app

## Core Concepts with Examples

- nodes: nodes are the machines in your cluster
  - each node runs:
    - Kubelet: Agent that communicates with the master
    - Container Runtime: Like Docker, to run containers
    - Kube-proxy: Manages networking
- YAML Manifest: Kubernetes resources are defined using YAML files
  - run YAML file with: `kubectl apply -f <filename>.yaml`

Pod YAML example:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - name: nginx-container
    image: nginx
```

## kubectl

- view nodes: `kubectl get nodes`
- view pods: `kubectl get pods`

```bash
kubectl scale deployment my-deployment --replicas=5
```

## YAML

scale applications up or down by adjusting the replicas in the deployment

```yaml
spec:
  replicas: 5
```
