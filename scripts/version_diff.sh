#!/bin/sh
set -eu

if [ "$#" -ne 3 ]; then
  echo "usage: version_diff.sh <generated resource list> <base JSON schema path> <bumped JSON schema path>" >&2
  exit 2
fi

resources_path=$1
base_path=$2
bumped_path=$3

echo "Reporting schema changes between \"${base_path}\" as base version and \"${bumped_path}\" as bumped version"

provider_name=$(jq -r '.provider_schemas | keys_unsorted[0] // empty' "${base_path}")
if [ -z "${provider_name}" ]; then
  echo "Cannot extract the provider name from the base schema: ${base_path}" >&2
  exit 1
fi

jq -n -r \
  --arg provider "${provider_name}" \
  --slurpfile resources "${resources_path}" \
  --slurpfile base "${base_path}" \
  --slurpfile bumped "${bumped_path}" '
    $resources[0][] as $name |
    ($base[0].provider_schemas[$provider].resource_schemas[$name].version // null) as $basever |
    ($bumped[0].provider_schemas[$provider].resource_schemas[$name].version // null) as $bumpver |
    if $basever == null or $bumpver == null then
      "\($name) is not found in schema"
    elif $basever != $bumpver then
      "\($name):\($basever)-\($bumpver)"
    else
      empty
    end
  '
