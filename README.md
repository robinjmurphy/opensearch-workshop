# opensearch-workshop

> A set of small challenges to help build familiarity with OpenSearch

## Setup

This repo uses Docker to spin up a local OpenSearch server. Make sure that Docker is running on your Mac before you begin!

You'll also need `docker-compose`:

```
brew install docker-compose
```

And `jq` for handling JSON data in the bulk indexing scripts:

```
brew install jq
```

Make sure there aren't already any other OpenSearch or ElasticSearch instances running in Docker using port `9200`:

```
docker ps --filter expose=9200
```

If there are, stop the containers:

```
docker stop $(docker ps --filter expose=9200 -q)
```

Now you can start the Docker containers for this project:

```
docker-compose up -d
```

Check you can access the OpenSearch REST API:

```
export OPENSEARCH_ADMIN_PASSWORD="yRPHP_Fsw6G2KAHCGsm"
```

```
curl -k https://localhost:9200/ -u admin:$OPENSEARCH_ADMIN_PASSWORD
{
  "name" : "opensearch-node1",
  "cluster_name" : "opensearch-cluster",
  "cluster_uuid" : "KlTNJQ4wQ_qMb167nxAqbg",
  "version" : {
    "distribution" : "opensearch",
    "number" : "2.17.1",
    "build_type" : "tar",
    "build_hash" : "1893d20797e30110e5877170e44d42275ce5951e",
    "build_date" : "2024-09-26T21:59:52.691008096Z",
    "build_snapshot" : false,
    "lucene_version" : "9.11.1",
    "minimum_wire_compatibility_version" : "7.10.0",
    "minimum_index_compatibility_version" : "7.0.0"
  },
  "tagline" : "The OpenSearch Project: https://opensearch.org/"
}
```

And check you can access the Dashboards web UI at http://localhost:5601/.

The username is `admin` and the password is `yRPHP_Fsw6G2KAHCGsm`, which is configured in the `docker-compose.yml` file.

## Exercises

1. [Introduction](exercises/1-intro.md)
2. [Mappings](exercises/2-mappings.md)
3. [Text analysis](exercises/3-text-analysis.md)
4. [Queries](exercises/4-queries.md)
5. [Aggregations](exercises/5-aggregations.md)
6. [Putting it all together](exercises/6-transactions.md)

Ideas for a future exercise: shards, replication and routing. 

## Data sources

* Netflix shows (https://www.kaggle.com/datasets/shivamb/netflix-shows)