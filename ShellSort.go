package sortings

import "fmt"

func ShellSort(n []int){
	var Len = len(n)
	var gap int = Len/2
	for gap > 0{
		var i = gap
		for i < Len{
			temp := n[i]
			j := i-gap
			for j >= 0 && n[j] > temp{
				n[j+gap] = n[j]
				j = j - gap
			}
			n[j+gap] = temp
			i = i+1
		}
		gap = gap/2
	}

	fmt.Println(n)
}