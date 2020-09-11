eval $(minikube docker-env)

docker build -t local/dummy -f dockerfile/dummy.Dockerfile .

kubectl apply -f deployment/dummydeployment.yaml
kubectl apply -f deployment/dummy-apiauth-q-c.yaml
kubectl apply -f deployment/dummy-apiauth-q-d.yaml
kubectl apply -f deployment/dummy-apiauth-h-c.yaml
kubectl apply -f deployment/dummy-apiauth-h-d.yaml
kubectl apply -f deployment/dummy-base-auth.yaml

kubectl apply -f keda_config/hpa_noauth.yaml
kubectl apply -f keda_config/hpa_basicauth.yaml
kubectl apply -f keda_config/hpa_apiauth_query_d.yaml
kubectl apply -f keda_config/hpa_apiauth_query_c.yaml
kubectl apply -f keda_config/hpa_apiauth_header_c.yaml
kubectl apply -f keda_config/hpa_apiauth_header_d.yaml