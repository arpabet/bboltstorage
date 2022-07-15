/**
  Copyright (c) 2022 Arpabet, LLC. All rights reserved.
*/

package bboltstorage

import (
	bolt "go.etcd.io/bbolt"
	"os"
)

func OpenDatabase(dataFile string, dataFilePerm os.FileMode, options ...Option) (*bolt.DB, error) {

	opts := &bolt.Options{}
	for _, opt := range options {
		opt.apply(opts)
	}

	return bolt.Open(dataFile, dataFilePerm, opts)
}

