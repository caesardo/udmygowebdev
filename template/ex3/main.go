package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

//Menu is
type Menu struct {
	Session string
	Dishes  []Dish
}

//Dish is
type Dish struct {
	Title        string
	Price        float64
	Availability bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menus := []Menu{
		{
			Session: "Breakfast",
			Dishes: []Dish{
				{"Bubur Ayam", 15000, true},
				{"Ketoprak Lontong", 12000, true},
				{"Nasi Kuning", 20000, false},
			},
		},
		{
			Session: "Lunch",
			Dishes: []Dish{
				{"Soto Ayam", 25000, true},
				{"Nasi Padang", 30000, false},
				{"Empal Goreng", 28000, false},
			},
		},
		{
			Session: "Dinner",
			Dishes: []Dish{
				{"Salad Thai", 40000, true},
				{"Fettucine Carbonara", 36000, true},
			},
		},
	}
	err := tpl.Execute(os.Stdout, menus)
	if err != nil {
		log.Fatalln(err)
	}
}
