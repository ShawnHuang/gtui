//usr/bin/env go run "$0" "$@"; exit "$?"
package main

import (
	"bufio"
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"os"
)

/*
 Three ways of taking input
   1. fmt.Scanln(&input)
   2. reader.ReadString()
   3. scanner.Scan()

   Here we recommend using bufio.NewScanner
*/

func main() {
	// To create dynamic array
	arr := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			arr = append(arr, text)
		} else {
			break
		}

	}

	app := tview.NewApplication()
	list := tview.NewList().
		SetWrapAround(false)
	frame := tview.NewFrame(list)
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		_, _, _, height := list.GetRect()
		numofmove := (height - height%2) / 4
		switch key := event.Key(); key {
		case tcell.KeyPgUp:
			if (list.GetCurrentItem() - numofmove + 1) < 0 {
				list.SetCurrentItem(0)
			} else {

				list.SetCurrentItem(list.GetCurrentItem() - numofmove + 1)
			}
			return nil
		case tcell.KeyPgDn:
			list.SetCurrentItem(list.GetCurrentItem() + numofmove - 1)
			return nil
		case tcell.KeyCtrlU:
			//return tcell.NewEventKey(tcell.KeyPgUp, 0, tcell.ModNone)
			if (list.GetCurrentItem() - numofmove + 1) < 0 {
				list.SetCurrentItem(0)
			} else {

				list.SetCurrentItem(list.GetCurrentItem() - numofmove + 1)
			}
			return nil
		case tcell.KeyCtrlD:
			//return tcell.NewEventKey(tcell.KeyPgDn, 0, tcell.ModNone)
			list.SetCurrentItem(list.GetCurrentItem() + numofmove - 1)
			return nil
		case tcell.KeyRune:
			ch := event.Rune()
			switch ch {
			case 'j':
				return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
			case 'k':
				return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
			case 'q':
				app.Stop()
			}
		}
		return event
	})
	for k, v := range arr {
		//fmt.Println(k, v)
		k = k
		list.AddItem(v, "", 0, func() {
			fmt.Println(list.GetItemText(list.GetCurrentItem()))
			app.Stop()
		})
	}

	//AddItem("List item 2", "Some explanatory text", 'b', nil).
	//AddItem("List item 3", "Some explanatory text", 'c', nil).
	//AddItem("List item 4", "Some explanatory text", 'd', nil).
	//list.AddItem("Quit", "Press to exit", 'q', func() {
	//	app.Stop()
	//})
	if err := app.SetRoot(frame, true).SetFocus(frame).Run(); err != nil {
		panic(err)
	}
	// Use collected inputs
	//fmt.Println(arr)
}
