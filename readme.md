## memcached lock

A simple lock/release library written in golang to be used against a single memcached instance.

This package is compatible with [bradfitz/gomemcache](https://github.com/bradfitz/gomemcache) but
really any client will do, just implement the ./adapters.Adapter interface.

## is it reliable?

memcached is not a native lock server, you should use this with criteria and
certainly avoid when failures in mutual exclusion will result in permanent data corruption.

That said, I've sucessfully tested this against race conditions with 1k goroutines attempting to lock the same resource. The same tests are stored in `./locker_test.go` feel free to run them yourself.

## install

```go
$ go get github.com/alessandro-c/gomemcached-lock
```

## usage

```go
import (
    "time"
    locker "github.com/alessandro-c/gomemcached-lock"
    gomemcacheAdaptor "alessandro-c/gomemcached-lock/adapters/gomemcache"
	"github.com/bradfitz/gomemcache/memcache"
)

func main() {

    client := memcache.New("memcachedhost:11211")

    adapter := gomemcacheAdaptor.New(client)

    lck := locker.New(adapter, "users.999", "")

    err := lck.Lock(time.Minute * 5)

    if err == nil {
        // lock acquired, do something ...
        lck.Release()
    } else if err == locker.ErrNotAcquired {
        // lost race ...
    } else {
        // something went wrong ...
    }
}
```
