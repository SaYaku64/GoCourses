package main

// methods for interface for every professions
func (d doctor) GetNamePosition() string {
	return d.FULLNAME + " : Doctor : " + d.POSITION
}
func (t teacher) GetNamePosition() string {
	return t.FULLNAME + " : Teacher : " + t.POSITION
}
func (m musician) GetNamePosition() string {
	return m.FULLNAME + " : Musician : " + m.POSITION
}
func (fm fireman) GetNamePosition() string {
	return fm.FULLNAME + " : Fireman : " + fm.POSITION
}
func (pr priest) GetNamePosition() string {
	return pr.FULLNAME + " : Priest : " + pr.POSITION
}

// Human - used for Employee
type Human struct {
	FULLNAME string
	AGE      int
}

// Employee - used for professions
type Employee struct {
	POSITION string
	SALARY   string
	*Human
}

type doctor struct {
	SPECIALITY string `json:"speciality"`
	*Employee
}

type teacher struct {
	MAINSUBJECT string `json:"subject"`
	*Employee
}

type musician struct {
	PLAYINGSTYLE string `json:"playingStyle"`
	*Employee
}
type fireman struct {
	DEPARTUREONCALLS int `json:"departureOnCalls"`
	*Employee
}
type priest struct {
	FAITH string `json:"faith"`
	*Employee
}

// Worker - interface, that is used to show name, profession and position
type Worker interface {
	GetNamePosition() string
}

// GetInfo function used to get info with interface
func GetInfo(w Worker) (info string) {
	info = w.GetNamePosition()
	return
}
