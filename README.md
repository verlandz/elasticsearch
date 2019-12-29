# Elasticsearch
tutorial of basic elastic + sample web project with `Go`


### Introduction
- **What's Elasticsearch ?**
<br>an open source `NoSQL` with high speed and scalability for searching, [ref](https://www.elastic.co/what-is/elasticsearch).

- **Why we use Elasticsearch ?**
<br>for faster searching (of course), complex query, etc.

- **When we use Elasticsearch ?**
<br>when you have complex requirement or the data's too big.

- **What's pros and cons using Elasticsearch ?**
<br>you can read [here](https://interviewbubble.com/elasticsearch-pros-and-cons-advantages-and-disadvantages-of-elasticsearch/)

### Prerequisites
- [JDK](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/setup.html#jvm-version), for running elasticsearch
- [Golang](https://golang.org/dl/), for running scripts (generate data, implementation, etc)

### Tools
- [elasticsearch-7.5.0](https://www.elastic.co/downloads/past-releases/elasticsearch-7-5-0), for elastic
- [cerebro-0.8.5](https://github.com/lmenezes/cerebro/releases/tag/v0.8.5), for elastic's UI
- [postman-7.14.0](https://www.getpostman.com/downloads/release-notes), helper for API request

### Structure
- **cluster**
    - group for communication between nodes
- **node**
    - min `node` required = num of `replica` + 1
    - like instance
- **sharding / partition**
    - for data distribution
    - inc/dec nodes will `rebalance` the shards
    - useful for avoid max memory limit, faster searching, etc
    - divided by 2;
        - `primary(p) / master`
            - the core
        - `replica(r) / slave`
            - duplicated shard
            - useful for backup if node is dead
    - primary and replica will be placed in diff nodes
    - ex:
        ```
        3 node 3 shard 1 replica
        [node 1] shard 1 replica 3
        [node 2] shard 2 replica 1
        [node 3] shard 3 replica 2
        ```
- **index**
    - the place where you store the data
    - store as documents
    - like table in db
- **document**
    - 1 document = 1 data
    - contains field/attribute by mapping
    - store in json format
    - it's [lucene document](http://www.lucenetutorial.com/basic-concepts.html)
    - like field in db
- **mapping**
    - describing as `field`=>`value` of the field
    - existing map field `cannot be updated`

### Behaviour
- **add**
    - add data to primary
    - primary is selected by [routing](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-routing-field.html)
    - distribute to replica
- **delete**
    - "Mark" the data as delete ("Mark" only mark not delete)
    - distribute the mark to replica
- **update**
    - delete data
    - add data (Why delete then add? because data is immutable by default)
- **partial Update**
    - update data (only partial, not delete then add, but data become mutable)
- **search**
    - jump to shard where data located (thanks to [routing](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-routing-field.html))
    - then searching the data in that shard

### Health Status
it divided by 3, `green`, `yellow` and `red` [read more](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html#cluster-health-api-desc).
<br> **case:**
<br> let's say you have `3 nodes, 3 primary, 1 replica`, and suddenly the node dead one by one, what will be happened ?
- *3 nodes*, balance `green`
```
[1p 3r][2p 1r][3p 2r]
```
- *2 nodes*, still balanced (p & r in diff place), at first it will `yellow` then after reblance become `green`
```
[1p 3p 2r][2p 3r 1r]
```
- *1 nodes*, not balanced (p & r in same place), should be `red`
```
[1p 2p 3p 1r 2r 3r]
```


# Getting Started
- [Setup](https://github.com/verlandz/elasticsearch/blob/master/docs/SETUP.md), how to setup used tools
- [Walkthrough](https://github.com/verlandz/elasticsearch/blob/master/docs/WALKTHOUGH.md), how to use this project + basic elastic
- [Query](https://github.com/verlandz/elasticsearch/blob/master/docs/QUERY.md), query knowledge


# References
- [Logo](https://i.pinimg.com/600x315/2e/7e/e5/2e7ee557deb8cd3c9ded97c99ae1858e.jpg)
- [Elastic Docs](https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index.html)
- [Cerebro](https://github.com/lmenezes/cerebro)
- [Postman](https://www.getpostman.com/downloads/)