#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
  "all" \
  "github.com/Arvintian/demo-k8s-crd/pkg/generated" \
  "github.com/Arvintian/demo-k8s-crd/pkg/apis" \
  "example.com:v1stable" \
  --go-header-file "./hack/boilerplate.go.txt"