/*
#!/usr/bin/env gorun
# _author_ lbbxsxlz@gmail.com
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

import "mycompress"
import "elfreader"
import "filelist"

const UUID = ".uuid"
const ELFFILE = "/bin/BigBang"

const version = "1.0.0"

var (
	h bool
	c bool
	d bool

	p string
	a string
	f string
	o string
)

func usage() {
	fmt.Println("gzCoredump version: ", version)
	fmt.Println("Usage: ")
	fmt.Println("gzCoredump -c [-p path] [-a argument]")
	fmt.Println("gzCoredump -d -f inputfile -o outputfile")

	fmt.Println("Options:")
	flag.PrintDefaults()

	os.Exit(0)
}

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&c, "c", false, "compress file")
	flag.BoolVar(&d, "d", false, "decompress file")

	flag.StringVar(&p, "p", "/home", "with -c, compress file path")
	flag.StringVar(&a, "a", "", "with -c, elf's infomation")

	flag.StringVar(&f, "f", "", "with -d, compressed coredump file")
	flag.StringVar(&o, "o", "", "with -d, output file")

	flag.Usage = usage
}


func main() {
	flag.Parse()

	if h {
		flag.Usage()
	} else if !c && !d {
		flag.Usage()
	}

	if c {
		year := time.Now().Year()
		month := time.Now().Month()
		day := time.Now().Day()
		hour := time.Now().Hour()
		minute := time.Now().Minute()
		second := time.Now().Second()

		timeStr := fmt.Sprintf("%04d%02d%02d_%02d%02d%02d", year, month, day, hour, minute, second)
		var fileName string
		if a == "" {
			fileName = fmt.Sprintf("%s/Core_%s.gz", p, timeStr)
		} else {
			fileName = fmt.Sprintf("%s/Core_%s_%s.gz", p, timeStr, a)
		}

		if strings.Contains(p, "/home/Coredump") {
			err := filelist.DeleteFileAndDir("/home/Coredump")
			if err != nil {
				log.Fatal("DeleteFileAndDir fail \n")
			}
		}

		if runtime.GOARCH == "arm64" || runtime.GOARCH == "arm" {
			cmd := exec.Command("/bin/sh", "-c", "echo end > /proc/watchdog")
			_, err := cmd.Output()
			if err != nil {
				log.Fatal("exec commmand fail \n")
			}
		}

		reader := bufio.NewReader(os.Stdin)
		err := mycompress.GzipBestCompress(fileName, reader)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(2)

		if strings.Contains(fileName, "BigBang") {
			var file = strings.Replace(fileName, ".gz", ".txt", 1)
			err = elfreader.CreateUuidFile(ELFFILE, UUID, file)
			if err != nil {
				fmt.Println(err)
			}
		}

		time.Sleep(3)

		if runtime.GOARCH == "arm64" || runtime.GOARCH == "arm" {
			cmd := exec.Command("/bin/sh", "-c", "echo start 60000 > /proc/watchdog")
			_, err = cmd.Output()
			if err != nil {
				log.Fatal("exec commmand failed \n")
			}
		}

		os.Exit(0)
	} else if d {
		if f != "" && o != "" {
			err := mycompress.GzReader(o, f)
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		} else {
			flag.Usage()
		}
 	}
}
