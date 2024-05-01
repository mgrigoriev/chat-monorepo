# Blue-Green deployment

### Initial setup
```
# Blue
kubectl apply -f deployment-v1.yml

# Green
kubectl apply -f deployment-v2.yml

# Route traffic to v1 
kubectl apply -f service.yml
```

### Route traffic to v2 (after testing)

1. Edit `service.yml`: change `spec/selector/version` from `1.0` to `2.0`
2. `kubectl apply -f service.yml`
3. `kubectl delete deployment users-v1`

### Test traffic distribution
```
while true; do curl http://127.0.0.1:83; echo; sleep 0.5; done
```
