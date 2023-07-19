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
See [04-l2](./04-l2).

### BGP Advertisement
See [04-bgp](./04-bgp).

## Deploy application
```
oc create -f 05-application.yaml
```
