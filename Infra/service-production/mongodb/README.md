# install mongodb operator
helm repo add mongodb https://mongodb.github.io/helm-charts

cd Infra/service-production/mongodb

helm install community-operator mongodb/community-operator --namespace mongodb --create-namespace -f values.yaml

kubectl apply -f mongodb-db-virtual-pet.yaml -n mongodb

Refrence : https://github.com/mongodb/mongodb-kubernetes-operator/blob/master/README.md