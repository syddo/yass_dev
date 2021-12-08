package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
)

func main() {

	parser := argparse.NewParser("yaas", "YAAS - Yet Another Support Script. Takes care of the small things so you can focus on the big things.")

	// generate command
	// subcommand for generation of support files
	// ie: selftest data handling, hexfiles dsp send
	generateCmd := parser.NewCommand("generate", "Generate support files")
	tarprogCmd := parser.NewCommand("tarprog", "Create a tar ball of this test program")

	generate_what := generateCmd.String("f", "autogenfile", &argparse.Options{Required: true, Help: "create the Test Program Autogen Files"})

	tarOpt := tarprogCmd.String("o", "output", &argparse.Options{Required: true, Help: "tar with or without sources"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if generateCmd.Happened() {

		//generate_what := generateCmd.String("f", "autogenfile", &argparse.Options{Required: true, Help: "create the Test Program Autogen Files"})
		fmt.Println("generate command was given")
		fmt.Println(*generate_what)

		if *generate_what == "selftestdatahandling" {
			generateSelftestDataHandlingFiles()
		} else if *generate_what == "selftestdspsend" {
			generateSelftestHexFileDspSend()
		} else {
			fmt.Printf("option %s is not supported.", *generate_what)
		}

	} else if tarprogCmd.Happened() {
		fmt.Println("tar of test program")
		fmt.Println(*tarOpt)

	} else {
		log.Fatal("something weird happened")
	}
}

func generateSelftestDataHandlingFiles() {
	fmt.Println("Generating Selftest DataHandling Autogen Files")
}

func generateSelftestHexFileDspSend() {
	fmt.Println("Generating Selftest DspSend Blocks Autogen Files")
}
