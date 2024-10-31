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

## Exercies

1. [Introduction](exercises/1-intro.md)
2. [Mappings](exercises/2-mappings.md)
3. [Text analysis](exercises/3-text-analysis.md)
4. [Queries](exercises/4-queries.md)
5. [Aggregations](exercises/5-aggregations.md)
6. [Putting it all together](exercises/6-transactions.md)

## Demo

- Docker setup
- Dashboards, login
- Devtools (cat, indexes)
- Mappings - strict vs dynamic
- Little bit of each exercise

## Notes

Goals:
- Background on this session (BBC)

- 1. Build familiarity with OpenSearch Devtools and docs (you can do everything from there)
- 2. Understand some of the core concepts
- 3. Learn some specifics that might be useful for us
- 4. Practice some pairing and learn together - we need to develop some experience as a team - different to individual.
- 5. Search is fairly subjective. It's fun to play around an get a feel for things. This does that.

### Basic queries (learn the query DSL)

- Find some data
- Find matching films
- Find films with certain subtitles
- Find terms in any fields
- "Find films featuring both X and Y actor"
- Find all comedies staring Azis Ansari

### Mappings and ingestion ?

- Data types
- Keywords terms vs text
- Custom date formats?

### Text analysis

- Solve for partial matching
- First use ngram, then use search_as_you_type

### Advanced queries

- Boost/weighting in the query (e.g. prefer title over subtitle)
- Fuzzy search e.g. slop

### Aggregations

- Show how it works and get it going on some data
- e.g. aggregate across 

### Transaction search

- Take what you've learnt
- Build a good transaction search
- Things to think about
1. Fields to index and analyse
2. Relative weighting of different fields e.g. subtitle
3. Handling different account types - e.g. keyword
- Example searches/use cases to try
e.g. "Find all Tesco transactions above 5 pounds"

---

- sharding/replication?? If time
- Routing https://opensearch.org/docs/latest/search-plugins/searching-data/search-shard-routing/

TODO: Compare this list with the official training

## Data sources

* Netflix shows (https://www.kaggle.com/datasets/shivamb/netflix-shows)