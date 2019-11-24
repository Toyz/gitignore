package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/toyz/gitignore/github"
)

func main() {
	path := filepath.Base(os.Args[0])
	path = strings.Split(path, ".")[0]

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage:")
		fmt.Printf("%s --types [types | license] -- List supported gitignore types\n", path)
		fmt.Printf("%s --license [file] -- License file to download\n", path)
		fmt.Printf("%s [file] -- Gitignore file to download\n", path)
		return // TODO
	}

	switch args[0] {
	case "--types", "-t", "--type":
		if len(args) < 2 {
			args[1] = "types"
		}

		switch args[1] {
		case "types":
			handleTypeList()
		case "license":
			handleLicenseList()
		default:
			handleTypeList()
		}
	case "--license", "--lic", "-l":
		if len(args) < 2 {
			fmt.Printf("Invalid usage uages: %s --license [key]", path)
			return
		}
		lic, err := github.GetLicense(strings.ToLower(args[1]))
		if err != nil {
			fmt.Println(err)
			return
		}
		name := "[fullname]"
		user, err := user.Current()
		if err == nil {
			name = user.Name
		}

		lic.Body = strings.ReplaceAll(lic.Body, "[year]", strconv.Itoa(time.Now().Year()))
		lic.Body = strings.ReplaceAll(lic.Body, "[fullname]", name)
		ioutil.WriteFile("LICENSE", []byte(lic.Body), 0755)
		fmt.Printf("LICENSE has been created... Your good to go!")

	default:
		err := github.Download(fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/master/%s.gitignore", strings.Title(args[0])), ".gitignore")
		if err != nil {
			panic(err)
		}
		fmt.Printf(".gitignore has been created... Your good to go!")
	}
}

func handleTypeList() {
	files, err := github.List()

	if err != nil {
		fmt.Println(err)
		return
	}

	types := make([]string, len(files))
	for i, _ := range types {
		types[i] = files[i].FileName
	}

	fmt.Printf("Types: %s", strings.Join(types, ", "))
}

func handleLicenseList() {
	files, err := github.ListLicense()

	if err != nil {
		fmt.Println(err)
		return
	}

	f := make([]string, len(files))
	for i, _ := range files {
		f[i] = files[i].Key
	}

	fmt.Printf("License: %s", strings.Join(f, ", "))
}