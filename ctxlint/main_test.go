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

package main_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/sectioneight/ctxlint/ctxlint"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLintCurrentDir(t *testing.T) {
	verifyExitCode(t, 0, RunLint)
}

func TestNoGoCode(t *testing.T) {
	withTempDir(t, func(tmp string) {
		defer withArgs(".", tmp)()

		verifyExitCode(t, 0, RunLint)
	})
}

func TestPathExpansion(t *testing.T) {
	withTempDir(t, func(tmp string) {
		defer withArgs("./...", tmp)()

		verifyExitCode(t, 0, RunLint)
	})
}

func TestPackageExpansion(t *testing.T) {
	defer withArgs("github.com/sectioneight/ctxlint/...")()
	verifyExitCode(t, 0, RunLint)
}

func TestLint_WithValidFile(t *testing.T) {
	withTempDir(t, func(tmp string) {
		defer withTempFile(t, tmp, "foo.go", "package main")()
		defer withArgs(filepath.Join(tmp, "foo.go"))()
		verifyExitCode(t, 0, RunLint)
	})
}

func withTempFile(t testing.TB, dir, name, contents string) func() {
	fpath := filepath.Join(dir, name)
	require.NoError(t, ioutil.WriteFile(fpath, []byte(contents), os.ModePerm), "Error writing file")

	return func() {
		os.Remove(fpath)
	}
}

func withTempDir(t testing.TB, fn func(dir string)) {
	tmp, err := ioutil.TempDir("", "ctxlint")
	require.NoError(t, err, "Tempdir creation failed")

	defer os.RemoveAll(tmp)

	fn(tmp)
}

func withArgs(args ...string) func() {
	oldArgs := os.Args
	os.Args = append([]string{os.Args[0]}, args...)

	return func() {
		os.Args = oldArgs
	}
}

func verifyExitCode(t testing.TB, expected int, fn func()) {
	oldExiter := Exiter
	defer func() {
		Exiter = oldExiter
	}()
	foundCode := 0
	Exiter = func(code int) {
		foundCode = code
	}

	fn()
	assert.Equal(t, foundCode, expected, "Exit code mismatch")
}
