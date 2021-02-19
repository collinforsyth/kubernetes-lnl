# part-3

Let's add another service!


Let's recreate our foo service
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

And let's create a bar service that calls the foo service (check out `main.go`)


Our bar service calls the foo service
```yaml
---
apiVersion: apps/v1
kind: Deployment
# This is how a unique Kubernetes resource can be made
metadata:
  name: bar
spec:
  replicas: 2
  selector:
    matchLabels:
      app: bar
  template:
    metadata:
      labels:
        app: bar
    spec:
      containers:
        - name: bar
          image: 966775768564.dkr.ecr.us-east-1.amazonaws.com/bar:latest
          ports:
            - containerPort: 8080
          securityContext:
            readOnlyRootFilesystem: true
          # let's add some environment variables to the container
          env:
            - name: FOO_ADDRESS
              value: "http://10.81.20.227:8080/ping"
```

```
kubectl apply -f foo/deployment.yaml
# we need a way to talk to our bar pod
kubectl get pods -o wide
# Grab that pod IP and create a deployment for bar
kubectl apply -f bar/deployment.yaml
kubectl get pods bar
kubectl logs <bar-pod> -f
# open a new shell and kill a foo container
kubectl delete pod <foo-pod>
```
