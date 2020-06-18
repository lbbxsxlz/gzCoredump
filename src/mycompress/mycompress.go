/*
#!/usr/bin/env gorun
# _author_ lbbxsxlz@gmail.com
*/

package mycompress

import (
	"bufio"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"io"
	"log"
	"os"
)

func GzipCompress(filename string, reader *bufio.Reader)(error) {
	gzFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("create file fail \n")
		return errors.New("create file fail \n")
	}

	defer gzFile.Close()
	gzWrite := gzip.NewWriter(gzFile)

	io.Copy(gzWrite, reader)

	gzWrite.Flush()
	gzWrite.Close()

	return nil
}

func GzipBestCompress(filename string, reader *bufio.Reader)(error) {
	gzFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("create file fail \n")
		return errors.New("create file fail \n")
	}

	defer gzFile.Close()
	gzWrite, err := gzip.NewWriterLevel(gzFile, gzip.BestCompression)
	if err != nil {
		log.Printf("create gzip writer fail \n")
		return errors.New("create gzip writer fail \n")
	}

	io.Copy(gzWrite, reader)

	gzWrite.Flush()
	gzWrite.Close()

	return nil
}

func ZlibCompress(filename string, reader *bufio.Reader)(error) {
	zlibFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("create file fail \n")
		return errors.New("create file fail \n")
	}

	defer zlibFile.Close()

	zWrite := zlib.NewWriter(zlibFile)
	io.Copy(zWrite, reader)

	zWrite.Flush()
	zWrite.Close()

	return nil
}

func ZlibBestCompress(filename string, reader *bufio.Reader)(error) {
	zlibFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("create file fail \n")
		return errors.New("create file fail \n")
	}

	defer zlibFile.Close()

	zWrite, err := zlib.NewWriterLevel(zlibFile, zlib.BestCompression)
	if err != nil {
		log.Printf("create zlib writer fail \n")
		return errors.New("create zlib writer fail \n")
	}
	io.Copy(zWrite, reader)

	zWrite.Flush()
	zWrite.Close()

	return nil
}

func ZlibReader(file, zfile string ) (error) {
	fp, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Printf("create file fail \n")
		return errors.New("create file fail \n")
	}
	defer fp.Close()
	fw := bufio.NewWriter(fp)

	zfp, err := os.OpenFile(zfile, os.O_RDONLY, 0666)
	if err != nil {
		log.Printf("create zfile fail \n")
		return errors.New("create zfile fail \n")
	}
	defer zfp.Close()

	zReader, err := zlib.NewReader(zfp)
	if err != nil {
		log.Printf("create zlib reader fail \n")
		return errors.New("create zlib reader fail \n")
	}

	io.Copy(fw, zReader)
	fw.Flush()
	zReader.Close()
	return nil;
}

func GzReader(file, gzfile string ) (error) {
	fp, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Printf("create file %s fail \n", file)
		return errors.New("create file fail \n")
	}
	defer fp.Close()
	fw := bufio.NewWriter(fp)

	gzfp, err := os.OpenFile(gzfile, os.O_RDONLY, 0666)
	if err != nil {
		log.Printf("create zfile fail \n")
		return errors.New("create zfile fail \n")
	}
	defer gzfp.Close()

	gzReader, err := gzip.NewReader(gzfp)
	if err != nil {
		log.Printf("create zlib reader fail \n")
		return errors.New("create zlib reader fail \n")
	}

	io.Copy(fw, gzReader)
	fw.Flush()
	gzReader.Close()
	return nil;
}
