package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func help() {
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("river usage: river %s %s\n", cyan("[command]"), yellow("[arguments ...]"))
	fmt.Printf("    river %s %s           Installs streem package with same name\n", cyan("install"), yellow("author name"))
	fmt.Printf("    river %s                       Installs streem package that you are in\n", cyan("install"))
	fmt.Printf("    river %s %s                   Removes streem package that you name\n", cyan("remove"), yellow("name"))
	fmt.Printf("    river %s                         Interactive prompt to setup your project\n", cyan("setup"))
	fmt.Printf("    river %s                           Runs your project in the src/main.strm\n", cyan("run"))
	fmt.Printf("    river %s                           Creates an executable of your project in the src/main.strm\n", cyan("bin"))
	fmt.Printf("    river %s                       Version\n", cyan("version"))
}

func main() {
	const VERSION = "v0.01"
	if len(os.Args) <= 1 {
		color.Red("Not enough arguments supplied!")
		os.Exit(1)
	}
	args := os.Args[1:]

	switch args[0] {
	case "install":
		fmt.Println("Creating package directory..")
		if len(args) == 3 {
			dir := "/usr/local/bin/river-pkgs/" + args[2]
			err := os.MkdirAll(dir, 0777)
			if err == nil {
				fmt.Println("Downloading package from servers....")

				binary, err1 := exec.LookPath("git")
				if err1 != nil {
					log.Fatal(err1)
				}

				args := []string{"git", "clone", "http://github.com/" + args[1] + "/" + args[2], dir}
				err2 := syscall.Exec(binary, args, os.Environ())
				if err2 != nil {
					log.Fatal(err2)
				}
				color.Green("Done!")
			} else {
				color.Red("ERROR!")
				log.Fatal(err)
			}
		} else {
			color.Red("Not enough arguments supplied!")
			os.Exit(1)
		}
	case "remove":
		fmt.Print("Are you sure? ")
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

		if len(args) == 2 {
			input := make([]byte, 1)
			os.Stdin.Read(input)
			if string(input[:]) == "Y" || string(input[:]) == "y" {
				os.RemoveAll("/usr/local/bin/river-pkgs/" + args[1])
				color.Green("Done!")
				os.Exit(0)
			} else {
				color.Yellow("OK, exiting now...")
			}
		}
	case "version":
		fmt.Println(VERSION)
	case "run":
	case "bin":
	case "help":
		help()
	default:
		color.Red("Unknown command!\n\n")
		help()
	}
}
