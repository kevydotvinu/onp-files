#!/bin/bash
# This script watches MetalLB Pods

# KUBECONFIG=${HOME}/openshift-network-playground/sno4/auth/kubeconfig
watch -n 1 -t -x oc get pods -n metallb-system
