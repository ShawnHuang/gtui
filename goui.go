//usr/bin/env go run "$0" "$@"; exit "$?"
package main

import (
	"github.com/rivo/tview"
)

func main() {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}

//
//package main
//
//import (
//	"io"
//	"os"
//
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//	textView := tview.NewTextView().
//		SetDynamicColors(true).
//		SetChangedFunc(func() {
//			app.Draw()
//		})
//	textView.SetBorder(true).SetTitle("Stdin")
//	go func() {
//		w := tview.ANSIWriter(textView)
//		if _, err := io.Copy(w, os.Stdin); err != nil {
//			panic(err)
//		}
//	}()
//	if err := app.SetRoot(textView, true).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import "github.com/rivo/tview"
//
//func main() {
//	app := tview.NewApplication()
//	button := tview.NewButton("Hit Enter to close").SetSelectedFunc(func() {
//		app.Stop()
//	})
//	button.SetBorder(true).SetRect(0, 0, 22, 3)
//	if err := app.SetRoot(button, false).SetFocus(button).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import "github.com/rivo/tview"
//
//func main() {
//	app := tview.NewApplication()
//	checkbox := tview.NewCheckbox().SetLabel("Hit Enter to check box: ")
//	if err := app.SetRoot(checkbox, true).SetFocus(checkbox).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import "github.com/rivo/tview"
//
//func main() {
//	app := tview.NewApplication()
//	dropdown := tview.NewDropDown().
//		SetLabel("Select an option (hit Enter): ").
//		SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil)
//	if err := app.SetRoot(dropdown, true).SetFocus(dropdown).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//	flex := tview.NewFlex().
//		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
//		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
//			AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
//			AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
//			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
//		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
//	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//	form := tview.NewForm().
//		AddDropDown("Title", []string{"Mr.", "Ms.", "Mrs.", "Dr.", "Prof."}, 0, nil).
//		AddInputField("First name", "", 20, nil, nil).
//		AddInputField("Last name", "", 20, nil, nil).
//		AddCheckbox("Age 18+", false, nil).
//		AddPasswordField("Password", "", 10, '*', nil).
//		AddButton("Save", nil).
//		AddButton("Quit", func() {
//			app.Stop()
//		})
//	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
//	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"github.com/gdamore/tcell"
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//	frame := tview.NewFrame(tview.NewBox().SetBackgroundColor(tcell.ColorBlue)).
//		SetBorders(2, 2, 2, 2, 4, 4).
//		AddText("Header left", true, tview.AlignLeft, tcell.ColorWhite).
//		AddText("Header middle", true, tview.AlignCenter, tcell.ColorWhite).
//		AddText("Header right", true, tview.AlignRight, tcell.ColorWhite).
//		AddText("Header second middle", true, tview.AlignCenter, tcell.ColorRed).
//		AddText("Footer middle", false, tview.AlignCenter, tcell.ColorGreen).
//		AddText("Footer second middle", false, tview.AlignCenter, tcell.ColorGreen)
//	if err := app.SetRoot(frame, true).SetFocus(frame).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"github.com/rivo/tview"
//)
//
//func main() {
//	newPrimitive := func(text string) tview.Primitive {
//		return tview.NewTextView().
//			SetTextAlign(tview.AlignCenter).
//			SetText(text)
//	}
//	menu := newPrimitive("Menu")
//	main := newPrimitive("Main content")
//	sideBar := newPrimitive("Side Bar")
//
//	grid := tview.NewGrid().
//		SetRows(3, 0, 3).
//		SetColumns(30, 0, 30).
//		SetBorders(true).
//		AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
//		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)
//
//	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
//	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
//		AddItem(main, 1, 0, 1, 3, 0, 0, false).
//		AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)
//
//	// Layout for screens wider than 100 cells.
//	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
//		AddItem(main, 1, 1, 1, 1, 0, 100, false).
//		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)
//
//	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"github.com/gdamore/tcell"
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//	inputField := tview.NewInputField().
//		SetLabel("Enter a number: ").
//		SetFieldWidth(10).
//		SetAcceptanceFunc(tview.InputFieldInteger).
//		SetDoneFunc(func(key tcell.Key) {
//			app.Stop()
//		})
//	if err := app.SetRoot(inputField, true).SetFocus(inputField).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//	list := tview.NewList().
//		AddItem("List item 1", "Some explanatory text", 'a', nil).
//		AddItem("List item 2", "Some explanatory text", 'b', nil).
//		AddItem("List item 3", "Some explanatory text", 'c', nil).
//		AddItem("List item 4", "Some explanatory text", 'd', nil).
//		AddItem("Quit", "Press to exit", 'q', func() {
//			app.Stop()
//		})
//	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
//		panic(err)
//	}
//}

//package main
//
//import (
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//	modal := tview.NewModal().
//		SetText("Do you want to quit the application?").
//		AddButtons([]string{"Quit", "Cancel"}).
//		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
//			if buttonLabel == "Quit" {
//				app.Stop()
//			}
//		})
//	if err := app.SetRoot(modal, false).SetFocus(modal).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"strings"
//
//	"github.com/gdamore/tcell"
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//
//	// Returns a new primitive which puts the provided primitive in the center and
//	// sets its size to the given width and height.
//	//modal := func(p tview.Primitive, width, height int) tview.Primitive {
//	//	return tview.NewFlex().
//	//		AddItem(nil, 0, 1, false).
//	//		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
//	//			AddItem(nil, 0, 1, false).
//	//			AddItem(p, height, 1, false).
//	//			AddItem(nil, 0, 1, false), width, 1, false).
//	//		AddItem(nil, 0, 1, false)
//	//}
//	modal := func(p tview.Primitive, width, height int) tview.Primitive {
//		return tview.NewGrid().
//			SetColumns(0, width, 0).
//			SetRows(0, height, 0).
//			AddItem(p, 1, 1, 1, 1, 0, 0, true)
//	}
//	background := tview.NewTextView().
//		SetTextColor(tcell.ColorBlue).
//		SetText(strings.Repeat("background ", 2000))
//
//	box := tview.NewBox().
//		SetBorder(true).
//		SetTitle("Centered Box")
//
//	pages := tview.NewPages().
//		AddPage("background", background, true, true).
//		AddPage("modal", modal(box, 40, 10), true, true)
//
//	if err := app.SetRoot(pages, true).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"fmt"
//
//	"github.com/rivo/tview"
//)
//
//const pageCount = 5
//
//func main() {
//	app := tview.NewApplication()
//	pages := tview.NewPages()
//	for page := 0; page < pageCount; page++ {
//		func(page int) {
//			pages.AddPage(fmt.Sprintf("page-%d", page),
//				tview.NewModal().
//					SetText(fmt.Sprintf("This is page %d. Choose where to go next.", page+1)).
//					AddButtons([]string{"Next", "Quit"}).
//					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
//						if buttonIndex == 0 {
//							pages.SwitchToPage(fmt.Sprintf("page-%d", (page+1)%pageCount))
//						} else {
//							app.Stop()
//						}
//					}),
//				false,
//				page == 0)
//		}(page)
//	}
//	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"strings"
//
//	"github.com/gdamore/tcell"
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//	table := tview.NewTable().
//		SetBorders(true)
//	lorem := strings.Split("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.", " ")
//	cols, rows := 10, 40
//	word := 0
//	for r := 0; r < rows; r++ {
//		for c := 0; c < cols; c++ {
//			color := tcell.ColorWhite
//			if c < 1 || r < 1 {
//				color = tcell.ColorYellow
//			}
//			table.SetCell(r, c,
//				tview.NewTableCell(lorem[word]).
//					SetTextColor(color).
//					SetAlign(tview.AlignCenter))
//			word = (word + 1) % len(lorem)
//		}
//	}
//	table.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
//		if key == tcell.KeyEscape {
//			app.Stop()
//		}
//		if key == tcell.KeyEnter {
//			table.SetSelectable(true, true)
//		}
//	}).SetSelectedFunc(func(row int, column int) {
//		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
//		table.SetSelectable(false, false)
//	})
//	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"fmt"
//	"strconv"
//	"strings"
//	"time"
//
//	"github.com/gdamore/tcell"
//	"github.com/rivo/tview"
//)
//
//const corporate = `Leverage agile frameworks to provide a robust synopsis for high level overviews. Iterative approaches to corporate strategy foster collaborative thinking to further the overall value proposition. Organically grow the holistic world view of disruptive innovation via workplace diversity and empowerment.
//
//Bring to the table win-win survival strategies to ensure proactive domination. At the end of the day, going forward, a new normal that has evolved from generation X is on the runway heading towards a streamlined cloud solution. User generated content in real-time will have multiple touchpoints for offshoring.
//
//Capitalize on low hanging fruit to identify a ballpark value added activity to beta test. Override the digital divide with additional clickthroughs from DevOps. Nanotechnology immersion along the information highway will close the loop on focusing solely on the bottom line.
//
//[yellow]Press Enter, then Tab/Backtab for word selections`
//
//func main() {
//	app := tview.NewApplication()
//	textView := tview.NewTextView().
//		SetDynamicColors(true).
//		SetRegions(true).
//		SetChangedFunc(func() {
//			app.Draw()
//		})
//	numSelections := 0
//	go func() {
//		for _, word := range strings.Split(corporate, " ") {
//			if word == "the" {
//				word = "[red]the[white]"
//			}
//			if word == "to" {
//				word = fmt.Sprintf(`["%d"]to[""]`, numSelections)
//				numSelections++
//			}
//			fmt.Fprintf(textView, "%s ", word)
//			time.Sleep(200 * time.Millisecond)
//		}
//	}()
//	textView.SetDoneFunc(func(key tcell.Key) {
//		currentSelection := textView.GetHighlights()
//		if key == tcell.KeyEnter {
//			if len(currentSelection) > 0 {
//				textView.Highlight()
//			} else {
//				textView.Highlight("0").ScrollToHighlight()
//			}
//		} else if len(currentSelection) > 0 {
//			index, _ := strconv.Atoi(currentSelection[0])
//			if key == tcell.KeyTab {
//				index = (index + 1) % numSelections
//			} else if key == tcell.KeyBacktab {
//				index = (index - 1 + numSelections) % numSelections
//			} else {
//				return
//			}
//			textView.Highlight(strconv.Itoa(index)).ScrollToHighlight()
//		}
//	})
//	textView.SetBorder(true)
//	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
//		panic(err)
//	}
//}

// Demo code for a timer based update
//package main
//
//import (
//	"fmt"
//	"time"
//
//	"github.com/rivo/tview"
//)
//
//const refreshInterval = 500 * time.Millisecond
//
//var (
//	view *tview.Modal
//	app  *tview.Application
//)
//
//func currentTimeString() string {
//	t := time.Now()
//	return fmt.Sprintf(t.Format("Current time is 15:04:05"))
//}
//
//func updateTime() {
//	for {
//		time.Sleep(refreshInterval)
//		app.QueueUpdateDraw(func() {
//			view.SetText(currentTimeString())
//		})
//	}
//}
//
//func main() {
//	app = tview.NewApplication()
//	view = tview.NewModal().
//		SetText(currentTimeString()).
//		AddButtons([]string{"Quit", "Cancel"}).
//		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
//			if buttonLabel == "Quit" {
//				app.Stop()
//			}
//		})
//
//	go updateTime()
//	if err := app.SetRoot(view, false).Run(); err != nil {
//		panic(err)
//	}
//}
// Demo code for a timer based update
//package main
//
//import (
//	"time"
//
//	"github.com/gdamore/tcell"
//	"github.com/rivo/tview"
//)
//
//const refreshInterval = 500 * time.Millisecond
//
//var (
//	view *tview.Box
//	app  *tview.Application
//)
//
//func drawTime(screen tcell.Screen, x int, y int, width int, height int) (int, int, int, int) {
//	timeStr := time.Now().Format("Current time is 15:04:05")
//	tview.Print(screen, timeStr, x, height/2, width, tview.AlignCenter, tcell.ColorLime)
//	return 0, 0, 0, 0
//}
//
//func refresh() {
//	tick := time.NewTicker(refreshInterval)
//	for {
//		select {
//		case <-tick.C:
//			app.Draw()
//		}
//	}
//}
//
//func main() {
//	app = tview.NewApplication()
//	view = tview.NewBox().SetDrawFunc(drawTime)
//
//	go refresh()
//	if err := app.SetRoot(view, true).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"io/ioutil"
//	"path/filepath"
//
//	"github.com/gdamore/tcell"
//	"github.com/rivo/tview"
//)
//
//// Show a navigable tree view of the current directory.
//func main() {
//	rootDir := "."
//	root := tview.NewTreeNode(rootDir).
//		SetColor(tcell.ColorRed)
//	tree := tview.NewTreeView().
//		SetRoot(root).
//		SetCurrentNode(root)
//
//	// A helper function which adds the files and directories of the given path
//	// to the given target node.
//	add := func(target *tview.TreeNode, path string) {
//		files, err := ioutil.ReadDir(path)
//		if err != nil {
//			panic(err)
//		}
//		for _, file := range files {
//			node := tview.NewTreeNode(file.Name()).
//				SetReference(filepath.Join(path, file.Name())).
//				SetSelectable(file.IsDir())
//			if file.IsDir() {
//				node.SetColor(tcell.ColorGreen)
//			}
//			target.AddChild(node)
//		}
//	}
//
//	// Add the current directory to the root node.
//	add(root, rootDir)
//
//	// If a directory was selected, open it.
//	tree.SetSelectedFunc(func(node *tview.TreeNode) {
//		reference := node.GetReference()
//		if reference == nil {
//			return // Selecting the root node does nothing.
//		}
//		children := node.GetChildren()
//		if len(children) == 0 {
//			// Load and show files in this directory.
//			path := reference.(string)
//			add(node, path)
//		} else {
//			// Collapse if visible, expand if collapsed.
//			node.SetExpanded(!node.IsExpanded())
//		}
//	})
//
//	if err := tview.NewApplication().SetRoot(tree, true).Run(); err != nil {
//		panic(err)
//	}
//}
//package main
//
//import (
//	"fmt"
//
//	"github.com/rivo/tview"
//)
//
//func main() {
//	app := tview.NewApplication()
//	pages := tview.NewPages()
//
//	form := tview.NewForm()
//	form.AddDropDown("称谓", []string{"先生", "女士", "博士", "老师", "师傅"}, 0, nil).
//		AddInputField("姓名", "", 20, nil, nil).
//		AddCheckbox("年龄 18+", false, nil).
//		AddPasswordField("密码", "", 10, '*', nil).
//		AddButton("保存", func() {
//			_, title := form.GetFormItem(0).(*tview.DropDown).GetCurrentOption()
//			userName := form.GetFormItem(1).(*tview.InputField).GetText()
//
//			alert(pages, "alert-dialog", fmt.Sprintf("保存成功，%s %s！", userName, title))
//		}).
//		AddButton("退出", func() {
//			app.Stop()
//		})
//	form.SetBorder(true).SetTitle("输入一些内容").SetTitleAlign(tview.AlignLeft)
//	pages.AddPage("base", form, true, true)
//
//	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
//		panic(err)
//	}
//}
//
//// alert shows a confirmation dialog.
//func alert(pages *tview.Pages, id string, message string) *tview.Pages {
//	return pages.AddPage(
//		id,
//		tview.NewModal().
//			SetText(message).
//			AddButtons([]string{"确定"}).
//			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
//				pages.HidePage(id).RemovePage(id)
//			}),
//		false,
//		true,
//	)
//}

// vim: ft=go
