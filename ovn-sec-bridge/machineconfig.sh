#!/bin/bash
INTERFACE=ens3
ROLE=worker
cat << EOF > sec-bridge-cni.yaml
apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfig
metadata:
  labels:
    machineconfiguration.openshift.io/role: ${ROLE}
  name: 12-${ROLE}-sec-bridge-cni
spec:
  config:
    ignition:
      version: 3.2.0
    storage:
      files:
        - path: /etc/ovnk/extra_bridge
          mode: 0644
          overwrite: true
          contents:
            source: data:text/plain;charset=utf-8;base64,$(echo ${INTERFACE} | base64 -w0)
          filesystem: root
EOF
echo "The machineconfig file has been generated in the present directory!"
