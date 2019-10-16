package testing

import (
	"net"
	"testing"
)

// TestServer is the host that points to the development memcached instance
const TestServer = "localhost:11211"

// TestCase provides useful repetitive tests functions
type TestCase struct {
	t    *testing.T
	conn net.Conn
}

// Setup will prepare the environment to be used for every test.
func Setup(t *testing.T) *TestCase {
	c, err := net.Dial("tcp", TestServer)
	if err != nil {
		t.Errorf("no server running at %s", TestServer)
	}
	c.Write([]byte("flush_all\r\n"))
	return &TestCase{t: t, conn: c}
}

// Teardown the test environment
func (tc *TestCase) Teardown() {
	tc.conn.Write([]byte("flush_all\r\n"))
	tc.conn.Close()
}
