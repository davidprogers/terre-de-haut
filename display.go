package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
)

// template used for output
const weatherTemplate = `Current weather for {{.Name}}:
Conditions: {{range .Weather}} {{.Description}} {{end}}
Now:         {{.Main.Temp}} {{.Unit}}
High:        {{.Main.TempMax}} {{.Unit}}
Low:         {{.Main.TempMin}} {{.Unit}}
`

func main() {
	stations := getStationList()
	for {
		currentConditionsMap := getCurrentForEachStation(stations)
		for _, w := range currentConditionsMap {
			fmt.Println(time.Now().Format(time.UnixDate))
			tmpl, err := template.New("weather").Parse(weatherTemplate)
			if err != nil {
				log.Fatalln(err)
			}

			// Render the template and display
			err = tmpl.Execute(os.Stdout, w)
			if err != nil {
				log.Fatalln(err)
			}
		}
		fmt.Print("******************\n\n")

		time.Sleep(time.Duration(1) * time.Second)
	}
}
