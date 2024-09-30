#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
/**
Question: Given a set S of n integers and an integer T, give an O(n(k-1) log n)
algorithm to test whether k integers in S add up to T

*/

// Comparison function for integers
int compare(const void *a, const void *b) {
    // Cast the pointers to int pointers and dereference them
    const int int_a = *(int *)a;
    const int int_b = *(int *)b;

    // Return the difference between the two integers
    return int_a - int_b;
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
    if (S[relativeA] == lookingFor) {
        return 1;
    } else if (S[relativeA] < lookingFor) {
        return find(S, leftOffset + a, rightOffset, lookingFor, arraySize);
    } else
        return find(S, leftOffset, leftOffset + a,  lookingFor, arraySize);
}

int twoSum(int *S, const int k, const int T, const int arraySize) {//TODO make the algorithm work for k > 2
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

int hasSum(int *S, const int k, const int T, const int arraySize) {
    if (k == 1)
        return find(S, 0, arraySize-1, T, arraySize);
    if (k ==2) {
        return twoSum(S, k, T, arraySize);
    }
    for (int i = 0; i < arraySize; i++) {
        if (i == 0) {// || S[i] != S[i-1]) {
            //int pppp = i + 1;
            //int* piss = S[pppp];
            if (hasSum(&S[i + 1], k-1, T - S[i], arraySize)) {
                return 1;
            }
        }
    }
    return 0;
}

// Function to solve the 2-sum problem
int two_sum(int *S, int left, int right, int target) {
    while (left < right) {
        int sum = S[left] + S[right];
        if (sum == target)
            return 1;  // Found the pair
        else if (sum < target)
            left++;
        else
            right--;
    }
    return 0;  // No pair found
}

// Function to solve the k-sum problem recursively
int k_sum(int *S, int n, int k, int target) {
    if (k == 2) {
        // Base case: 2-sum problem
        return two_sum(S, 0, n - 1, target);
    }

    // Recursive case: reduce k-sum to (k-1)-sum
    for (int i = 0; i < n; i++) {
        // Check if (k-1) integers sum up to target - S[i]
        if (k_sum(S + i + 1, n - i - 1, k - 1, target - S[i])) {
            return 1;  // Found k integers that sum up to the target
        }
    }

    return 0;  // No solution found
}


int testKSum() {
    int S[] = {3,2,1,4,5,6, 63, 12, 23, 34, -19, 19};
    // Calculate the number of elements in the array
    int n = sizeof(S) / sizeof(S[0]);

    // Use qsort to sort the array
    qsort(S, n, sizeof(int), compare);

    int actual = k_sum(S, n, 2, 5);
    assert(actual == 1);

    actual = k_sum(S, n, 2, 0);
    assert(actual == 1);

    // k > 2
    actual = k_sum(S, n, 3, 200);
    assert(actual == 0);

    actual = k_sum(S, n, 3, 7);
    assert(actual == 1);

    actual = k_sum(S, n, 3, 0);
    assert(actual == 0);

    actual = k_sum(S, n, 3, 6);
    assert(actual == 1);

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

int main() { // Example array of integers
    testBinarySearchOddNumberOfElements();
    testBinarySearchEvenNumberOfElements();
    testKSum();
}
