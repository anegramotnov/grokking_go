package selectsort

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
    array := []int{7, 3, 5, 4, 0, 2, 1, 6, 8}

    minIndex, _ := findMin(array)
    assert.Equal(t, 4, minIndex)

    _, err := findMin([]int{})
    assert.NotEqual(t, nil, err)
}

type SortTestCase struct {
    array []int
    sorted []int
}

var SortCases = []SortTestCase{
    {[]int{}, []int{}},
    {[]int{1}, []int{1}},
    {[]int{1, 2}, []int{1, 2}},
    {[]int{2, 1}, []int{1, 2}},
    {[]int{1, 2, 3}, []int{1, 2, 3}},
    {[]int{3, 2, 1}, []int{1, 2, 3}},
    {[]int{9, 25, -1, 6, 84, 32, -50, 51, 0, 55}, []int{-50, -1, 0, 6, 9, 25, 32, 51, 55, 84}},
}

func TestSelectSort(t *testing.T) {
    for _, sCase := range SortCases {
        SelectSort(sCase.array)
        assert.Equal(t, sCase.sorted, sCase.array)
    }
}
