package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadAliases(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testrc")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	testContent := `alias a='git add'
alias g='git'
export PATH=$PATH:/usr/local/bin`

	if _, err := tmpfile.Write([]byte(testContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	aliases, err := readAliases(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"a": "git add",
		"g": "git",
	}

	if len(aliases) != len(expected) {
		t.Errorf("Got %d aliases, want %d", len(aliases), len(expected))
	}

	for k, v := range expected {
		if got, exists := aliases[k]; !exists || got != v {
			t.Errorf("aliases[%s] = %q; want %q", k, got, v)
		}
	}
}

func TestMain(t *testing.T) {
	tmpHome, err := os.MkdirTemp("", "ichimoji-test-home")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpHome)

	origHome := os.Getenv("HOME")
	defer os.Setenv("HOME", origHome)

	os.Setenv("HOME", tmpHome)

	zshrcPath := filepath.Join(tmpHome, ".zshrc")
	testContent := `
alias a='git add'
alias g='git'
`
	if err := os.WriteFile(zshrcPath, []byte(testContent), 0644); err != nil {
		t.Fatal(err)
	}

	origShell := os.Getenv("SHELL")
	defer os.Setenv("SHELL", origShell)

	os.Setenv("SHELL", "/bin/zsh")

	expected := `ichimoji alias list.
✅ a='git add'
🈳 b
🈳 c
🈳 d
🈳 e
🈳 f
✅ g='git'
🈳 h
🈳 i
🈳 j
🈳 k
🈳 l
🈳 m
🈳 n
🈳 o
🈳 p
🈳 q
🈳 r
🈳 s
🈳 t
🈳 u
🈳 v
🈳 w
🈳 x
🈳 y
🈳 z`

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	got := buf.String()

	if strings.TrimSpace(got) != strings.TrimSpace(expected) {
		t.Errorf("Unexpected output:\ngot:\n%s\nwant:\n%s", got, expected)
	}
}
