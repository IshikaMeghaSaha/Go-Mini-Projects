package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	pdf := gofpdf.New("P", "mm", "A4", "") //Arguments are Orientation,measurements,size,font
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)                         //Arguments are font,style(Bold,Italic etc.),Size
	pdf.MultiCell(190, 5, string(file), "0", "0", false) //Copies contents of files to new file. Arguments are width,height,contents of file,border,alignment,fill
	pdf.OutputFileAndClose("Converted.pdf")              //New PDF created
	fmt.Println("PDF Created")
}
