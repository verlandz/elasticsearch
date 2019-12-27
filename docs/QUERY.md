
<!-- 
    
    - [Search by URI](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-uri-request.html)
    - [Search by Body](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html)
    - [Search in 1 or more index](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-search.html#search-search-api-example)
    - [How to Explain](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-explain)
    - [How to Offset](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-from-size)
    - [How to Sort](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-sort)
    - [How to Get Mapping](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/indices-get-mapping.html)
    - [Data Types](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-types.html)
    - [Date Format](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-date-format.html)
 -->

- https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-search.html
- https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-uri-request.html
- https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html

- https://www.elastic.co/guide/en/elasticsearch/reference/current/search-uri-request.html
- https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-body.html

- https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html
- https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-match-query.html
- https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html
- https://www.elastic.co/guide/en/elasticsearch/reference/7.1/search-suggesters-completion.html
- https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-wildcard-query.html
- https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-edgengram-tokenizer.html
- https://www.elastic.co/guide/en/elasticsearch/reference/current/compound-queries.html

Aggregations
- https://www.elastic.co/guide/en/elasticsearch/reference/current/search-aggregations.html

Health
- https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html

Sample ElasticSearch UI
- Let's try https://github.com/lmenezes/cerebro
- download the .zip/.tgz from https://github.com/lmenezes/cerebro/releases
- then edit `conf/application.conf` in hosts part
hosts = [
  {
   host = "http://localhost:9201"
   name = "cluster-robot"
   headers-whitelist = [ "x-proxy-user", "x-proxy-roles", "X-Forwarded-For" ]
  },
  {
   host = "http://localhost:9202"
   name = "cluster-robot"
   headers-whitelist = [ "x-proxy-user", "x-proxy-roles", "X-Forwarded-For" ]
  },
  {
   host = "http://localhost:9203"
   name = "cluster-robot"
   headers-whitelist = [ "x-proxy-user", "x-proxy-roles", "X-Forwarded-For" ]
  }
]
- then run `./bin/cerebro` or `./bin/cerebro.bat` (you need to download JDE/JDK first)
- go to http://localhost:9000
- good explaination : https://www.youtube.com/watch?v=ZjVmg9fftUM&list=LLbjgbCc74UHvkqzFemEiBVQ&index=2&t=0s

Usefullinks (host:localhost:9201)
- http://localhost:9201
- http://localhost:9201/_cat/nodes?v
- http://localhost:9201/_cat/shards?v
- http://localhost:9201/_cat/indices?v
- http://localhost:9201/_cat/health?v
- http://localhost:9201/_cluster/health

> frontend
- autocomplete
- what's inverted index
- how to boosting > minimum_should_match

> SOURCE : 
-   IMG : https://i.pinimg.com/600x315/2e/7e/e5/2e7ee557deb8cd3c9ded97c99ae1858e.jpg
-   SEARCH : https://bootstrapious.com/p/bootstrap-search-bar
-   CARD : https://www.webhozz.com/code/bootstrap-card/
-   LABEL: https://www.w3schools.com/bootstrap/bootstrap_badges_labels.asp
-   BUTTONS : https://getbootstrap.com/docs/4.0/components/buttons/
