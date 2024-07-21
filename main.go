package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

const createName string = "N"
const createPath string = "P"
const defaultPath string = "./"
const mainText string = "package main \n\nfunc main() {}"

func main() {
	run()
}

func run() {
	help := flag.String("H", "", "Display commands")
	projectName := flag.String(createName, "", "Create new Go project")
	projectPath := flag.String(createPath, defaultPath, "Path to new project")
	flag.Parse()
	if *help == "" {
		flag.PrintDefaults()
	}
	if err := createFolder(*projectName, *projectPath); err != nil {
		log.Fatal(err)
	}
	if err := createProject(*projectName); err != nil {
		// log.Fatal(err)
		panic(err)
	}

}

func createProject(name string) error {
	createCmd := exec.Command("go", "mod", "init", name)
	if err := createCmd.Run(); err != nil {
		return err
	}
	var mainFile *os.File
	if file, err := os.Create("main.go"); err != nil {
		return err
	} else {
		mainFile = file
	}
	if _, err := mainFile.Write([]byte(mainText)); err != nil {
		log.Panic(err)
		log.Println(mainFile.Name())
		return err
	}
	log.Print("Project ", name, " is created.")
	return nil
}

func createFolder(name string, path string) error {
	switch path {
	case defaultPath:
		err := os.Mkdir(name, 0777)
		if err != nil {
			return err
		}
		if err := os.Chdir(path + name); err != nil {
			return err
		}
	default:
		if err := os.Chdir(path); err != nil {
			return err
		}
		if err := os.Mkdir(name, 0777); err != nil {
			return err
		}
		if err := os.Chdir(name); err != nil {
			return err
		}
	}

	return nil
}
