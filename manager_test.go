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

func TestManager(t *testing.T) {
	a := assert.New(t)
	key := "key"
	newKey := "new-key"

	res := Manager(key)
	a.Equal(&mgr{counters: map[string]Counter{}}, res)
	a.Len(managersList, 1)
	m, exists := managersList[key]
	a.True(exists)
	a.Equal(m, res)

	managersList[newKey] = &mgr{
		counters: map[string]Counter{
			"test": NewIntCounter(55, 100),
		},
	}
	res = Manager(newKey)
	a.Equal(managersList[newKey], res)
	a.Len(managersList, 2)
}

func TestAdd(t *testing.T) {
	a := assert.New(t)
	key := "key"
	m := &mgr{
		counters: map[string]Counter{
			"test": NewIntCounter(55, 100),
		},
	}

	m.Add(key, nil)
	a.Len(m.counters, 1)

	counter := NewIntCounter(77, 100)
	m.Add(key, counter)
	a.Len(m.counters, 2)
	counterRes := m.Get(key)
	a.Equal(counter, counterRes)

	newCounter := NewIntCounter(88, 100)
	m.Add(key, newCounter)
	a.Len(m.counters, 2)
	counterRes = m.Get(key)
	a.Equal(counter, counterRes)
}

func TestRemove(t *testing.T) {
	a := assert.New(t)
	key := "exists"
	notExists := "not-exists"
	start := map[string]Counter{
		key: NewIntCounter(55, 100),
	}
	m := &mgr{
		counters: start,
	}

	m.Remove("")
	a.Equal(m.counters, start)

	m.Remove(notExists)
	a.Equal(m.counters, start)

	m.Remove(key)
	a.Empty(m.counters)
}

func TestManagerResetAllExcept(t *testing.T) {
	a := assert.New(t)
	exceptions := []string{"one", "two"}
	start := map[string]Counter{
		exceptions[0]: NewIntCounter(50, 100),
		exceptions[1]: NewIntCounter(60, 100),
		"stays":       NewIntCounter(70, 100),
	}

	m := &mgr{
		counters: start,
	}

	m.ResetAllExcept(exceptions...)
	a.Equal(m.counters, map[string]Counter{
		exceptions[0]: NewIntCounter(50, 100),
		exceptions[1]: NewIntCounter(60, 100),
		"stays":       NewIntCounter(0, 100),
	})
}

func TestGet(t *testing.T) {
	a := assert.New(t)
	start := map[string]Counter{
		"one": NewIntCounter(70, 100),
	}

	m := &mgr{
		counters: start,
	}

	a.NotNil(m.Get("one"))
	a.Nil(m.Get("two"))
}
