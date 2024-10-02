#include <stdlib.h>
#include <assert.h>

/**
Question: Given a set S of n integers and an integer T give two different algorithms to determine if two elements exist whose sum is T:
    case 1: S is unsorted - should have cost of O(n log n)
    case 2: S is sorted - should have cost of O(n)
*/

//Binary search
int find(int* S, const int leftOffset, const int rightOffset, const int lookingFor, const int arraySize) {
    //If lookingFor is completely outside array return false right away
    if (S[0] > lookingFor || S[arraySize - 1] < lookingFor)
        return 0;
    //If lookingFor is first or last element return true right away
    if (S[arraySize - 1] == lookingFor || S[0] == lookingFor)//case where looking for is left-most or right-most element.
        return 1;

    //Otherwise recursively split array with binary search
    const int a = (rightOffset - leftOffset)/2;

    const int relativeA = leftOffset + a;
    if (S[relativeA] == lookingFor) {
        return 1;
    } else if (S[relativeA] < lookingFor) {
        return find(S, leftOffset + a, rightOffset, lookingFor, arraySize);
    } else
        return find(S, leftOffset, leftOffset + a,  lookingFor, arraySize);
}

int twoSum(int *S, const int T, const int arraySize) {
    int desiredDifference = 0;
    for (int j = 0; j < arraySize; j++) {
        desiredDifference = T - S[j];
        const int there = find(S, 0, arraySize-1, desiredDifference, arraySize);
        if (there)
            return 1;
        if (S[j] > T)
            return 0;//No need to look further since S[j] is too large to sum to T
    }
    return 0;
}

int testTwoSum() {
    int S[] = {3,2,1,4,5,6, 63, 12, 23, 34, -19, 19};
    // Calculate the number of elements in the array
    int n = sizeof(S) / sizeof(S[0]);

    // Use qsort to sort the array
    qsort(S, n, sizeof(int), compare);

    int actual = twoSum(S, 5, n);
    assert(actual == 1);

//    actual = k_sum(S, n, 2, 0);
//    assert(actual == 1);
//
//    // k > 2
//    actual = k_sum(S, n, 3, 200);
//    assert(actual == 0);
//
//    actual = k_sum(S, n, 3, 7);
//    assert(actual == 1);
//
//    actual = k_sum(S, n, 3, 0);
//    assert(actual == 0);
//
//    actual = k_sum(S, n, 3, 6);
//    assert(actual == 1);

}

int testBinarySearchOddNumberOfElements() {
    int S[] = {3,2,1,4,5};
    // Calculate the number of elements in the array
    const int n = sizeof(S) / sizeof(S[0]);

    // Use qsort to sort the array
    qsort(S, n, sizeof(int), compare);

    int r = find(S, 0, n-1, 3, n);
    assert(r == 1);

    r = find(S, 0, n-1, 2, n);
    assert(r == 1);

    r = find(S, 0, n-1, 0, n);
    assert(r == 0);

    r = find(S, 0, n-1, 1, n);
    assert(r == 1);

    r = find(S, 0, n-1, 50, n);
    assert(r == 0);

    r = find(S, 0, n-1, 4, n);
    assert(r == 1);

    r = find(S, 0, n-1, 5, n);
    assert(r == 1);
}

int testBinarySearchEvenNumberOfElements() {
    int S[] = {3,2,1,4,5,8};
    // Calculate the number of elements in the array
    int n = sizeof(S) / sizeof(S[0]);

    // Use qsort to sort the array
    qsort(S, n, sizeof(int), compare);

    int r = find(S, 0, n-1, 3, n);
    assert(r == 1);

    r = find(S, 0, n-1, 2, n);
    assert(r == 1);

    r = find(S, 0, n-1, 0, n);
    assert(r == 0);

    r = find(S, 0, n-1, 1, n);
    assert(r == 1);

    r = find(S, 0, n-1, 50, n);
    assert(r == 0);

    r = find(S, 0, n-1, 4, n);
    assert(r == 1);

    r = find(S, 0, n-1, 8, n);
    assert(r == 1);
}

