```bash
$ oc patch configs.imageregistry.operator.openshift.io cluster --type merge --patch '{"spec":{"managementState":"Managed"}}'
$ oc patch configs.imageregistry.operator.openshift.io cluster --type merge --patch '{"spec":{"storage":{"emptyDir":{}}}}'
$ oc new-project tcp
$ oc new-app --name=tcp --strategy=docker .
$ oc start-build tcp --from-dir=./ --follow=true --wait=true
$ oc scale --replicas=3 deploy/tcp
$ oc expose deploy/tcp --port 9900
$ oc run --rm --tty --stdin --restart=Never onp --image=quay.io/onp/onp-tools -- /bin/sh -c 'while true; do nc -w 1 tcp.tcp.svc 9900; sleep 1s; done'
```
