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

// To run just this test, from the root of the cockroach repo:
// CGO_CFLAGS="-I../gorocks/rocksdb/include" CGO_LDFLAGS="-L../gorocks/rocksdb -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy" go test ./disk/rocksdb
package rocksdb

import (
	"testing"

	"github.com/cockroachdb/cockroach/disk"
	"github.com/cockroachdb/gorocks"
)

func TestDisk(t *testing.T) {
	d := openTestDB()

	key := disk.Key([]byte("hello"))
	msg := "durable"
	d.Put(key, disk.Value{
		Bytes:      []byte(msg),
		Timestamp:  0,
		Expiration: 0})

	checkGet(t, d, key, msg)

	close(d)
	d = openTestDB()
	checkGet(t, d, key, msg)
	close(d)
}

func checkGet(t *testing.T, d disk.Disk, key disk.Key, expected string) {
	val, err := d.Get([]byte(key))
	if err != nil {
		t.Fatal(err)
	}
	if string(val.Bytes) != expected {
		t.Fatalf("expected %s got %s", expected, string(val.Bytes))
	}
}

func openTestDB() *Disk {
	db, err := gorocks.Open("testdb", gorocks.DBOptions{CreateIfMissing: true})
	if err != nil {
		panic(err)
	}

	return New(db)
}

func close(d *Disk) {
	gorocks.Close(d.db)
}
