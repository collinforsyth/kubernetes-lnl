# part-6

Alright let's deploy this.

```
cd foo/
touch values.yaml
```

```yaml
---
app: foo
image: 966775768564.dkr.ecr.us-east-1.amazonaws.com/foo:latest
```

```
helm install my-very-special-foo-release -f values.yaml ../mychart/mychart-0.1.0.tgz
```

Now let's do our bar service
```
---
app: bar
image: 966775768564.dkr.ecr.us-east-1.amazonaws.com/bar:latest
env:
  FOO_ADDRESS: foo.test-collin.svc.cluster.local
```

```
helm install my-very-special-bar-release -f values.yaml ../mychart/mychart-0.1.0.tgz
```

Cleaning up:
```
helm delete my-very-special-bar-release
helm delete my-very-special-foo-release
kubectl delete namespace <your-namespace>
```
