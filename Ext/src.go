package main

import (
	"log"
	"strconv"

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

	// pdf.SetTopMargin(10)
	// pdf.AddHeader()
	pdf.SetXY(180, 15) //move current location
	err = pdf.SetFont("wts11", "", 20)
	errorFind(err)

	pdf.Cell(nil, "Chaos Mesh Daily Report\n")
	pdf.Br(30)

	err = pdf.SetFont("wts11", "", 14)
	errorFind(err)

	reportDate := 256
	pdf.Cell(nil, "Report Date : ")
	pdf.Cell(nil, strconv.Itoa(reportDate))

	pdf.Br(40)
	pdf.SetX(440)
	pdf.Cell(nil, "NAME")
	pdf.SetX(220)
	pdf.Cell(nil, "READY")
	pdf.SetX(300)

	pdf.Cell(nil, "STATUS")
	pdf.SetX(370)
	pdf.Cell(nil, "RESTARTS")
	pdf.SetX(500)
	pdf.Cell(nil, "AGE")

	// demoooo

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
