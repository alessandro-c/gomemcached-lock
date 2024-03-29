package gomemcache

import (
	"github.com/alessandro-c/gomemcached-lock/adapters"
	owntesting "github.com/alessandro-c/gomemcached-lock/testing"
	"github.com/bradfitz/gomemcache/memcache"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	tc := owntesting.Setup(t)

	mc := memcache.New(owntesting.TestServer)

	c := New(mc)

	duration := 1 * time.Minute

	err := c.Add("foo", "bar", duration)

	if err != nil {
		t.Errorf("should have addedd successfully")
	}

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
	err := c.Add("foo", "bar", 0)

	if err != nil {
		t.Errorf("should have addedd successfully")
	}

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

	err := c.Add("foo", "bar", 0)

	if err != nil {
		t.Errorf("should have addedd successfully")
	}

	err = c.Delete("foo")

	if err != nil {
		t.Errorf("should have deleted successfully")
	}

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
