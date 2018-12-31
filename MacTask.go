// PACKAGE DEFINITION
package main

// IMPORTS
import (
	"fmt"
	"log"
	"strings"
	"os"
	"os/exec"
	"strconv"
	"flag"
	//"time"
	"prockill"
)

// Get System Information
func productName() string {
	Cmd := exec.Command("sw_vers", "-productName")
	Out, err := Cmd.Output()
    if err != nil {panic(err)}
	output := strings.TrimSpace(string(Out))
	return output
}
func productVersion() string {
	Cmd := exec.Command("sw_vers", "-productVersion")
	Out, err := Cmd.Output()
    if err != nil {panic(err)}
	output := strings.TrimSpace(string(Out))
	return output
}
func buildVersion() string {
	Cmd := exec.Command("sw_vers", "-buildVersion")
	Out, err := Cmd.Output()
    if err != nil {panic(err)}
	output := strings.TrimSpace(string(Out))
	return output
}

// Get User Info
func loggedInUser() string {
	// https://gobyexample.com/spawning-processes
	Cmd := exec.Command("stat", "-f%Su", "/dev/console")
	Out, err := Cmd.Output()
    if err != nil {panic(err)}
	output := strings.TrimSpace(string(Out))
	return output
}
func CurrentUser() string {return os.Getenv("USER")}
func UserHomeDir() string {return os.Getenv("HOME")}
func userInfo() {
	liu := loggedInUser()
	cu := CurrentUser()
	home := UserHomeDir()

	fmt.Println("Logged in User: " + liu)
	fmt.Println("Current User: " + cu)
	fmt.Println("Home Directory: " + home)
}

// Empty Trash
func emptyTrash() {
	homeDir := UserHomeDir()
	trashDir := homeDir + "/" + ".Trash" + "/"

	fmt.Println("Emptying Trash at: " + trashDir)
	os.RemoveAll(trashDir)
	os.MkdirAll(trashDir, 0777)
	fmt.Println("Trash Emptied!")
}

// Reset Launchpad
func resetLP() {
	cmd1 := exec.Command("defaults", "write", "com.apple.dock", "ResetLaunchPad", "-bool", "true")
	cmd2 := exec.Command("pkill", "Dock")
	cmd1.Start()
	cmd2.run()
}

// Reset Dock
func resetDock() {
	cmd1 := exec.Command("defaults", "delete", "com.apple.dock")
	cmd2 := exec.Command("pkill", "Dock")
	cmd1.Start()
	cmd2.Start()
}

// Elevated Rights Check
func checkSudo() {
    cmd := exec.Command("id", "-u")
    output, err := cmd.Output()
    if err != nil {log.Fatal(err)}
    i, err := strconv.Atoi(string(output[:len(output)-1]))
    if err != nil {log.Fatal(err)}
    if i == 0 {
    	log.Println("Awesome! You are now running this program with root permissions!")
    } else {
		log.Fatal("This program must be run as root! (sudo)")
    }
}

// MANAGE
func pctl() {
	var processes[11]string
	processes[0] = "Game Center"
	processes[1] = "Messages"
	processes[2] = "FaceTime"
	processes[3] = "iTunes"
	processes[4] = "Terminal"
	processes[5] = "Adobe Application Manager"
	processes[6] = "PDApp"
	processes[7] = "steam_osx"
	processes[8] = "App Store"
	processes[9] = "Self Service"
	processes[10] = "Console"

	for _, process := range processes {
        Cmd := exec.Command("pkill", process, "&&", "sleep", "1")
		Cmd.Run()
    }
}
func manage() {
	for {
		//time.Sleep(time.Second)
		pctl()
	}
}
// THE MAIN FUNCTION
func main() {
	// CLI INTERFACE
	boolPtr := flag.Bool("resetLaunchpad", false, "Reset the Launchpad")
	boolPtr1 := flag.Bool("emptyTrash", false, "Empty the Trash")
	boolPtr2 := flag.Bool("resetDock", false, "Reset the Dock")
	boolPtr3 := flag.Bool("userInfo", false, "Display information about the current user")
	boolPtr4 := flag.Bool("manage", false, "Perform management tasks")
	flag.Parse()

	// Conditional Statements
	if *boolPtr == true {resetLP()}
	if *boolPtr1 == true {emptyTrash()}
	if *boolPtr2 == true {resetDock()}
	if *boolPtr3 == true {userInfo()}
	if *boolPtr4 == true {manage()}
}
