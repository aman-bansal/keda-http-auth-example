eval $(minikube docker-env)

docker build -t local/app -f dockerfile/app.Dockerfile .

kubectl apply -f deployment/appdeployment.yaml

#kubectl delete pod

#minikube start
#minikube tunnel