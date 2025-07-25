[docker](/docker.md)

# kubernetes

> kubernetes (K8s) is an open-source platform for managing containerized applications at scale

- automates tasks like deployment, scaling and management of containerized applications
- self-healing: if a container/pod fails, kubernetes automatically replaces it
- declarative configuration and automation
- able to work with multiple cloud-providers simultaneously

[kubernetes documentation](https://kubernetes.io/docs/)
[kubernetes tutorials by katacoda](https://www.katacoda.com/courses/kubernetes)

> [!TIP]
> The easiest way to install Kubernetes on a low-spec computer is by using **`kubectl` + `minikube`**.
> `kubectl` is the command-line tool to manage Kubernetes and `minikube` runs a lightweight local cluster

- resources in kubernetes run on a private, isolated network by default
  - they can communicate with other resources within kubernetes unless explicitly configured
  - they are **not** exposed to anything outside kubernetes (e.g. from your browser or mobile app)

## key components

- pods: smallest deployable unit in kubernetes
  - group of one or more containers
  - shared storage, network, and configuration for containers
  - e.g. a web application might have a frontend pod (running nginx) and a backend pod (running node.js with postgresql)
- node: physical or virtual machine that runs components required to execute and manage containers within pods
  - managed by the control plane
  - each node can host multiple pods
  - each node runs:
    - kubelet: communicates with the control plane, manages pod lifecycle
    - kube-proxy: manages networking, forwards traffic to pods
    - container runtime (e.g., containerd, Docker): runs the actual containers
  - node roles: all machines in the cluster are nodes
    - control plane components typically run on designated nodes (control plane nodes)
    - worker node
      - runs user applications (pods)
      - runs workloads
    - control plane (master) node
      - runs core Kubernetes components (API server, scheduler, controller manager, etc.)
      - usually replicated (at least 3 replicas for HA in production)
- cluster: group of nodes, highest-level structure in kubernetes
- object: persistent entity representing the desired state of a component in the cluster (e.g., Pod, Deployment, Service)
- workload: application(s) running on the cluster
- service: method for exposing a network application that is running as one or more pods in your cluster

- imagine a kubernetes cluster as a city:
  - cluster = the entire city
  - node = a building in the city
  - pod = a single apartment in the building where people (containers) live

![single kubernetes node](./images/kubernetes_node.png)

kubernetes cluster architecture:

- master node (control plane): manages the cluster and schedules workloads
  - kube-apiserver: provides the interaction for management tools (e.g. kubectl or kubernetes dashboard)
  - etcd: key-value store that holds the **desired** state of all kubernetes objects
    - used by control plane to ensure current state matches desired state
    - can also be backed up to restore the cluster state
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

## kubectl

> CLI tool that enables cluster management

- view nodes: `kubectl get nodes`
- apply YAML file: `kubectl apply -f <filename>.yaml`
- open shell inside running container: `kubectl exec -it <nextjs-pod-name> -- /bin/sh`
- create Secret: `sudo kubectl create secret generic peso-secrets --from-env-file=.env`
- generate Secret yaml file: `kubectl get secret nextjs-env -o yaml`
- ensure Kubernetes cluster is running: `kubectl cluster-info`
- check the pod logs: `kubectl logs <pod-name>`
- check the service status: `kubectl describe service nginx-service`
- create Kubernetes Secret from .env: `kubectl create secret generic peso-secrets --from-env-file=.env.local`
- update Kubernetes Secret :`kubectl delete secret peso-secrets && kubectl create secret generic peso-secrets --from-env-file=.env.local`
- delete resources created by YAML files: `kubectl delete -f file1.yaml -f file2.yaml`
- delete pod: `kubectl delete pod test-nginx`

## kinds

> type of kubernetes object, used in manifest files

- kubernetes object: persistent entity in kubernetes system that represents the desired state of the cluster
- manifest: declarative configuration file (YAML or JSON) that describes the desired state of a kubernetes object
  - used to manage and configure the cluster

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
    - name: nginx
      image: nginx
```

## resources

> building blocks used to define the desired state of the cluster

- each resource is defined by `kind`

> [!TIP]
> run `kubectl api-resources` to show a list of all api resources

### workload resources

- Pod
- Deployment
  - most common way to manage pods
  - ensures specified number of pod replicas are running
  - supports rolling updates and rollbacks
  - used for stateless applications
  - example: deploy a web application with 3 replicas
- ReplicaSet
  - maintains a stable set of replica pods
  - usually managed by deployments
  - ensures specified number of pods are running at all times
- StatefulSet
  - for applications that need stable, unique network identifiers
  - persistent storage per pod
  - example: databases like mongodb, mysql
- DaemonSet
  - ensures all nodes run a copy of a specific pod
  - useful for node monitoring, log collection
  - example: running a logging agent on every node
- Job
  - for one-time tasks
  - ensures pods run to completion
  - example: batch processing, data migration
- CronJob
  - creates jobs on a schedule
  - like unix crontab
  - example: regular backup jobs, scheduled reports

### resources for networking and service discovery

- Service: exposes a group of pods as a network service
  - types:
    - ClusterIP (default): internal communication inside the cluster
    - NodePort: Exposes port on each node
    - LoadBalancer: External load balancer
    - ExternalName: DNS alias
- Ingress: manages external HTTP/HTTPS traffic to services
  - provides routing rules and supports SSL termination
- NetworkPolicy: controls traffic between pods based on labels and rules
  - doesn't control external access like traffic from example.com

### configuration and secrets management

- ConfigMap: stores non-sensitive configuration data as key-value pairs
  - example: application settings
- Secret: stores sensitive data
  - example: API keys, passwords, certificates

### storage management

- PersistentVolume: physical storage resource in the cluster
- PersistentVolumeClaim: request for storage by an application
- StorageClass: defines different storage types (e.g. SSD, HDD)

### cluster management

- Namespace: creates isolated environments within the cluster
  - groups related resources (pods, services, deployments, etc)
  - resources go to `default` namespace if a namespace isn't specified
- ResourceQuota: limits CPI, memory and storage per namespace
- LimitRange: sets minimum and maximum resource limits for pods

## manifest

> declarative configuration file written in yaml

every manifest must include at least these top‑level fields:

- `apiVersion`: kubernetes api version
- `kind`: type of resource you're defining
- `metadata`: data about the resource
  - `name` (required)
  - `namespace` (optional)
  - `labels` and `annotations` (optional)
- `spec`: specification of your desired state, the fields here vary depending on `kind`
  - `replicas`: desired number of copies of a pod
  - `selector`: defines which pods the controller will manage
    - `matchLabels`: matches `labels` in `metadata`
  - `containers`: list of containers that will run inside the pod
  - `name`: container name
  - `ports`:
    - `containerPort`: port that the container listens on internally
    - `protocol`: protocol to be used by container port, usually `TCP` or `UDP` (default)

Pod YAML example:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
  labels:
    app: nginx
spec:
  containers:
    - name: nginx
      image: nginx:1.25
      ports:
        - containerPort: 80
```

deployment + service example:

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:1.25
          ports:
            - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
spec:
  type: LoadBalancer
  selector:
    app: nginx
  ports:
    - port: 80
      targetPort: 80
```

- deployment: ensures 3 replicas of your pod are always running and handles rolling updates
- service: exposes those pods on port 80, optionally as a cloud load‑balancer if your environment supports it

## networking

- pod-to-pod communication: all pods within the cluster can communicate directly with each other
  - each pod has its own private ip
  - pods **don't** have their own public ip by default
- pod-to-service communication:
  - `DNS`
  - `ClusterIP`
  - a service provides a stable ip and a dns name to a group of pods
    - dns name can't be accessed externally
- service-to-service communication: use internal DNS
- external-to-service communication
  - `NodePort`: exposes a service on a specific port of each worker node
  - `LoadBalancer`: assigns a public IP (if running on a cloud provider)
  - `Ingress`: routes traffic to services based on URL paths or domains

## pod lifecycle

> phases a pod goes through from creation to termination

- pending: pod has been accepted by the kubernetes cluster, but one or more of its containers hasn't started yet
- running: pod has been bound to a node and all containers have been created
- succeeded: all containers in the pod have terminated successfully
- failed: all containers have terminated, but at least one has failed
- unknown: the state of the pod couldn't be determined

## helm

> package manager for kubernetes

- chart: helm package
- repository: place where charts can be collected and shared
- release: instance of chart running in kubernetes cluster

- one chart can be installed multiple times into the same cluster
  - each time it is installed, a new release is created

### cert-manager

> kubernetes add-on that automates the management of TSL/SSL certificates

- can request, renew and issue certificates from different Certificate Authorities (CAs), like Let's Encrypt
- to secure an application with HTTPS, an SSL/TLS certificate is required
- cert-manager is responsible for:
  - creation of certificates
  - uses an Issuer or ClusterIssuer to obtain a certificate
    - Issuer: resource that defines how certificates are requested from a Certificate Authority (e.g., Let’s Encrypt)
  - stores the certificate in a kubernetes Secret
- how the certificate is used:
  - the Secret containing the certificate is referenced in an Ingress resource
  - the Ingress resource is configured to use TLS (HTTPS)
  - when a user accesses the domain, the browser verifies the certificate
  - if everything is valid, the website loads over HTTPS

- install `cert-manager`: `kubectl apply -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml`
- check if it's running: `kubectl get pods -n cert-manager`
- check if certificate was issued: `kubectl get certificate`
- check if Secret was created: `kubectl get secret pesodevops-tls`

## tools

- kubeadm: set up and manage kubernetes clusters easily

### k3s

> lightweight kubernetes that's easy to install

- to run any kubernetes command, add `k3s`
  - e.g.: `sudo k3s kubectl get nodes`
