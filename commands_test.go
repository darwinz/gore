package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAction_Doc(t *testing.T) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	s, err := NewSession(stdout, stderr)
	defer s.Clear()
	require.NoError(t, err)

	err = s.Eval(":import encoding/json")
	require.NoError(t, err)
	err = s.Eval(":i fmt")
	require.NoError(t, err)

	test := func() {
		err = s.Eval(":doc fmt")
		require.NoError(t, err)

		err = s.Eval(":doc fmt.Print")
		require.NoError(t, err)

		err = s.Eval(":d json.NewEncoder(nil).Encode")
		require.NoError(t, err)
	}

	test()

	// test :doc works after some code

	s.Eval("a := 1")
	s.Eval("fmt.Print()")

	test()

	assert.Contains(t, stdout.String(), "package fmt")
	assert.Contains(t, stdout.String(), "func Printf")
	assert.Equal(t, "", stderr.String())
}

func TestAction_Import(t *testing.T) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	s, err := NewSession(stdout, stderr)
	defer s.Clear()
	require.NoError(t, err)

	err = s.Eval(":import encoding/json fmt")
	require.NoError(t, err)

	err = s.Eval("fmt.Print")
	require.NoError(t, err)

	err = s.Eval("json.Encoder{}")
	require.NoError(t, err)

	assert.Contains(t, stdout.String(), "(func(...interface {}) (int, error))")
	assert.Contains(t, stdout.String(), "json.Encoder")
	assert.Equal(t, "", stderr.String())
}

func TestAction_Clear(t *testing.T) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	s, err := NewSession(stdout, stderr)
	defer s.Clear()
	require.NoError(t, err)

	codes := []string{
		`x := 10`,
		`x`,
		`:clear`,
		`x := "foo"`,
		`x`,
		`:clear`,
		`x`,
	}

	for _, code := range codes {
		_ = s.Eval(code)
	}

	assert.Equal(t, `10
10
"foo"
"foo"
`, stdout.String())
	assert.Equal(t, "undefined: x\n", stderr.String())
}

func TestAction_Help(t *testing.T) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	s, err := NewSession(stdout, stderr)
	defer s.Clear()
	require.NoError(t, err)

	err = s.Eval(": :  :   help  ")
	require.NoError(t, err)

	assert.Contains(t, stdout.String(), ":import <package>")
	assert.Contains(t, stdout.String(), ":write [<file>]")
	assert.Contains(t, stdout.String(), "show this help")
	assert.Contains(t, stdout.String(), "quit the session")
	assert.Equal(t, "", stderr.String())

	err = s.Eval(":h")
	require.NoError(t, err)
}

func TestAction_Quit(t *testing.T) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	s, err := NewSession(stdout, stderr)
	defer s.Clear()
	require.NoError(t, err)

	err = s.Eval(" :\t: quit")
	require.Equal(t, ErrQuit, err)

	assert.Equal(t, "", stdout.String())
	assert.Equal(t, "", stderr.String())

	err = s.Eval(":q")
	require.Equal(t, ErrQuit, err)
}

func TestAction_CommandNotFound(t *testing.T) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	s, err := NewSession(stdout, stderr)
	defer s.Clear()
	require.NoError(t, err)

	err = s.Eval(":::")
	require.NoError(t, err)

	err = s.Eval(":foo")
	require.Error(t, err)

	err = s.Eval(":ii")
	require.Error(t, err)

	err = s.Eval(":docc")
	require.Error(t, err)

	err = s.Eval(":help]")
	require.Error(t, err)

	assert.Equal(t, "", stdout.String())
	assert.Equal(t, `command not found: foo
command not found: ii
command not found: docc
command not found: help]
`, stderr.String())
}
