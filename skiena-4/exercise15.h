#include <assert.h>
#include <stdlib.h>

/**
Question: Given a list I of n intervals, specified as (xi, yi) pairs, return a list where the overlapping
intervals are merged into a single pair. For I ={(1,3), (2,6),(8,10), (7,18)}, the correct answer is:
{(1,6), (7,18)}.  Algorithm should be no worse than O(n log n).  Invariants are:
1. Array has even number of elements
2. Adjacent pairs of elements are merged into a single element.

Algorithm Efficiency - O(n).
*/

bool makeArray(const int arraySize, int** result, int** value1);

struct Pair {
    int a;
    int b;
};

struct Pair* mergeIntervals(struct Pair* pairs, const int arraySize) {
	struct Pair* result;
	struct Pair* value1;
    if (!makeArray(arraySize, &result, &value1)) return value1;

    for (int i = 0, j = 1, k = 0; i < arraySize; k++) {
    	if (pairs[i].a < pairs[j].a && pairs[i].b < pairs[j].b) {
        	result[k].a = pairs[i].a;
            result[k].b = pairs[j].b;
        } else if (pairs[i].a > pairs[j].a && pairs[i].b < pairs[j].b) {
        	result[k].a = pairs[j].a;
            result[k].b = pairs[j].b;
        } else if (pairs[i].a < pairs[j].a && pairs[i].b > pairs[j].b) {
        	result[k].a = pairs[i].a;
            result[k].b = pairs[i].b;
		} else if (pairs[i].a > pairs[j].a && pairs[i].b < pairs[j].b) {
        	result[k].a = pairs[j].a;
            result[k].b = pairs[j].b;
        }
		j = j + 2;
		i = i + 2;
    }
    return result;
}

int testMergeIntervals() {
	struct Pair pairs[] = {
        {1, 3},
        {2, 6},
        {8, 10},
        {7, 18}
    };

    // Calculate the number of elements in the array
    int arraySize = sizeof(pairs) / sizeof(pairs[0]);


    struct Pair *result = mergeIntervals(pairs, arraySize);
    assert(result[0].a == 1);
    assert(result[0].b == 6);
    assert(result[1].a == 7);
    assert(result[1].b == 18);

    free(result);
}

