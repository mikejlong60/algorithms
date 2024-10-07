#include <stdlib.h>
#include <assert.h>
#include "inthashmap.h"

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

int moreThanK(int* xs, const int T, const int arraySize)
{
    for (int j = 0; j < arraySize; j++)
    {
        install(xs[j], xs[j]);
    }
    return 0;
}


int testMoreThanK()
{
    int xs[] = {3, 2, 1, 4, 5, 6, 63, 12, 23, 34, 19, 19, 23};
    // Calculate the number of elements in the array
    int n = sizeof(xs) / sizeof(xs[0]);

    // Use qsort to sort the array
    int actual = moreThanK(xs, 1, n);

    assert(actual == 0);
}
