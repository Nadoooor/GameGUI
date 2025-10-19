package main

import (
	"BlueGUI/JSON"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("android.gamegui")
	w := a.NewWindow("GameGUI")
	var namegame string
	var bannername string
	var gamescon *fyne.Container = container.NewGridWrap(fyne.NewSize(198, 150))
	var load []JSON.His = JSON.Load()
	var gamescroll fyne.CanvasObject = container.NewScroll(gamescon)
	back, err := fyne.LoadResourceFromPath("GameCenter.png")
	if err != nil {
		println(err.Error())
	}
	Backg := canvas.NewImageFromResource(back)
	Backg.FillMode = canvas.ImageFillContain
	Backg.Resize(fyne.NewSize(1000, 600))
	Backg.Move(fyne.NewPos(-8, 0))

	back2, err := fyne.LoadResourceFromPath("AddNew.png")
	if err != nil {
		println(err.Error())
	}
	NewBack := canvas.NewImageFromResource(back2)
	NewBack.FillMode = canvas.ImageFillContain
	NewBack.Resize(fyne.NewSize(1000, 600))
	NewBack.Move(fyne.NewPos(-8, 0))

	Name := widget.NewEntry()
	Name.PlaceHolder = "Game Name"

	EXE := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		namegame = reader.URI().Path()

	}, w)
	Banner := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		bannername = reader.URI().Path()

	}, w)
	exeB := widget.NewButton("Browse file...", func() {
		EXE.Show()
	})
	bannerB := widget.NewButton("Browse file...", func() {
		Banner.Show()
	})
	Submit := widget.NewButton("Submit", func() {
		gamescon.Add(container.NewVBox(canvas.NewImageFromFile(bannername), widget.NewButton(Name.Text, func() {
			exec.Command(namegame).Start()
		})))

		newrow := JSON.His{
			Name:   Name.Text,
			Exe:    namegame,
			Banner: bannername,
		}
		load = append(load, newrow)
		JSON.Save(load)
	})
	for _, Each := range load {
		nnname := Each.Name
		Baanner := Each.Banner
		EEXXEE := Each.Exe
		gamescon.Add(container.NewVBox(canvas.NewImageFromFile(Baanner), widget.NewButton(nnname, func() {
			exec.Command(EEXXEE).Start()
		})))
	}

	gamescroll.Resize(fyne.NewSize(800, 400))
	gamescroll.Move(fyne.NewPos(70, 180))
	Games := container.NewWithoutLayout(Backg, gamescroll)

	AddNew := container.NewWithoutLayout(NewBack, exeB, bannerB, Submit, Name)

	Name.Resize(fyne.NewSize(785, 45))
	Name.Move(fyne.NewPos(100, 108))
	Submit.Resize(fyne.NewSize(182, 45))
	Submit.Move(fyne.NewPos(703, 422))
	exeB.Resize(fyne.NewSize(182, 45))
	exeB.Move(fyne.NewPos(100, 265))
	bannerB.Resize(fyne.NewSize(182, 45))
	bannerB.Move(fyne.NewPos(100, 422))
	Tabs := container.NewAppTabs(
		container.NewTabItem("Games", Games),
		container.NewTabItem("AddNew", AddNew),
	)
	w.SetContent(Tabs)
	w.SetFixedSize(false)
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()

}
