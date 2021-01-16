// Copyright © 2021 Elis Lulja
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

package hitcount

import (
	"reflect"
	"sync"
)

type intCounter struct {
	val    int64
	target int64

	lock sync.Mutex
}

// NewIntCounter returns a counter working with integers.
func NewIntCounter(initial, target int64) Counter {
	return &intCounter{
		val:    initial,
		target: target,
	}
}

func (i *intCounter) Reset() Counter {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.val = 0
	return i
}

func (i *intCounter) Increase() Counter {
	i.lock.Lock()
	defer i.lock.Unlock()

	if i.val == i.target {
		i.val = 0
		return i
	}

	i.val++
	return i
}

func (i *intCounter) Decrease() Counter {
	i.lock.Lock()
	defer i.lock.Unlock()

	if i.val == 0 {
		return i
	}

	i.val--
	return i
}

func (i *intCounter) Value(out interface{}) Counter {
	i.lock.Lock()
	defer i.lock.Unlock()

	v := reflect.ValueOf(out)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	v.Set(reflect.ValueOf(i.val))
	return i
}

func (i *intCounter) Hit() bool {
	i.lock.Lock()
	defer i.lock.Unlock()

	return i.val == i.target
}
