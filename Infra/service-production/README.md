# Runbook Step Production env on K8s
## Requirements
- Setup and running k8s cluster 
- Makesure cluster ready to use
- Install helm cli

## Step 1 (install nginx ingress controller)
> Install nginx ingress controller
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.1/deploy/static/provider/cloud/deploy.yaml
```
Refrence : [Link](https://kubernetes.github.io/ingress-nginx/deploy/#quick-start)

## Step 2 (install kubernetes metrics-server)
> Install kubernetes metrics-server
```
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/high-availability-1.21+.yaml
```
Refrence : [Link](https://github.com/kubernetes-sigs/metrics-server)

## Step 3 (install & setup mongodb server)
> Add helm chart repo
```
helm repo add mongodb https://mongodb.github.io/helm-charts
helm repo update
```
> Enter to config location
```
cd Infra/service-production/mongodb
```
> Install mongodb operator
```
helm install community-operator mongodb/community-operator --namespace mongodb --create-namespace -f values.yaml
```
> Create mongodb server
```
kubectl apply -f mongodb-db-virtual-pet.yaml -n mongodb
```
> Noted

- For change credential config on mongodb server, update on this file `Infra/service-production/mongodb/mongodb-db-virtual-pet.yaml`, find `users,username`, and `Secret`declaration.
- users:
```
  users:
    - name: pet
```
- username:
```
  prometheus:
    username: admin
```
- Secret:
```
---
apiVersion: v1
kind: Secret
metadata:
  name: virtual-pet-password
type: Opaque
stringData:
  password: "password123"

---
# Secret holding the prometheus metrics endpoint HTTP Password.
---
apiVersion: v1
kind: Secret
metadata:
  name: metrics-endpoint-password
type: Opaque
stringData:
  password: "password123"
```

Refrence : [link](https://github.com/mongodb/mongodb-kubernetes-operator/blob/master/README.md)

## Step 4 (install virtual-pet app)
> Copy and update Credential for app
- Copy env file for virtual-pet app 
```
cp Infra/service-production/virtual-pet/config/.env-template Infra/service-production/virtual-pet/config/.env
```
- Check credential mongodb, copy value `connectionString.standard`, and update credential on `Infra/service-production/virtual-pet/config/.env`
```
kubectl get secrets -n mongodb virtual-pet-mongodb-virtual-pet-pet -o json | jq -r '.data | with_entries(.value |= @base64d)'
```
- Install virtual-pet app
```
cd Infra/service-production/virtual-pet
helm upgrade --install virtual-pet .
```

> Noted

- Update domain on ingress config on this file `Infra/service-production/virtual-pet/values.yaml` find this part and update `host`.
```
ingress:
  enabled: true
  className: "nginx"
  annotations:
    kubernetes.io/ingress.class: nginx
  hosts:
    - host: virtual-pet.k8s.orb.local
      paths:
        - path: /
          pathType: Prefix
```

## Step 5 (install prometheus operator)
> Add helm chart repo
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```
> Enter to config location
```
cd Infra/service-production/prometheus
```
> Install prometheus-server
```
helm install prometheus-server prometheus-community/prometheus -f values.yaml
```
Refrence : [link](https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/README.md)