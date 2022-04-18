package commonlogic

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const RootDir = "../"

type AppConfig struct {
}

func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
		//panic(e)
	}
}

func GetCurrentPath(bPrint bool) string {
	currentPath, err := os.Getwd()
	CheckError(err)

	if bPrint {
		fmt.Printf("Current Path: %s\n", currentPath)
	}

	return currentPath
}

func FindFiles(root, ext string) []string {
	var files []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			files = append(files, filepath.Base(s))
		}
		return nil
	})
	return files
}

func FilterFilesInList(files *[]string, filter string) []string {
	var sFilteredFileList []string

	for _, entry := range *files {
		if strings.Contains(entry, filter) {
			sFilteredFileList = append(sFilteredFileList, entry)
		}
	}

	/*fmt.Println("The Following should be unique items?")
	for _, item := range sFilteredFileList {
		fmt.Println(item)
	}*/

	return sFilteredFileList
}

func RemoveDuplicateValues(stringSlice []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// InCorrectDirectory / Checks that the executable in running from the Supports folder.
func InCorrectDirectory() bool {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	fmt.Printf("directory contains: %t\n", strings.Contains(exPath, "Support"))
	return true
}

// IsinRootDirectory / Check that the *.una file in the current directory.
func IsinRootDirectory() bool {

	return false
}

func ChangeToRootDirectory() {

	fmt.Println("called ChangeToRootDirectory")
	os.Chdir(RootDir)
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current Directory: %s", currentDirectory)

	// validate by checking the presence of UNA files.
	files, err := ioutil.ReadDir(currentDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
