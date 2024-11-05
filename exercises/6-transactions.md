# 6. Transactions

Now that you're familiar with the main features of OpenSearch, it's time to put them together into a useful search application.

We're going to build a simple transaction search using some dummy data from the last 90 days.

It's up to you how to index the data, but you should be able to handle the following queries:

* Searching for all transactions from a single merchant (e.g. `Uber`)
* Searching for transactions based on a note or reference in the `subtitle` field
* Filtering for transactions above or below a given `amount`
* Filtering for transactions associated with a particular `account_id`
* Returning results in a "search as you type"/autocomplete style (i.e. `Tes` should return matching transactions at Tesco)
* Fuzzy matching
* Sorting by transaction date (i.e. most recent first)
* (Bonus) aggregating transaction counts and total/average spend amount on a per-category/per-account basis

You might want to think about how your search index performs when:

* Two transactions share the same title (which one should appear first?)
* One transaction matches the `title` field and one matches the `subtitle` field (which should be considered more relevant?)

The data you need for this exercise can be found in `data/transactions.json`.

You can use the `bin/upsert.sh` tool to ingest the data into an index:

```
./bin/upsert.sh <index-name> data/transactions.json
```

Remember to save your final DevTools input and any index mappings/settings. 

The data in `data/transactions.json` is generated using the script in [`bin/transactions/main.go`](../bin/transactions/main.go). Feel free to tweak it if you find yourself limited by the test data!