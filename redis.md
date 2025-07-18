
# redis

> in-memory key-value store, often used as cache, message broker or lightweight database

- in-memory: all data lives in RAM, which makes it very fast
- not structured
- REmote DIctionary Server
- use cases:
  - caching: fast response for API data, DB queries
  - message broker
  - simple database (for fast reads/writes)

## basic operations

- start redis CLI: `redis-cli`

- set a key: `SET name "John"`
- get a key: `GET name`
- increment a key: `INCR counter`
- decrement a key: `DECR counter`
- delete a key: `DEL name`
- set a key with expiration: `EXPIRE mykey 10` (expires in 10 seconds)
- check if key exists: `EXISTS keyname`
- time to live: `TTL keyname`

- add a value to the beginning of a list: `LPUSH tasks "Task1"`
- add a value to the end of a list: `RPUSH tasks "Task1"`
- return a range of values from a list: `LRANGE list_name start_position stop_position`
  - e.g. print all elements in a list: `LRANGE tasks 0 -1`
    - `0`: first position
    - `-1`: last position
- remove and return the first element: `LPOP mylist`
- remove and return the last element: `RPOP mylist`

- add member(s) to a set: `SADD myset "apple" "banana" "cherry"`
- remove member(s) to a set: `SREM myset "apple" "banana" "cherry"`
- get all members of a set: `SMEMBERS myset`
- check if a value is in a set: `SISMEMBER set1 "a"`

- get all keys that match a pattern: `KEYS my*` => example output: "myKey", "myCounter", "myExample"
  - time complexity O(n)
- checks if one or more keys exist: `EXISTS myKEY` => example output: (integer) 1
  - time complexity O(n)


> [!NOTE]
> redis doesn't support chaining commands like SET counter 10 INCR counter in a single line
> each command must be sent separately

## data types

| Type       | Example Use                |
|------------|----------------------------|
| String     | SET key "value"            |
| List       | LPUSH list "a", "b", "c"   |
| Set        | SADD myset "apple" "pear"  |
| Hash       | HSET user:1 name "Arthur"  |
| Sorted Set | Leaderboards               |
| Streams    | Event queues               |

