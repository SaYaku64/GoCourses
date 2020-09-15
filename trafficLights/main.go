package main

import (
	"fmt"
	"math/rand"
	"time"
)

var light bool // false - зелёный, true - красный

// Тип структуры, который описывает направление пути
type way struct {
	name     string
	gotCars  int
	sentCars int
}

var north = way{name: "north"}
var east = way{name: "east"}
var south = way{name: "south"}
var west = way{name: "west"}

// Тип структуры, которая описывает маршрут машины
type car struct {
	startPoint string
	endPoint   string
}

// Очередь из машин
var standingCars []car

// Генерирует начальную и конечную точки пути для машины
func genWay() (string, string) {
	ways := [4]string{"north", "east", "south", "west"}

	rand.Seed(time.Now().Unix())
	start := ways[rand.Intn(len(ways))]
	end := ways[rand.Intn(len(ways))]
	for {
		if start == end {
			end = ways[rand.Intn(len(ways))]
		} else {
			return start, end
		}
	}

}

// Создаёт машину с маршрутом
func initCar() car {
	car := car{}
	car.startPoint, car.endPoint = genWay()
	return car
}

// "Ловит" изменение цвета светофора
// По сути просто меняет цвет
func lightsOn(change <-chan time.Time) {
	for {
		select {
		case <-change:
			if light == false {
				showWays()
				fmt.Println("==========||Red||==========")
				light = true
				getRidOfQueue()
			} else {
				showWays()
				fmt.Println("==========||Green||==========")
				light = false
				getRidOfQueue()
			}
		}
	}
}

// Начало движения в зависимости от цвета
func ridingByLight(car car) {
	if light == false {
		// Зелёный светит
		greenRiding(car)
		time.Sleep(getRandomDuration() * time.Millisecond) // Время, которое тратит машина на проезд или остановку
	} else {
		// Красный светит
		redRiding(car)
		time.Sleep(getRandomDuration() * time.Millisecond)
	}
}

// Избавление от очередей
func getRidOfQueue() {
	for _, car := range standingCars {
		ridingByLight(car)
		if len(standingCars) != 0 {
			standingCars = standingCars[1:]
		}
	}
}

// Зелёный для Севера (North) и Юга (South)
// Красный для Востока (East) и Запада (West)
func greenRiding(car car) {
	switch car.startPoint {
	case "north":
		switch car.endPoint {
		case "west":
			north.sentCars++
			west.gotCars++

		case "south":
			north.sentCars++
			south.gotCars++
		case "east":
			north.sentCars++
			east.gotCars++
		}
	case "east":
		switch car.endPoint {
		case "west":
			standingCars = append(standingCars, car)
		case "south":
			standingCars = append(standingCars, car)
		case "north":
			east.sentCars++
			north.gotCars++
		}
	case "south":
		switch car.endPoint {
		case "west":
			south.sentCars++
			west.gotCars++
		case "east":
			south.sentCars++
			east.gotCars++
		case "north":
			south.sentCars++
			north.gotCars++
		}
	case "west":
		switch car.endPoint {
		case "south":
			west.sentCars++
			south.gotCars++
		case "east":
			standingCars = append(standingCars, car)
		case "north":
			standingCars = append(standingCars, car)
		}
	}
}

// Красный для Севера (North) и Юга (South)
// Зелёный для Востока (East) и Запада (West)
func redRiding(car car) {
	switch car.startPoint {
	case "north":
		switch car.endPoint {
		case "west":
			north.sentCars++
			west.gotCars++
		case "south":
			standingCars = append(standingCars, car)
		case "east":
			standingCars = append(standingCars, car)
		}
	case "east":
		switch car.endPoint {
		case "west":
			east.sentCars++
			west.gotCars++
		case "south":
			east.sentCars++
			south.gotCars++
		case "north":
			east.sentCars++
			north.gotCars++
		}
	case "south":
		switch car.endPoint {
		case "west":
			standingCars = append(standingCars, car)
		case "east":
			south.sentCars++
			east.gotCars++
		case "north":
			standingCars = append(standingCars, car)
		}
	case "west":
		switch car.endPoint {
		case "south":
			west.sentCars++
			south.gotCars++
		case "east":
			west.sentCars++
			east.gotCars++
		case "north":
			west.sentCars++
			north.gotCars++
		}
	}
}

// Возвращает псевдо-рандомную переменную-множитель для создания разного времени проезда машины или остановки на светофоре
func getRandomDuration() time.Duration {
	r := 10 * (rand.Intn(50) + 50)
	fmt.Println("Auto passes road in ", r, "mSec")
	return time.Duration(r)
}

// Показывает сколько выехало и въехало в каждый путь и очередь из машин
func showWays() {
	fmt.Printf("North || Got cars: %v Sent cars: %v\n", north.gotCars, north.sentCars)
	fmt.Printf("South || Got cars: %v Sent cars: %v\n", south.gotCars, south.sentCars)
	fmt.Printf("East  || Got cars: %v Sent cars: %v\n", east.gotCars, east.sentCars)
	fmt.Printf("West  || Got cars: %v Sent cars: %v\n", west.gotCars, west.sentCars)
	fmt.Println("Queue of standing cars: ", standingCars)
}

// Создаёт машины и отправляет их в путь по своему маршруту, если свет позволяет
func checkColor() {
	for {
		car := initCar()

		ridingByLight(car)
	}
}

func main() {

	fmt.Println("==========||Green||==========")
	change := time.Tick(5000 * time.Millisecond) // период смены цвета
	go checkColor()
	go lightsOn(change) // смена цвета

	fmt.Scanln()
}
