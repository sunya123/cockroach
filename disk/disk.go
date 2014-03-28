// Copyright 2014 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.  See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Shawn Morel (shawn@strangemonad.com)

// Package disk implements local (single-node) storage.
// There are no cluster-level concernces like replication at this level.
package disk

type Key []byte

// Value specifies the value at a key. Multiple values at the same key
// are supported based on timestamp. Values which have been overwritten
// have an associated expiration, after which they will be permanently
// deleted.
type Value struct {
	// Bytes is the byte string value.
	Bytes []byte
	// Timestamp of value in nanoseconds since epoch.
	Timestamp int64
	// Expiration in nanoseconds.
	Expiration int64

	// XXX (shawn) does a a map of timestamp -> bytes make more sense than timestamped values?
}

type Capacity struct {
}

// A KV provides Key-Value storage.
//
// Implementations are *NOT* expected to be thread-safe - i.e.
// callers are expected to serialize access.
//
// Different implementations will obviously have different efficiency and performance
// characteristics e.g. control over flushing, compaction, durability etc.
type KV interface {

	// Trying to reconcile what I had in mind with the Engine API
	// but some things don't make sense
	// Questions:
	// - would put(k, t, v) or put(k, v, [t]) work better than put(k,v) where v has timestamp
	// - how does del(k) work for multi-valued keys?
	// - for get(k) how do we represent multiple value results?
	// - can we simply not fetch a specific value at a timestamp?
	// - is this making too many assumptions

	// Put sets the given key to the value provided.
	Put(key Key, val Value) error
	// Get returns the value for the given key, nil otherwise.
	Get(key Key) (Value, error)
	// Delete removes the item from the db with the given key.
	Del(key Key) error

	// TODO(shawn) expose interfaces for batch gets since different engines
	// might be able to support that more efficiently.
	// same for range scans
}

type Info interface {
	// Capacity returns capacity details for the engine's available storage.
	Capacity() (*Capacity, error)
}
