#!/usr/bin/env bash

set -ex

# checkout
(cd beats && git checkout -b ${FILEBEAT_VERSION} ${FILEBEAT_VERSION})

# add custom processors
cat filebeat-processors/import.txt >> beats/libbeat/cmd/instance/imports_common.go

# build
(cd beats/filebeat && make crosscompile && mv build/bin/* ../../build)

# test processors
./build/filebeat-linux-amd64 test config -c filebeat-processors/filebeat-test.yml
