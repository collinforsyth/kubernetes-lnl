# Kubernetes L&L

### Before beginning:
* Install Kubectx `brew install kubectx`
* Have `helm` and `kubectl` installed

Let's do some file clean up.
```
find . -name "*.yaml"
find ./ -name "*.yaml" -delete
```

Now let's set up our environment:
```
kubectx <dev-ctx>
```

Create a custom namespace for us to play around in
```
kubectl create ns {your-name}-test
```

```
kubens {your-name}-test
```

WARNING: If you do not do this, I will not be held liable for anything that breaks in dev :)
