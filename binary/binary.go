package binary

import (
    "errors"
    "fmt"
)

func BinarySearch(array []int, value int) (int, error) {
    startIndex := 0
    endIndex := len(array)-1

    fmt.Printf("Find %d in %v\n", value, array)
    // tick := 1
    for startIndex <= endIndex {
        // fmt.Println("")
        // fmt.Printf("Tick %d\n", tick)
        // tick++

        mid := (endIndex + startIndex) / 2
        // fmt.Printf("Index state: startIndex=%d, endIndex=%d, mid=%d\n", startIndex, endIndex, mid)
        guess := array[mid]
        // fmt.Printf("guess=%d\n", guess)
        if guess == value {
            fmt.Printf("Index: %d\n\n", mid)
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
    fmt.Printf("Not found\n\n")
    return 0, errors.New("Not Found")
}