ROOT_DIR:=../..
include $(ROOT_DIR)/Makefile.env
include $(ROOT_DIR)/hack/make-rules/tools.mk

WITH_OPENSHIFT ?= false

.PHONY: deploy
deploy:
	@echo "Installing vault ..."
	$(TOOLBIN)/kubectl create namespace m4d-system || true
	$(TOOLBIN)/helm repo add hashicorp https://helm.releases.hashicorp.com
	$(TOOLBIN)/helm dependency update $(ROOT_DIR)/charts/vault
	$(TOOLBIN)/helm install vault $(ROOT_DIR)/charts/vault \
		--set "vault.global.openshift=$(WITH_OPENSHIFT)" \
                --set "vault.injector.enabled=false" \
                --set "vault.server.dev.enabled=true" \
		--values $(ROOT_DIR)/charts/vault/env/dev/plugin-secrets-kubernetes-reader-values.yaml \
		--namespace m4d-system \
		--wait --timeout 120s

.PHONY: wait-for-vault
wait-for-vault:
	$(TOOLBIN)/kubectl wait --for=condition=ready --all pod -n m4d-system --timeout=120s
