# MetalLB
## Install MetalLB
```
oc create -f 01-install.yaml
```
## Watch installation
```
bash 02-watch-install.sh
```
## Create instance
```
oc create -f 03-instance.yaml
```
## Watch instance
```
bash 04-watch-instance.sh
```
## Configure MetalLB
### L2 Advertisement
See [05-l2](./05-l2).

### BGP Advertisement
See [05-bgp](./05-bgp).

## Deploy application
```
oc create -f 06-application.yaml
```
