package main

import (
	"fmt"
	// there should be "encoding/json", but if I add it, all tags automatically deleting
)

// premium - maximum value of premium. It multiplies to coefficient, but in this
// program - didn't used. Saved for keeping "coefficient" field in employee type
const premium = 3000

// human - structure, that keeps information about person
// personal identificator (id) - number, that helps to distinguish people
// name - person's name
// surname - person's surname
type human struct {
	id      int
	name    string
	surname string
}

// employee - structure, that keeps information about working time of employees
// normal hours (nH) - normal count of working hours per month
// actual hours (aH) - actual count of working hours per month
// coefficient (c) - (0.0-1.0) coefficient that will be multiplied to PREMIUM
// salary - amount of salary
// human - "human" structure, keeps personal information
type employee struct {
	nH, aH int
	c      float64
	salary int
	human  human
}

// pediatrician - structure, that keeps count of serviced pacients (sP),
// count of prescribed medication (pM) and
// employee info (eI) - "employee" structure now and in other types
type pediatrician struct {
	sP int `json:"ServicedPacients"`
	pM int `json:"PrescribedMedication"`
	eI employee
}

// chef - structure, that keeps count of made dishes (mD) and count of thanks from
// the guests specially to chef
type chef struct {
	mD     int `json:"MadeDishes"`
	thanks int `json:"Thanks"`
	eI     employee
}

// tutor - structure, that keeps count of children and price per hour of teaching
type tutor struct {
	children int `json:"ChildrenAmount"`
	price    int `json:"PricePerHour"`
	eI       employee
}

// fireman - structure, that keeps count of departures on calls (dC)
type fireman struct {
	dC int `json:"DeparturesOnCalls"`
	eI employee
}

// gardener - structure, that keeps working area,
// the square of the cultivated plot
type gardener struct {
	area float64 `json:"CultivatedP"`
	eI   employee
}

func main() {

	// slicing game)
	namesSurnames := [10]string{
		"Alexander",
		"Tony",
		"Tommy",
		"Markus",
		"Joe",
		"Yakushev",
		"Montana",
		"Vercetti",
		"Persson",
		"Barbaro",
	}
	names := namesSurnames[:5]
	surnames := namesSurnames[5:]
	// these "slicing" strokes were added because of hometask, it should be simpler

	// creating human types for employees
	human1 := human{id: 1, name: names[0], surname: surnames[0]}
	human2 := human{id: 2, name: names[1], surname: surnames[1]}
	human3 := human{id: 3, name: names[2], surname: surnames[2]}
	human4 := human{id: 4, name: names[3], surname: surnames[3]}
	human5 := human{id: 5, name: names[4], surname: surnames[4]}

	//
	// creating some employee types for professions
	//

	// type of employee for pediatrician
	empPed := employee{
		nH:     180,
		aH:     182,
		c:      0.75,
		salary: 17500,
		human:  human1,
	}

	// type of employee for chef
	empChef := employee{
		nH:     160,
		aH:     157,
		c:      0,
		salary: 16000,
		human:  human2,
	}

	// type of employee for tutor
	empTutor := employee{
		aH:    42,
		human: human3,
	}

	//
	// creating all professions using humans and employees
	//

	// creating pediatrician
	mrPediatrician := pediatrician{
		sP: 520,
		pM: 450,
		eI: empPed,
	}

	// creating chef
	mrChef := chef{
		mD:     375,
		thanks: 100,
		eI:     empChef,
	}

	// creating tutor
	mrTutor := tutor{
		children: 42,
		price:    250,
		eI:       empTutor,
	}

	// creating fireman
	mrFireman := fireman{
		dC: 10,
		eI: employee{
			nH:     180,
			aH:     180,
			c:      0.5,
			salary: 20000,
			human:  human4,
		},
	}

	// creating gardener
	mrGardener := gardener{
		area: 205.5,
		eI: employee{
			nH:     126,
			aH:     126,
			c:      0.25,
			salary: 10000,
			human:  human5,
		},
	}

	// determining the tutor's salary using actual hours and price per hour
	mrTutor.eI.salary = mrTutor.eI.aH * mrTutor.price

	// displaying everything (drop the hashes)
	fmt.Printf("%#v\n", mrPediatrician)
	fmt.Printf("%#v\n", mrChef)
	fmt.Printf("%#v\n", mrTutor)
	fmt.Printf("%#v\n", mrFireman)
	fmt.Printf("%#v\n", mrGardener)
}
