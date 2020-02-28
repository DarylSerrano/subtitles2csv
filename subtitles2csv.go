package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/martinlindhe/subtitles"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	infilePtr := flag.String("infile", "", "Path to the subtitle file to convert into csv")
	outpathPtr := flag.String("outpath", dir, "Output path")
	flag.Parse()

	if *infilePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *outpathPtr != dir {
		absPath, err := filepath.Abs(*outpathPtr)
		if err != nil {
			log.Fatal(err)
		}
		outpathPtr = &absPath
	}

	convertToCsv(*infilePtr, *outpathPtr)
}

func convertToCsv(fileIn string, outdir string) {
	extension := filepath.Ext(fileIn)
	filename := filepath.Base(fileIn)
	content, err := ioutil.ReadFile(fileIn)
	if err != nil {
		log.Fatal(err)
	}

	outFilePath := path.Join(outdir, "out.csv")
	outFile, err := os.OpenFile(outFilePath, os.O_RDWR|os.O_CREATE, 0666) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	writterOutFile := csv.NewWriter(outFile)

	contentString := subtitles.ConvertToUTF8(content)

	var res subtitles.Subtitle
	if strings.Contains(extension, "srt") {
		res, err = subtitles.NewFromSRT(contentString)
		if err != nil {
			log.Fatal(err)
		}
	} else if strings.Contains(extension, "ssa") {
		res, err = subtitles.NewFromSSA(contentString)
		if err != nil {
			log.Fatal(err)
		}
	} else if strings.Contains(extension, "vtt") {
		res, err = subtitles.NewFromVTT(contentString)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatalln("Please provide a supported extension file: ", extension)
	}

	for _, caption := range res.Captions {
		for _, line := range caption.Text {
			record := []string{filename, line, ""}
			if err := writterOutFile.Write(record); err != nil {
				log.Fatalln("error writing record to csv:", err)
			}
		}
	}

	writterOutFile.Flush()
	if err := writterOutFile.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Saved into: ", outFilePath)

}
