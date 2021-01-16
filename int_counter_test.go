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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntCounter(t *testing.T) {
	a := assert.New(t)

	res := NewIntCounter(0, 0)
	a.NotNil(res)
	a.Equal(&intCounter{val: 0, target: 0}, res)
}

func TestReset(t *testing.T) {
	a := assert.New(t)
	c := NewIntCounter(55, 100)

	c.Reset()

	var count int64
	c.Value(&count)

	a.Zero(count)
}

func TestIncrease(t *testing.T) {
	a := assert.New(t)
	var initial int64 = 55
	var target int64 = 100

	c := NewIntCounter(initial, target)
	c.Increase()
	var count int64
	c.Value(&count)
	a.Equal(initial+1, count)

	c = NewIntCounter(target, target)
	c.Increase()
	c.Value(&count)
	a.Zero(count)
}

func TestDecrease(t *testing.T) {
	a := assert.New(t)
	var initial int64 = 55
	var target int64 = 100

	c := NewIntCounter(initial, target)
	c.Decrease()
	var count int64
	c.Value(&count)
	a.Equal(initial-1, count)

	c = NewIntCounter(0, target)
	c.Decrease()
	c.Value(&count)
	a.Zero(count)
}

func TestHit(t *testing.T) {
	a := assert.New(t)
	var target int64 = 100

	c := NewIntCounter(target-1, target)
	c.Increase()
	var count int64
	c.Value(&count)
	a.True(c.Hit())
}

// func TestHit() bool {
// 	i.lock.Lock()
// 	defer i.lock.Unlock()

// 	return i.val == i.target
// }
