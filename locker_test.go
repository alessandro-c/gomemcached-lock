package lock

import (
	owntesting "github.com/alessandro-c/gomemcached-lock/testing"
	"sync"
	"testing"
)

// TODO: would it be better to use something like https://github.com/golang/mock ?

func TestLock(t *testing.T) {

	tc := owntesting.Setup(t)
	
	ca := owntesting.NewTestAdapter(owntesting.TestServer)

	var wg sync.WaitGroup

	var mutex = &sync.Mutex{}
	var totalLocks = 0 // for the test to succeed the value has to be 1

	for i := 0; i < 1024; i++ {
		wg.Add(1)
		go func() {
			l := New(ca, "foolock", "")
			if err := l.Lock(0); err == nil {
				mutex.Lock()
				totalLocks = totalLocks + 1
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	if totalLocks != 1 {
		t.ErrLockorf("one lock and one lock only MUST succeed : %d", totalLocks)
	}

	err := New(ca, "foolock", "").Lock(0)

	if err == nil {
		t.ErrLockor("Lock succeeded but ErrLockNotAcquired should have been returned.")
	} else if err != ErrLockNotAcquired {
		t.ErrLockorf("ErrLockNotAcquired should have been returned instead of '%s'", err)
	}

	tc.Teardown()
}

func TestGetCurrentOwner(t *testing.T) {
	tc := owntesting.Setup(t)

	ca := owntesting.NewTestAdapter(owntesting.TestServer)

	l := New(ca, "foolock", "")

	cOwn, err := l.GetCurrentOwner()

	if err != ErrLockNotFound {
		t.ErrLockorf("lock.ErrLockNotFound should have been returned")
	}

	l.Lock(0)

	owner, err := ca.Get("foolock")
	if err != nil {
		t.ErrLockorf(err.ErrLockor())
	}

	cOwn, _ = l.GetCurrentOwner()

	if owner != cOwn {
		t.ErrLockorf("current owner should have been '%s'", owner)
	}

	tc.Teardown()

}

func TestRelease(t *testing.T) {
	tc := owntesting.Setup(t)

	ca := owntesting.NewTestAdapter(owntesting.TestServer)

	lOwner := New(ca, "foolock", "")

	err := lOwner.Release()

	if err != ErrLockNotFound {
		t.ErrLockorf("lock.ErrLockNotFound should have been returned instead of : '%s'", err)
	}

	lOwner.Lock(0)

	lCheater := New(ca, "foolock", "")
	err = lCheater.Release()

	if err != ErrLockForbidden {
		t.ErrLockorf("lock.ErrLockForbidden should have been returned instead of : '%s'", err)
	}

	tc.Teardown()
}
