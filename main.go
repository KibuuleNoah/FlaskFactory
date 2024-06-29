package main

import (
	"FlaskFactory/cmd/simpleflask"
	"fmt"
)

func main() {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Cyan = "\033[36m"

	var Yellow = "\033[33m"
	var Green = "\033[32m"

	fmt.Println(Yellow + `
 ________ ___       ________  ________  ___  __       
|\  _____\\  \     |\   __  \|\   ____\|\  \|\  \     
\ \  \__/\ \  \    \ \  \|\  \ \  \___|\ \  \/  /|_   
 \ \   __\\ \  \    \ \   __  \ \_____  \ \   ___  \  
  \ \  \_| \ \  \____\ \  \ \  \|____|\  \ \  \\ \  \ 
   \ \__\   \ \_______\ \__\ \__\____\_\  \ \__\\ \__\
    \|__|    \|_______|\|__|\|__|\_________\|__| \|__|
                                \|_________|          
                                                      
 ________ ________  ________ _________  ________  ________      ___    ___ 
|\  _____\\   __  \|\   ____\\___   ___\\   __  \|\   __  \    |\  \  /  /|
\ \  \__/\ \  \|\  \ \  \___\|___ \  \_\ \  \|\  \ \  \|\  \   \ \  \/  / /
 \ \   __\\ \   __  \ \  \       \ \  \ \ \  \\\  \ \   _  _\   \ \    / / 
  \ \  \_| \ \  \ \  \ \  \____   \ \  \ \ \  \\\  \ \  \\  \|   \/  /  /  
   \ \__\   \ \__\ \__\ \_______\  \ \__\ \ \_______\ \__\\ _\ __/  / /    
    \|__|    \|__|\|__|\|_______|   \|__|  \|_______|\|__|\|__|\___/ /     
                                                              \|___|/      
  ` + Reset)

	var projectName string
	var needDB bool
	var dbRes string

	fmt.Print(Green + "Enter Project Name\n ~# " + Reset + Cyan)
	fmt.Scan(&projectName)

	fmt.Print(Green + "Do you need a DataBase(sqlite3) y/n\n ~# " + Reset + Cyan)
	fmt.Scan(&dbRes)
	if dbRes == "y" || dbRes == "yes" {
		needDB = true
	}

	if len(projectName) > 0 {

		if err := simpleflask.CreateProject(projectName, needDB); err != nil {
			fmt.Println(Red, "Error: ", err, Reset)
			return
		}
	}

	fmt.Printf(Yellow+"Project %s Created"+Reset, projectName)
}
