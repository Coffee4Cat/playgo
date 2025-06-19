package user_interface

import (
	"github.com/rivo/tview"
	"fmt"
)


func InitializeUI() {
	fmt.Println("playgo [version 1.0]")
	box := tview.NewBox().SetBorder(true).SetTitle("PlayGo")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}



