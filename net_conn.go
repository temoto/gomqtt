// Copyright (c) 2014 The gomqtt Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transport

import (
	"net"

	"github.com/gomqtt/packet"
)

// A NetConn is a wrapper around a basic TCP connection.
type NetConn struct {
	BaseConn

	// The underlying net.Conn.
	Conn net.Conn
}

// NewNetConn returns a new NetConn.
func NewNetConn(conn net.Conn) *NetConn {
	return &NetConn{
		BaseConn: BaseConn{
			carrier: conn,
			stream:  packet.NewStream(conn, conn),
			conn:    conn,
		},
		Conn: conn,
	}
}
