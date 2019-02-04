package selectsort

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
)


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
    {
        []int{9, 25, -1, 6, 84, 32, -50, 51, 0, 55},
        []int{-50, -1, 0, 6, 9, 25, 32, 51, 55, 84},
    },

}

func TestQuickSort(t *testing.T) {
    for _, sCase := range SortCases {
        QuickSort(sCase.array)
        assert.Equal(t, sCase.sorted, sCase.array)
    }
}
