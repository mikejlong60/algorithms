#include <stdlib.h>
#include <assert.h>

/**
Question: Give O(n) algorithm that, given a list of elements(xs), finds all the elements that appear more than n/k times in the list.
  Then do the same for all the elements tha appear more than n/4 times in the list.

Pseudocode:
Loop over array xs
  lookup x in hashmap
  if exists
     increment counter
     if counter >= k
        add x to result array
  else
     add x to hashmap and set counter to 1
end loop
return result array

*/

//Binary search
int afind(int* S, const int leftOffset, const int rightOffset, const int lookingFor, const int arraySize)
{
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
    } else if (rightOffset - leftOffset == 1) {
        return 0;
    } else if (S[relativeA] < lookingFor) {
        return find(S, leftOffset + a, rightOffset, lookingFor, arraySize);
    } else
        return find(S, leftOffset, leftOffset + a,  lookingFor, arraySize);
}

///Cost of O(n log n).  Might not be good enough. Book says you can do this in O(n). O(log n) is cost of binary search when array is sorted.
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
    int S[] = {3,2,1,4,5,6, 63, 12, 23, 34, -19, 19, 23};
    // Calculate the number of elements in the array
    int n = sizeof(S) / sizeof(S[0]);

    // Use qsort to sort the array
    qsort(S, n, sizeof(int), compare);

    int actual = twoSum(S, 5, n);
    assert(actual == 1);

    actual = twoSum(S, 2, n);
    assert(actual == 1);

    actual = twoSum(S, 75, n);
    assert(actual == 1);

    actual = twoSum(S, 112, n);
    assert(actual == 0);

    actual = twoSum(S, -112, n);
    assert(actual == 0);
}

int testBinarySearchOddNumberOfElements() {
    int S[] = {3,2,1,4,5,12,32,-200, 37};
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
    int S[] = {3,2,1,4,5,8,-32,34,12,203};
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

