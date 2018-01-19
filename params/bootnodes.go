// Copyright 2015 The go-ethereum Authors
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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Trinity Foundation (soon) Go Bootnodes
	"enode://ae808feecd7c92e4355c782e811729ef1419e7bbcf225fbf03b726a7eaa1e2a4c22fca527befdb80abcba5c4b0a53e4237a9be2268d706e5322977b59e0f4057@40.79.84.146:30303",
	"enode://874c03b1db99e213d7a2a772029fce677ec00f8520e1362ff07c68063b0d5a74a3cb8fd87b1583b9ea7ce2b010c334b992255f7b30161197328abf2dc8767dd3@52.191.171.167:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
