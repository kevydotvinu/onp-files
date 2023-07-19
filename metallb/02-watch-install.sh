#!/bin/bash
# This script watches MetalLB Operator deployment

# KUBECONFIG=${HOME}/openshift-network-playground/sno4/auth/kubeconfig
watch -n 1 'oc get clusterserviceversion -n metallb-system -o custom-columns=Name:.metadata.name,Phase:.status.phase'
