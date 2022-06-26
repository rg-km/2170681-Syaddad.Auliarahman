package sortKM

// TODO: answer here

//concrete strategy implementation
type DescendingSort struct{}

func (ds *DescendingSort) Sort(array []int) {
	//choose any sort algo you want
	// TODO: answer here
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
