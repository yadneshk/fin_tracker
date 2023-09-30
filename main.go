package main

import (
	"flag"
	"fmt"
	"github.com/yadneshk/fin_tracker/read_reports"
)

func main() {
	file_path := flag.String("p", "", "Path to balance sheet")
	file_type := flag.String("t", "", "Type of the file (pdf,csv)")
	flag.Parse()
	fmt.Println("filepath : ", *file_path)
	fmt.Println("filetype : ", *file_type)

	switch *file_type {
	case "pdf":
		read_pdf.ReadPdf(*file_path)
	default:
		fmt.Println("invalid file type")
	}

}