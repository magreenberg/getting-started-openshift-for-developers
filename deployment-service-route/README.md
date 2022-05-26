# Start Deployment
```bash
oc create -f deployment.yaml
```
View using 'oc' and via web console.

# Create a Service
```bash
oc create -f service.yaml
```
Discuss
* What is a service
* Types of service


# Scaling and Self Healing
* Kill one pod
* Scale the deployment: `oc scale deployment/httpd-deployment --replicas=4` and using web console

# Working with Routes
```bash
oc expose service httpd-service
```

Discuss:
* Termination
* Cluster Resource

#  