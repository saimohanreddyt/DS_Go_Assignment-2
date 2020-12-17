//Quick-Sort
//-x-x-x-x-x-

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {

    slice := generateSlice(20)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    quicksort(slice)
    fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}

func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }

    left, right := 0, len(a)-1

    pivot := rand.Int() % len(a)

    a[pivot], a[right] = a[right], a[pivot]

    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }

    a[left], a[right] = a[right], a[left]

    quicksort(a[:left])
    quicksort(a[left+1:])

    return a
}




/*	p @ pivot
	----------------------------------
			p
		<=p	p	>p
	-----------------------------------=


					6 4 5 1 8 2 9 7 10 3
					-
				4 5 1 3 2 	6		9 7 10 8
				-				-
			2 1 3	4	5		7 8 	9	10
			-
		1	2	3

					1 2 3 4 5 6 7 8 9 10


Quick sort is also a recursive algorithm.
· In each step, we select a pivot (let us say first element of list).
. Then we traverse the rest of the list and copy all the elements of the list which are smaller than the pivot to the left side of list
. We copy all the elements of the list, which are greater than pivot to theright side of the list. Obviously, the pivot is at its sorted position.
· Then we sort the left and right subarray separately.
· When the algorithm returns the whole list is sorted.

*/
