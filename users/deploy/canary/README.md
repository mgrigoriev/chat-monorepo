# Canary deployment

### Initial setup

```
# Current deployment (docker image :v1, 10 replicas)
kubectl apply -f deployment-v1.yml

# Canary deployment (docker image :v2, 0 replicas)
kubectl apply -f deployment-v2.yml

kubectl apply -f service.yml
```

### Route traffic to canary release

```
# 10%
kubectl scale --replicas=1 deployment/users-v2
kubectl scale --replicas=9 deployment/users-v1


# 50%
kubectl scale --replicas=5 deployment/users-v2
kubectl scale --replicas=5 deployment/users-v1


# 100%
kubectl scale --replicas=10 deployment/users-v2
kubectl scale --replicas=0 deployment/users-v1
```

### Test traffic distribution
```
while true; do curl http://127.0.0.1:83; echo; sleep 0.5; done
```
