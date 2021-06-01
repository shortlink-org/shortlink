# Create K8S cluster

```
gcloud beta container \
  --project "fit-fin-309510" clusters create "shortlink" \
  --zone "us-central1-c" \
  --no-enable-basic-auth \
  --cluster-version "1.20.6-gke.1000" \
  --release-channel "rapid" \
  --machine-type "e2-medium" \
  --image-type "COS_CONTAINERD" \
  --disk-type "pd-standard" \
  --disk-size "100" \
  --metadata disable-legacy-endpoints=true \
  --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
  --num-nodes "4" \
  --enable-stackdriver-kubernetes \
  --enable-ip-alias \
  --network "projects/fit-fin-309510/global/networks/default" \
  --subnetwork "projects/fit-fin-309510/regions/us-central1/subnetworks/default" \
  --no-enable-intra-node-visibility \
  --default-max-pods-per-node "110" \
  --enable-dataplane-v2 \
  --no-enable-master-authorized-networks \
  --addons HorizontalPodAutoscaling,HttpLoadBalancing,GcePersistentDiskCsiDriver \
  --enable-autoupgrade \
  --enable-autorepair \
  --max-surge-upgrade 1 \
  --max-unavailable-upgrade 0 \
  --enable-shielded-nodes \
  --node-locations "us-central1-c"
```
