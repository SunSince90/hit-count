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

import "sync"

var (
	managersList map[string]CounterManager = map[string]CounterManager{}
)

// CounterManager is in charge of handling different counters.
type CounterManager interface {
	Add(string, Counter) CounterManager
	Remove(string) CounterManager
	ResetAllExcept(...string) CounterManager
	Get(string) Counter
}

type mgr struct {
	counters map[string]Counter
	lock     sync.Mutex
}

// Manager returns a CounterManager or creates a new one if it doesn't exist.
func Manager(key string) CounterManager {
	m, exists := managersList[key]
	if !exists {
		managersList[key] = &mgr{
			counters: map[string]Counter{},
		}
		m = managersList[key]
	}

	return m
}

func (m *mgr) Add(key string, count Counter) CounterManager {
	m.lock.Lock()
	defer m.lock.Unlock()

	if len(key) == 0 || count == nil {
		return m
	}

	if _, exists := m.counters[key]; exists {
		return m
	}

	m.counters[key] = count
	return m
}

func (m *mgr) Remove(key string) CounterManager {
	m.lock.Lock()
	defer m.lock.Unlock()

	if len(key) == 0 {
		return m
	}

	if _, exists := m.counters[key]; !exists {
		return m
	}

	delete(m.counters, key)
	return m
}

func (m *mgr) ResetAllExcept(except ...string) CounterManager {
	m.lock.Lock()
	defer m.lock.Unlock()

	for key, counter := range m.counters {
		ignore := false

		for _, exp := range except {
			if key == exp {
				ignore = true
			}
		}
		if ignore {
			continue
		}

		counter.Reset()
	}

	return m
}

func (m *mgr) Get(key string) Counter {
	m.lock.Lock()
	defer m.lock.Unlock()

	counter, exists := m.counters[key]
	if !exists {
		return nil
	}

	return counter
}
