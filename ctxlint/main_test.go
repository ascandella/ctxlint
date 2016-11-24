package main_test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/sectioneight/ctxlint/ctxlint"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLintCurrentDir(t *testing.T) {
	verifyExitCode(t, 0, RunLint)
}

func TestNoGoCode(t *testing.T) {
	tmp, err := ioutil.TempDir("", "ctxlint")
	require.NoError(t, err)

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()

	os.Args = []string{os.Args[0], ".", tmp}
	verifyExitCode(t, 0, RunLint)
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
