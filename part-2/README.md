# part-2

```yaml
---
apiVersion: apps/v1
kind: Deployment
# This is how a unique Kubernetes resource can be made
metadata:
  name: foo
spec:
  replicas: 2
  selector:
    matchLabels:
      app: foo
  template:
    metadata:
      labels:
        app: foo
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
kubectl apply -f deployment.yaml
```

```
# in another terminal
kubectl get pods -w
```

```
kubectl delete pod <pod>
```

```
kubectl delete deployment foo
```
