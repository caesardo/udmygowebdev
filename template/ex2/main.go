package main

import (
	"log"
	"os"
	"text/template"
)

//Hotel is
type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

//Region is
type Region struct {
	Name   string
	Hotels []Hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	hotels := []Region{
		Region{
			Name: "Southern",
			Hotels: []Hotel{
				{"Hotel California", "Jbawokie Street Block D number 1A", "Los Angeles", "9392092"},
				{"Hotel South", "Hangret Street 2331R", "Los Angeles", "320923"},
			},
		},
		Region{
			Name: "Central",
			Hotels: []Hotel{
				{"Hotel Central Boulevard", "Central Junction Mall floor 7", "San Diego", "39202"},
				{"Hotel Holiday", "Central Beach Boulevard 87T", "San Diego", "2323"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
