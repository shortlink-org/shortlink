# CERT TASKS ===========================================================================================================

# Prepare variables
export CERT_PATH=./ops/cert
export CERT_CONF_PATH=${CERT_PATH}/conf

cert-up: ## Generate all certificates
	# Generate the CA configuration file, certificate, and private key:
	@cfssl gencert -initca ${CERT_CONF_PATH}/ca-csr.json | cfssljson -bare ${CERT_PATH}/ca

	@cfssl gencert -initca ${CERT_CONF_PATH}/intermediate-ca.json | cfssljson -bare ${CERT_PATH}/intermediate_ca
	@cfssl sign -ca ${CERT_PATH}/ca.pem -ca-key ${CERT_PATH}/ca-key.pem -config ${CERT_CONF_PATH}/cfssl.json -profile intermediate_ca ${CERT_PATH}/intermediate_ca.csr | cfssljson -bare ${CERT_PATH}/intermediate_ca

	# Host Certificates
	@cfssl gencert -ca ${CERT_PATH}/intermediate_ca.pem -ca-key ${CERT_PATH}/intermediate_ca-key.pem -config ${CERT_CONF_PATH}/cfssl.json -profile=peer   ${CERT_CONF_PATH}/traefik.local.json | cfssljson -bare ${CERT_PATH}/shortlink-peer
	@cfssl gencert -ca ${CERT_PATH}/intermediate_ca.pem -ca-key ${CERT_PATH}/intermediate_ca-key.pem -config ${CERT_CONF_PATH}/cfssl.json -profile=server ${CERT_CONF_PATH}/traefik.local.json | cfssljson -bare ${CERT_PATH}/shortlink-server
	@cfssl gencert -ca ${CERT_PATH}/intermediate_ca.pem -ca-key ${CERT_PATH}/intermediate_ca-key.pem -config ${CERT_CONF_PATH}/cfssl.json -profile=client ${CERT_CONF_PATH}/traefik.local.json | cfssljson -bare ${CERT_PATH}/shortlink-client


cert-down: ## Delete generated cert-files
	-rm ${CERT_PATH}/*.csr
	-rm ${CERT_PATH}/*.pem
