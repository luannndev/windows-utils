package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application

func initializeApp() {
	app = tview.NewApplication()

	list := createActionList()

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(list, 0, 2, true)

	if err := app.SetRoot(flex, true).SetFocus(list).Run(); err != nil {
		fmt.Println("Error starting the application:", err)
	}
}

func createActionList() *tview.List {
	return tview.NewList().SetSecondaryTextColor(tcell.ColorBlue).
		AddItem("Disk Cleanup", "Deletes all your temp folders and runs Windows Disk cleanup", '1', func() {
			RunDiscCleanup()
		}).
		AddItem("Power Plan", "Applies the Ultimate Powerplan to gain more performance", '2', func() {
			UltimatePowerPlan()
		}).
		AddItem("Update All Programs", "Updates all programs on your computer (a reboot may be required)", '3', func() {
			RunUpdateAllPrograms()
		}).
		AddItem("God Mode", "Access all system settings in one place", '4', func() {
			OpenGodMode()
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
}
