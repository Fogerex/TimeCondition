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
			if datetime.Year() != datetime2.Year() || datetime.Month() != datetime2.Month() || datetime.Day() != datetime2.Day() {
				if datetime.Format("15:04") == "00:00" && datetime2.Format("15:04") == "00:00" {
					output.SetText(fmt.Sprint("(date()=date(", datetime.Year(), ",", int(datetime.Month()), ",", datetime.Day(), "))||\n(date()>date(", datetime.Year(), ",", int(datetime.Month()), ",", datetime.Day(), ")&&date()<date(", datetime2.Year(), ",", int(datetime2.Month()), ",", datetime2.Day(), "))||\n(date()=date(", datetime2.Year(), ",", int(datetime2.Month()), ",", datetime2.Day(), "))"))
				}
				if !(datetime.Format("15:04") == "00:00" && datetime2.Format("15:04") == "00:00") {
					output.SetText(fmt.Sprint("(date()=date(", datetime.Year(), ",", int(datetime.Month()), ",", datetime.Day(), ")&&time()>=time(", datetime.Hour(), ",", datetime.Minute(), "))||\n(date()>date(", datetime.Year(), ",", int(datetime.Month()), ",", datetime.Day(), ")&&date()<date(", datetime2.Year(), ",", int(datetime2.Month()), ",", datetime2.Day(), "))||\n(date()=date(", datetime2.Year(), ",", int(datetime2.Month()), ",", datetime2.Day(), ")&&time()<=time(", datetime2.Hour(), ",", datetime2.Minute(), "))"))
				}
				//output.SetText(fmt.Sprint("(date()=date(", datetime.Year(), ",", int(datetime.Month()), ",", datetime.Day(), ")&&time()>=time(", datetime.Hour(), ",", datetime.Minute(), "))||\n(date()>date(", datetime.Year(), ",", int(datetime.Month()), ",", datetime.Day(), ")&&date()<date(", datetime2.Year(), ",", int(datetime2.Month()), ",", datetime2.Day(), "))||\n(date()=date(", datetime2.Year(), ",", int(datetime2.Month()), ",", datetime2.Day(), ")&&time()<=time(", datetime2.Hour(), ",", datetime2.Minute(), "))"))

			} else {
				output.SetText(fmt.Sprint("(date()=date(", datetime.Year(), ",", int(datetime.Month()), ",", datetime.Day(), ")&&time()>=time(", datetime.Hour(), ",", datetime.Minute(), "))||\n(date()=date(", datetime2.Year(), ",", int(datetime2.Month()), ",", datetime2.Day(), ")&&time()<=time(", datetime2.Hour(), ",", datetime2.Minute(), "))"))
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
