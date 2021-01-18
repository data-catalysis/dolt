// Copyright 2019 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This file incorporates work covered by the following copyright and
// permission notice:
//
// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package functions

import (
	"sync"
	"sync/atomic"
)

// All runs all functions in |fs| in parallel, and returns when all functions have returned.
func All(fs ...func() error) error {
	var res atomic.Value
	wg := &sync.WaitGroup{}
	wg.Add(len(fs))
	for _, f := range fs {
		capf := f
		go func() {
			defer wg.Done()
			if err := capf(); err != nil {
				res.Store(err)
			}
		}()
	}
	wg.Wait()
	if err := res.Load(); err != nil {
		return err.(error)
	}
	return nil
}
