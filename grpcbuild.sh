#!/bin/sh
set -e -o pipefail

cd `dirname $0`
CURRENT=`pwd`

echo $CURRENT

#rm core/**/*.pb.go || true

protoc -I grpc --go_out=plugins=grpc:core grpc/*.proto && echo Done