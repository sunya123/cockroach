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

// Package mem implements an in-memory disk
// Not thread safe and not durable (obviously)
//
// One notable caveat is that key's end up being represented as strings
// rather than byte slices
package mem

import "github.com/cockroachdb/cockroach/disk"

type MemDisk struct {
	cacheSize int64
	data      map[string][]byte
}

func New(cacheSize int64) *MemDisk {
	return &MemDisk{
		cacheSize: cacheSize,
		data:      make(map[string][]byte),
	}
}

func (b *MemDisk) Put(key disk.Key, value disk.Value) error {
	b.data[string(key)] = value.Bytes
	return nil
}

func (b *MemDisk) Get(key disk.Key) (disk.Value, error) {
	return disk.Value{
		Bytes: b.data[string(key)],
	}, nil
}

func (b *MemDisk) Del(key disk.Key) error {
	delete(b.data, string(key))
	return nil
}

func (r *MemDisk) Capacity() (*disk.Capacity, error) {
	panic("not implemented")
}
