# Walkthrough
a guideline step by step to do this project + basic elastic<br>

### 0. Setup
- [read here](https://github.com/verlandz/elasticsearch/blob/master/docs/SETUP.md)

### 1. Create Index 
- run `Create Index` in postman. If success, it will return:
```
{
    "acknowledged": true,
    "shards_acknowledged": true,
    "index": "robots_v1"
}
```
- You can modify `number_of_shards` and `number_of_replicas`, but this project use 3 shards and 1 replica.
- now check agian `http://localhost:9201/_cat/shards?v` and `http://localhost:9201/_cat/indices?v`.


### 2. Mapping
- run `Get Mapping` in postman. You will get empty like this:
```
{
    "robots_v1": {
        "mappings": {}
    }
}
```
- run `Insert Mapping`. This is manual mapping, the elastic will automatically create new mapping when you insert data.
- run `Get Mapping`. New map is added, and should be like this
```
{
    "robots_v1": {
        "mappings": {
            "properties": {
                "test": {
                    "type": "date",
                    "format": "MMM dd, yyyy"
                }
            }
        }
    }
}
```
- change `"type": "date"` to `"type": "string"` in `Insert Mapping`, then run again. You will get error like this.
```
{
    "error": {
        "root_cause": [
            {
                "type": "mapper_parsing_exception",
                "reason": "No handler for type [string] declared on field [test]"
            }
        ],
        "type": "mapper_parsing_exception",
        "reason": "No handler for type [string] declared on field [test]"
    },
    "status": 400
}
```
- the error means, once you already declared the type of field, it can't be updated.
- to solve this, you need to do `reindex` by create `new index` and new mapping then inject the data again.
- injecting data not from old index, but from your `Main DB`.
- in this situation we don't have to `reindex`, we can just recreate it.
- let's recreate it by run `Delete Index` then `Create Index`
- for mapping types, you can [read here](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-types.html)

### 3. Alias
- run `Alias` in postman.
- this one will be useful later when you want to change the index without change the config. Here's the result in cerebro.
![alias](https://i.ibb.co/fSThndQ/3.png)

### 4. Insert data by generate
- cd `scripts/seed` and run `seed.go`
- check `Get Mapping` in postman and you'll see the new mapping is automatically inserted
- search `Search w/o param/body` in postman and you'll get the results from generator
- check `http://localhost:9201/_cat/shards?v` and you will see docs distribution
- check `Cerebro` and you will see docs is increased
![alias](https://i.ibb.co/8rWwgQh/x2.png)

### 5. Specific docs - insert, update, update partial, search, delete
- run `Insert/Update Doc` in postman for insert
- then run `Search by Doc` to see the result
- back to `Insert/Update Doc`, change the body into `{ "name": "gone", "ability": [ "disappear" ] }` then run it again.
- run `Search by Doc` to see the result.
- in here you will know that update is basically `FULL REPLACE` the doc, so be carefull.
- now run `Partial Update Doc`, to try partial update.
- run `Search by Doc` again and you'll see only desired filed that got updated
- run `Delete Doc`, to delete the doc

### 6. Web
- cd `scripts/web` and run `web.go`
- go to `http://localhost:8080/search` and search something, you'll see like this
![web](https://i.ibb.co/4NCr2kZ/x4.png)
- if you want to increased the data, simply just change `const N_ROBOT = 10` to something in `scripts/seed/seed.go` then re-run