//NOT DONE TODO

package main

import (
	"sort"
	"fmt"
)

type car struct {
	position int
	speed int 
}

type cars []car

func (c cars) Len() int {
	return len(c)
}

func (c cars) Less(first, second int) bool {
	return c[first].position > c[second].position
}

func (c cars) Swap(first, second int) {
	c[first], c[second] = c[second], c[first]
}

func carFleet(target int, position []int, speed []int) int {
	var cars cars
    for index, _ := range position {
		cars = append(cars, car{position: position[index], speed: speed[index]})
	}

	sort.Sort(cars)
	fmt.Println(cars)

	var stack []car
	
	for _, car := range cars {
		if (len(stack) == 0) {
			stack = append(stack, car)
			continue
		}

		top := stack[len(stack)-1]
		if (top.speed > car.speed)
	}
	return 0
}

func main() {
	carFleet(12, []int{10,8,0,5,3}, []int{2,4,1,1,3})
}