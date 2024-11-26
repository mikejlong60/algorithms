#include <assert.h>
#include <stdlib.h>

/**
Question: You are given a set S of n intervals on a line, with the ith interval described by its left and right endpoints(li, ri).
Give a O(n log n) algorithm to identify a point p on the line which is in the largest number of intervals.
As an example, for S = {(10,40), (20,60), (50,90), (15,70)}, no point exists in all 4 intervals.
But 50 exists in 3 intervals. You can assume an endpoint counts as being in its own interval(inclusive).

Algorithm Efficiency -  O(n^2) -- it sucks I know.

*/

struct Max {
	int l;
  	int r;
  	int numOverlaps;
};

struct Interval {
    int l;
    int r;
};

char mostFrequentPoint(const struct Interval* S, const int n)  {
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

    return  "fred";
}
  // Comparison function for intervals
int compareIntervals(const void *a, const void *b) {
    // Cast the pointers to Interval pointers and dereference them
    const struct Interval *int_a = (struct Interval *)a;
    const struct Interval *int_b = (struct Interval *)b;

    // Return the difference between the two integers
    return int_a->l - int_b->l;
}

int testMostFrequentPoint() {
    struct Interval suck[] = {
        {10, 40},
        {15, 70},
        {20, 60},
        {50, 90},
    };
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
