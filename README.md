# configfile

![Build Status](https://github.com/acoshift/configfile/actions/workflows/test.yaml/badge.svg?branch=master)
[![codecov](https://codecov.io/gh/acoshift/configfile/branch/master/graph/badge.svg)](https://codecov.io/gh/acoshift/configfile)
[![Go Report Card](https://goreportcard.com/badge/github.com/acoshift/configfile)](https://goreportcard.com/report/github.com/acoshift/configfile)
[![GoDoc](https://godoc.org/github.com/acoshift/configfile?status.svg)](https://godoc.org/github.com/acoshift/configfile)

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

## Example YAML

```go
package main

import (
    "log"
    "net/http"

    "github.com/acoshift/configfile"
)

func main() {
    var config = configfile.NewReader("testdata/config.yaml")
    // or use NewYAMLReader
    var config = configfile.NewYAMLReader("testdata/config.yaml")

    log.Println(config.Bool("data1")) // true
    log.Println(config.String("data2")) // false
    log.Println(config.Int("data3")) // 9
    log.Println(config.Int("data4")) // 0
    log.Println(config.String("empty")) // ""
}
```
