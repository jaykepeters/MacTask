package main

import (
    "fmt"
	"log"
	"os/exec"
	"time"
)

// Global Variables
var (
	status int = 1
	errorMessage string = "THERE WAS AN ERROR!"
)

func killProcess() {
	// Kill Safari Process Until Complete
	for status > 0 {
		command := exec.Command("pkill", "Safari")
		command.Run()
		time.Sleep(250 * time.Millisecond)
	}

	for status > 0 {
		command := exec.Command("pkill", "Google Chrome")
		command.Run()
		time.Sleep(250 * time.Millisecond)
	}
}

func openApplication() {
	var code = `# Bash Shell Script
	osascript <<EOF
	tell application "Safari"
		activate
	end tell
	EOF
	`

	cmd := exec.Command("/bin/bash", "-c", code)

	fmt.Println("Opening Safari")

	_ /* out */, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(errorMessage)
		log.Fatal(err)
	}
}
func ResetSafari() {
	go killProcess()
	var code = `# Bash Shell Script
	rm -Rf ~/Library/Caches/Apple\ -\ Safari\ -\ Safari\ Extensions\ Gallery; \
	rm -Rf ~/Library/Caches/Metadata/Safari; \
	rm -Rf ~/Library/Caches/com.apple.Safari; \
	rm -Rf ~/Library/Caches/com.apple.WebKit.PluginProcess; \
	rm -Rf ~/Library/Cookies/Cookies.binarycookies; \
	rm -Rf ~/Library/Preferences/Apple\ -\ Safari\ -\ Safari\ Extensions\ Gallery; \
	rm -Rf ~/Library/Preferences/com.apple.Safari.LSSharedFileList.plist; \
	rm -Rf ~/Library/Preferences/com.apple.Safari.RSS.plist; \
	rm -Rf ~/Library/Preferences/com.apple.Safari.plist; \
	rm -Rf ~/Library/Preferences/com.apple.WebFoundation.plist; \
	rm -Rf ~/Library/Preferences/com.apple.WebKit.PluginHost.plist; \
	rm -Rf ~/Library/Preferences/com.apple.WebKit.PluginProcess.plist; \
	rm -Rf ~/Library/PubSub/Database; \
	rm -Rf ~/Library/Safari/Extensions; \
	rm -Rf ~/Library/Saved\ Application\ State/com.apple.Safari.savedState; \

	## ChillTab Extras
	launchctl unload ~/Library/LaunchAgents/com.MacCheck.plist; \
	launchctl unload ~/Library/LaunchAgents/macsearch.plist; \
	rm -Rf ~/Library/LaunchAgents/com.MacCheck.plist; \
	rm -Rf ~/Library/LaunchAgents/macsearch.plist
	sleep 5`
    cmd := exec.Command("/bin/bash", "-c", code)

	fmt.Println("Resetting Safari")
	_ /* out */, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(errorMessage)
		log.Fatal(err)
	} else {
		fmt.Println("Safari Reset")
		status = 0
		openApplication()
	}
}

func main() {
	ResetSafari()
}
