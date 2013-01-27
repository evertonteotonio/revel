package main

import (
	"fmt"
	"github.com/robfig/revel"
	"io/ioutil"
	"os"
	"path"
)

var cmdPackage = &Command{
	UsageLine: "package [import path]",
	Short:     "package a Revel application (e.g. for deployment)",
	Long: `
Package the Revel web application named by the given import path.
This allows it to be deployed and run on a machine that lacks a Go installation.

For example:

    revel package github.com/robfig/revel/samples/chat
`,
}

func init() {
	cmdPackage.Run = packageApp
}

func packageApp(args []string) {
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, cmdPackage.Long)
		return
	}

	appImportPath := args[0]
	rev.Init("", appImportPath, "")

	// Remove the archive if it already exists.
	destFile := path.Base(rev.BasePath) + ".zip"
	os.Remove(destFile)

	// Collect stuff in a temp directory.
	tmpDir, err := ioutil.TempDir("", path.Base(rev.BasePath))
	panicOnError(err, "Failed to get temp dir")

	buildApp([]string{args[0], tmpDir})

	// Create the zip file.
	zipName := mustZipDir(destFile, tmpDir)

	fmt.Println("Your archive is ready:", zipName)
}
