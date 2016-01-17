/*
Hash - Guillermo Estrada

Simple utility to obtain the MD5 and/or SHA-1
of a file from the command line.
*/

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	md5f := flag.Bool("md5", false, "-md5 calculate md5 hash of file")
	sha1f := flag.Bool("sha1", false, "-sha1 calculate sha1 hash of file")
	flag.Parse()

	if !*md5f && !*sha1f {
		fmt.Println("err: No hash specified. Use -md5 or -sha1 or both.")
		os.Exit(1)
	}

	infile, inerr := os.Open(flag.Arg(0))
	if inerr == nil {
		if *md5f {
			md5h := md5.New()
			io.Copy(md5h, infile)
			fmt.Printf("%x  %s\n", md5h.Sum(nil), flag.Arg(0))
		}
		if *sha1f {
			sha1h := sha1.New()
			io.Copy(sha1h, infile)
			fmt.Printf("%x  %s\n", sha1h.Sum(nil), flag.Arg(0))
		}

	} else {
		fmt.Println(inerr)
		os.Exit(1)
	}
}
