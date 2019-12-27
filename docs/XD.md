- Why 2 nodes still "green" from 3 nodes, 3 primary, 1 replica ?
3 nodes = [1p 3r][2p 1r][3p 2r]
2 nodes = [1p 3p 2r][2p 3r 1r] => still balance (p & r in diff place)
1 nodes = [1p 2p 3p 1r 2r 3r] => not balance (p & r in same place)

- How to know what's shard this doc located ?
use explain:
{
	"explain": true,
    "query": {
        //...
    }
}

- How to search by document id ?
{
    "query": {
        "bool": {
            "must": [
                {
                    "match": {
                        "_id": "1"
                    }
                }
            ]
        }
    }
}

- What to do if I want to add new field ?
reindex


- How to alias and what for ?
[POST] http://localhost:9201/_aliases
{
    "actions": [
        {
            "add": {
                "index": "robots_v1",
                "alias": "robots"
            }
        }
    ]
}
for migration


- .keyword what's that?
- what happend if 1 node hold 2 index


{
                "_index": "robots_v1",
                "_type": "_doc",
                "_id": "1806",
                "_score": 3.6108608,
                "_source": {
                    "name": "Aea",
                    "release": {
                        "date": "September 28, 2013",
                        "location": "Australia/Perth",
                        "version": 3.21
                    },
                    "ability": [
                        "Detail Orientation",
                        "Raise Money"
                    ],
                    "attribute": [
                        {
                            "name": "Ice",
                            "power": 71
                        }
                    ]
                }
            },


- garbage collection