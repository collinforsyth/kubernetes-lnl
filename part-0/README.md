# part-0

Create a pod and run it in Kubernetes

```yaml
---
apiVersion: v1
kind: Pod
metadata:
  # metadata.name is the unique identifier for all Kubernetes resources
  name: foo
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
kubectl apply -f pod.yaml
```

```
k get pods
```
