package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

func main() {
	layout := "15:04 02.01.2006"
	myApp := app.New()
	myWindow := myApp.NewWindow("Date and Time Condition for ICM")
	now := time.Now().Local()
	nowdate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	startbox := widget.NewEntry()
	start := nowdate.Format(layout)
	startbox.SetText(start)
	endbox := widget.NewEntry()
	end := nowdate.Format(layout)
	endbox.SetText(end)
	output := widget.NewMultiLineEntry()
	go func() {
		for {
			datetime, err := time.Parse(layout, startbox.Text)
			if err != nil {
				fmt.Println(err)
			}
			datetime2, err := time.Parse(layout, endbox.Text)
			if err != nil {
				fmt.Println(err)
			}
			string1 := fmt.Sprint("(date()=date(", datetime.Year(), ",", int(datetime.Month()), ",", datetime.Day(), ")&&time()>=time(", datetime.Hour(), ",", datetime.Minute(), "))")
			string2 := fmt.Sprint("(date()>date(", datetime.Year(), ",", int(datetime.Month()), ",", datetime.Day(), ")&&date()<date(", datetime2.Year(), ",", int(datetime2.Month()), ",", datetime2.Day(), "))")
			string3 := fmt.Sprint("(date()=date(", datetime2.Year(), ",", int(datetime2.Month()), ",", datetime2.Day(), ")&&time()<=time(", datetime2.Hour(), ",", datetime2.Minute(), "))")
			if datetime.Format("02.01.2006") == datetime2.Format("02.01.2006") {
				output.SetText(fmt.Sprint(string1, "||\n", string3))
			} else {
				output.SetText(fmt.Sprint(string1, "||\n", string2, "||\n", string3))
			}
			time.Sleep(time.Second)
		}
	}()
	btn := widget.NewButton("Копировать", func() {
		clipboard.WriteAll(output.Text)
	})
	dateTimeBox := container.NewVBox(startbox, endbox, output, btn)
	content := container.NewVBox(dateTimeBox)
	frame := widget.NewCard("Enter Date and Time             ", "", content)
	myWindow.SetContent(frame)
	myWindow.ShowAndRun()
}
