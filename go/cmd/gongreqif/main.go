package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"strconv"

	gongreqif_models "github.com/thomaspeugeot/gongreqif/go/models"
	gongreqif_stack "github.com/thomaspeugeot/gongreqif/go/stack"
	gongreqif_static "github.com/thomaspeugeot/gongreqif/go/static"

	"github.com/thomaspeugeot/gongreqif/go/schema"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	filename = flag.String("filename", "../../../test/sample.reqif", "filename of the requif file")
)

func main() {

	log.SetPrefix("gongreqif: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongreqif_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := gongreqif_stack.NewStack(r, "gongreqif", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// Replace "path/to/file.reqif" with the actual path to your file
	data, err := readFile(*filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var reqif schema.REQ_IF
	err = xml.Unmarshal(data, &reqif)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	var _ gongreqif_models.REQIF

	stack.Stage.Commit()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err = r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
