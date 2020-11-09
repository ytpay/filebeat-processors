#!/usr/bin/env bash

set -ex

# create source dir
mkdir -p /go/src/github.com/elastic
cd /go/src/github.com/elastic

# get source code
apt update \
  && apt upgrade -y \
  && apt install git -y \
  && git clone https://github.com/elastic/beats.git

# checkout
cd /go/src/github.com/elastic/beats
git checkout -b ${VERSION} ${VERSION}

# add custom processors
cat /tmp/import.txt >> libbeat/cmd/instance/imports_common.go

# build
cd filebeat && make crosscompile && mv build/bin/* /build

# test processors
/build/filebeat-linux-amd64 test config -c /filebeat-test.yml
