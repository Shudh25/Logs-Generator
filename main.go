package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/signintech/gopdf"
)

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := pdf.AddTTFFont("wts11", "wts11.ttf")
	errorFind(err)

	err = pdf.SetFont("wts11", "", 14)
	errorFind(err)

	pdf.SetXY(180, 15) //move current location
	err = pdf.SetFont("wts11", "", 20)
	errorFind(err)

	pdf.Cell(nil, "Chaos Mesh Daily Report\n")
	pdf.Br(30)

	err = pdf.SetFont("wts11", "", 14)
	errorFind(err)

	//READ FILE CODE
	filePath := "final.log"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Extract data from the line and store it in a variable.
		variable := line

		//exiting the loop after the cluster status
		if strings.Contains(variable, "Warnings") {
			break
		}
		pdf.Cell(nil, variable)
		pdf.Br(20)
		// You can do further processing with the data as needed.
		// fmt.Println(variable)
	}

	// READING ERRORS AND WARNINGS
	for scanner.Scan() {
		line := scanner.Text()
		// Extract data from the line and store it in a variable.
		variable := line
		pdf.Cell(nil, variable)
		pdf.Br(20)
		// You can do further processing with the data as needed.
		// fmt.Println(variable)
	}

	type warnErrors struct {
		PD   string
		TiDB string
		TIKV string
	}

	var arr []warnErrors
	a := warnErrors{"20", "20", "20"}
	arr = append(arr, a)

	fmt.Print(arr)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	//end read file

	//Exporting PDF
	pdf.WritePdf("report.pdf")

}

// Error Function
func errorFind(err error) {
	if err != nil {
		log.Print(err.Error())
		return
	}
}