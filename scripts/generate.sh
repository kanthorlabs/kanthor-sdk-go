#!/bin/sh
set -e

CHECKSUM_FILE=checksum
SCANNING_DIR=.

CHECKSUM_NEW=$(find $SCANNING_DIR -type f -name '*.go' -exec sha256sum {} \; | sort -k 2 | sha256sum | cut -d  ' ' -f1)
CHECKSUM_OLD=$(cat $CHECKSUM_FILE || true)

if [ "$CHECKSUM_NEW" != "$CHECKSUM_OLD" ];
then
	openapi-generator-cli generate -i openapi.json \
		-g go \
		-c openapi-generator-config.json \
		-o internal/openapi \
		--ignore-file-override .openapi-generator-ignore \
		--type-mappings int32=int64

	echo "generating checksum ...";
	find $SCANNING_DIR -type f -name '*.go' -exec sha256sum {} \; | sort -k 2 | sha256sum | cut -d  ' ' -f1 > $CHECKSUM_FILE;
fi