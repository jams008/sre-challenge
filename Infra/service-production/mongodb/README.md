# if see error
```
Error: INSTALLATION FAILED: An error occurred while checking for chart dependencies. You may need to run `helm dependency build` to fetch missing dependencies: found in Chart.yaml, but missing in charts/ directory: community-operator-crds
```
fix using : `helm dependency build`


helm install mongodb-ce-operator . --create-namespace --namespace mongodb


mongodb+srv://admin:password123@localhost:49688/virtual-pet

mongodb://admin:password123@198.19.249.3:30000/virtual-pet?replicaSet=virtual-pet-mongodb&ssl=false
mongodb://admin:password123@virtual-pet-mongodb-svc.mongodb.svc.cluster.local:27017/admin?replicaSet=virtual-pet-mongodb&ssl=false

virtual-pet-mongodb-svc.mongodb.svc.cluster.local

192.168.194.12 virtual-pet-mongodb-svc.mongodb.svc.cluster.local