package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.NewWithID("android.gamegui")
	w := a.NewWindow("GameGUI")

	back, err := fyne.LoadResourceFromPath("GameCenter.png")
	if err != nil {
		println(err.Error())
	}
	Backg := canvas.NewImageFromResource(back)
	Backg.FillMode = canvas.ImageFillContain
	Backg.Resize(fyne.NewSize(1000, 600))
	Backg.Move(fyne.NewPos(-8, 0))
	Container := container.NewWithoutLayout(Backg)
	Tabs := container.NewAppTabs(
		container.NewTabItem("Games", Container),
		container.NewTabItem("AddNew", container.NewVBox()),
	)
	w.SetContent(Tabs)
	w.SetFixedSize(false)
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()

}
