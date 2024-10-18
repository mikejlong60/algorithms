#include <stdlib.h>
#include <assert.h>
//#include "util.h"

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

int* setUnion(const int* A, const int* B, const int arrayASize, const int arrayBSize) {
    int* result;
    int* value1;
    if (!makeArray(arrayASize + arrayBSize, &result, &value1)) return value1;

    int i = 0, j = 0, k = 0;
    while (i < arrayASize, j < arrayBSize)
    {
        if (A[i] < B[j])
        {
            result[k] = A[i];
            i++;
            k++;

            if (i == arrayASize -1)
            {
                for (; j < arrayBSize; j++) {
                    result[k] = B[j];
                    return result;
                }
            }
        }
        else if (A[i] > B[j])
        {
            result[k] = B[j];
            j++;
            k++;
            if (j == arrayBSize -1)
            {
                for (; i < arrayASize; i++) {
                    result[k] = A[i];
                    return result;
                }
            }
        }
        else
        {//they are equal
            j++;
            k++;
            i++;
        }
    }
        return result;
}

int testSetUnion() {
    const int A[] = {1,2,3,4};
    const int B[] = {3, 6};

    //Arrays re already sorted

    // Calculate the number of elements in the array
    int arrayASize = sizeof(A) / sizeof(A[0]);
    int arrayBSize = sizeof(B) / sizeof(B[0]);

    int *result = setUnion(A, B, arrayASize, arrayBSize);
    printf("result[0]:%d\n", result[0]);
    printf("result[1]:%d\n", result[1]);
    assert(result[0] == 3);
    assert(result[1] == 1);

    free(result);
}
