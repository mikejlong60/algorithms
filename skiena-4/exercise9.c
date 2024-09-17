#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
/**
Question: Given a set S of n positive(simplifies for now) integers and an integer T, give an O(n(k-1) log n)
algorithm to test whether k of three distinct integers in S add up to T

Answer:
    1. Sort the set of integers
    2. Find the closest spot j in S using a binary search where s[j] < T. Example
        Example data set: S = {1,2,3,4,5}. k = 3, T = 8
        2.1 Find closest spot in array with binary search less than 8(T).  It's 5, position 4.
        2.2 Work backwards 2 spots from 5 since you have three numbers you have to add and you cannot repeat usage
        of the same number.
        2.3 recursively function f-
            2.3.1 start at position 3.  Add position 2 to working sum(ws). Does ws exceed T?
                Yes - stop algorithm - no such set of k integers exits
                No - call f with


func f(int[] S, int T, int k, int workingIdx, int workingSum) bool {


}
*/

// Comparison function for integers
int compare(const void *a, const void *b) {
    // Cast the pointers to int pointers and dereference them
    int int_a = *(int *)a;
    int int_b = *(int *)b;

    // Return the difference between the two integers
    return (int_a - int_b);
}

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
    printf("S[relativeA] = %d\n", S[relativeA]);
    if (S[relativeA] == lookingFor) {
        return 1;
    } else if (S[relativeA] < lookingFor) {
        return find(S, leftOffset + a, rightOffset, lookingFor, arraySize);
    } else
        return find(S, leftOffset, leftOffset + a,  lookingFor, arraySize);
}


int main() { // Example array of integers
    int S[] = {3,2,1,4,5};
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

    r = find(S, 0, n-1, 5, n);
    assert(r == 1);

    return r;
}
