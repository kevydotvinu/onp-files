#!/bin/bash
set -x
NODEIP=$(oc get pod foo -o json | jq -r '.status.hostIP')
sudo ip route del 192.168.126.0/24 via ${NODEIP}
