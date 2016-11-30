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

package ctxlint_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/sectioneight/ctxlint"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testData = "testdata"

func TestIntegrationExamples(t *testing.T) {
	assert.NoError(t, filepath.Walk(testData, func(path string, info os.FileInfo, err error) error {
		assert.NoError(t, err, "Unexpected error walking testdata tree")

		if info.IsDir() {
			_, file := filepath.Split(path)
			if file == testData {
				// root of the folder, continue
				return nil
			}
			t.Run(file, func(t *testing.T) {
				runTest(t, path, file)
			})
		}

		return nil
	}))
}

func runTest(t testing.TB, path, testName string) {
	sourceFiles, err := filepath.Glob(filepath.Join(path, "*.go"))
	require.NoError(t, err, "Unexpected error globbing test directory")
	sources := make(map[string][]byte, len(sourceFiles))
	for _, source := range sourceFiles {
		sourceBytes, err := ioutil.ReadFile(source)
		require.NoError(t, err, "Unable to read file")
		sources[source] = sourceBytes
	}

	l := Linter{}

	problems, err := l.LintFiles(sources)
	assert.NoError(t, err, "Unexpected error parsing files")
	assert.Empty(t, problems)
	// TODO add test case for err return from LintFiles (e.g. package mismatch)
}
