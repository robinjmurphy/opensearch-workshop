# 6. Transactions

Now that we've seen some of the main features of OpenSearch, it's time to put them all together into a useful search application.

You're going to build a transaction search application that lets you query for things you've paid for.

It's up to you how to index the data, but you should be able to handle the following queries:

*

You might want to think about how your index performs in these cases:

* e.g. title/vs subtitle
* sorting when two transactions match
* stopwords and synonyms
* partial matching
* filtering on accounts