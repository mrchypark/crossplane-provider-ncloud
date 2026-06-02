#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
: "${UPTEST_NCLOUD_REGION:?UPTEST_NCLOUD_REGION must be set, for example KR}"
: "${UPTEST_NCLOUD_SITE:=public}"

echo "Creating cloud credential secret..."
${KUBECTL} -n crossplane-system create secret generic provider-secret --from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default cluster provider config..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: ncloud.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  region: ${UPTEST_NCLOUD_REGION}
  site: ${UPTEST_NCLOUD_SITE}
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: crossplane-system
      key: credentials
EOF

${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m
