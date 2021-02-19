# part-5

Creating a reusable helm chart for this.

```
rm -rf mychart
helm create mychart
```

And now let's delete pretty much everything because its an unreadable mess
```
cd templates
# check out some templates and lets collectively hurt our brains
rm -rf templates
```
```
# charts is used to store "subcharts" but we don't need to worry about that
rm -rf charts
```
```
# Let's also trim the values.
```

```
cd templates
# grab our current stuff and put in our directory
cp ../../foo/service.yaml ./
cp ../../foo/deployment.yaml ./
```

```yaml
---
apiVersion: apps/v1
kind: Deployment
# This is how a unique Kubernetes resource can be made
metadata:
  name: {{ required "requires Values.app to be set" .Values.app }}
  labels:
    app: {{ required "requires Values.app to be set" .Values.app }}
spec:
  replicas: {{ required "requires Values.replicas to be set" .Values.replicas }}
  selector:
    matchLabels:
      app: {{ required "requires Values.app to be set" .Values.app }}
  template:
    metadata:
      labels:
        app: {{ required "requires Values.app to be set" .Values.app }}
    spec:
      containers:
        - name: foo
          image: {{ required "requires Values.image to be set" .Values.image }}
          ports:
            - containerPort: 8080
          securityContext:
            readOnlyRootFilesystem: true
          env:
          # Here's something that's a bit more nuanced
          {{- range $k, $v := .Values.env }}
            - name: {{ $k }}
              value: {{ $v }}
          {{ end }}
```

```
---
apiVersion: v1
kind: Service
metadata:
  name: {{ required "requires Values.app to be set" .Values.app }}
spec:
  selector:
    # This is the label to route all traffic to
    app: {{ required "requires Values.app to be set" .Values.app }}
  ports:
    - name: http
      protocol: TCP
      port: 8080
      targetPort: 8080
```

Let's update `values.yaml`

Ok now let's package this up into something.
```
helm package -d ../ ./ 
```

What all is in here?
```
mkdir -p whatsinthebox && tar -xvf mychart-0.1.0.tgz -C whatsinthebox
cd whatsinthebox/mychart
```
It's all just some tarred up yaml files *yawn*
