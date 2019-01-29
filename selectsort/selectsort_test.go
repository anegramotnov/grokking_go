package selectsort

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
)

func TestMin(t *testing.T) {
    array := []int{7, 3, 5, 4, 0, 2, 1, 6, 8}

    minIndex, _ := findMin(array)
    assert.Equal(t, 4, minIndex)

    _, err := findMin([]int{})
    assert.NotEqual(t, nil, err)
}

func TestSelectSort(t *testing.T) {
    array := []int{7, 3, 5, 8, -10, 2, 1, 6, 4}
    sorted_array := []int{-10, 1, 2, 3, 4, 5, 6, 7, 8}

    SelectSort(array)
    fmt.Printf("Sorted array: %v\n", array)


    assert.Equal(t, sorted_array, array)

}

func TestSelectSortEmpty(t *testing.T) {
    SelectSort([]int{})
}