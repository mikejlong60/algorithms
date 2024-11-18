#include <assert.h>
#include <stdlib.h>

/**
Question: You are given a set S of n intervals on a line, with the ith interval described by its left and right endpoints(li, ri).
Give a O(n log n) algorithm to identify a point p on the line which is in the largest number of intervals.
As an example, for S = {(10,40), (20,60), (50,90), (15,70)}, no point exists in all 4 intervals.
But 50 exists in 3 intervals. You can assume an endpoint counts as being in its own interval(inclusive).

Algorithm Efficiency - 

S sorted by l: S = {(10,40), (15, 70), (20,60), (50, 90)}
S sorted by r: S = {(10,40), (20,60), (15, 70), (50, 90)}
S l sorted: Sl = {10, 15, 20, 50}
S r sorted: Sl = {40, 60, 70, 90}


Algorithm:
struct currentMax {
   l int
   r int
   numOverlaps int
}
1. sort S by l
2. Set resultMax = currentMax{}
3. resultMax.l = S[0].l;
4. resultMax.r = s[0].r;
5. resultMax.numOverlaps = 1;
6. for int i := 1; i < len(s); i++ {
   maybeMax = currentMax{};
   if S[i].l <= currentMax.r && currentMax{
      currentMax.numOverlaps++;
   } else {
      currentMax.l = S[i].l
      currentMax.r = S[i].r
      currentMax.numOverlaps = 1
   }
}
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

struct Max mostFrequentPoint(const struct Interval* S, const int n) {
    struct Max currentOverlap; //the interval that has an overlap with subsequent intervals
    currentOverlap.l = 0;
    currentOverlap.r = 0;
    currentOverlap.numOverlaps = 0;
    struct Max maxOverlap;//the current maximum overlap
    maxOverlap.l = 0;
    maxOverlap.r = 0;
    maxOverlap.numOverlaps = 0;
    for (int i = 0; i < n; i++) {
        if (i == 0) {//First element
            currentOverlap.l = S[i].l;
            currentOverlap.r = S[i].r;
            currentOverlap.numOverlaps = 1;
            maxOverlap.l = S[i].l;
            maxOverlap.r = S[i].r;
            maxOverlap.numOverlaps = 1;
        } else
        {
            //Not at first element
            if (S[i].l <= currentOverlap.r) {//Is there an overlap with next element from currentOverlap
                currentOverlap.numOverlaps += 1;
            } else {// begin new overlap calculation
                if (currentOverlap.numOverlaps > maxOverlap.numOverlaps) {//new max overlap
                    maxOverlap = currentOverlap;
                }
                currentOverlap.l = S[i].l;
                currentOverlap.r = S[i].r;
                currentOverlap.numOverlaps = 1;
            }
        }
    }

    return  maxOverlap;
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
    qsort(S, n, sizeof(struct Interval), compareIntervals);

    struct Max result = mostFrequentPoint(S, n);
    assert(result.l == 50);
    assert(result.r == 90);
}
