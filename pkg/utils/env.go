/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
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

package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetEnvString(key, v string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return v
}

func GetEnvBool(key string, v bool) bool {
	if value := os.Getenv(key); value != "" {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return v
		}

		return b
	}

	return v
}

func GetEnvStringSlice(name string, v []string) []string {
	if value := os.Getenv(name); value != "" {
		return strings.Split(value, ",")
	}

	return v
}
