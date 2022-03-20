package main

import (
	"flag"
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

	showList *bool
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

	showList = flag.Bool("ls", false, "show list")
	flag.Parse()
}

func main() {
	crs, err := ini.Load(credPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ps := crs.SectionStrings()

	if *showList {
		showProfiles(ps)
	} else {
		setProfile(ps)
	}
}

func showProfiles(ps []string) {
	for _, s := range ps {
		fmt.Println(s)
	}
}

func setProfile(ps []string) {
	var t string

	args := flag.Args()
	if len(args) < 1 {
		// todo fzf
		fmt.Println("must specify target")
		return
	} else {
		t = args[0]
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

	cmd := fmt.Sprintf("export AWS_PROFILE=%s", t)
	_, err = f.Write([]byte(cmd))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Current profile is '%s' \n", t)
}
