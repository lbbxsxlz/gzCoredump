/*
#!/usr/bin/env gorun
# _author_ lbbxsxlz@gmail.com
*/

package mycompress

import (
	"bufio"
	"io"
	"os"
	"testing"
	"time"
)

var file1 = "test1.gz"
var file2 = "test2level.gz"
var file3 = "test3zlib.gz"
var file4 = "test4zlib.gz"

func TestMyCompress(t *testing.T) {
	var err error
	rdfile, err := os.OpenFile("/home/lbbxsxlz/BigBang", os.O_RDONLY, 0666)
	if err != nil {
		t.Error("OpenFile fail")
	}

	defer rdfile.Close()
	reader := bufio.NewReader(rdfile)

	err = GzipCompress(file1, reader)
	if err != nil {
		t.Error("gzipCompress fail")
	}

	rdfile.Seek(0, io.SeekStart)
	reader2:= bufio.NewReader(rdfile)
	err = GzipBestCompress(file2, reader2)
	if err != nil {
		t.Error("gzipCompress fail")
	}

	rdfile.Seek(0, io.SeekStart)
	reader3:= bufio.NewReader(rdfile)
	err = ZlibCompress(file3, reader3)
	if err != nil {
		t.Error("zlibCompress fail")
	}

	rdfile.Seek(0, io.SeekStart)
	reader4:= bufio.NewReader(rdfile)
	err = ZlibBestCompress(file4, reader4)
	if err != nil {
		t.Error("zlibCompress fail")
	}

	time.Sleep(3)
	err= GzReader("BigBang-gz", file2)
	if err != nil {
		t.Error("GzReader fail")
	}
	time.Sleep(1)
	err = ZlibReader("BigBang-zlib", file4)
	if err != nil {
		t.Error("GzReader fail")
	}
}
