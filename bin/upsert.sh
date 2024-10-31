#! /usr/bin/env bash

if [ $# -lt 2 ]; then
    echo "Usage: ./bin/upsert.sh <index-name> <path-to-json-file>"
    exit 1
fi

OPENSEARCH_ADMIN_PASSWORD="yRPHP_Fsw6G2KAHCGsm"

INDEX=$1

# FILEPATH should be the path to a JSON file containing an array of objects
# to be indexed with at least an `id` field (this is used as the OpenSearch
# document ID and so must be present).
FILEPATH=$2

cat $FILEPATH |
    jq -c '.[] | {"update": {"_index": "'$INDEX'", "_id": .id}}, { "doc": ., "doc_as_upsert": true }' |
    curl -i https://localhost:9200/_bulk -X POST -ku "admin:$OPENSEARCH_ADMIN_PASSWORD" -H 'Content-Type: application/json' --data-binary @- > /dev/null
