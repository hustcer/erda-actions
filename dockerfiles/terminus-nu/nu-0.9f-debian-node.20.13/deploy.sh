#!/usr/bin/env bash
set -eo pipefail

image=registry.erda.cloud/erda-actions/terminus-debian-nu:0.9f-n20.13

docker buildx build --platform linux/amd64 -t ${image} --push . -f Dockerfile

echo "action meta: terminus-debian-nu-0.9f-n20.13=$image"
