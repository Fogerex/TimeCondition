package main

import (
	"fmt"
	"time"
	"flag"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	version     = "0.1.0" // версия приложения
	showVersion bool      // флаг, указывающий на необходимость вывода версии
)

func init() {
	flag.BoolVar(&showVersion, "v", false, "вывести версию приложения и выйти")
	flag.Parse()
}

func main() {
		// версия приложения
		if showVersion {
			fmt.Println(version)
			return
		}

	layout := "15:04 02.01.2006"
	myApp := app.New()
	myWindow := myApp.NewWindow( fmt.Sprintf("Date and Time Condition for ICM %s",version))
	now := time.Now().Local()
	nowdate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	startbox := widget.NewEntry()
	start := nowdate.Format(layout)
	startbox.SetText(start)
	endbox := widget.NewEntry()
	end := nowdate.Format(layout)
	endbox.SetText(end)
	output := widget.NewMultiLineEntry()

	startbox.OnChanged = func(text string) {
		output.SetText(changetime(layout, startbox, endbox, output))
	}
	endbox.OnChanged = func(text string) {
		output.SetText(changetime(layout, startbox, endbox, output))
	}
	btn := widget.NewButton("Копировать", func() {
		// Записываем содержимое в буфер обмена
		myWindow.Clipboard().SetContent(output.Text)
	})
	dateTimeBox := container.NewVBox(startbox, endbox, output, btn)
	content := container.NewVBox(dateTimeBox)
	frame := widget.NewCard("Enter Date and Time             ", "", content)
	myWindow.SetContent(frame)
	myWindow.ShowAndRun()
}

func changetime(layout string, startbox *widget.Entry, endbox *widget.Entry, output *widget.Entry) string {

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
		output.SetText(fmt.Sprint(string1, "&&", "\n", string3))
	} else {
		output.SetText(fmt.Sprint(string1, "||", "\n", string2, "||", "\n", string3))
	}
	return output.Text

	//time.Sleep(time.Second)

}
