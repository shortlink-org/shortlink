#!/bin/bash
k8s_delete_ns=$1
echo "Provided namesapce: ${k8s_delete_ns}..."
echo "Exporting namespace configuration..."
kubectl get namespaces -o json | grep "${k8s_delete_ns}"
kubectl get namespace ${k8s_delete_ns} -o json > temp.json
echo "Opening editor..."
wait 3
vi temp.json
echo "Sending configuration to k8s master for processing..."
curl -H "Content-Type: application/json" -X PUT --data-binary @temp.json http://127.0.0.1:8001/api/v1/namespaces/${k8s_delete_ns}/finalize
echo "Waiting for namespace deletion to process..."
wait 12
kubectl get namespaces
echo "...done."
