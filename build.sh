#!/usr/bin/env bash

VERSION=${1:-"v7.10.1"}

rm -rf build

docker run --rm -it -e VERSION=${VERSION} \
                    -v `pwd`/filebeat-test.yml:/filebeat-test.yml \
                    -v `pwd`/import.txt:/tmp/import.txt \
                    -v `pwd`/.compile.sh:/compile.sh \
                    -v `pwd`/build:/build \
                    golang:1.15 /compile.sh
