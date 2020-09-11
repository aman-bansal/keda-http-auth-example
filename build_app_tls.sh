eval $(minikube docker-env)

docker build -t local/tlsapp -f dockerfile/tlsapp.Dockerfile .

kubectl apply -f deployment/tlsappdeployment.yaml

docker build -t local/dummy -f dockerfile/dummy.Dockerfile

kubectl apply -f deployment/dummy-tls-auth.yaml

kubectl apply -f keda_config/hpa_tlsauth.yaml