Create the StatefulSet
```bash
oc create -f pvc.yaml -f service.yaml -f statefulset.yaml
```

Check the PVC Contents
```
oc rsh sample-ss-0 cat /path/in/pod/events.log
```bash

Scale the Number of Replicas
```bash
oc scale statefulset sample-ss --replicas=3
```

Check the number of pods:
```bash
oc get pods
```

Check the PVC Contents
```bash
oc rsh sample-ss-0 cat /path/in/pod/events.log
```