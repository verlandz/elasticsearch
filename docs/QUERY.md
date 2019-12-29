# Query
brief of queries information.


### 1. Request by URI vs BODY
in elastic, there's 2 ways to query:
- by [URI](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-uri-request.html)
- by [BODY](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html) (with JSON Body)

but I prefer you stick with BODY ways, because it support more queries than URI.


### 2. Mapping & Data Types
mapping consist of field => value. If you want to know index's map list, try [here](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/indices-get-mapping.html). You must remember that existing mapping cannot be deleted or updated. So if somehow you want to change field `ID` from `int` to `string`, you got 2 choices; 
- just make 1 new mapping with new type (and ofc new field name too ex. `IDstr`)
  - I don't really recommended because the mapping become dirty (but up to you).
  - strategy:
    - for old data, make a script to iterate all data who `IDstr` still empty and fill it
    - for new data, code fill `ID` and `IDstr`
    - after done, code read from `ID` to `IDstr`
- make new index where `ID` already set as `string` (new index = new fresh data = data still empty)
  - this one is recommended but you need to duplicate the data. How? just re-import from your main DB (not elastic) to your elastic.
  - strategy:
    - for old data, make script for import from db to new elastic
    - for new data, code not only fill in old elastic but also in new elastic
    - after done, delete old elastic

for list of mapping types can read [here](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/mapping-types.html)


### 3. Search in 1 or more index
index is like table in DB, if you want to select some specific index in query can try [here](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-search.html#search-search-api-example)


### 4. How to Explain
want to know whether query bad/not ? try [explain](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-explain)


### 5. How to Offset & Sort
this combo usually use for pagination, you can try here for [offset](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-from-size) and for [sort](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-sort)


### 6. Query DSL
query DSL(Domain Specific Language) is high-level library to manipulate queries. It [consist of](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/query-dsl.html):
- *Leaf query clauses*
  - [match](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/query-dsl-match-query.html)
  - [term](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/query-dsl-term-query.html)
  - [range](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/query-dsl-range-query.html)
- *Compound query clauses*
  - [Boolean query](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/query-dsl-bool-query.html)
    - `must`, like AND
    - `filter`, also like AND but score is ignored
    - `should`, like OR
    - `must_not`, also like AND with not
  - [Disjunction max query](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/query-dsl-dis-max-query.html)
    - like subqueries
  - and [others](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/compound-queries.html)

***TRIVIA***
- [match vs term](https://discuss.elastic.co/t/term-query-vs-match-query/14455)
- [must vs filter](https://stackoverflow.com/questions/43349044/what-is-the-difference-between-must-and-filter-in-query-dsl-in-elasticsearch)


### 7. Wildcard
like `LIKE` in DB, read [here](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/query-dsl-wildcard-query.html). But be carefull for [performance problem](https://discuss.elastic.co/t/wildcard-query-performance/4348/8)


### 8. Autocomplete
for autocomplete / suggestion related from your keyword, read [here](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-suggesters.html#completion-suggester)
  

### 9. Aggregations
[read here](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-aggregations.html)
`Bucketing` is like `GROUP BY` in DB. The simple one can follow [here](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-aggregations-bucket-terms-aggregation.html)


### 10. Health Check
[read here](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/cluster-health.html)


### x. You can explore more from [here](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index.html). There's a lot off stuff