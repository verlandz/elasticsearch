# Setup
these guide only change necessary things and all setup's results will be put into `tools`

## Elasticsearch
using `elasticsearch-7.5.0`

### Config
- [download here](https://www.elastic.co/downloads/past-releases/elasticsearch-7-5-0)
- put it into `tools`
- copy 3x and set with diff name
  - why 3x? because we want to up with `3 nodes`. Each node 1 folder.
  - eg: `node-robot-1`, `node-robot-2` and `node-robot-3`
  <br>![folder-es](https://i.ibb.co/gZ3dYF3/5.png)
- for each node, edit *these files* that located in `tools/{YOUR_NODE}/config`:
  - `elasticsearch.yml`
  - `jvm.options`
- You can override it from the sample:
  - `tools/sample_conf/node-robot-1/config` for `node-robot-1`
  - `tools/sample_conf/node-robot-2/config` for `node-robot-2`
  - `tools/sample_conf/node-robot-3/config` for `node-robot-3`
- changes in `elasticsearch.yml`:
  - <i>**cluster.name**, define name of cluster.</i>
    - each nodes must have the `same` name
    - eg: `cluster-robot`
  - <i>**node.name**, define name of node.</i>
    - better use `different` name for each nodes to ease tracking.
    - eg: `node-robot-1`, `node-robot-2`, and `node-robot-3`
  - <i>**http.port**, define port of node.</i>
    - by default `:9000`
    - since we want to test in localhost (1 IP with 3 nodes), we should differentiated the port for each node
    - eg: 
      - `:9201` for `node-robot-1`
      - `:9202` for `node-robot-2`
      - `:9203` for `node-robot-3`
  - <i>**discovery.seed_hosts**, where to host.</i>
    - set `["127.0.0.1"]` for localhost
  - <i>**cluster.initial_master_nodes**, nodes in the cluster.</i>
    - each nodes must have the `same` value.
    - set `["node-robot-1","node-robot-2","node-robot-3"]`, to connect those nodes in this cluster.
- changes in `jvm.options`:
  - <i>**JVM heap size**, size of heap.</i>
    - the min and max heap should be same, read more on these desc.
    - change the value to suit you RAM capacity
    - *remember, the size will be produced is <b>total_node * heap_size</b> (if every node have the same size)*

### Run
- open 3 terminals.
- cd to `tools/node-robot-1`, `tools/node-robot-2` and `tools/node-robot-3`
- for each, run `./bin/elasticsearch`
- check `http://localhost:9201/`, `http://localhost:9202/` and `http://localhost:9203/` to ensure node activated.
![_](https://i.ibb.co/BrrvYqL/7.png)
- check `http://localhost:9201/_cat` to get list all info related to nodes.
![_cat](https://i.ibb.co/wd47bQ6/8.png )
```
You can focus here:
http://localhost:9201/_cat/allocation?v
http://localhost:9201/_cat/nodes?v
http://localhost:9201/_cat/shards?v
http://localhost:9201/_cat/indices?v

/shards and /indicies will be unlocked once you have at least 1 index
```


## Cerebro
using `cerebro-0.8.5`, for elastic UI

### Config
- [download here](https://github.com/lmenezes/cerebro/releases/tag/v0.8.5)
- put it into `tools`
  <br>![folder-cerebro](https://i.ibb.co/gR7Ry8W/6.png)
- edit `tools/cerebro/conf/application.conf`
- You can override it from the sample `tools/sample_conf/cerebro/conf/application.conf`
- changes in `application.conf`:
  - <i>**hosts**, list of know hosts.</i>
    - simply just change `host` and `name` for `each` of your node
    
    ```
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
    ```

### Run
- open 1 new terminal.
- cd to `tools/cerebro` and run `./bin/cerebro`
- go to `http://localhost:9000` and if run correctly you will see this
![cerebro-home](https://i.ibb.co/N2YkmKm/1.png)
- click the `cluster-robot`, and it will direct you to dashboard
![cerebro-dashboard](https://i.ibb.co/4WnXRcY/2.png)
- you can look [here](https://youtu.be/ZjVmg9fftUM?list=LLbjgbCc74UHvkqzFemEiBVQ&t=115) for explanation of cerebro's feature