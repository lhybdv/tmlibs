// Copyright 2018 The Trias Co., LTD. All rights reserved.
// Use of this source code is governed by a GNU General Public License v3.0
// license that can be found in the LICENSE file.

package db

import (
	"github.com/spf13/viper"
	"path"
	"sync"
	"github.com/trias-lab/gondwana/file"
)

func init() {
	registerDBCreator(TriasDBBackendStr, func(name string, dir string) (DB, error) {
		// just handle block store with TriasDB, others with GoLevelDB
		if name == "blockstore" {
			return NewTriasDB(dir), nil
		}
		// return NewGoLevelDB(name, dir)
		return NewGoLevelDB(name, dir)
	}, false)
}

// TriasDB is used for Trias system
type TriasDB struct {
	mtx   sync.Mutex
	dir   string
	store *file.Store
}

// NewTriasDB create a TriasDB instance
func NewTriasDB(dir string) *TriasDB {
	workDir := path.Join(dir, "store")
	vbackend := viper.Get("warehouse_backend")
	vaddress := viper.Get("warehouse_address")
	backend := vbackend.(string)
	address := vaddress.(string)
	opt := file.StoreOptions{
		FileSize: (1 << 20) * 10,
		WarehouseBackend: backend,
		WarehouseAddress:  address,
	}
	store, err := file.NewFileStore(workDir, opt)
	if err != nil {
		panic(err)
	}
	database := &TriasDB{
		dir:   dir,
		store: store,
	}
	return database
}

// Get implemented for interface DB
func (db *TriasDB) Get(key []byte) []byte {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	// TODO: unimplement

	value, err := db.store.Get(key)
	if err != nil {
		panic(err)
	}
	return value
}

// Has implemented for interface DB
func (db *TriasDB) Has(key []byte) bool {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	// TODO: unimplement
	return false
}

// Set implemented for interface DB
func (db *TriasDB) Set(key []byte, value []byte) {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	db.SetNoLock(key, value)
}

// SetSync implemented for interface DB
func (db *TriasDB) SetSync(key []byte, value []byte) {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	db.SetNoLock(key, value)
}

// SetNoLock implemented for interface DB
func (db *TriasDB) SetNoLock(key []byte, value []byte) {
	err := db.store.Set(key, value)
	if err != nil {
		panic(err)
	}
}

// Delete implemented for interface DB
func (db *TriasDB) Delete(key []byte) {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	db.DeleteNoLock(key)
}

// DeleteSync implemented for interface DB
func (db *TriasDB) DeleteSync(key []byte) {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	db.DeleteNoLock(key)
}

// DeleteNoLock implemented for interface DB
func (db *TriasDB) DeleteNoLock(key []byte) {
	// TODO: unimplement
	panic("TriasDB.DeleteNoLock not yet implemented")
}

// Close implemented for interface DB
func (db *TriasDB) Close() {
	// TODO: unimplement
	panic("TriasDB.Close not yet implemented")
}

// Print implemented for interface DB
func (db *TriasDB) Print() {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	// TODO: unimplement
	panic("TriasDB.Print not yet implemented")
}

// Stats implemented for interface DB
func (db *TriasDB) Stats() map[string]string {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	// TODO: unimplement
	panic("TriasDB.Stats not yet implemented")
}

// NewBatch implemented for interface DB
func (db *TriasDB) NewBatch() Batch {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	panic("TriasDB.NewBatch not yet implemented")
}

// Mutex implemented for interface DB
func (db *TriasDB) Mutex() *sync.Mutex {
	return &(db.mtx)
}

// Iterator implemented for interface DB
func (db *TriasDB) Iterator() Iterator {
	return db.MakeIterator(false)
}

// MakeIterator implemented for interface DB
func (db *TriasDB) MakeIterator(isReversed bool) Iterator {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	// TODO: unimplement
	panic("TriasDB.MakeIterator not yet implemented")
}

// ReverseIterator implemented for interface DB
func (db *TriasDB) ReverseIterator(start, end []byte) Iterator {
	return db.MakeIterator(true)
}
