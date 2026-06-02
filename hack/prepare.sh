#!/usr/bin/env bash
set -euox pipefail

: ${PROVIDER_NAME_LOWER:=$(read -r -p "Lower case provider name (ex. github): " PROVIDER_NAME_LOWER; echo -n "${PROVIDER_NAME_LOWER}")}
: ${PROVIDER_NAME_NORMAL:=$(read -r -p "Normal case provider name (ex. GitHub): " PROVIDER_NAME_NORMAL; echo -n "${PROVIDER_NAME_NORMAL}")}
: ${ORGANIZATION_NAME:=$(read -r -p "Organization (e.g., my-org-name): " ORGANIZATION_NAME; echo -n "${ORGANIZATION_NAME}")}
: ${CRD_ROOT_GROUP:=$(read -r -p "CRD rootGroup (e.g., crossplane.io): " CRD_ROOT_GROUP; echo -n "${CRD_ROOT_GROUP}")}

REPLACE_FILES='./* ./.github :!build/** :!go.* :!hack/prepare.sh'
# shellcheck disable=SC2086
git grep -l 'ncloud' -- ${REPLACE_FILES} | xargs sed -i.bak "s/crossplane-provider-ncloud/provider-${PROVIDER_NAME_LOWER}/g"
# shellcheck disable=SC2086
git grep -l 'ncloud' -- ${REPLACE_FILES} | xargs sed -i.bak "s/ncloud/${PROVIDER_NAME_LOWER}/g"
# shellcheck disable=SC2086
git grep -l "crossplane/provider-${PROVIDER_NAME_LOWER}" -- ${REPLACE_FILES} | xargs sed -i.bak "s|crossplane/provider-${PROVIDER_NAME_LOWER}|${ORGANIZATION_NAME}/provider-${PROVIDER_NAME_LOWER}|g"
# shellcheck disable=SC2086
git grep -l 'Ncloud' -- ${REPLACE_FILES} | xargs sed -i.bak "s/Ncloud/${PROVIDER_NAME_NORMAL}/g"
# shellcheck disable=SC2086
git grep -l "crossplane.io" -- "apis/cluster/v1*" | xargs sed -i.bak "s|crossplane.io|${CRD_ROOT_GROUP}|g"
git grep -l "crossplane.io" -- "apis/namespaced/v1*" | xargs sed -i.bak "s|crossplane.io|${CRD_ROOT_GROUP}|g"
# shellcheck disable=SC2086
git grep -l "crossplane.io" -- "cluster/test/setup.sh" | xargs sed -i.bak "s|crossplane.io|${CRD_ROOT_GROUP}|g"
# shellcheck disable=SC2086
git grep -l "ujconfig\.WithRootGroup(\"${PROVIDER_NAME_LOWER}\.crossplane\.io\")" -- "config/provider.go" | xargs sed -i.bak "s|ujconfig.WithRootGroup(\"${PROVIDER_NAME_LOWER}\.crossplane\.io\")|ujconfig.WithRootGroup(\"${PROVIDER_NAME_LOWER}.${CRD_ROOT_GROUP}\")|g"

# We need to be careful while replacing "ncloud" keyword in go.mod as it could tamper
# some imported packages under require section.
sed -i.bak "s|crossplane/crossplane-provider-ncloud|${ORGANIZATION_NAME}/provider-${PROVIDER_NAME_LOWER}|g" go.mod
sed -i.bak -e "s|PROJECT_REPO ?= github.com/crossplane/|PROJECT_REPO ?= github.com/${ORGANIZATION_NAME}/|g" -e "s|\(blob/main/internal/\)${PROVIDER_NAME_LOWER}s|\1nclouds|g" Makefile
sed -i.bak "s/\[YEAR\]/$(date +%Y)/g" LICENSE

# Clean up the .bak files created by sed
git clean -fd

git mv "internal/clients/ncloud.go" "internal/clients/${PROVIDER_NAME_LOWER}.go"
git mv "cluster/images/crossplane-provider-ncloud" "cluster/images/provider-${PROVIDER_NAME_LOWER}"

# We need to remove this api folder otherwise first `make generate` fails with
# the following error probably due to some optimizations in go generate with v1.17:
# generate: open /Users/hasanturken/Workspace/crossplane-contrib/crossplane-provider-ncloud/apis/null/v1alpha1/zz_generated.deepcopy.go: no such file or directory
rm -rf apis/cluster/null
rm -rf apis/namespaced/null
# remove the sample directory which was a configuration in the ncloud
rm -rf config/cluster/null
rm -rf config/namespaced/null
# remove the sample MR example from the ncloud
rm -rf examples/cluster/null
rm -rf examples/namespaced/null
