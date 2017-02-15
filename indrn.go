/*
  Simple file renamer.  Indrn renames the content of a directory and
  outputs the association of old and new file names, so the inverse
  process can be executed. */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
)

func main() {
	var e error
	var l []string
	if len(os.Args) == 2 {
		l, e = readLs(os.Stdin)
		if e == nil {
			k := genLs(l, os.Args[1])
			oR := &osRenamer{}
			e = renLs(l, k, oR)
			if e == nil {
				e = csvAssoc(os.Stdout, l, k)
			}
		}
	} else {
		fmt.Fprint(os.Stderr, "Expected base name\n")
	}
	r := 0
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		r = 1
	}
	os.Exit(r)
}

type osRenamer struct {
}

func (o *osRenamer) Rename(oldname, newname string) (e error) {
	e = os.Rename(oldname, newname)
	return
}

func readLs(r io.Reader) (l []string, e error) {
	l = make([]string, 0)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		l = append(l, sc.Text())
	}
	e = sc.Err()
	return
}

func csvAssoc(w io.Writer, l, k []string) (e error) {
	// { len(l) < len(k) }
	for i := 0; e == nil && i != len(l); i++ {
		_, e = fmt.Fprintf(w, "%s,%s\n", l[i], k[i])
	}
	return
}

func renLs(l, k []string, r Renamer) (e error) {
	// { len(l) < len(k) }
	for i := 0; e == nil && i != len(l); i++ {
		e = r.Rename(l[i], k[i])
	}
	// { (A i: 0 <= i < len(l): name.(file.i) = k.i)}
	return
}

type Renamer interface {
	Rename(oldname, newname string) error
}

func genLs(l []string, base string) (k []string) {
	k = make([]string, len(l))
	for i, n := 0, 0; i != len(l); i++ {
		k[i], n = genNN(l[i], base, n)
	}
	return
}

func genNN(old, base string, n int) (k string, r int) {
	//creates a name that doesn't exists
	//when the file is renamed
	ex := path.Ext(old)
	b, k, r := true, "", n
	// the loop terminates since there is a
	// finite number of files and k always
	// has a new file name
	for b {
		k, r = fmt.Sprintf("%s%d%s", base, n, ex), r+1
		b = existsOrError(k)
	}
	return
}

func existsOrError(s string) (b bool) {
	_, e := os.Stat(s)
	b = !os.IsNotExist(e) && e != nil
	return
}
