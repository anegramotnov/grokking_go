package binary

import (
    "errors"
    "fmt"
)


func factorial(n int64) (result int64) {
    if n == 0 {
        return 1
    } else {
        return n * factorial(n - 1)
    }
}

func Factorial(n int64) (result int64, err error) {
    if n < 0 {
        return 0, errors.New("Factorial not defined for negative values")
    } else {
        return factorial(n), nil
    }

}

func Sum(array []int) int {
    if len(array) == 0 {
        return 0
    } else {
        return array[0] + Sum(array[1:len(array)])
    }
}

func Count(array []int) int {
    if len(array) == 0 {
        return 0
    } else {
        return 1 + Count(array[1:len(array)])
    }

}

func max(array []int) int {
    if len(array) == 1 {
        return array[0]
    } else {
        interMax := max(array[1:len(array)])
        if array[0] > interMax {
            return array[0]
        } else {
            return interMax
        }
    }
}

func Max(array []int) (int, error) {
    if len(array) < 1 {
        return 0, errors.New("Empty slices not supported")
    } else {
        return max(array), nil
    }
}

func recursiveBinarySearch(array []int, start, end, value int) int {
    fmt.Printf("Search %d in %v (%d:%d)\n", value, array, start, end)
    if end >= start {
        mid := (end + start) / 2
        // mid := start + (end - start) / 2
        guess := array[mid]
        fmt.Printf("mid=%d, guess=%d\n", mid, guess)

        if guess == value {
            fmt.Printf("Bingo!\n\n")
            return mid
        }

        if value < guess {
            return recursiveBinarySearch(array, start, mid-1, value)
        } else {
            return recursiveBinarySearch(array, mid+1, end, value)
        }
    }
    fmt.Printf("Not found :(\n\n")
    return -1
}

func RecursiveBinarySearch(array []int, value int) (int, error) {
    searchResult := recursiveBinarySearch(array, 0, len(array)-1, value)
    if searchResult == -1 {
        return 0, errors.New("Not Found")
    } else {
        return searchResult, nil
    }
} 

