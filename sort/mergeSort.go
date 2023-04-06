package sort

import (
	"sync"
)

// to run code RunMultiMergesortWithSem([]int)
// from https://medium.com/@yliu224/goroutines-on-mergesort-12a2a7a43cc2
func MultiMergeSortWithSem(data []float64, sem chan struct{}) []float64 {
	if len(data) < 2 {
		return data
	}

	middle := len(data) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var ldata []float64
	var rdata []float64

	select {
	case sem <- struct{}{}:
		go func() {
			ldata = MultiMergeSortWithSem(data[:middle], sem)
			<-sem
			wg.Done()
		}()
	default:
		ldata = SingleMergeSort(data[:middle])
		wg.Done()
	}

	select {
	case sem <- struct{}{}:
		go func() {
			rdata = MultiMergeSortWithSem(data[middle:], sem)
			<-sem
			wg.Done()
		}()
	default:
		rdata = SingleMergeSort(data[middle:])
		wg.Done()
	}

	wg.Wait()
	return Merge(ldata, rdata)
}

func RunMultiMergesortWithSem(data []float64) []float64 {
	sem := make(chan struct{}, 4)
	return MultiMergeSortWithSem(data, sem)
}

func SingleMergeSort(data []float64) []float64 {
	if len(data) < 2 {
		return data
	}
	middle := len(data) / 2
	return Merge(SingleMergeSort(data[:middle]), SingleMergeSort(data[middle:]))
}

func Merge(ldata []float64, rdata []float64) (result []float64) {
	result = make([]float64, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx] < rdata[ridx]:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}
	return
}
