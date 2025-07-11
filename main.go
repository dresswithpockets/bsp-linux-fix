package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// bsplinuxfix expects vpkeditcli v4.4 or higher to be installed and on path
	// bsplinuxfix input-directory output-directory
	if len(os.Args) != 3 {
		fmt.Printf("Expected 3 args, got %d\n", len(os.Args))
		os.Exit(1)
		return
	}

	inputDirectory := os.Args[1]
	outputDirectory := os.Args[2]

	cleanOutputErr := os.RemoveAll(outputDirectory)
	if cleanOutputErr != nil && !os.IsNotExist(cleanOutputErr) {
		panic(cleanOutputErr)
	}

	mkdirErr := os.Mkdir(outputDirectory, 0777)
	if mkdirErr != nil && !os.IsExist(mkdirErr) {
		panic(mkdirErr)
	}

	fixFile, createErr := os.CreateTemp("", "bsplinuxfix_file")
	if createErr != nil {
		panic(createErr)
	}

	defer os.Remove(fixFile.Name())

	// the fix file needs some stuff in it or vpkeditcli wont add it to the bsp
	if _, err := fixFile.WriteString("some stuff"); err != nil {
		panic(err)
	}

	intermediateDir, intermediateDirErr := os.MkdirTemp("", "bsplinuxfix_intermediate")
	if intermediateDirErr != nil {
		panic(intermediateDirErr)
	}

	defer os.RemoveAll(intermediateDir)

	fixFileName := fixFile.Name()
	vpkFixFilePath := fmt.Sprintf("%s.txt", filepath.Base(fixFileName))

	items, readDirErr := os.ReadDir(inputDirectory)
	if readDirErr != nil {
		panic(readDirErr)
	}

	for _, item := range items {
		// if any ModeType bits are set on the iterated item's type, then we know its not a normal file
		if (item.Type() & fs.ModeType) != 0 {
			continue
		}

		if !strings.HasSuffix(item.Name(), ".bsp") {
			continue
		}

		fullItemPath := filepath.Join(inputDirectory, item.Name())

		fmt.Printf("Step 1: %s\n", fullItemPath)

		command := exec.Command("vpkeditcli", fullItemPath, "-o", intermediateDir, "--add-file", fixFileName, vpkFixFilePath)
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			panic(err)
		}
	}

	items, readDirErr = os.ReadDir(intermediateDir)
	if readDirErr != nil {
		panic(readDirErr)
	}

	for _, item := range items {
		fullItemPath := filepath.Join(intermediateDir, item.Name())

		fmt.Printf("Step 2: %s\n", fullItemPath)

		command := exec.Command("vpkeditcli", fullItemPath, "-o", outputDirectory, "--remove-file", vpkFixFilePath)
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			panic(err)
		}
	}
}
