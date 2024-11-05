# 2. Mappings

In the first exercise we added some data into a new OpenSearch index, but we didn't tell OpenSearch anything about the structure of our data or how it should be indexed. Instead, we took advantage of a feature called [dynamic mappings](https://opensearch.org/docs/latest/field-types/#dynamic-mapping) to automatically detect the fields and types of the data we were indexing.

We can see what OpenSearch inferred about the structure of our data using the mappings API ("[mapping](https://opensearch.org/docs/latest/field-types)" is the OpenSearch term for the schema or structure of a dataset).

```
GET netflix/_mapping
```

You should see that most of the fields have been inferred as `text` fields ("`"type": "text"`). `text` is a common data type that allows for full text search ([full list of supported types](https://opensearch.org/docs/latest/field-types/supported-field-types/index/)).

By default, OpenSearch indexes `text` fields in a way that allows them to be queried without requiring an exact match (for example in the `title` field `Queer Eye` OpenSearch will index the individual _terms_ `queer` and `eye` in a case-insensitive way). We'll look at this process in more detail in the next exercise, but for now it's just important to know that `text` is a flexible string type that triggers some [extra processing](https://opensearch.org/docs/latest/analyzers/) of the indexed data.

[`keyword`](https://opensearch.org/docs/latest/field-types/supported-field-types/keyword/) on the other hand is much simpler string type. Unlike `text`, it isn't processed in any way and so only supports exact matches:

> A keyword field type contains a string that is not analyzed. It allows only exact, case-sensitive matches.

The `keyword` type is often used for IDs, email addresses and phone numbers where partial matches could be confusing.

Whilst you won't see any directly inferred `keyword` types at the top level of our mapping, you will see that the `keyword` type appears as a _separate field_ in every one of the fields in our mapping.

For example, the title has a type of `text` but it also has a nested `keyword` field of type `keyword`.

```json
{
    "title": {
    "type": "text",
    "fields": {
        "keyword": {
            "type": "keyword",
            "ignore_above": 256
        }
    }
},
```

This is because for any `text` fields that OpenSearch infers using dynamic mapping it also indexes a copy of the original (unprocessed) data in a `keyword` field that allows for strict matching.

The `fields` property can be used more broadly to instruct OpenSearch to index the same data in a source document in multiple different ways. This can be useful when you want to do more processing at indexing time that allows for more flexible querying later.

To make use of a nested field in a search query you use the dot `.` notation:

```
GET netflix/_search
{
  "query": {
    "match": {
      "title.keyword": "Queer Eye"
    }
  }
}
```

Notice how only an exact match works on the `title.keyword` field here. Even `Queer eye` won't return a result because the query is case-sensitive.

Whilst dynamic mapping can be helpful to get something up and running quickly, the general guidance is to instead rely on explicit mappings when indexing data. This ensures that search behaviour remains consistent and provides more predictable performance.

> While dynamic mappings automatically add new data and fields, using explicit mappings is recommended. Explicit mappings let you define the exact structure and data types upfront. This helps to maintain data consistency and optimize performance, especially for large datasets or high-volume indexing operations.

We're going to update our index to use explicit mappings. There are some cases where an existing mapping can be updating (e.g. when two data types are compatible) but it isn't always possible, so making a mapping change usually involves creating a _new index_ and re-indexing your existing documents.

Create a new index with the following explicit mappings:

```
PUT netflix-explicit-mappings
{
  "mappings": {
    "properties": {
      "id": {
        "type": "keyword"
      },
      "title": {
        "type": "text"
      },
      "type": {
        "type": "keyword"
      },
      "release_year": {
        "type": "date",
        "format": "yyyy"
      },
      "rating": {
        "type": "keyword"
      },
      "duration": {
        "type": "text"
      },
      "director": {
        "type": "text"
      },
      "cast": {
        "type": "text"
      },
      "description": {
        "type": "text"
      },
      "date_added": {
        "type": "text"
      },
      "country": {
        "type": "text"
      },
      "categories": {
        "type": "text"
      }
    }
  }
}
```

You can see that we've used the `keyword` type for some of the properties where we don't need full text search (e.g. on the `id` and `type` fields).

We can now [re-index](https://opensearch.org/docs/latest/api-reference/document-apis/reindex/) the existing data in our `netflix` index to populate the new index:

```
POST _reindex
{
  "source": {
    "index": "netflix"   
  },
  "dest": {
    "index": "netflix-explicit-mappings"
  }
}
```

Looking at the cat endpoint you should see that our new index now also has `8807` documents:

```
GET _cat/indices/netflix*
```

The `*` in the URL here will query all matching indices.

You should also notice that our new (explicit mappings) index is using less space on disk. That's because we have fewer full text fields and we've removed the additional `keyword` fields on each property.

This process of indexing and re-indexing data is fairly common, and so OpenSearch has a feature called _index aliases_ that allows you to point your application at a fixed index name (the alias) which can be updated to point to a different underlying index over time.

Let's create an alias for our index and point it to the index with explicit mapping:

```
PUT netflix-explicit-mappings/_aliases/netflix-latest
```

Now we can query the `netflix-latest` index, even if we later update the alias to point to a new underlying index.

```
GET netflix-latest/_search
{
  "query": {
    "match": {
      "title": "Squid"
    }
  }
}
```