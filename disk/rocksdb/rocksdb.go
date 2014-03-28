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

// Package rocksdb implements the Disk interface backed by RocksDB.
// Clients should only create a single rocks DB instance per physical disk.
package rocksdb

import (
	"github.com/cockroachdb/cockroach/disk"
	"github.com/cockroachdb/gorocks"
)

type Disk struct {
	db *gorocks.DB
}

// TODO(shawn) think through how snapshot / iterator control is exposed
// For now just expose the most minimal Disk interface.
func New(db *gorocks.DB) *Disk {
	return &Disk{db}
}

// An implementation that flushes with every put for now
// TODO(shawn) expose multi-key put batch / flush control
func (d *Disk) Put(key disk.Key, val disk.Value) error {
	return gorocks.Put(d.db, key, val.Bytes, gorocks.WriteOptions{Sync: true})
}

func (d *Disk) Get(key disk.Key) (disk.Value, error) {
	bytes, err := gorocks.Get(d.db, key, gorocks.ReadOptions{})
	if err != nil {
		return disk.Value{}, err
	}

	result := disk.Value{}
	result.Bytes = bytes
	// XXX timestamp / expiration
	return result, nil
}

func (d *Disk) Del(key disk.Key) error {
	panic("not implemented")
}

func (d *Disk) Capacity() (*disk.Capacity, error) {
	panic("not implemented")
}
