package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"

	"fknsrs.biz/p/xml/c14n"
)

var (
	flag_withComments = flag.Bool("c", false, "use WithComments option")
)

func main() {
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := c14n.Canonicalise(xml.NewDecoder(os.Stdin), os.Stdout, *flag_withComments); err != nil {
		log.Fatal(err)
	}
}
