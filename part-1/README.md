# part-1


```yaml
---
apiVersion: v1
kind: Pod
metadata:
  name: foo-2
spec:
  containers:
    - name: foo
      image: 966775768564.dkr.ecr.us-east-1.amazonaws.com/foo:latest
      ports:
        - containerPort: 8080
      securityContext:
        readOnlyRootFilesystem: true
```

```yaml
---
apiVersion: v1
kind: Pod
metadata:
  name: foo-1
spec:
  containers:
    - name: foo
      image: 966775768564.dkr.ecr.us-east-1.amazonaws.com/foo:latest
      ports:
        - containerPort: 8080
      securityContext:
        readOnlyRootFilesystem: true
```


```
kubectl apply -f pod-1.yaml
kubectl apply -f pod-2.yaml
```


```
kubectl delete pods --all
```
