# configfile
Read config from file, useful when read data from kubernetes configmaps, and secret.

## Example

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/acoshift/configfile"
	"github.com/garyburd/redigo/redis"
)

var config = configfile.NewReader("config")

var (
	addr      = config.StringDefault("addr", ":8080")
	redisAddr = config.MustString("redis_addr")
	redisPass = config.String("redis_pass")
	redisDB   = config.Int("redis_db")
)

func main() {
	pool := redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				redisAddr,
				redis.DialPassword(redisPass),
				redis.DialDatabase(redisDB),
			)
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := pool.Get()
		defer c.Close()
		cnt, err := redis.Int64(c.Do("INCR", "cnt"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "count: %d", cnt)
	})

	http.ListenAndServe(addr, nil)
}
```
