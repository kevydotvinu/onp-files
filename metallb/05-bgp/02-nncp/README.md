# NetworkPolicy
## Install NMState Operator
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
oc create -f 04-instance.sh
```
## Configure second interface
```
oc create -f 05-onp2-ethernet.yaml
```
