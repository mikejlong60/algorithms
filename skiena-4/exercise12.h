#include <stdlib.h>
#include <assert.h>
#include <stdbool.h>

#include "inthashmap.h"

/**
Question: Give an efficient algorithm to compute the union of two sets A and B, where n = max(|A|,|B|). The output
should be an array of distict elements that form the union of the sets.  Write two different algorithms:
    1. one that assumes that the arrays are sorted and performs with O(n)
    2. one that assumes the arrays are unsorted and performs with O(n log n0.

Pseudocode for 1:
for (int i = 0, j = 0, k = 0; i < arrayASize, j < arrayBSize;) {
    if A[i] < B[j]
        result[k] = A[i];
        i++;
        k++;
        if i == arrayASize -1
            for (; j < arrayBSize; j++ {
                result[k] = B[j];
                return result;
            }
    else if A[i] > B[j]
        result[k] = B[j];
        j++;
        k++;
        if j == arrayBSize -1
            for (; i < arrayASize; i++ {
                result[k] = A[i];
                return result;
            }
    else //they are equal
        result[k] = B[j];
        j++;
        k++;
        i++;
}
*/

bool makeArray(const int arraySize, int** result, int** value1)
{
    *result = malloc(arraySize * sizeof(int));
    if (*result == NULL) {
        // Handle allocation failure
        *value1 = NULL;
        return false;
    }
    return true;
}

int* moreThanK(const int* xs, const int k, const int arraySize) {
    const int moreThan = arraySize/k;

    int* result;
    int* value1;
    if (!makeArray(arraySize, &result, &value1)) return value1;

    for (int j = 0, i = 0; j < arraySize; j++) {
        const struct Nlist* a = lookup(xs[j]);//lookup integer at position j in hashmap
        if (a == NULL) {  //key not there so add a new entry to hashmap
            install(xs[j], 1);
        } else {  //found key
            if (a->value > moreThan) {
                result[i] = a->key;
                i++;
            }
            //update hashmap to increment number of appearances of int in xs array
            install(xs[j], a->value + 1);
        }
    }
    return result;
}


int testMoreThanK()
{
    const int xs[] = {3, 3,3,3,1,1,1,1};
    // Calculate the number of elements in the array
    int arraySize = sizeof(xs) / sizeof(xs[0]);

    int *result = moreThanK(xs, 3, arraySize);
    printf("result[0]:%d\n", result[0]);
    printf("result[1]:%d\n", result[1]);
    assert(result[0] == 3);
    assert(result[1] == 1);

    free(result);
}
