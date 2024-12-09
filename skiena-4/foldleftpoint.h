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

bool makePointCountArray(const int arraySize, PointCount** result, PointCount** value1)
{
    *result = malloc(arraySize * sizeof(int));
    if (*result == NULL) {
        // Handle allocation failure
        *value1 = NULL;
        return false;
    }
    return true;
}

int FoldLeft(PointCount* array, int size, int accumulator, int (*func)(int, PointCount)) {
    if (size == 0) {
        return accumulator;
    }
    // Process the first element and recursively call for the rest
    return FoldLeft(array + 1, size - 1, func(accumulator, array[0]), func);
}

// Example callback function: sum values in the array
int sumValues(int accumulator, PointCount element) {
    return accumulator + element.c;
}

// Example callback function: find max value in the array
int maxValues(int accumulator, PointCount element) {
    return accumulator > element.c ? accumulator : element.c;
}
struct Interval {
    int l;
    int r;
};

struct PointCount* mostFrequentPoint(const struct Interval* S, const int n)  {
    struct PointCount* result;
    struct PointCount* value1;
    if (!makePointCountArray(n * 2, &result, &value1)) return value1;


    for (int a = 0; a < n; a++) {
        for (int b = 0; b < n; b++)
        {
            if (S[b].l >= S[a].l && S[b].r > S[a].r) {
                //increment b.l point total
                printf("\n1 - S[b].l point total:%d ",S[b].l);
            } else if(S[b].l >= S[a].l && S[b].r <= S[a].r) {
                //increment b.l point total
                printf("\n2 - S[b].l point total:%d ",S[b].l);
            } else if (S[b].l < S[a].l && S[b].r <= S[a].r) {
                //increment b.r point total
                printf("\n3 - S[b].r point total:%d ",S[b].r);
            } else if (S[b].l > S[a].r) {
                //can't increment anything, no overlap
                printf("\n4 - no overlap");

            } else if (S[b].l < S[a].l && S[b].r > S[a].r) {
                //increment a.l point total
                printf("\n5 - S[a].l point total:%d ",S[b].l);
            }
        }
    }

    return  FoldLeft(result, n, 0, maxValues);
}

int testMostFrequentPoint() {
    struct Interval S[] = {
        {10, 40},
        {20, 60},
        {50, 90},
        {15, 70}
    };
         // Calculate the number of elements in the array
    int n = sizeof(S) / sizeof(S[0]);

    // Use qsort to sort the array
    //qsort(S, n, sizeof(struct Interval), compareIntervals);

    char result = mostFrequentPoint(S, n);
    assert(result == 50);
    //assert(result.r == 90);
}
