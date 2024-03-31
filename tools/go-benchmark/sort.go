package sort

func InsertionSort(a []int) {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
}

func QuickSort(a []int) {
	doQuickSort(a, 0, len(a))
}

func doQuickSort(a []int, p int, r int) {
	if p < r {
		q := partition(a, p, r)
		doQuickSort(a, p, q)
		doQuickSort(a, q+1, r)
	}
}

func partition(a []int, p int, r int) int {
	x := a[r-1]
	i := p - 1
	for j := p; j < r-1; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r-1] = a[r-1], a[i+1]
	return i + 1
}
