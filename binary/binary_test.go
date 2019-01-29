package binary

import (
    "testing"
    "github.com/stretchr/testify/assert"
    // "fmt"
)

func TestBinarySearch(t *testing.T) {
    even_arr := []int{1, 3, 5, 7, 9, 11}
    odd_arr := []int{1, 3, 5, 7, 9, 11, 13}
    // odd_arr := []int{2, 4, 6, 8, 10, 12, 14}

    res, err := BinarySearch(even_arr, 3)
    assert.Equal(t, nil, err)
    assert.Equal(t, 1, res)

    res, err = BinarySearch(even_arr, 1)
    assert.Equal(t, nil, err)
    assert.Equal(t, 0, res)

    res, err = BinarySearch(even_arr, 7)
    assert.Equal(t, nil, err)
    assert.Equal(t, 3, res)

    res, err = BinarySearch(even_arr, 11)
    assert.Equal(t, nil, err)
    assert.Equal(t, 5, res)

    res, err = BinarySearch(even_arr, 13)
    assert.NotEqual(t, nil, err)

    res, err = BinarySearch(even_arr, -1)

    res, err = BinarySearch(odd_arr, 3)
    assert.Equal(t, nil, err)
    assert.Equal(t, 1, res)

    res, err = BinarySearch(odd_arr, 1)
    assert.Equal(t, nil, err)
    assert.Equal(t, 0, res)

    res, err = BinarySearch(odd_arr, 7)
    assert.Equal(t, nil, err)
    assert.Equal(t, 3, res)

    res, err = BinarySearch(odd_arr, 11)
    assert.Equal(t, nil, err)
    assert.Equal(t, 5, res)

    res, err = BinarySearch(odd_arr, 13)
    assert.Equal(t, nil, err)
    assert.Equal(t, 6, res)

    res, err = BinarySearch([]int{}, 1)
    assert.NotEqual(t, nil, err)

    res, err = BinarySearch([]int{1}, 1)
    assert.Equal(t, nil, err)
    assert.Equal(t, 0, res)

}