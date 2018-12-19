/*
 * Copyright 2018 Dgraph Labs, Inc. and Contributors
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
 */

package x

import (
	"bytes"
)

// SafeMessage is a raw encoded JSON value. Almost identical to json.RawMessage
// but it double-escapes sequences that might get marshalled into non-UTF8 control sequences.
type SafeMessage []byte

// MarshalJSON returns m as the JSON encoding of m.
func (m SafeMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	// Dont change the message.
	b := bytes.Replace(m, []byte(`\x`), []byte(`\u00`), -1)
	return b, nil
}
