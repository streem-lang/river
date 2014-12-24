package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

func help() {
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("river usage: river %s %s\n", cyan("[command]"), yellow("[arguments ...]"))
	fmt.Printf("    river %s %s           Installs streem package with same name\n", cyan("install"), yellow("name"))
	fmt.Printf("    river %s                Installs streem package that you are in\n", cyan("install"))
	fmt.Printf("    river %s %s            Removes streem package that you name\n", cyan("remove"), yellow("name"))
	fmt.Printf("    river %s                  Interactive prompt to setup your project\n", cyan("setup"))
	fmt.Printf("    river %s                    Runs your project in the src/main.strm\n", cyan("run"))
	fmt.Printf("    river %s                    Creates an executaple of your project in the src/main.strm\n", cyan("bin"))
	fmt.Printf("    river %s                Version\n", cyan("version"))
}
func main() {
	VERSION := "v0.01"
	if len(os.Args) <= 1 {
		color.Red("Not enough arguments supplied!")
		log.Fatal("\n")
	}
	args := os.Args[1:]

	switch args[0] {
	case "install":
		fmt.Println("Creating package directory..")
		if len(args) == 2 {
			err := os.MkdirAll("~/.river-pkgs/"+args[1], 0777)
			if err == nil {
				fmt.Println("Downloading package from servers....")
				color.Green("Done!")
			} else {
				color.Red("ERROR!")
				log.Fatal(err)
			}
		} else {
			color.Red("Not enough arguments supplied!")
			log.Fatal("\n")
		}
	case "remove":
		color.Red("Are you sure? ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		if input == "Y" || input == "y" {
			os.RemoveAll("~/.river-pkgs/" + args[1])
			color.Green("Done!")
		} else {
			color.Yellow("OK, exiting now...")
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
