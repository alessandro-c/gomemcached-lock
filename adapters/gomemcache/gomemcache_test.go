package gomemcache

import (
	"alessandro-c/gomemcached-lock/adapters"
	owntesting "alessandro-c/gomemcached-lock/testing"
	"github.com/bradfitz/gomemcache/memcache"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	tc := owntesting.Setup(t)

	mc := memcache.New(owntesting.TestServer)

	c := New(mc)

	duration := 1 * time.Minute

	c.Add("foo", "bar", duration)

	item, err := mc.Get("foo")

	if err != nil {
		t.Error(err.Error())
	}

	if item.Key != "foo" {
		t.Error("wrong key set")
	}

	if string(item.Value) != "bar" {
		t.Error("wrong value set")
	}

	err = c.Add("foos", "bar", 0)

	if err != nil {
		t.Error(err.Error())
	}

	// retry again, already existing keys should not be allowed
	err = c.Add("foo", "bars", 0)

	if err != adapters.ErrNotStored {
		t.Errorf("adapters.ErrNotStored should have been returned")
	}

	tc.Teardown()
}

func TestGet(t *testing.T) {
	tc := owntesting.Setup(t)
	mc := memcache.New(owntesting.TestServer)

	c := New(mc)
	c.Add("foo", "bar", 0)
	value, err := c.Get("foo")

	if err != nil {
		t.Error(err.Error())
	}

	if value != "bar" {
		t.Errorf("obtained value should have been 'bar' instead of '%s'", value)
	}

	_, err = c.Get("nonexistent")

	if err != adapters.ErrNotFound {
		t.Errorf("adapters.ErrNotFound should have been returned")
	}

	tc.Teardown()
}

func TestDelete(t *testing.T) {
	tc := owntesting.Setup(t)

	mc := memcache.New(owntesting.TestServer)

	c := New(mc)

	c.Add("foo", "bar", 0)

	c.Delete("foo")

	item, err := mc.Get("foo")

	if err == nil {
		t.Error("key \"foo\" should have been deleted", item)
	}

	err = c.Delete("nonexistent")

	if err != adapters.ErrNotFound {
		t.Errorf("adapters.ErrNotFound should have been returned")
	}

	tc.Teardown()
}
