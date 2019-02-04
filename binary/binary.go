package binary

import (
    "errors"
)

func BinarySearch(array []int, value int) (int, error) {
    startIndex := 0
    endIndex := len(array)-1

    for startIndex <= endIndex {
        // TODO: How will mid work with "Search 13 in [1 3 5 7 9 11]"?
        mid := (endIndex + startIndex) / 2
        guess := array[mid]
        if guess == value {
            return mid, nil
        }
        if value < guess {
            endIndex = mid - 1 
            continue
        } else {
            startIndex = mid + 1
            continue
        }
    }
    return 0, errors.New("Not Found")
}