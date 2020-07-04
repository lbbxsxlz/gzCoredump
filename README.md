# gzCoredump
coredump实时压缩工具，提供压缩core文件、core文件与APP一一对应匹配，压缩文件解压等功能

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
&nbsp;&nbsp;&nbsp;&nbsp;使用 mycompress.GzipBestCompress 接口压缩的文件可以被7zip解压。
&nbsp;&nbsp;&nbsp;&nbsp;使用 mycompress.ZlibBestCompress 接口压缩的文件无法被7-zip解压，可当加密使用。
&nbsp;&nbsp;&nbsp;&nbsp;对应的解压接口亦有实现。<br>
&nbsp;&nbsp;&nbsp;&nbsp;详见[code](https://github.com/lbbxsxlz/gzCoredump/blob/master/src/mycompress/mycompress.go)

## ELF文件接口	
&nbsp;&nbsp;&nbsp;&nbsp;core文件与APP的一一对应关系通过解析ELF中section来实现。
&nbsp;&nbsp;&nbsp;&nbsp;编译ELF可执行文件时使用objcopy创建特殊的section，在生成core文件时使用ELF文件接口读取APP中的特殊的section，并保存成文件。<br>
&nbsp;&nbsp;&nbsp;&nbsp;详见[code](https://github.com/lbbxsxlz/gzCoredump/blob/master/src/elfreader/elfreader.go)

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

