package adapter

import (
	"github.com/alessandro-c/gomemcached-lock/adapters"
	"github.com/alessandro-c/gomemcached-lock/adapters/gomemcache"
	"github.com/bradfitz/gomemcache/memcache"
)

// NewTestAdapter instantiate an adapter implementation to be used for testing purposes
func NewTestAdapter(tserv string) adapters.Adapter {
	mc := memcache.New(tserv)
	return gomemcache.New(mc)
}
