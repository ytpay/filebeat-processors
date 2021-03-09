#!/usr/bin/env bash

set -ex

# create source dir
mkdir -p ${HOME}/go/src/github.com/elastic
cd ${HOME}/go/src/github.com/elastic

# get source code
apt update \
  && apt upgrade -y \
  && apt install git -y \
  && git clone https://github.com/elastic/beats.git

# checkout
cd ${HOME}/go/src/github.com/elastic/beats
git checkout -b ${VERSION} ${VERSION}

# add custom processors
cat ${HOME}/filebeat-processors/import.txt >> libbeat/cmd/instance/imports_common.go

# build
cd filebeat && make crosscompile && mv build/bin/* ${HOME}/build

# test processors
${HOME}/build/filebeat-linux-amd64 test config -c ${HOME}/filebeat-processors/filebeat-test.yml
