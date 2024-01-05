// Copyright 2015 Matthew Holt and The Caddy Authors
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

// Even though the filename ends in _unix.go, we still have to specify the
// build constraint here, because the filename convention only works for
// literal GOOS values, and "unix" is a shortcut unique to build tags.
//go:build unix && solaris

package caddy

import (
	"syscall"
)

// reusePort sets SO_REUSEPORT on other Unix.
// Not present on Solaris 10 or Illumos.
func reusePort(network, address string, conn syscall.RawConn) error {
	return nil
}
