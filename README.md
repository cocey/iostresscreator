# iostresscreator
This application just write, read and delete random files to create I/O stress on server. Please be aware of that It could be dangerous for disks.

You can set proccess number and package size.

Application Language
===========
go version go1.10.2 linux/amd64
https://golang.org/doc/

User Guide
===========
iostress = Creating stress on disks
``` sh
usage: iostress [options]

    -a  int
            byte array length (default 16)
    -b  int
            if 0 not use buffered writer
    -h	display this help dialog
    -l  string
            level of stress on memory (default "hard")
    -r	use random bytes.
    -v	output version information and exit.
```

You can see more detail in docs folder.