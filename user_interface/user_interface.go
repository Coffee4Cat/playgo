package user_interface

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"playgo/structures"
)

var cmd chan structures.PlayerCommand
var volumeText *tview.TextView
var volume float64 = 0.5

func InitializeUI(folders *[]structures.AudioFolder, command chan structures.PlayerCommand) {
	cmd = command

    tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault
    tview.Styles.ContrastBackgroundColor = tcell.ColorDefault
    tview.Styles.MoreContrastBackgroundColor = tcell.ColorDefault
	
	pages := tview.NewPages()
	fileList := initFileList(pages)
	folderList := makeFolderList(folders, fileList, pages)
	
	pages.AddPage("main", folderList, true, true).
		AddPage("tracks", fileList, true, false)
		
	volumeText = tview.NewTextView().
		SetText(makeVolumeBar(volume)).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 4, true).
		AddItem(volumeText, 0, 1, false)

	mainFrame := tview.NewFrame(layout).
		SetBorders(0, 0, 0, 0, 1, 1).
		AddText("<<PLAYGO>>", true, tview.AlignCenter,tcell.ColorKhaki)

	mainFrame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRight:
			setVolumeUp()
			return nil
		case tcell.KeyLeft:
			setVolumeDown()
			return nil
		}
		return event
	}) 


	if err := tview.NewApplication().SetRoot(mainFrame, true).Run(); err != nil {
		panic(err)
	}
}




// utilities

func setVolumeUp() {
	cmd <- structures.PlayerCommand{Action: structures.ActionSetLevel, Level: true}
	if volume < 0.9 {
		volume += 0.1
	}
	if volumeText != nil {
		volumeText.SetText(makeVolumeBar(volume))
	}
}
func setVolumeDown() {
	cmd <- structures.PlayerCommand{Action: structures.ActionSetLevel, Level: false}
	if volume > 0.0 {
		volume -= 0.1
	}
	if volumeText != nil {
		volumeText.SetText(makeVolumeBar(volume))
	}
}


func initFileList(pages *tview.Pages) *tview.List {
	list := tview.NewList()
	list.SetBorder(true).
		SetTitle("Tracks").
		SetTitleColor(tcell.ColorKhaki).
		SetBorderStyle(tcell.StyleDefault.Foreground(tcell.ColorLightGreen).Bold(true))
		
	list.SetMainTextColor(tcell.ColorIndianRed).
		SetSelectedTextColor(tcell.ColorLightGreen).
		SetSelectedBackgroundColor(tcell.ColorIndianRed)

	list.SetInputCapture(func (event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			pages.SwitchToPage("main")
			return nil
		}
		return event
	})
	return list
}




func makeFileList(filelist *tview.List, folder *structures.AudioFolder) {
	for _, file := range folder.AudioFiles {
		f := file
		filelist.AddItem(file.Repr(),"",0, func () {
			cmd <- structures.PlayerCommand{ Action: structures.ActionSetTrack, Track: &f}
		})	
	}

	filelist.SetBorder(true).
		SetTitle("Tracks").
		SetTitleColor(tcell.ColorKhaki).
		SetBorderStyle(tcell.StyleDefault.Foreground(tcell.ColorLightGreen).Bold(true))
		
	filelist.SetMainTextColor(tcell.ColorIndianRed).
		SetSelectedTextColor(tcell.ColorLightGreen).
		SetSelectedBackgroundColor(tcell.ColorIndianRed)
}

func makeFolderList(folders *[]structures.AudioFolder, fileList *tview.List, pages *tview.Pages) *tview.List {
	selectCallback := func (folder *structures.AudioFolder) {
		fileList.Clear()
		makeFileList(fileList, folder)
		pages.SwitchToPage("tracks")
	}
	
	list := tview.NewList()
	for _, folder := range *folders {
		f := folder
		list.AddItem(folder.Repr(),"",0, func () {selectCallback(&f)})	
	}
	list.SetBorder(true).
		SetTitle("Albums").
		SetTitleColor(tcell.ColorKhaki).
		SetBorderStyle(tcell.StyleDefault.Foreground(tcell.ColorLightGreen).Bold(true))
		
	list.SetMainTextColor(tcell.ColorIndianRed).
		SetSelectedTextColor(tcell.ColorLightGreen).
		SetSelectedBackgroundColor(tcell.ColorIndianRed)

	list.SetInputCapture(func (event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			pages.SwitchToPage("main")
			return nil
		}

		return event
	})
	return list
}


func makeVolumeBar(level float64) string {
	var maxx int = 28
	var filled int = int(level * float64(maxx))
	var bar string = "[lightgreen]VOLUME\n[[-] "
	for i := 0; i < maxx ; i++ {
		if i <= filled && i <= int(0.33 * float64(maxx)) {
			bar += "[green]#[-]"
		} else if i <= filled && i <= int(0.66 * float64(maxx)){
			bar += "[yellow]#[-]"
		}else if i <= filled && i <= int(1.0 * float64(maxx)){
			bar += "[red]#[-]"
		} else {
			bar += "[gray]0[-]"
		}
	}
	bar += "[lightgreen] ][-]"
	return bar
}