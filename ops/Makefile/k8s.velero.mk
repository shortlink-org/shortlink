# VELERO TASKS =========================================================================================================
velero-up:
	@kubectl apply -f ops/Helm/velero/minio/00-minio-deployment.yaml
	@velero install \
		--provider aws \
		--plugins velero/velero-plugin-for-aws:v1.0.0 \
		--bucket velero \
		--secret-file ./ops/Helm/velero/credentials-velero \
		--use-volume-snapshots=false \
		--backup-location-config region=minio,s3ForcePathStyle="true",s3Url=http://minio.velero.svc:9000 \
		--wait
	@kubectl get deployments -l component=velero --namespace=velero

velero-backup:
	@velero backup create shortlink-backup --include-namespaces shortlink

velero-restore:
	@velero restore create --from-backup shortlink-backup
