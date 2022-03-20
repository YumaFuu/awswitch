package main

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/exp/slices"
	"gopkg.in/ini.v1"
)

const (
	awsDir    = ".aws"
	awswiFile = "awswi"
	credFile  = "credentials"
)

var (
	awswiPath string
	credPath  string
)

func init() {
	var err error

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	awswiPath = filepath.Join(home, awsDir, awswiFile)
	credPath = filepath.Join(home, awsDir, credFile)
}

func main() {
	crs, err := ini.Load(credPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ps := crs.SectionStrings()

	var t string

	args := os.Args
	if len(args) < 2 {
		// todo
		fmt.Println("must specify target")
		return
	} else {
		t = args[1]
	}

	if !slices.Contains(ps, t) {
		fmt.Println(fmt.Sprintf("profile %s not found", t))
		return
	}

	f, err := os.Create(awswiPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c := fmt.Sprintf("export AWS_PROFILE=%s", t)

	_, err = f.Write([]byte(c))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Current profile is '%s' \n", t)
}
