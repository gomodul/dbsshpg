package dbsshpg

import (
	"database/sql/driver"
	"net"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/ssh"
)

// Dialer ...
type Dialer struct {
	Client *ssh.Client
}

// Open ...
func (dialer *Dialer) Open(dsn string) (driver.Conn, error) {
	return pq.DialOpen(dialer, dsn)
}

// Dial ...
func (dialer *Dialer) Dial(network, address string) (net.Conn, error) {
	return dialer.Client.Dial(network, address)
}

// DialTimeout ...
func (dialer *Dialer) DialTimeout(network, address string, _ time.Duration) (net.Conn, error) {
	return dialer.Client.Dial(network, address)
}
