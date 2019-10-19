package testing

import (
	lock "github.com/alessandro-c/gomemcached-lock"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/alessandro-c/gomemcached-lock/adapters/gomemcache"
)

// NewTestAdapter instantiate an adapter implementation to be used for testing purposes
func NewTestAdapter(tserv string) lock.Adapter {
	mc := memcache.New(tserv)
	return gomemcache.New(mc)
}
