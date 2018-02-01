kubectl deploy -f deploy-zookeeper.yaml
kubectl deploy -f deploy-redis-pod.yaml
kubectl deploy -f deploy-redis-service.yaml
kubectl create config mysql-schema --from-file=schema.sql:schema.sql
kubectl deploy -f deploy-mysql-pod.yaml
kubectl deploy -f deploy-mysql-service.yaml
kubectl deploy -f deploy-reader.yaml
kubectl deploy -f deploy-writer.yaml
kubectl deploy -f deploy-ingress.yaml
