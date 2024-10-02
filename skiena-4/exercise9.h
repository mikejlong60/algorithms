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


