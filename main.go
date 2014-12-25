package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func help() {
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("river usage: river %s %s\n", cyan("[command]"), yellow("[arguments ...]"))
	fmt.Printf("    river %s %s           Installs streem package with same name\n", cyan("install"), yellow("author name"))
	fmt.Printf("    river %s                       Installs streem package that you are in\n", cyan("install"))
	fmt.Printf("    river %s %s                   Removes streem package that you name\n", cyan("remove"), yellow("name"))
	fmt.Printf("    river %s %s                    Interactive prompt to setup your project\n", cyan("setup"), yellow("name"))
	fmt.Printf("    river %s                           Runs the code specified in the main_file option in the strm.json file\n", cyan("run"))
	fmt.Printf("    river %s                           Creates an executable of the code specified in the main_file option in the strm.json file\n", cyan("bin"))
	fmt.Printf("    river %s                       Version\n", cyan("version"))
}

func read(p string) (error, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(p)
	dat, err := reader.ReadString('\n')
	return err, dat
}

func read_license(license string) (lic []byte, valed bool) {
	switch license {
	case "GPL":
		lic = []byte(`
		one line to give the program's name and a brief description.
		Copyright (C) 2014 John Doe

		This program is free software: you can redistribute it and/or modify
		it under the terms of the GNU General Public License as published by
		the Free Software Foundation, either version 3 of the License, or
		(at your option) any later version.

		This program is distributed in the hope that it will be useful,
		but WITHOUT ANY WARRANTY; without even the implied warranty of
		MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
		GNU General Public License for more details.

		You should have received a copy of the GNU General Public License
		along with this program.  If not, see <http://www.gnu.org/licenses/>.
		`)
		break
	case "MIT":
		lic = []byte(`
		one line to give the program's name and a brief description
		Copyright (C) 2014 John Doe

		Permission is hereby granted, free of charge, to any person obtaining
		a copy of this software and associated documentation files (the "Software"),
		to deal in the Software without restriction, including without limitation
		the rights to use, copy, modify, merge, publish, distribute, sublicense,
		and/or sell copies of the Software, and to permit persons to whom the
		Software is furnished to do so, subject to the following conditions:

		The above copyright notice and this permission notice shall be included
		in all copies or substantial portions of the Software.

		THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
		EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
		OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
		IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
		DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
		TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
		OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
		`)
		break
	case "AGPL3":
		lic = []byte(`
		one line to give the program's name and a brief description.
		Copyright (C) 2014 John Doe

		This program is free software: you can redistribute it and/or modify
		it under the terms of the GNU Affero General Public License as
		published by the Free Software Foundation, either version 3 of the
		License, or (at your option) any later version.

		This program is distributed in the hope that it will be useful,
		but WITHOUT ANY WARRANTY; without even the implied warranty of
		MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
		GNU Affero General Public License for more details.

		You should have received a copy of the GNU Affero General Public License
		along with this program.  If not, see <http://www.gnu.org/licenses/>.
		`)
		break
	case "GPL2":
		lic = []byte(`
		One line to give the program's name and a brief description.
		Copyright (C) 2014 John Doe

		This program is free software; you can redistribute it and/or modify
		it under the terms of the GNU General Public License as published by
		the Free Software Foundation; either version 2 of the License, or
		(at your option) any later version.

		This program is distributed in the hope that it will be useful,
		but WITHOUT ANY WARRANTY; without even the implied warranty of
		MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
		GNU General Public License for more details.

		You should have received a copy of the GNU General Public License
		along with this program; if not, see <http://www.gnu.org/licenses/>.
		`)
		break
	case "GPL3":
		lic = []byte(`
		one line to give the program's name and a brief description.
		Copyright (C) 2014 John Doe

		This program is free software: you can redistribute it and/or modify
		it under the terms of the GNU General Public License as published by
		the Free Software Foundation, either version 3 of the License, or
		(at your option) any later version.

		This program is distributed in the hope that it will be useful,
		but WITHOUT ANY WARRANTY; without even the implied warranty of
		MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
		GNU General Public License for more details.

		You should have received a copy of the GNU General Public License
		along with this program.  If not, see <http://www.gnu.org/licenses/>.


		`)
	case "APACHE":
		lic = []byte(`
		one line to give the program's name and a brief description
		Copyright 2014 John Doe

		Licensed under the Apache License, Version 2.0 (the "License");
		you may not use this file except in compliance with the License.
		You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

		Unless required by applicable law or agreed to in writing, software
		distributed under the License is distributed on an "AS IS" BASIS,
		WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
		See the License for the specific language governing permissions and
		limitations under the License.
		`)
		break
	default:
		color.Red("Unknown licence version.")
		return lic, false
		break
	}
	return lic, true
}

func main() {
	const VERSION = "v0.03"
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
				url := "http://github.com/" + args[1] + "/" + args[2]
				cyan := color.New(color.FgCyan).SprintFunc()

				fmt.Println("Downloading package from " + cyan(url) + "....")

				_, err1 := exec.LookPath("git")
				if err1 != nil {
					log.Fatal(err1.Error())
				}

				cmd := exec.Command("git", "clone", url, dir)
				err2 := cmd.Run()

				if err2 != nil {
					color.Red(err2.Error())
					color.Green("BTW: You probably want to remove this package!")
					log.Fatal("Cannot clone from github!")
				}

				b, err3 := ioutil.ReadFile(dir + "/strm.json")

				if err3 != nil {
					color.Red(err3.Error())
					color.Green("BTW: You probably want to remove this package!")
					log.Fatal("Cannot read config file.")
				}
				var m map[string]interface{}

				err4 := json.Unmarshal(b, &m)
				if err4 != nil {
					color.Red(err4.Error())
					color.Green("BTW: You probably want to remove this package!")
					log.Fatal("Cannot parse config file.")
				}
				fmt.Println(m["main_file"])

				color.Green("Done!")
			} else {
				color.Red("ERROR!")
				log.Fatal(err)
			}
		} else {
			color.Red("Not enough arguments supplied!")
			os.Exit(1)
		}
		break
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
		break
	case "version":
		fmt.Println(VERSION)
		break
	case "run":
		color.Red("Sorry, the streem language isn't even implemented yet.")
		break
	case "bin":
		color.Red("Sorry, the streem language isn't even implemented yet.")
		break
	case "setup":
		if len(args) != 2 {
			color.Red("Wrong number of arguments!")
			os.Exit(1)
		}
		name := args[1]
		os.MkdirAll(name+"/src/", 0777)

		nf, _ := os.Create(name + "/src/main.strm")
		nf.Write([]byte(`"Hello, Project!" | STDOUT`))
		defer nf.Close()

		_, license := read("Enter project license type (GPL, AGPL3, GPL2, GPL3, APACHE, MIT): ")
		lic, valed := read_license(strings.TrimSpace(license))
		for !valed {
			_, license = read("Enter project license type (GPL, AGPL3, GPL2, GPL3, APACHE, MIT): ")
			lic, valed = read_license(strings.TrimSpace(license))
		}
		f, err5 := os.Create(name + "/LICENSE.txt")
		f.Write(lic)
		if err5 != nil {
			color.Red("Could not open license file for writing!")
			os.Exit(1)
		}
		defer f.Close()

		_, main_file := read("Enter location of main file in your project (src/main.strm): ")
		cf, _ := os.Create(name + "/strm.json")
		if main_file == "\n" {
			main_file = "src/main.strm"
		}
		cf.Write([]byte("{\n" + "  \"main_file\": \"" + strings.TrimSpace(main_file) + "\",\n  \"license\": \"" + strings.TrimSpace(license) + "\"\n}\n"))
		defer cf.Close()

		break
	case "help":
		help()
		break
	default:
		color.Red("Unknown command!\n")
		help()
		break
	}
}
