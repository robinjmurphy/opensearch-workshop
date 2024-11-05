# 1. Introduction

Start by making sure you've completed the [setup steps](../README.md#setup) and can access the [OpenSearch Dahsboards UI locally](http://localhost:5601/).

Open the DevTools console at http://localhost:5601/app/dev_tools#/console (the username is `admin` and the password is `yRPHP_Fsw6G2KAHCGsm`).

Start by listing the indices that are already initialised in the cluster using the [cat indices API](https://opensearch.org/docs/latest/api-reference/cat/cat-indices/). Enter the following request on the left hand side of the DevTools UI and hit `cmd` + `enter` to execute it.

```
GET _cat/indices
```

You should see a list of built-in indices like `.opensearch-observability` and `.kibana_1`.

You can add column titles to the output using the `?v` param.

```
GET _cat/indices?v
```

This will return column headings like `pri.store.size` for the size of each index's primary shards (i.e. without replicas) on disk.

You can also get an explanation of the columns in the response using the `?help` param.

```
GET _cat/indices?help
```

At this point you may have noticed that the DevTools UI helpfully offers suggestions for API calls and parameters as you're typing. This means you don't have to remember the (fairly sizeable!) OpenSearch API.

Let's create a new empty index called `netflix` using the [create index API](https://opensearch.org/docs/latest/api-reference/index-apis/create-index/). We're going to use this to store some data on the Netflix video catalog:

```
PUT netflix
```

This API call might look a bit odd because all of the operations on a specific index use the index name without a prefix in the URL path.

Listing the indices again shoud show our empty index with 0 documents (`docs.count`) and minimal storage usage (`store.size`):

```
GET _cat/indices?v
```

We can also cat the index directly:

```
GET _cat/indices/netflix?v
```

We can also look at the settings for the index:

```
GET netflix
```

You should see that the index has inherited the default values for `number_of_shards` and `number_of_replicas` (which will both be set to `1`).

We can change the replication factor at any time using the [index settings API](https://opensearch.org/docs/latest/api-reference/index-apis/update-settings/), but some settings (like the shard count) cannot be updated once the index is created and can only be set as part of the initial create index API call.

If you do need to change the number of shards in an index after it's created you will typically need to create a new index with the correct settings and reindex your existing data into it ([example](https://opster.com/guides/opensearch/opensearch-operations/how-to-increase-primary-shard-count-in-opensearch/)).

Now that our index exists, we can ingest some data. Run the following command in your terminal from the root of this project:

```
./bin/upsert.sh netflix data/netflix.json
```

This uses the [upsert.sh](../bin/upsert.sh) script to parse and bulk ingest an array of JSON objects stored in the [netflix.json](../data/netflix.json) file using the OpenSearch [bulk API](https://opensearch.org/docs/latest/api-reference/document-apis/bulk/).

You can check that the ingest was successful back in the DevTools:

```
GET _cat/indices/netflix?v
```

You should now see `8807` documents in the `docs.count`.

We can take a look at this data by using the [search API](https://opensearch.org/docs/latest/api-reference/search/) and querying for all documents in the index with a [match all query](https://opensearch.org/docs/latest/query-dsl/match-all/). This isn't a query you'd expect to use in a real application, but it can be useful for debugging when you're not sure what data is an index:

```
GET netflix/_search
{
  "query": {
    "match_all": {}
  }
}
```

You should see a list of results or "`hits`" as the OpenSearch response calls them. Each hit includes some metadata about the index `_index`, an id `_id` and `_score`. These internal OpenSearch fields use underscores by convention.

The actual document itself is returned in the `_source` field.

10 results are returned by default, but this can be cusotmised using the `?size` field and [paginated](https://opensearch.org/docs/latest/search-plugins/searching-data/paginate/) using the `from` field.

```
GET netflix/_search?size=50
{
  "query": {
    "match_all": {}
  }
}
```

We can search for a specific TV show by its title using a [match query](https://opensearch.org/docs/latest/query-dsl/full-text/match/):


```
GET netflix/_search
{
  "query": {
    "match": {
      "title": "Queer"
    }
  }
}
```

You should see two results `Queer Eye` and `Queer Eye - We're in Japan!` are returned.

Interestingly, if we amend the search term and match on `Queer eye` you'll see that even more results are returned (which may be a surprise!):

```
GET netflix/_search
{
  "query": {
    "match": {
      "title": "Queer eye"
    }
  }
}
```

This is because by default OpenSearch will break your search query up into its separate words `queer` and `eye` using a process called [tokenization](https://opensearch.org/docs/latest/analyzers/) (we'll cover this in more detail in the [text analysis](./4-text-analysis.md) exercise).

This means we're now seeing results containing _either_ the word `queer` or the word `eye`.

However, `Queer Eye` is still at the top of the `hits` list because OpenSearch considers it to be the best (most relevant) match. This is reflected in the `_score` for that document, which is higher than the documents with a `title` that only contains the word `eye`. 

If we include the full title of the Japanese spin off `Queer eye we're in Japan` you'll see that it's now the top hit (with the highest score) because it matches all of the words in our query:

```
GET netflix/_search
{
  "query": {
    "match": {
      "title": "Queer eye we're in japan"
    }
  }
}
```

`Queer Eye` is second (matching two words) and the other documents are below.

If we want to only return results that match _all_ of the words in our query we can replace the match query with a [match phrase query](https://opensearch.org/docs/latest/query-dsl/full-text/match-phrase/):

```
GET netflix/_search
{
  "query": {
    "match_phrase": {
      "title": "Queer eye"
    }
  }
}
```

This will only match documents that contain an exact phrase in a specified order.

You can customise the behaviour of the match phrase query using concepts like [`slop`](https://opensearch.org/docs/latest/query-dsl/full-text/match-phrase/#slop)(great name!) to allow for some fuzziness in the presence and ordering of words.

For example, `queer eye japan` won't return `Queer Eye - We're in Japan!` by default because it's not an exact match:

```
GET netflix/_search
{
  "query": {
    "match_phrase": {
      "title": "Queer eye japan"
    }
  }
}
```

However, by using a `slop` value of 2 (which allows for 2 extra words to appear between words in my search phrase) it will be returned:

```
GET netflix/_search
{
  "query": {
    "match_phrase": {
      "title": {
        "query": "Queer eye japan",
        "slop": 2
      }
    }
  }
}
```

If you try searching for a partial word here e.g. `Quee` you may be surprised to see that no results are returned:

```
GET netflix/_search
{
  "query": {
    "match": {
      "title": "Quee"
    }
  }
}
```

That's because by default OpenSearch only considers _whole_ words as part of the tokenisation process that happens when a document is indexed. We'll look at how we can customise our index to support partial matches in the [text analysis](./3-text-analysis.md) exercise later on.

This should give you a feel for the way queries are structured in the OpenSearch API. We've only touched on two query types `match` and `match_phrase` here but one of the powerful (and sometimes confusing!) features of OpenSearch is its wide range of [different query types](https://opensearch.org/docs/latest/query-dsl/full-text/index/).

Now that you've had a chance to get to grips with the DevTools UI and ingested some data, it's time to learn more about how OpenSearch makes sense of the data you're indexing in the [mappings excerise](./3-mappings.md).