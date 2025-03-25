# install prometheus operator
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

cd Infra/service-production/prometheus

helm install prometheus-server prometheus-community/prometheus -f values.yaml

Refrence : https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/README.md