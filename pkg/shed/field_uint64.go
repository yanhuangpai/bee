// Copyright 2018 The go-ethereum Authors
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

package shed

import (
	"encoding/binary"

	"github.com/dgraph-io/badger/v2"
	"github.com/ethersphere/bee/pkg/logging"
)

// Uint64Field provides a way to have a simple counter in the database.
// It transparently encodes uint64 type value to bytes.
type Uint64Field struct {
	db     *DB
	key    []byte
	logger logging.Logger
}

// NewUint64Field returns a new Uint64Field.
// It validates its name and type against the database schema.
func (db *DB) NewUint64Field(name string, logger logging.Logger) (f Uint64Field, err error) {
	key, err := db.schemaFieldKey(name, "uint64")
	if err != nil {
		return f, err
	}
	return Uint64Field{
		db:     db,
		key:    key,
		logger: logger,
	}, nil
}

// Get retrieves a uint64 value from the database.
// If the value is not found in the database a 0 value
// is returned and no error.
func (f Uint64Field) Get() (val uint64, err error) {
	b, err := f.db.Get(f.key)
	if err != nil {
		if err == ErrNotFound {
			f.logger.Errorf("key %s not found", string(f.key))
			return 0, nil
		}
		return 0, err
	}
	return binary.BigEndian.Uint64(b), nil
}

// Put encodes uin64 value and stores it in the database.
func (f Uint64Field) Put(val uint64) (err error) {
	return f.db.Put(f.key, encodeUint64(val))
}

// PutInBatch stores a uint64 value in a batch
// that can be saved later in the database.
func (f Uint64Field) PutInBatch(batch *badger.Txn, val uint64) (err error){
	return batch.Set(f.key, encodeUint64(val))
}

// Inc increments a uint64 value in the database.
// This operation is not goroutine save.
func (f Uint64Field) Inc() (val uint64, err error) {
	val, err = f.Get()
	if err != nil {
		return 0, err
	}
	val++
	return val, f.Put(val)
}

// IncInBatch increments a uint64 value in the batch
// by retreiving a value from the database, not the same batch.
// This operation is not goroutine save.
func (f Uint64Field) IncInBatch(batch *badger.Txn) (val uint64, err error) {
	val, err = f.Get()
	if err != nil {
		return 0, err
	}
	val++
	err = f.PutInBatch(batch, val)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// Dec decrements a uint64 value in the database.
// This operation is not goroutine save.
// The field is protected from overflow to a negative value.
func (f Uint64Field) Dec() (val uint64, err error) {
	val, err = f.Get()
	if err != nil {
		return 0, err
	}
	if val != 0 {
		val--
	}
	return val, f.Put(val)
}

// DecInBatch decrements a uint64 value in the batch
// by retreiving a value from the database, not the same batch.
// This operation is not goroutine save.
// The field is protected from overflow to a negative value.
func (f Uint64Field) DecInBatch(batch *badger.Txn) (val uint64, err error) {
	val, err = f.Get()
	if err != nil {
		return 0, err
	}
	if val != 0 {
		val--
	}
	err = f.PutInBatch(batch, val)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// encode transforms uint64 to 8 byte long
// slice in big endian encoding.
func encodeUint64(val uint64) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, val)
	return b
}