# Problem - There are 25 horses that can race 5 at a time. Find the minimum number of races it takes to determine the first three fastest horses.
## Assumptions:
  1.  You do not know the fastest three horses in advance.  And the horses do not get tired. They will always race at the same rate.
  2. In the races the only time you know is the time of the winner, no other horse's time.
  3. So you must race them together. Sooo.

## Solution:
  1. Race them in groups of 5. That takes 5 races.  And from those races you have five winners. 5 races so far.
  2. Race the 5 winners together.  6 races so far.  That horse is fastest
  3. Remove the winner from that race and race the four remaining horses.  7 races so far. That winner is second fastest.
  4. Remove the winner from that rae and race the three remaining horses.  8 races so far. That winner is third fastest.