# 5. Aggregations

If you're short on time feel free to skip this exercise and go straight into the transaction search challenge. [Aggregations](https://opensearch.org/docs/latest/aggregations/) are an interesting feature of OpenSearch that allow you to extract aggregate statistics from your data, but they aren't technically required when implementing core search functionality.

A common feature of faceted search applications is showing how many results share a particular field value. For example, when searching for a washing machine on an e-commerce site you might expect to see a breakdown by brand or size in the sidebar that allows for further filtering (e.g. AEG 5, Bosch 2).

OpenSearch supports this kind of functionality with [aggregations](https://opensearch.org/docs/latest/aggregations/).

We can start by returning the number of matching titles in our Netflix catalog, broken down by rating:

```
GET netflix-latest/_search
{
  "size": 0,
  "aggs": {
    "response_codes": {
      "terms": {
        "field": "rating",
        "size": 10
      }
    }
  }
}
```

This query returns the top 10 unique `terms` in the `ratings` field by document count. Crucially, aggregations operate on the entire dataset, and not just the results on the current page. This is why we can specify a page `size` of `0` and still get back accurate aggregations.

Aggregations (aggs) can be combined with a search query to return only the statistics for documents that match a search term:

```
GET netflix-latest/_search
{
  "size": 0,
  "query": {
    "match": {
      "title": "Family"
    }
  }, 
  "aggs": {
    "response_codes": {
      "terms": {
        "field": "release_year",
        "size": 10
      }
    }
  }
}
```

This query returns the breakdown of TV shows and movies matching the search term `Family` by the year of release.

Aggregations don't work on `text` fields, only `keyword`, `date` and other simple data types that can be easily aggregated. If you need to aggregate on a text field, you should also keep a separate `keyword` version of the field in your mappings. However, be aware that this only works well for fields with a bounded cardinality.