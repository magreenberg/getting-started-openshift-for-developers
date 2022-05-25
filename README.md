# Getting Started with OpenShift for Developers
https://github.com/magreenberg/getting-started-openshift-for-developers


# Starter Lab
See https://openshift-labs.github.io/starter-guides-html/common-explore.html

# Run a Pod in the Namespace/Project
```bash
oc run rhel8 --image=registry.access.redhat.com/ubi8/ubi --command -- bash -c 'sleep infinity'
```

# Additional Topics
* Liveness/Readiness probe
* debugging failed pods
  * cpu limit too high
  * memory too high
  * crash loop backoff
  * image not found
