@echo off
kubectl delete deployment yblog-deployment
kubectl delete service yblog-service 
kubectl delete statefulset mysql
kubectl delete service mysql-service
kubectl delete deployment cloudflared
kubectl delete configmap yblog-config
kubectl delete secret db-secret
