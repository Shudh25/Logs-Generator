package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {

	pdf := gopdf.GoPdf{}
	mm6ToPx := 22.68

	// Base trim-box
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4, //595.28, 841.89 = A4
		TrimBox:  gopdf.Box{Left: mm6ToPx, Top: mm6ToPx, Right: 595 - mm6ToPx, Bottom: 842 - mm6ToPx},
	})

	// Page trim-box
	opt := gopdf.PageOption{
		PageSize: gopdf.PageSizeA4, //595.28, 841.89 = A4
		TrimBox:  &gopdf.Box{Left: mm6ToPx, Top: mm6ToPx, Right: 595 - mm6ToPx, Bottom: 842 - mm6ToPx},
	}
	pdf.AddPageWithOption(opt)

	if err := pdf.AddTTFFont("wts11", "wts11.ttf"); err != nil {
		log.Print(err.Error())
		return
	}

	if err := pdf.SetFont("wts11", "", 14); err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Cell(nil, "Hi")
	pdf.WritePdf("hello2.pdf")
}
