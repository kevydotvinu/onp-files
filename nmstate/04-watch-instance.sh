#!/bin/bash
# This script watches NMState Operator deployment

# KUBECONFIG=${HOME}/openshift-network-playground/sno4/auth/kubeconfig
watch -n 1 -t -x oc get pods -n openshift-nmstate
