# part-4

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: bar
spec:
  selector:
    # This is the label to route all traffic to
    app: bar
  ports:
    - name: http
      protocol: TCP
      port: 8080
      targetPort: 8080
```

```
kubectl apply -f service.yaml
```

We need to create a label for our deployment object, so that a service can forward to all pods with that service.
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
        # Add this to our pod template
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
          env:
          # add after the fact
            - name: FOO_ADDRESS
              value: "http://foo.collin-test.svc.cluster.local:8080/ping"

```

```
kubectl apply -f deployment.yaml
```

Update to use the pod service DNS name.

Wait, what is this thing?

We can use DNS to automically assign dns names to, this is provided by an add-on we use on Kubernetes called `CoreDNS`.

```
k get svc foo -o wide
k get pods -l app=bar
k exec -it <bar-pod> -- ash
nslookup bar.test-collin.svc.clusterl.local
```


Ok, so we can use a DNS value to target this service Virtual IP
```
          # add after the fact
            - name: FOO_ADDRESS
              value: "http://foo.collin-test.svc.cluster.local:8080/ping"
```
Now deploy that
```
kubectl apply -f bar/deployment.yaml
```

Let's watch an `endpoint` object, which is a Kubernetes abstraction over a pod IP, grouped by labels.
```
kubectl get endpoints -w
```

In another terminal
```
kubectl delete pods -l app=foo
```
