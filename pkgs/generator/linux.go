package generator

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Sunling472/gopm/pkgs/libs"
)

const (
	NewArg  = "new"
	HelpArg = "help"
	MaxArgs = 2
)

func flagInit() string {
	pathFlag := flag.String("P", "./", "Path to new project")
	flag.Parse()
	return *pathFlag
}

func ParseCommands() (string, string) {
	pathFlag := flagInit()
	var projectNameCmdArg string

	switch firstArg := flag.Arg(0); firstArg {
	case NewArg:
		if len(flag.Args()) < 2 {
			panic("Error new")
		}
		projectNameCmdArg = flag.Arg(1)
		return projectNameCmdArg, pathFlag
	case HelpArg:
		fmt.Println("Example: gopm new <project_name> -P <target_path>")
		flag.PrintDefaults()
		return "", pathFlag
	default:
		flag.PrintDefaults()
		return "", pathFlag
	}
}

func LinuxRun() {
	name, path := ParseCommands()
	if name == "" {
		return
	}
	CreateFolder(name, path)
	CreateProject(name)
}

func CreateProject(name string) error {
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
	if _, err := mainFile.Write([]byte(libs.MainText)); err != nil {
		log.Panic(err)
		log.Println(mainFile.Name())
		return err
	}
	log.Print("Project ", name, " is created.")
	return nil
}

func CreateFolder(name string, path string) error {
	switch path {
	case libs.DefaultPath:
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
