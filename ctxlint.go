// Copyright 2016 Aiden Scandella
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

package ctxlint

import "go/token"

// A Linter handles walking the AST looking for context violations
type Linter struct{}

// A Problem represents a violation in source code
type Problem struct {
	Position   token.Position
	Text       string
	SourceLine string
}

// LintFiles runs through a map of files (represented as byte slices)
// and checks for context violations
func (l Linter) LintFiles(files map[string][]byte) ([]Problem, error) {
	// TODO
	return nil, nil
}
