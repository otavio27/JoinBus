/*
 MODULE: TTKDb.go
 AUTHOR: Leo Schneider <schleo@outlook.com>
 DATE  : March 2018
 INFO  : This module handles read/write to Key-Value store (BoltDB)
*/

package ttktools

import (
	"context"
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
)

// ttkdb ...
type ttkdb struct {
	module string
	ctx    context.Context
	log    *ttklog
	db     *bolt.DB
	dbfile string
}

// Db Returns an instance of ttkdb
func Db(ctx context.Context, dbfile string, logger *ttklog) *ttkdb {
	var err error
	t := ttkdb{}
	t.module = "ttkdb"
	t.ctx = ctx
	t.log = logger
	t.dbfile = dbfile
	t.db, err = bolt.Open(dbfile, 0600, nil)
	if err != nil {
		t.log.Msg(ctx, t.module, err.Error(), DEBUG)
	}
	return &t
}

// Write a key in a bucket
func (t *ttkdb) Write(bucket string, key string, value []byte) {
	err := t.db.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("Failed to create bucket: %v", err)
		}
		err = bk.Put([]byte(key), value)
		if err != nil {
			return fmt.Errorf("Failed to insert/update '%s': %v", key, string(value))
		}
		return nil
	})
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
	}
}

// Read a key from a bucket
func (t *ttkdb) Read(bucket string, key string) []byte {
	var ret []byte
	err := t.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(bucket))
		if bk == nil {
			return fmt.Errorf("Failed to open bucket: %v", bucket)
		}
		ret = bk.Get([]byte(key))
		if ret == nil {
			return fmt.Errorf("Failed to retrieve '%s'", key)
		}
		return nil
	})
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return nil
	}
	return ret
}

// List the keys from a bucket
func (t *ttkdb) List(bucket string) map[string]string {
	ret := make(map[string]string)
	err := t.db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(bucket))
		if bk == nil {
			return fmt.Errorf("Failed to open bucket: %v", bucket)
		}
		c := bk.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			ret[string(k)] = string(v)
		}
		return nil
	})
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return nil
	}
	return ret
}

// Scan for a key in a bucket
func (t *ttkdb) Scan(bucket string, query string) map[string]string {
	ret := make(map[string]string)
	err := t.db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(bucket))
		if bk == nil {
			return fmt.Errorf("Failed to open bucket: %v", bucket)
		}
		c := bk.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if strings.Contains(string(v), query) {
				ret[string(k)] = string(v)
			}
		}
		return nil
	})
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return nil
	}
	return ret
}

// Delete a Key from a bucket
func (t *ttkdb) Delete(bucket string, key string) error {
	err := t.db.Update(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(bucket))
		if bk == nil {
			return fmt.Errorf("Failed to open bucket: %v", bucket)
		}
		err := bk.Delete([]byte(key))
		if err != nil {
			return fmt.Errorf("Failed to delete '%s'", key)
		}
		return nil
	})
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
	}
	return nil
}

// ExistBucket checks is a bucket exists
func (t *ttkdb) ExistBucket(bucket string) bool {
	err := t.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(bucket))
		if bk == nil {
			return fmt.Errorf("Failed to open bucket: %v", bucket)
		}
		return nil
	})
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return false
	}
	return true
}

// CreateBucket creates an empty bucket
func (t *ttkdb) CreateBucket(bucket string) bool {
	err := t.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("Failed to create bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return false
	}
	return true
}

// Count the keys in a bucket
func (t *ttkdb) Count(bucket string) int {
	var ret int
	err := t.db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(bucket))
		if bk == nil {
			return fmt.Errorf("Failed to open bucket: %v", bucket)
		}
		x := bk.Stats()
		ret = x.KeyN
		return nil
	})
	if err != nil {
		t.log.Msg(t.ctx, t.module, err.Error(), DEBUG)
		return 0
	}
	return ret
}
