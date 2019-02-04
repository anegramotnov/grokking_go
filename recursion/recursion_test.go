package binary

import (
    "testing"
    "github.com/stretchr/testify/assert"
    // "fmt"
)

func TestFactorial(t *testing.T) {
    fac, _ := Factorial(5)
    assert.Equal(t, int64(120), fac)

    fac, _ = Factorial(1)
    assert.Equal(t, int64(1), fac)

    fac, _ = Factorial(0)
    assert.Equal(t, int64(1), fac)

    _, err := Factorial(-1)
    assert.NotEqual(t, err, nil)
}

func TestSum(t *testing.T) {
    sum := Sum([]int{})
    assert.Equal(t, 0, sum)

    sum = Sum([]int{0, 1, 2, 3, 4})
    assert.Equal(t, 10, sum)
}

func TestCount(t *testing.T) {
    count := Count([]int{})
    assert.Equal(t, 0, count)

    count = Count([]int{0, 1, 2, 3, 4})
    assert.Equal(t, 5, count)
}

func TestMax(t *testing.T) {
    _, err := Max([]int{})
    assert.NotEqual(t, nil, err)

    max, _ := Max([]int{0})
    assert.Equal(t, 0, max)

    max, _ = Max([]int{5, 4, 3, 2, 1, 0})
    assert.Equal(t, 5, max)

    max, _ = Max([]int{0, 1, 2, 3, 4, 5})
    assert.Equal(t, 5, max)

    max, _ = Max([]int{3, 2, 5, 4, 0, 1})
    assert.Equal(t, 5, max)

}

func TestRecursiveBinarySearch(t *testing.T) {
    // TODO: Case Pattern! 
    even_arr := []int{1, 3, 5, 7, 9, 11}
    odd_arr := []int{1, 3, 5, 7, 9, 11, 13}

    res, err := RecursiveBinarySearch(even_arr, 3)
    assert.Equal(t, nil, err)
    assert.Equal(t, 1, res)

    res, err = RecursiveBinarySearch(even_arr, 1)
    assert.Equal(t, nil, err)
    assert.Equal(t, 0, res)

    res, err = RecursiveBinarySearch(even_arr, 7)
    assert.Equal(t, nil, err)
    assert.Equal(t, 3, res)

    res, err = RecursiveBinarySearch(even_arr, 11)
    assert.Equal(t, nil, err)
    assert.Equal(t, 5, res)

    res, err = RecursiveBinarySearch(even_arr, 13)
    assert.NotEqual(t, nil, err)

    res, err = RecursiveBinarySearch(even_arr, -1)

    res, err = RecursiveBinarySearch(odd_arr, 3)
    assert.Equal(t, nil, err)
    assert.Equal(t, 1, res)

    res, err = RecursiveBinarySearch(odd_arr, 1)
    assert.Equal(t, nil, err)
    assert.Equal(t, 0, res)

    res, err = RecursiveBinarySearch(odd_arr, 7)
    assert.Equal(t, nil, err)
    assert.Equal(t, 3, res)

    res, err = RecursiveBinarySearch(odd_arr, 11)
    assert.Equal(t, nil, err)
    assert.Equal(t, 5, res)

    res, err = RecursiveBinarySearch(odd_arr, 13)
    assert.Equal(t, nil, err)
    assert.Equal(t, 6, res)

    res, err = RecursiveBinarySearch([]int{}, 1)
    assert.NotEqual(t, nil, err)

    res, err = RecursiveBinarySearch(even_arr, -5)
    assert.NotEqual(t, nil, err)


    res, err = RecursiveBinarySearch([]int{1}, 1)
    assert.Equal(t, nil, err)
    assert.Equal(t, 0, res)

}