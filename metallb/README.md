# MetalLB
## Install MetalLB
```
oc create -f 01-install.yaml
```
## Watch deployment
```
bash 02-watch-deployment.sh
```
## Create instance
```
oc create -f 03-instance.yaml
```
## Configure MetalLB
### L2 Advertisement
See [02-l2](./02-l2).

### BGP Advertisement
See [02-bgp](./02-bgp).

## Deploy application
```
oc create -f 05-application.yaml
```
