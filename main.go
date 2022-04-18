package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/syddo/yaas_dev/selftestdatahandling"
	"github.com/syddo/yaas_dev/selftestdspsend"
)

func main() {

	parser := argparse.NewParser("yaas", "YASS - Yet Another Support Script. Takes care of the small things so you can focus on the big things.")

	// generate command
	// subcommand for generation of support files
	// ie: selftest data handling, hexfiles dsp send
	generateCmd := parser.NewCommand("generate", "Generate support files")
	generateWhat := generateCmd.String("f", "autogenfile", &argparse.Options{Required: true, Help: "create the Test Program Autogen Files"})
	generateWhatInputFilesDir := generateCmd.String("i", "input-dir", &argparse.Options{Required: false, Help: "Specify Input Files Location."})
	//generateSHFDspSendCfg := generateCmd.String("c", "config", &argparse.Options{Required: false, Help: "Specify DSP Send Block Config Name."})

	tarprogCmd := parser.NewCommand("tarprog", "Create a tar ball of this test program")
	tarOpt := tarprogCmd.String("o", "output", &argparse.Options{Required: true, Help: "tar with or without sources"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if generateCmd.Happened() {

		//generate_what := generateCmd.String("f", "autogenfile", &argparse.Options{Required: true, Help: "create the Test Program Autogen Files"})
		fmt.Println("generate command was given")
		fmt.Println(*generateWhat)

		if *generateWhat == "selftestdatahandling" {
			selftestdatahandling.GenerateSDHFiles(*generateWhatInputFilesDir)
			fmt.Println("Success...")
		} else if *generateWhat == "selftestdspsend" {
			// selftestdspsend.GenerateSHFAutogen(*generateWhatInputFilesDir, *generateSHFDspSendCfg)
			selftestdspsend.GenerateSHFAutogen(*generateWhatInputFilesDir, "default") //default config
			fmt.Println("Success...")
		} else {
			fmt.Printf("option %s is not supported.", *generateWhat)
		}

	} else if tarprogCmd.Happened() {
		fmt.Println("tar of test program")
		fmt.Println(*tarOpt)

	} else {
		//log.Fatal("something weird happened")
	}
}
