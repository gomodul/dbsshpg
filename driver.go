package dbsshpg

import (
	"database/sql"

	"golang.org/x/crypto/ssh"
)

// Driver ...
type Driver struct {
	name string
}

// Name ...
func (d *Driver) Name() string {
	return d.name
}

// Register ...
func (d *Driver) Register(client *ssh.Client) {
	d.name = "postgres+" + client.LocalAddr().Network() + "+" + client.LocalAddr().String()

	for _, driver := range sql.Drivers() {
		if driver == d.name {
			return
		}
	}

	sql.Register(d.name, &Dialer{Client: client})
}
