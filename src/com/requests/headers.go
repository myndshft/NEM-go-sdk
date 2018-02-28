// Copyright 2018 Myndshft Technologies, Inc.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package requests

import (
	"strconv"
)

// URLEncoded is a url encoded header
var URLEncoded = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}

// JSON creates an application/json header
func JSON(data []byte) map[string]string {
	header := map[string]string{"Content-Type": "application/json",
		"Content-Length": strconv.Itoa(len(data))}
	return header
}