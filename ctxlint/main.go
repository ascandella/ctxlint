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

// ctxlint lints Go source files looking for context propagation
// Most of this code is adapted directly from
// https://github.com/golang/lint/blob/master/golint/golint.go
// Which is released under a BSD licence:
// https://developers.google.com/open-source/licenses/bsd
package main

import (
	"flag"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sectioneight/ctxlint"
)

func main() {
	RunLint()
}

func RunLint() {
	// TODO(ai) flags for custom context types
	flag.Parse()
	switch flag.NArg() {
	case 0:
		lintDir(".")
	case 1:
		// TODO
	default:
		lintFiles(flag.Args()...)
	}
}

func lintDir(dir string) {
	pkg, err := build.ImportDir(dir, 0)
	lintImportedPackage(pkg, err)
}

func lintFiles(filenames ...string) {
	files := make(map[string][]byte)
	for _, filename := range filenames {
		src, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		files[filename] = src
	}

	l := new(ctxlint.Linter)
	ps, err := l.LintFiles(files)
	if err != nil {
		exitWithError(err, 1)
		return
	}

	for _, p := range ps {
		fmt.Printf("%v: %s\n", p.Position, p.Text)
	}
}

func lintImportedPackage(pkg *build.Package, err error) {
	if err != nil {
		if _, nogo := err.(*build.NoGoError); nogo {
			// Ignore errors about "no go source"
			return
		}
		exitWithError(err, 0)
	}

	var files []string
	files = append(files, pkg.GoFiles...)
	files = append(files, pkg.CgoFiles...)
	files = append(files, pkg.TestGoFiles...)
	if pkg.Dir != "." {
		for i, f := range files {
			files[i] = filepath.Join(pkg.Dir, f)
		}
	}

	lintFiles(files...)
}

func exitWithError(msg interface{}, code int) {
	fmt.Fprintln(os.Stderr, msg)
	Exiter(code)
}

// silly hacks to allow testing exit codes easily in the test suite.

type exitFn func(int)

var Exiter = os.Exit
