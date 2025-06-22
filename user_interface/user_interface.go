package user_interface

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"fmt"
	"playgo/structures"
)


func InitializeUI(folders *[]structures.AudioFolder) {
	fmt.Println("DEBUG")

    tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault
    tview.Styles.ContrastBackgroundColor = tcell.ColorDefault
    tview.Styles.MoreContrastBackgroundColor = tcell.ColorDefault


	mainBox := tview.NewBox().
		SetBorder(true).
		SetTitle("PlayGo").
		SetTitleColor(tcell.ColorLightGreen).
		SetBorderStyle(tcell.StyleDefault.Foreground(tcell.ColorLightGreen).Bold(true))
	_ = mainBox

	bottomBox := tview.NewBox().
		SetBorder(true).
		SetTitle("Controller").
		SetTitleColor(tcell.ColorKhaki).
		SetBorderStyle(tcell.StyleDefault.Foreground(tcell.ColorLightGreen).Bold(true))

	folderList := makeFolderList(folders)

	pages := tview.NewPages().
    	AddPage("main", folderList, true, true)

	









	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 4, true).
		AddItem(bottomBox, 0, 1, true)

	mainFrame := tview.NewFrame(layout).
		SetBorders(0, 0, 0, 0, 1, 1).
		AddText("<<PLAYGO>>", true, tview.AlignCenter,tcell.ColorKhaki)



	if err := tview.NewApplication().SetRoot(mainFrame, true).Run(); err != nil {
		panic(err)
	}
}






// utilities
func makeFolderList(folders *[]structures.AudioFolder) *tview.List {
	list := tview.NewList()
	for _, folder := range *folders {
		list.AddItem(folder.Repr(),"",0,nil)	
	}
	list.SetBorder(true).
		SetTitle("Albums").
		SetTitleColor(tcell.ColorKhaki).
		SetBorderStyle(tcell.StyleDefault.Foreground(tcell.ColorLightGreen).Bold(true))
		
	list.SetMainTextColor(tcell.ColorIndianRed).
		SetSelectedTextColor(tcell.ColorLightGreen).
		SetSelectedBackgroundColor(tcell.ColorIndianRed)
	return list
}

