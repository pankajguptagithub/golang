// Ok idom and checking if a given element exits in the map

package main

import "fmt"

func main() {
    // Creating a map
    myMap := map[string]int{
        "apple":  5,
        "banana": 10,
        "orange": 15,
    }

    // Check if an element exists in the map
    key := "banana"
    if _, ok := myMap[key]; ok {
        fmt.Printf("Element with key '%s' exists in the map.\n", key)
    } else {
        fmt.Printf("Element with key '%s' does not exist in the map.\n", key)
    }

    // Another example
    key = "grape"
    if _, ok := myMap[key]; ok {
        fmt.Printf("Element with key '%s' exists in the map.\n", key)
    } else {
        fmt.Printf("Element with key '%s' does not exist in the map.\n", key)
    }
}
