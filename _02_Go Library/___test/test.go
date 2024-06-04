package main

import "flag"

func main() {
	flag.Int("n", 1234, "help message for flag n")
	flag.String("s", "flag", "help message for flag s")
	flag.Usage()
}

/*
Usage of test.exe:
  -n int
        help message for flag n (default 1234)
  -s string
        help message for flag s (default "flag")
*/
