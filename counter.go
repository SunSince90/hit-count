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

// Counter takes care managing counting operations
type Counter interface {
	// Increase the counter
	Increase() Counter
	// Decrease the counter
	Decrease() Counter
	// Reset the counter
	Reset() Counter
	// Value puts the current value on the provided variable
	Value(interface{}) Counter
	// Hits returns true if the target was hit
	Hit() bool
}
