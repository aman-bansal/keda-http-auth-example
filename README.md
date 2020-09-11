
minikube start
minikube tunnel

bash build_app.sh

bash build_dummy.sh

make run ARGS="--zap-log-level=debug"

make deploy