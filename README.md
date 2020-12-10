# gzCoredump
coredump实时压缩工具，Linux内核产生的异常程序的coredump文件将以压缩包的形式保存。
为了确保coredump文件与异常程序的对应关系，提供了一种匹配的方法，程序生成后需要对ELF中追加.uuid的section。
另外该工具同样提供了压缩文件、解压文件等功能，支持*.gz

## build
本工程的目录结构按照一般工程的工程目录，存在src，bin，pkg目录

把程序当前目录加入到GOPATH中

```
cd bin;go build gzCoredump
```

即会在bin目录下生成可执行文件

本工程同样支持go test命令。例如

```
go test elfreader
```

"
ok      elfreader       0.001s
"

本工程同样支持go install命令，运行之后即会在pkg目录中生成对应的库文件，例如
```
go build filelist
go install filelist
ls -l pkg/linux_amd64/
```

"
-rw-rw-r-- 1 lbbxsxlz lbbxsxlz 5330 10月 22 15:44 filelist.a
"

## Usage
```	
	gzCoredump version: 1.0.0
	
	Usage:
	gzCoredump -c [-p path] [-a argument]
	gzCoredump -d -f inputfile -o outputfile
	
	Options:
		-h    this help
		-c    compress file
		-d    decompress file")

		-p    with -c, compress file path
		-a    with -c, elf's infomation

		-f    with -d, compressed coredump file
		-o    with -d, output file
```

## 加解密接口	
	使用 mycompress.GzipBestCompress 接口压缩的文件可以被7zip解压。
	使用 mycompress.ZlibBestCompress 接口压缩的文件无法被7-zip解压，可当加密使用。
	对应的解压接口亦有实现。
	
&nbsp;&nbsp;&nbsp;&nbsp;[code](https://github.com/lbbxsxlz/gzCoredump/blob/master/src/mycompress/mycompress.go)

## ELF文件接口	
	coreump文件与应用程序的一一对应关系通过解析ELF中section来实现。
	编译ELF可执行文件时使用objcopy创建特殊的section，在生成core文件时使用ELF文件接口读取应用程序中的特殊的section，并保存成文件。
&nbsp;&nbsp;&nbsp;&nbsp;[code](https://github.com/lbbxsxlz/gzCoredump/blob/master/src/elfreader/elfreader.go)

## core文件压缩[ref](https://linux.die.net/man/5/core):

Piping core dumps to a program

Since kernel 2.6.19, Linux supports an alternate syntax for the /proc/sys/kernel/core_pattern file. If the first character of this file is a pipe symbol (|), then the remainder of the line is interpreted as a program to be executed. Instead of being written to a disk file, the core dump is given as standard input to the program. Note the following points:
*
The program must be specified using an absolute pathname (or a pathname relative to the root directory, /), and must immediately follow the '|' character.

*

The process created to run the program runs as user and group root.

*

Command-line arguments can be supplied to the program (since kernel 2.6.24), delimited by white space (up to a total line length of 128 bytes).

*

The command-line arguments can include any of the % specifiers listed above. For example, to pass the PID of the process that is being dumped, specify %p in an argument.

