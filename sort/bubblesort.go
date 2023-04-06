package sort

// from https://dejavuqa.tistory.com/353
func BubbleSort(tosort []int) {
	size := len(tosort)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if tosort[j] < tosort[j-1] {
				tosort[j], tosort[j-1] = tosort[j-1], tosort[j]
			}
		}
	}
}
