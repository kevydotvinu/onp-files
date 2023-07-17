# MetalLB
## Install MetalLB
```
oc create -f install.yaml
```

## Configure MetalLB
### L2 Advertisement
See [02-l2](metallb/02-l2)

### BGP Advertisement
See [02-bgp](metallb/02-bgp)

## Deploy application
```
oc create -f application.yaml
```
