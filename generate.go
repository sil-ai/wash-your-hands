package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/fogleman/gg"
)

func main() {

	// Open the elements file.
	f, err := os.Open("poster_elements.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Read in all of the CSV records.
	elementData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range elementData {
		if err := generate(record[0], "output/"+record[0]+".png"); err != nil {
			log.Fatal(err)
		}
	}
}

func generate(iso string, outFile string) error {

	im, err := gg.LoadPNG("assets/how_to_wash.png")
	if err != nil {
		return err
	}
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	dc := gg.NewContext(w, h)
	dc.DrawImage(im, 0, 0)

	im, err = gg.LoadPNG("poster_text_render/" + iso + ".png")
	if err != nil {
		return err
	}
	iw := im.Bounds().Size().X
	//h := im.Bounds().Size().Y
	dc.DrawImage(im, w/2-iw/2, 230)

	// Output
	dc.SavePNG(outFile)

	return nil
}
