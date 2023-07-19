#!/bin/bash
# This script watches nmstate Operator deployment

# KUBECONFIG=${HOME}/openshift-network-playground/sno4/auth/kubeconfig
watch 'oc get clusterserviceversion -n metallb-system -o custom-columns=Name:.metadata.name,Phase:.status.phase'
