### Deploy UDP server
```bash
$ oc patch configs.imageregistry.operator.openshift.io cluster --type merge --patch '{"spec":{"managementState":"Managed"}}'
$ oc patch configs.imageregistry.operator.openshift.io cluster --type merge --patch '{"spec":{"storage":{"emptyDir":{}}}}'
$ oc new-project udp
$ oc new-app --name=udp --strategy=docker .
$ oc start-build udp --from-dir=./ --follow=true --wait=true
$ oc scale --replicas=3 deploy/udp
$ oc expose deploy/udp --port 9900 --protocol UDP
$ oc logs -f deploy/udp
```
### Connect using nc tool
```
$ oc run --rm --tty --stdin --restart=Never onp --image=quay.io/onp/onp-tools -- /bin/sh -c 'while true; do echo request | nc -u -w 1 udp.udp.svc 9900; sleep 1s; done'
```
### Connect using Golang client
```
$ oc patch svc udp --type merge --patch '{"spec":{"type":"NodePort"}}'
$ go run client.go --server <node-hostname>:<nodeport> --message hello --count 100 --sleep 100
```
