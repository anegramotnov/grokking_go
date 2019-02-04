package selectsort

import (
    "errors"
)

func findMin(array []int) (minIndex int, err error) {
    if len(array) == 0 {
        return 0, errors.New("Empty slice")
    }

    minIndex = 0
    min := array[minIndex]

    for i := 1; i < len(array); i++ {
        if array[i] < min {
            min = array[i]
            minIndex = i
        }
    }

    return minIndex, nil
}

func SelectSort(array []int) {
    for i := 0; i < len(array); i++ {
        minIndex, _ := findMin(array[i:len(array)])
        array[i], array[minIndex+i] = array[minIndex+i], array[i]
    }
}

func QuickSort(array[]int) {
    if len(array) < 2 {
        return
    }
}