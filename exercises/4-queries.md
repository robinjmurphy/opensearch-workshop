# 4. Queries

So far we've seen how some simple text queries work but OpenSearch offers a powerful (and sometimes confusing!) query API to find relevant documents.

For example, to query across multiple fields at once you can use a [`multi_match`](https://opensearch.org/docs/latest/query-dsl/full-text/multi-match) query:

```
GET netflix-latest/_search
{
  "query": {
    "multi_match": {
      "query": "Aziz Ansari",
      "fields": ["title", "cast"]
    }
  }
}
```

This query will return all results where there's a match for `Aziz Ansari` in either the `title` or `cast` fields.

Using the OpenSearch query API, see if you can find results for the following queries:

1. Find TV shows (not movies) with Aziz Ansari as a cast member
2. Find any movies about dinosaurs (anything with the word "dinosaur" in the title should rank higher)
3. Find some suggestions for true crime stories to watch
4. Find any TV shows or movies where Jason Bateman stars alongside Vince Vaughn
5. Find comedies about dating released in 2018
6. Find TV shows or movies starring or directed by George Clooney
7. Find some popular Netflix titles but allow for typos (e.g. `Breeking Bad`, `Arrested Dvelopment`)
8. Find movies with a cast member whose first name is Bianca

### Hints 
* You'll find that there's sometimes more than one way to construct an OpenSearch query that may impact the relevance of your results, so this exercise is as much about exploring the query API as it is finding a particular "right" answer!
* Use Google and [the OpenSearch query docs](https://opensearch.org/docs/latest/query-dsl/) to help you
* When working on each query, try to look for any unexpected hits and see if you can iterate on your query until it only returns the results you're expecting (this isn't always possible, but see how good - subjectively! - you can make each one).
* If you think the results could be improved by tweaking the mappings or text analysis settings on the index, try making some changes and see what impact it has!
* Keep each query in your DevTools window and export them using the `Export` button when you're done (so you can share them or access them later)
* Keep a copy of your final index mapping (`GET <index-name>`) (if you make any changes) so we can compare notes