#include <assert.h>
#include <stdlib.h>

/**
Question: You are given a set S of n intervals on a line, with the ith interval described by its left and right endpoints(li, ri).
Give a O(n log n) algorithm to identify a point p on the line which is in the largest number of intervals.
As an example, for S = {(10,40), (20,60), (50,90), (15,70)}, no point exists in all 4 intervals.
But 50 exists in 3 intervals. You can assume an endpoint counts as being in its own interval(inclusive).

Algorithm Efficiency -  O(n^2) -- it sucks I know.

*/
typedef struct {
    int p;
    int c;
} PointCount;


bool find20(PointCount s) {
    if (s.p == 20) {
        return true;
    } else {
        return false;
    }
}

// Function that returns a pointer to an array of Data structures
PointCount* makePointCountArray(int size) {
    // Allocate memory for the array
    PointCount* array = (PointCount*)malloc(size * sizeof(PointCount));
    if (!array) {
        printf("Memory allocation failed!\n");
        return NULL;  // Return NULL on allocation failure
    }

    // Initialize the array
    for (int i = 0; i < size; i++) {
        array[i].p = 0;
        array[i].c = 0;
    }
    return array;
}

PointCount* FoldLeft(PointCount* array, int size, PointCount* accumulator, PointCount* (*f)(PointCount*, PointCount)) {
    if (size == 0) {
        return accumulator;
    }
    // Process the first element and recursively call for the rest
    return FoldLeft(array + 1, size - 1, f(accumulator, array[0]), f);
}

// The efficiency of this algorithm is O(N)
PointCount* Filter(PointCount* as, bool (*p)(PointCount)) {
    // var g = func(accum []T, h T) []T {
    //     if p(h) {
    //         return append(accum, h)
    //     } else {
    //         return accum
    //     }
    // }
    // return FoldLeft(as, []T{}, g)
}




// func TestFilterIntArray(t *testing.T) {
//     arr := []int{1, 2, 3, 3, 3, 3}
//     var p = func(s int) bool {
//         if s == 3 {
//             return true
//         } else {
//             return false
//         }
//     }
//
//     bigarray := Filter(arr, p)
//     if diff := deep.Equal(bigarray, []int{3, 3, 3, 3}); diff != nil {
//         t.Error(diff)
//     }
// }

// Define a function pointer type for the operation
typedef int (*thefunc)(int acc, int element);

// Recursive foldLeft function
int foldLeft(int *arr, int size, int index, int acc, thefunc f) {
    // Base case: when we've processed all elements
    if (index >= size) {
        return acc;
    }

    // Recursive case: apply the function and recurse
    return foldLeft(arr, size, index + 1, f(acc, arr[index]), f);
}

// Example operation: sum function
int lsumOperation(int acc, int element) {
    return acc + element;
}

int testFoldLeft() {
    PointCount S[] = {
        {10,0},
        {20,0},
        {50,0},
        {15,0},
        {40,0},
        {60,0},
        {90,0},
        {70,0}
    };

    PointCount *result = makePointCountArray(8);
    if (!result) {
        return 1;  // Exit if allocation failed
    }

    // Print the result
    for (int i = 0; i < 8; i++) {
        printf("Pointcount p = %d, c = %d\n", result[i].p, result[i].c);
    }

    free(result);


    //struct PointCount* result2 = FoldLeft(S, 8, value1);

    //char result = mostFrequentPoint(S, n);
    //assert(result == 50);
    //assert(result.r == 90);
}
