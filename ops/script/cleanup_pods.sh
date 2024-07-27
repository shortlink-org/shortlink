#!/bin/bash

# Script to delete pods with specific statuses in all namespaces
# This script deletes pods with the following statuses: Evicted, Error, Completed, ContainerStatusUnknown
# It retrieves all namespaces and iterates over each namespace to find and delete the matching pods

# Get all namespaces
namespaces=$(kubectl get ns --no-headers -o custom-columns=:metadata.name)

# Loop through each namespace
for ns in $namespaces; do
  echo "Processing namespace: $ns"
  kubectl get pod -n $ns --no-headers | grep -E 'Evicted|Error|CrashLoopBackOff|ContainerStatusUnknown' | awk '{print $1}' | xargs -r kubectl delete pod -n $ns
done
