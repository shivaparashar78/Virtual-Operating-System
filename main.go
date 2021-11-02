package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myapp fyne.App = app.New();
var myWindow fyne.Window = myapp.NewWindow("Virtual OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget

var img fyne.CanvasObject;
var DeskBtn fyne.Widget
var panelContent *fyne.Container

func main(){
	img = canvas.NewImageFromFile("icon\\wallpaper.png")
	i1, _ := os.Open("icon\\Weather.png")
	r1 := bufio.NewReader(i1)
	b1, _ := ioutil.ReadAll(r1)
	btn1 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b1), func() {
		WeatherApp()
	})
	i2, _ := os.Open("icon\\calculator.png")
	r2 := bufio.NewReader(i2)
	b2, _ := ioutil.ReadAll(r2)
	btn2 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b2), func() {
		Calc()
	})
	i3, _ := os.Open("icon\\textEditor.jpg")
	r3 := bufio.NewReader(i3)
	b3, _ := ioutil.ReadAll(r3)
	btn3 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", b3), func() {
		TextEditor()
	})

	DeskBtn = widget.NewButtonWithIcon("", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(nil,nil , panelContent,  nil,img))
	})

	panelContent = container.NewVBox(container.NewGridWithRows(4, DeskBtn, btn1, btn2, btn3))

	myWindow.Resize(fyne.NewSize(1280,720))
	myWindow.CenterOnScreen()
	r, _ := fyne.LoadResourceFromPath("icon\\windowIcon.png")

	myWindow.SetIcon(r)
	myWindow.SetContent(
		container.NewBorder(nil,nil , panelContent,  nil,img),
	)
	myWindow.ShowAndRun()
}
