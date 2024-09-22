package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func RunDiscCleanup() {
	clearScreen()
	fmt.Println("Cleaning your Drive...")

	//get the temp folder path
	tempPath := filepath.Join(os.Getenv("TEMP"))

	//delete all files in the temp directory
	deleteFilesInDirectory(tempPath)

	//run windows disk cleanup on the c drive
	runCommand("cleanmgr /d C /sagerun:1")

	//clear the recycle bin
	runCommand("powershell", "-Command", "Clear-RecycleBin -Confirm:$false")

	fmt.Println("Successfully cleaned your Drive!")

	reset()
}

func UltimatePowerPlan() {
	clearScreen()

	//just apply the ultimate power plan by using its guid
	runCommand("powercfg /setactive cc78e803-7c76-48ea-aa70-e2756c919b18")

	fmt.Println("Successfully switched to the Ultimate Power Plan!")

	reset()
}

func RunUpdateAllPrograms() {
	clearScreen()

	if !IsWingetInstalled() {
		fmt.Println("Winget is not installed! Packages can't be updated.")
		reset()
		return
	}

	fmt.Println("Updating all packages...")
	if output, err := runCommand("winget", "upgrade", "--all"); err != nil {
		fmt.Println("Error updating packages:", err)
		fmt.Println("Output:", output)
	} else {
		fmt.Println("Successfully updated all packages!")
	}

	reset()
}

func OpenGodMode() {
	godModePath := "shell:::{ED7BA470-8E54-465E-825C-99712043E01C}"

	runCommand("explorer", godModePath)
}

func IsWingetInstalled() bool {
	output, err := runCommand("winget", "--version")
	if err != nil {
		return false
	}
	return strings.Contains(strings.TrimSpace(output), "v")
}

func reset() {
	time.Sleep(2 * time.Second)
	clearScreen()
	app.Stop()
	main()
}

func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
