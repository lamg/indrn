package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"path"
	"strings"
	"testing"
)

func TestExistsOrError(t *testing.T) {
	assert.False(t, existsOrError("home0.gz"))
}

func TestGenNN(t *testing.T) {
	k, _ := genNN("coc.tar.gz", "home", 0)
	assert.EqualValues(t, "home0.gz", k, "Not Equal names")
}

func TestGenLs(t *testing.T) {
	l := []string{"a", "b", "c", "d", "e", "f"}
	for i := 0; i != len(l); i++ {
		l[i] = path.Join(l[i], ".txt")
	}
	k := genLs(l, "home")
	assert.True(t, len(l) == len(k))
	for i := 0; i != len(k); i++ {
		s := fmt.Sprintf("home%d.txt", i)
		assert.EqualValues(t, s, k[i])
	}
}

func TestReadLs(t *testing.T) {
	l := []string{"a", "b", "c", "d oe", "e", "f"}
	s := strings.Join(l, "\n")
	b := bytes.NewBufferString(s)
	r, e := readLs(b)
	if assert.NoError(t, e, "Read Error") {
		t.Logf("%v", r)
	}
}

func TestCSVAssoc(t *testing.T) {
	l := []string{"a", "b", "c", "d", "e", "f"}
	k := genLs(l, "home")
	b := bytes.NewBufferString("")
	e := csvAssoc(b, l, k)
	assert.NoError(t, e)
	t.Logf("%s", b.String())
}

func TestRename(t *testing.T) {
	l := []string{"a", "b", "c", "d", "e", "f"}
	k := genLs(l, "home")
	m := &mockRenamer{t: t}
	e := renLs(l, k, m)
	assert.NoError(t, e)
}

type mockRenamer struct {
	t *testing.T
}

func (m *mockRenamer) Rename(old, new string) (e error) {
	m.t.Logf("old: %s, new: %s", old, new)
	return
}
