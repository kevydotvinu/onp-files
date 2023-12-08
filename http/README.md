```bash
$ oc patch configs.imageregistry.operator.openshift.io cluster --type merge --patch '{"spec":{"managementState":"Managed"}}'
$ oc patch configs.imageregistry.operator.openshift.io cluster --type merge --patch '{"spec":{"storage":{"emptyDir":{}}}}'
$ oc new-project http
$ oc new-app --name=http --strategy=docker .
$ oc start-build http --from-dir=./ --follow=true --wait=true
$ oc scale --replicas=3 deploy/http
$ oc expose deploy/http --port 9900
$ oc run --rm --tty --stdin --restart=Never onp --image=quay.io/onp/onp-tools -- /bin/sh -c 'while true; do curl http.http.svc:9900; sleep 1s; done'
```
