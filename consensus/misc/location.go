// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package misc

import (
	"fmt"

	"github.com/spruce-solutions/go-quai/core/types"
	"github.com/spruce-solutions/go-quai/params"
)

// VerifyLocation verifies the location of an incoming header
func VerifyLocation(config *params.ChainConfig, header *types.Header) error {
	if len(config.MapLocation) < int(header.Location[0]) {
		return fmt.Errorf("invalid location mined for header: have 0x%x, want 0x%x", header.Location, config.MapLocation)
	}
	if config.MapLocation[int(header.Location[0])] < int(header.Location[1]) {
		return fmt.Errorf("invalid location mined for header: have 0x%x, want 0x%x", header.Location, config.MapLocation)
	}
	// All ok, return
	return nil
}
