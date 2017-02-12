Renamer and indexer
=====
This program renames the files in a directory to a
supplied string followed by a number, conserving the
extension (characters after the last dot in the file name).
It reads the file names from Stdin, and writes the old
file names associated with new file names separated
by commas, each in a new line.

For example, a directory with files `["a.txt",
"b.tar.gz"]`, after `ls|indrn coco` is `["coco0.txt",
"coco1.gz"]`. The output is:

```
a.txt,coco0.txt
b.tar.gz,coco1.gz
```

Installation
====
`go get github.com/lamg/indrn`