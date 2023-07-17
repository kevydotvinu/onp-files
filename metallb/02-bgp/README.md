# BGP
## Architecture
![Architecture](metallb/02-bgp/.architecture.png?raw=true)
## Configure VyOS
See [vyos](metallb/02-bgp/vyos)

## Configure NNCP
```
oc create -f nncp.yaml
```

## Configure BGP
```
oc create -f 03-configure.yaml
```
