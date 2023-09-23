package main

import (
	"flag"
	"fmt"
)

func main() {
	file_path := flag.String("p", "", "Path to balance sheet")
	file_type := flag.String("t", "", "Type of the file (pdf,csv)")
	flag.Parse()
	fmt.Println("filepath : ", *file_path)
	fmt.Println("filetype : ", *file_type)

}