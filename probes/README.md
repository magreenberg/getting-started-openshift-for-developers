# Build
```bash
REGISTRY="$(oc get routes -n openshift-image-registry -o json | jq -r '.items[].spec | select(.to.name=="image-registry") | .host')"
PROJECT=$(oc project -q)
podman login --tls-verify=false -u unused -p $(oc whoami -t) ${REGISTRY}
```

# Build the Image
```bash
podman build -t ${REGISTRY}/${PROJECT}/server:1 .
```

# Push the Image to the Registry
```bash
podman push ${REGISTRY}/${PROJECT}/server:1
```

# Create the Deployment
```bash
sed "s/__PROJECT__/${PROJECT}/" deployment.yaml | oc create -f -
```

# Check the Logs
```bash
oc logs - httpserver -c httpserver

```
# Access the Application
curl $(oc get route httpserver -o jsonpath='{.spec.host}')/

# Runtime
## Readiness Probe
Note that requests to the application will not be accepted until ???

# Toggle the Readiness Probe
```bash
curl $(oc get route httpserver -o jsonpath='{.spec.host}')/togglereadiness
```
Why did this command fail?
Now try this:
```bash
oc rsh httpserver curl http://localhost:8080/togglereadiness
```
