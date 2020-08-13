package main

import (
	"sync"
)

// Types used in our program
type human struct {
	birthday string
	name     string
	surname  string
}

// experience - number of working experience in years (4.5  -  4 yrs and 6 mnths)
type employee struct {
	jobName    string
	experience float64
	info       human
}

// type head struct {
// 	departmentName string
// 	subordinates   int
// 	info           human
// }

type myMutex struct {
	mutex sync.RWMutex
}
