#!/bin/sh
set -eu

cred=${NCLOUD_CREDENTIALS_FILE:-.work/ncloud-credentials.json}
site=${NCLOUD_SITE:-gov}
region=${NCLOUD_REGION:-KR}
outdir=${NCLOUD_PREFLIGHT_OUTDIR:-.work/acceptance}
ts_label=$(date -u +%Y%m%dT%H%M%SZ)

if [ "$site" != gov ]; then
  echo "unsupported site for this preflight: $site" >&2
  exit 2
fi

if [ ! -f "$cred" ]; then
  echo "missing credentials file: $cred" >&2
  exit 2
fi

access_key=$(jq -r '.access_key // empty' "$cred")
secret_key=$(jq -r '.secret_key // empty' "$cred")
if [ -z "$access_key" ] || [ -z "$secret_key" ]; then
  echo "missing access_key or secret_key in credentials JSON" >&2
  exit 2
fi

mkdir -p "$outdir"

hex() {
  od -An -tx1 | tr -d ' \n'
}

hmac_hex() {
  key_hex=$1
  msg=$2
  printf '%s' "$msg" | openssl dgst -sha256 -mac HMAC -macopt "hexkey:$key_hex" -binary | hex
}

hmac_b64() {
  key=$1
  msg=$2
  printf '%s' "$msg" | openssl dgst -sha256 -hmac "$key" -binary | openssl base64 | tr -d '\n'
}

case "$region" in
  KR)
    s3_signing_region=gov-standard
    s3_endpoint=https://kr.object.gov-ncloudstorage.com
    cla_region=kr
    ;;
  KRS)
    s3_signing_region=gov2-standard
    s3_endpoint=https://krs.object.gov-ncloudstorage.com
    cla_region=krs
    ;;
  *)
    echo "unsupported gov region for preflight: $region" >&2
    exit 2
    ;;
esac

run_billing() {
  if [ -x .work/acceptance/billing-cost.sh ]; then
    .work/acceptance/billing-cost.sh "preflight-$ts_label" || true
  fi
}

run_live_count() {
  if command -v kubectl >/dev/null 2>&1 && kubectl cluster-info >/dev/null 2>&1; then
    kubectl get managed -n ncloud-accept --no-headers 2>/dev/null | wc -l | tr -d ' '
  else
    printf 'unknown'
  fi
}

run_s3_list_buckets() {
  amz_date=$(date -u +%Y%m%dT%H%M%SZ)
  date_stamp=$(date -u +%Y%m%d)
  host=${s3_endpoint#https://}
  payload_hash=UNSIGNED-PAYLOAD
  canonical_headers="host:${host}
x-amz-content-sha256:${payload_hash}
x-amz-date:${amz_date}
"
  signed_headers='host;x-amz-content-sha256;x-amz-date'
  canonical_request="GET
/

${canonical_headers}
${signed_headers}
${payload_hash}"
  algorithm=AWS4-HMAC-SHA256
  credential_scope="${date_stamp}/${s3_signing_region}/s3/aws4_request"
  canonical_hash=$(printf '%s' "$canonical_request" | openssl dgst -sha256 -binary | hex)
  string_to_sign="${algorithm}
${amz_date}
${credential_scope}
${canonical_hash}"
  secret_hex=$(printf 'AWS4%s' "$secret_key" | hex)
  date_key=$(hmac_hex "$secret_hex" "$date_stamp")
  region_key=$(hmac_hex "$date_key" "$s3_signing_region")
  service_key=$(hmac_hex "$region_key" s3)
  signing_key=$(hmac_hex "$service_key" aws4_request)
  signature=$(hmac_hex "$signing_key" "$string_to_sign")
  authorization="${algorithm} Credential=${access_key}/${credential_scope}, SignedHeaders=${signed_headers}, Signature=${signature}"
  body_file="$outdir/preflight-s3-listbuckets-$ts_label.xml"
  http_code=$(curl -sS -w '%{http_code}' -o "$body_file" "$s3_endpoint/" \
    -H "Host: $host" \
    -H "x-amz-date: $amz_date" \
    -H "x-amz-content-sha256: $payload_hash" \
    -H "Authorization: $authorization")
  error_code=$(sed -n 's:.*<Code>\([^<]*\)</Code>.*:\1:p' "$body_file" | head -1)
  [ -n "$error_code" ] || error_code=none
  printf 'objectstorage_s3_listbuckets_http=%s error_code=%s file=%s\n' "$http_code" "$error_code" "$body_file"
  [ "$http_code" = 200 ]
}

run_cla_capacity() {
  uri="/api/${cla_region}-v1/capacity"
  timestamp=$(($(date +%s) * 1000))
  msg="GET ${uri}
${timestamp}
${access_key}"
  sig=$(hmac_b64 "$secret_key" "$msg")
  body_file="$outdir/preflight-cla-capacity-$ts_label.json"
  header_file="$outdir/preflight-cla-capacity-$ts_label.headers"
  http_code=$(curl -sS -D "$header_file" -w '%{http_code}' -o "$body_file" \
    "https://cloudloganalytics.apigw.gov-ntruss.com${uri}" \
    -H "x-ncp-apigw-timestamp: ${timestamp}" \
    -H "x-ncp-iam-access-key: ${access_key}" \
    -H "x-ncp-apigw-signature-v2: ${sig}" \
    -H 'Content-Type: application/json')
  bytes=$(wc -c < "$body_file" | tr -d ' ')
  code=empty
  message=empty
  if [ "$bytes" != 0 ]; then
    code=$(jq -r '.code // empty' "$body_file" 2>/dev/null || true)
    message=$(jq -r '.message // empty' "$body_file" 2>/dev/null || true)
    [ -n "$code" ] || code=unknown
    [ -n "$message" ] || message=unknown
  fi
  printf 'cla_capacity_http=%s body_bytes=%s code=%s message=%s file=%s\n' "$http_code" "$bytes" "$code" "$message" "$body_file"
  [ "$http_code" = 200 ] && [ "$code" = 0 ]
}

printf 'ncloud_gov_acceptance_preflight timestamp=%s site=%s region=%s\n' "$ts_label" "$site" "$region"
run_billing
printf 'live_managed_resources=%s\n' "$(run_live_count)"

s3_ok=false
if run_s3_list_buckets; then
  s3_ok=true
fi

cla_ok=false
if run_cla_capacity; then
  cla_ok=true
fi

printf 'preflight_summary objectstorage_ready=%s cloud_log_analytics_ready=%s\n' "$s3_ok" "$cla_ok"

if [ "$s3_ok" = true ] && [ "$cla_ok" = true ]; then
  exit 0
fi
exit 1
