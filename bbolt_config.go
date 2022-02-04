/*
 *
 * Copyright 2020-present Arpabet, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package bboltstorage

import (
	"errors"
	"os"
	"time"
)

var (

	DataFilePerm  = os.FileMode(0755)

	OpenTimeout = time.Minute
	NoGrowSync = false
	MmapFlags  = 0
    InitialMmapSize = 0

	ErrDatabaseReadOnly = errors.New("readonly")
	ErrInvalidSeek = errors.New("invalid seek")

)


type BoltConfig struct {
	DataFile    string
	ReadOnly    bool
}
