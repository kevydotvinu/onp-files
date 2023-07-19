# NetworkPolicy
## Install NMState Operator
```
oc create -f 01-install.yaml
```
## Watch deployment
```
bash 02-watch-deployment
```
## Create instance
```
oc create -f 03-instance.yaml
```
## Configure second interface
```
oc create -f 04-onp2-ethernet.yaml
```
