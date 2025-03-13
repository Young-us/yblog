@echo off
kubectl apply -f configmap.yaml
kubectl apply -f secret.yaml
kubectl apply -f mysql/
kubectl apply -f yblog/
kubectl apply -f cloudflare/