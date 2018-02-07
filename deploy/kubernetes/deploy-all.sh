kubectl create -f deploy-zookeeper.yaml
kubectl create -f deploy-redis-pod.yaml
kubectl create -f deploy-redis-service.yaml
kubectl create config mysql-config --from-file=schema.sql:schema.sql
kubectl create -f deploy-mysql-pod.yaml
kubectl create -f deploy-mysql-service.yaml
kubectl create -f deploy-reader.yaml
kubectl create -f deploy-writer.yaml
kubectl create -f deploy-ingress.yaml
kubectl create -f deploy-prometheus.yaml
