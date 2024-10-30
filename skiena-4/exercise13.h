#include <assert.h>
#include <stdlib.h>


/**
Question: A camera at the door tracks the entry time a, and the exit time b,
(assume bi < ai) for each of n persons p, attending a party.  Give a O(n log n)
algorithm that analyzes this data to determine the time when the most people were
at the party. Assume that all party entry and exit times are distinct.


Pseudocode:
Loop over array of integers that consists of 1 and -1.  1 represents entering the party. -1 represents leaving the party.
Keep a global variable that is the max and the current time(a go timestamp nano).
Just add the number 1 or -1 to the current max inside the loop. Reset the current max every time it grows higher and record the timestamp.
It's like your parent's thermometer that records the max temperature for a given period.

This algorithm is O(n).
*/


typedef struct {
    int max;
    int time;
} MostPeople;

MostPeople biggestCrowd(const int* people, const int peopleSize) {
    MostPeople result;
    result.max = 0;
    result.time = 0;

    int i = 0;
    int currentMax = 0;
    while (i < peopleSize)
    {
        currentMax += people[i];
        if (currentMax > result.max) {
            result.max = currentMax;
            result.time = i;
        }
        i++;
    }
    return result;
}

int testBiggestCrowd() {
    const int people[] = {1,-1, 1,1,1,1,-1,-1,-1,1,1,1,1,-1};
    int peopleSize = sizeof(people) / sizeof(people[0]);

    MostPeople result = biggestCrowd(people, peopleSize);
    assert(result.max == 5);
    assert(result.time == 12);
}
