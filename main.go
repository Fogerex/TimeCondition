package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	layout := "15:04 02.01.2006"
	myApp := app.New()
	myWindow := myApp.NewWindow("Date and Time Condition for ICM")
	now := time.Now().Local()
	nowdate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	// Create a label to display the selected date and time
	startbox := widget.NewEntry()
	start := nowdate.Format(layout)
	startbox.SetText(start)
	datetime, err := time.Parse(layout, start)
	if err != nil {
		fmt.Println(err)
		return
	}
	endbox := widget.NewEntry()
	end := nowdate.Format(layout)
	endbox.SetText(nowdate.Format(layout))
	datetime2, err := time.Parse(layout, end)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Create MultiLineEntry
	output := widget.NewMultiLineEntry()
	if datetime.Year() != datetime2.Year() || datetime.Month() != datetime2.Month() || datetime.Day() != datetime2.Day() {
		output.SetText(fmt.Sprint("(date()=date(", datetime.Year(), ",", datetime.Month(), ",", datetime.Day(), ")&&time()>=time(", datetime.Hour(), ",", datetime.Minute(), "))||\n(date()>date(", datetime.Year(), ",", datetime.Month(), ",", datetime.Day(), ")&&date()<date(", datetime2.Year(), ",", datetime2.Month(), ",", datetime2.Day(), "))||\n(date()=date(", datetime2.Year(), ",", datetime2.Month(), ",", datetime2.Day(), ")&&time()<=time(", datetime2.Hour(), ",", datetime2.Minute(), "))"))
	} else {
		output.SetText(fmt.Sprint("(date()=date(", datetime.Year(), ",", datetime.Month(), ",", datetime.Day(), ")&&time()>=time(", datetime.Hour(), ",", datetime.Minute(), "))||\n(date()=date(", datetime2.Year(), ",", datetime2.Month(), ",", datetime2.Day(), ")&&time()<=time(", datetime2.Hour(), ",", datetime2.Minute(), "))"))
	}

	// Create a horizontal box container to hold the date and time picker and the label
	dateTimeBox := container.NewVBox(startbox, endbox, output)

	// Create a vertical box container to hold the horizontal box container and the frame
	content := container.NewVBox(dateTimeBox)

	// Create a frame to display the selected date and time
	frame := widget.NewCard("Enter Date and Time", "", content)

	// Update the label with the selected date and time whenever the date and time picker value changes
	// dateTimePicker.OnChanged = func(dateTime time.Time) {
	// 	dateTimeLabel.SetText(dateTime.Format("2006-01-02 15:04:05"))
	// }

	// Update the label with the current date and time every second
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
				output.SetText(fmt.Sprint("(date()=date(", datetime.Year(), ",", datetime.Month(), ",", datetime.Day(), ")&&time()>=time(", datetime.Hour(), ",", datetime.Minute(), "))||\n(date()>date(", datetime.Year(), ",", datetime.Month(), ",", datetime.Day(), ")&&date()<date(", datetime2.Year(), ",", datetime2.Month(), ",", datetime2.Day(), "))||\n(date()=date(", datetime2.Year(), ",", datetime2.Month(), ",", datetime2.Day(), ")&&time()<=time(", datetime2.Hour(), ",", datetime2.Minute(), "))"))
			} else {
				output.SetText(fmt.Sprint("(date()=date(", datetime.Year(), ",", datetime.Month(), ",", datetime.Day(), ")&&time()>=time(", datetime.Hour(), ",", datetime.Minute(), "))||\n(date()=date(", datetime2.Year(), ",", datetime2.Month(), ",", datetime2.Day(), ")&&time()<=time(", datetime2.Hour(), ",", datetime2.Minute(), "))"))
			}
			time.Sleep(time.Second)

		}
	}()

	// Set the window content to the frame
	myWindow.SetContent(frame)

	// Show the window
	myWindow.ShowAndRun()
}
