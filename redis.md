
# redis

> in-memory key-value store, often used as cache, message broker or lightweight database

- in-memory: all data lives in RAM, which makes it very fast
- REmote DIctionary Server
- use cases:
  - caching: fast response for API data, DB queries
  - message broker
  - simple database (for fast reads/writes)

## basic operations

- start redis CLI: `redis-cli`

- set a key: `SET name "John"`
- get a key: `GET name`
- delete a key: `DEL name`
- set a key with expiration: `SET temp "123" EX 10` (expires in 10 seconds)
- check if key exists: `EXISTS keyname`
- time to live: `TTL keyname`

## data types

| Type       | Example Use                |
|------------|----------------------------|
| String     | SET key "value"            |
| List       | LPUSH list "a", "b", "c"   |
| Set        | SADD myset "apple" "pear"  |
| Hash       | HSET user:1 name "Arthur"  |
| Sorted Set | Leaderboards               |
| Streams    | Event queues               |

