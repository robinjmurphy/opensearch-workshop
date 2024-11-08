# 3. Text analysis

When OpenSearch indexes a document it performs some extra processing (called "text analysis") on data stored in any full text fields.

This is designed to provide more relevant results and optimise performance at query time.

Text analysis is highly configurable and OpenSearch provides different kinds of ["analyzers"](https://opensearch.org/docs/latest/analyzers/supported-analyzers/index/) that can be customised to a suite a particular dataset.

The "standard" analyzer will parse a string and break it up into separate words (or "tokens" in OpenSearch speak), remove most punctuation and convert tokens (words) to lowercase. This same process happens at query time when you provide a search term or phrase, which allows OpenSearch to then compare the tokens in the query with the tokens in the index to find relevant documents. This process is described in more detail in the OpenSearch Docs ([Text Analysis](https://opensearch.org/docs/latest/analyzers/)).

You can see how OpenSearch will process any given block of text using the analysis API:

```
POST _analyze
{
  "analyzer": "standard",
  "text": "Breaking Bad"
}
```

Here you can see that using the standard analyzer the text "Breaking Bad" is split into two separate tokens (words) `breaking` and `bad`. Notice they've been converted to lowercase. We can also see by adding a `!` that trailing punctuation is stripped from the string:

```
POST _analyze
{
  "analyzer": "standard",
  "text": "Breaking Bad!"
}
```

We can also see that punctuation between two words is treated as a boundary between tokens:

```
POST _analyze
{
  "analyzer": "standard",
  "text": "Breaking-bad"
}
```

Some analyzers are more "aggressive" with their processing than others. The `stop` analyzer, for example, will strip out English stop words like "the" or "an" in an effort to improve results. As an example, use the analysis API to compare the output for `The King's Speech` using the `standard` analyzer vs the `stop` analyzer.

Some analysers (like the [language-specific analyzers](https://opensearch.org/docs/latest/analyzers/language-analyzers/)) will even go as far as to carry out "stemming" on a string where individual words are reduced to their "stem" form (.e.g `running` becomes `run`).

You can try this out with the analysis API using the `english` analyzer on a string like `Cool Runnings`.

As well as the built-in analyzers, it's also possible to build up your own analyzer by combining individual [tokenizers](https://opensearch.org/docs/latest/analyzers/tokenizers/index/) and [filters](https://opensearch.org/docs/latest/analyzers/token-filters/index/) to suite your dataset.

In the first exercise we learnt that by default OpenSearch doesn't support partial matches on `text` fields. We can now use text analysis with a custom analyzer to overcome this.

One way to solve this problem in OpenSearch is to use an "ngram" analyzer whereby a string is broken down not just into its constituent words, but also into smaller clusters of characters that make up those words.

```
POST _analyze
{
  "text": "Queer Eye",
  "tokenizer": {
    "type": "ngram",
    "min_gram": 2,
    "max_gram": 3,
    "token_chars": [
      "letter",
      "digit"
    ]
  }
}
```

In the example above we're breaking the string `Queer Eye` down into all of its possible 2-3 character substrings. If we were to use this custom analyzer in our index we would now return `Queer Eye` as a match on any query containing the letters `Qu`, `ue`, `eer` etc.

A variant of this that works well for "search as you type"/autocomplete use cases is an "edge ngram" tokenizer, which can be helpful in reducing false positives from pairs of common letters in the middle of the words. With this analyzer, only matching _prefixes_ are considered (i.e. `Que` but not `er`) :

```
POST _analyze
{
  "text": "Queer Eye",
  "tokenizer": {
    "type": "edge_ngram",
    "min_gram": 2,
    "max_gram": 3,
    "token_chars": [
      "letter",
      "digit"
    ]
  }
}
```

The analysis API can be helpful for debugging how a field is processed and iterating on analyzer settings, but once we're happy with the results we can update the index mapping to instruct OpenSearch to use these settings when indexing a field.

Let's create a new index where the `title` field makes use of the `edge_ngram` analyzer to unlock partial matching for "search as you type" behaviour:

```
PUT netflix-partial-matches
{
  "settings": {
    "analysis": {
      "analyzer": {
        "edge_ngram_analyzer": {
          "tokenizer": "edge_ngram_tokenizer"
        }
      },
      "tokenizer": {
        "edge_ngram_tokenizer": {
          "type": "edge_ngram",
          "min_gram": 2,
          "max_gram": 10,
          "token_chars": [
            "letter",
            "digit"
          ]
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "id": {
        "type": "keyword"
      },
      "title": {
        "type": "text",
        "fields": {
            "ngram": {
                "type": "text",
                "analyzer": "edge_ngram_analyzer"
            }
        }
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

Here we've defined a custom analyzer as part of our index and instructed OpenSearch to use that analyzer in place of the default `standard` analyzer in an additional `title.ngram` subfield that can be queried alongside the original `title` field. It's common practice to add additional indexed versions of a field like this rather than replacing them so that you can still access the `text` version of the field for queries where it's a better fit.

Let's re-index our data into the new index:

```
POST _reindex
{
  "source": {
    "index": "netflix"   
  },
  "dest": {
    "index": "netflix-partial-matches"
  }
}
```

Searching for a partial match on the `title` field should still return no matches:
```
GET netflix-partial-matches/_search
{
  "query": {
    "match": {
      "title": "Quee"
    }
  }
}
```


But if we search for a partial title in our new `ngram` subfield we should see some results:

```
GET netflix-partial-matches/_search
{
  "query": {
    "match": {
      "title.ngram": "Quee"
    }
  }
}
```

Unlike in our original index, this should now return the result `Queer Eye`.

There's a helpful feature in the analysis API that lets you test how a string would be analysed in a particular field in your index (rather than specifying the analyzer directly). You just need to provide the `field` name in the query:

```
POST netflix-partial-matches/_analyze
{
  "text": "Queer Eye",
  "field": "title.ngram"
}
```

Using ngrams can help improve results (especially for autocomplete applications) but it doesn't come without cost. See [this discussion of some tradeoffs](https://bigdataboutique.com/blog/dont-use-n-gram-in-elasticsearch-and-opensearch-6f0b48) as an example.

An alternative to a custom ngram analyzer is the [`search_as_you_type` field type](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-as-you-type.html). This is a more recent addition to the OpenSearch API and combines multiple ngram analyzers in a way that should result in better results and more efficient matches.

Try using the `search_as_you_type` field type in the Netflix data set and compare the results to the edge ngram approach we took above.

Once you have a preferred approach, update the `netflix-latest` alias to point to the index you'd like to use for future queries.
