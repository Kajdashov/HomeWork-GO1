package main

import "fmt"

func main() {
	sort()
}

func sort() {
	unsortedarray := [...]int{44, 3, 5, 17, 1, 33, 6, 6, 54, 13}
	//sortedarray := [len(unsortedarray)]int{}

	for n := 1; n < len(unsortedarray); n++ {
		curentElement := unsortedarray[n]
		l := n
		for ; l >= 1 && unsortedarray[l-1] > curentElement; l-- {
			unsortedarray[l] = unsortedarray[l-1]
		}
		unsortedarray[l] = curentElement

	}
	fmt.Print(unsortedarray)
	/*array = unsortedarray
	n := 1
	l := array[n]
	newElement := array[n-1]
	targetPosition := n - 1
	array[targetPosition] = newElement

	n = 2
	newElement = array[n-1]
	targetPosition = n - 1
	x := array[n-2] < newElement
	if x {

	} else {
		targetPosition = targetPosition - 1
	}
	array[n-1] = array[n-2] //123455
	array[n-2] = array[n-3] // 123445
	array[n-3] = array[n-4] //123345

	n = 2
	x := array[n-2] < array[n-1]
	if x == false {
		//array[1] = array[0]
		var a = array[n-2]
		array[n-2] = array[n-1]
		array[n-1] = a
	}
	n = 3
	y := array[n-2] < array[n-1]
	if y == false {
		c := array[n-3] < array[n-1]
		if c == true {
			var d = array[n-1]
			array[n-1] = array[n-2]
			array[n-2] = d
		} else if c == false {
			var f = array[n-1]
			array[n-1] = array[n-2]
			array[n-2] = array[n-3]
			array[n-3] = f
		}
	}
	*/
	//fmt.Print(sortedarray)

}
