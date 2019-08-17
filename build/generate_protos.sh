#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail
shopt -s nullglob

root="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
xds="${root}/data-plane-api"

imports=(
  ${xds}
  "${root}/gogo/"
  "${root}/vendor/github.com/envoyproxy/protoc-gen-validate"
  "${root}/vendor/github.com/gogo/protobuf"
  "${root}/vendor/istio.io/gogo-genproto/googleapis"
  "${root}/vendor/istio.io/gogo-genproto/prometheus"
  "${root}/vendor/istio.io/gogo-genproto/opencensus/proto/trace/v1"
  "${root}/vendor/istio.io/gogo-genproto"
)

protoc="protoc"
protocarg=""
for i in "${imports[@]}"
do
  protocarg+="--proto_path=$i "
done

mappings=(
  "google/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations"
  "google/api/http.proto=google.golang.org/genproto/googleapis/api/annotations"
  "google/rpc/code.proto=google.golang.org/genproto/googleapis/rpc/code"
  "google/rpc/error_details.proto=google.golang.org/genproto/googleapis/rpc/errdetails"
  "google/rpc/status.proto=google.golang.org/genproto/googleapis/rpc/status"
  "gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto"
  "trace.proto=github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
  "opencensus/proto/trace/v1/trace.proto=github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
  "opencensus/proto/trace/v1/trace_config.proto=github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
  "metrics.proto=github.com/prometheus/client_model/go"
  "validate/validate.proto=github.com/envoyproxy/protoc-gen-validate/validate"
)

goprotoarg="plugins=grpc"

# assign importmap for canonical protos
for mapping in "${mappings[@]}"
do
  goprotoarg+=",M$mapping"
done

# assign importmap for all referenced protos in data-plane-api
for path in $(find ${xds}/envoy -type d)
do
  path_protos=(${path}/*.proto)
  if [[ ${#path_protos[@]} > 0 ]]
  then
    for path_proto in "${path_protos[@]}"
    do
      mapping=${path_proto##${xds}/}=github.com/envoyproxy/go-control-plane/${path##${xds}/}
      goprotoarg+=",M$mapping"
    done
  fi
done

for path in $(find ${xds}/envoy -type d)
do
  path_protos=(${path}/*.proto)
  if [[ ${#path_protos[@]} > 0 ]]
  then
    echo "Generating protos ${path} ..."
    $protoc ${protocarg} ${path}/*.proto \
      --go_out=${goprotoarg}:. \
      --validate_out="lang=go:."
  fi
done
