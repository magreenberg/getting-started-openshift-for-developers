# Create ConfigMap for Development
```bash
oc create configmap httpconfig --from-file=config.properties=develop.properties
```
# Start the Application
```bash
oc create -f deployment-with-cm.yaml -f route-with-cm.yaml -f service-with-cm.yaml
```

# Test the Mount
```bash
curl $(oc get route httpd-service-with-cm -o jsonpath='{"http://"}{.spec.host}{"\n"}')/config.properties
```

# Create the ConfigMap for Production
```bash
oc delete configmap httpconfig
oc create configmap httpconfig --from-file=config.properties=production.properties
oc scale deployment httpd-deployment-with-cm --replicas=0
oc scale deployment httpd-deployment-with-cm --replicas=1
```