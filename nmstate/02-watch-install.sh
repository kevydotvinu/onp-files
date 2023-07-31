#!/bin/bash
# This script watches nmstate Operator deployment

# KUBECONFIG=${HOME}/openshift-network-playground/sno4/auth/kubeconfig
watch -n 1 -t -x oc get clusterserviceversion -n openshift-nmstate -o custom-columns=Name:.metadata.name,Phase:.status.phase
