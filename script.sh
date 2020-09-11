make deploy


kubectl delete pod -n keda
kubectl scale deployment/keda-operator --replicas=0 -n keda

make run ARGS="--zap-log-level=debug"