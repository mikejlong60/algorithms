#include <assert.h>
#include <stdlib.h>

/**
Question: You are given a set S of n intervals on a line, with the ith interval described by its left and right endpoints(li, ri).
Give a O(n log n) algorithm to identify a point p on the line which is in the largets number of intervals.
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

