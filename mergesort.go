package main

import "fmt"

func fes(a []int, b []int) {
	t := []int{}

	i := 0
	j := 0
	fmt.Print(4545)

	for k := 0; k < len(a); k += 1 {
		u := k
		for h := k + 1; h < len(a); h++ {
			if a[u] > a[h] {
				u = h
			}

		}
		a[k], a[u] = a[u], a[k]
	}

	fmt.Print(a)
	//	l := 0

	//h := 0
	k := 0

	for ; k < len(b); k += 1 {
		u := k + 1
		for ; u < len(b); u++ {
			if b[u] < b[k] {
				b[u], b[k] = b[k], b[u]
			}
		}
	}

	//b = append(b[len(b)-1:], b[:len(b)-1]...)
	fmt.Print(b)
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			t = append(t, a[i])
			i++
		} else {
			t = append(t, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		t = append(t, a[i])
	}
	for ; j < len(b); j++ {
		t = append(t, b[j])
	}
	fmt.Println(t)

}
func main() {
	f := []int{3, 10, 5, 3, 1, 2, 0, 1, 0, 6, 89, 78}
	d := f[:len(f)/2]
	h := f[len(f)/2:]
	fes(d, h)
}
