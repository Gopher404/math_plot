package mt

func getMaxValY(values []DataValue) int {
	var m int
	for i, v := range values {
		if v.Y > values[m].Y {
			m = i
		}
	}
	return m
}

func getMaxValX(values []DataValue) int {
	var m int
	for i, v := range values {
		if v.X > values[m].X {
			m = i
		}
	}
	return m
}

func sortValues(arr []DataValue, by string) []DataValue {
	switch by {
	case "Y":
		return sortedValuesY(arr, 0, len(arr)-1)
	case "X":
		return sortedValuesX(arr, 0, len(arr)-1)
	default:
		return arr
	}
}

func sortedValuesY(arr []DataValue, low int, high int) []DataValue {
	if low < high {
		var p int
		arr, p = partitionY(arr, low, high)
		arr = sortedValuesY(arr, low, p-1)
		arr = sortedValuesY(arr, p+1, high)
	}
	return arr
}

func partitionY(arr []DataValue, low, high int) ([]DataValue, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j].Y < pivot.Y {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func sortedValuesX(arr []DataValue, low int, high int) []DataValue {
	if low < high {
		var p int
		arr, p = partitionX(arr, low, high)
		arr = sortedValuesX(arr, low, p-1)
		arr = sortedValuesX(arr, p+1, high)
	}
	return arr
}

func partitionX(arr []DataValue, low, high int) ([]DataValue, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j].X < pivot.X {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
