#include <stdlib.h>
#include <assert.h>
#include <stdbool.h>

#include "inthashmap.h"
//#include "util.h"

/**
Question: Give O(n) algorithm that, given a list of elements(xs), finds all the elements that appear more than n/2 times in the list.
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

int* moreThanK(const int* xs, const int k, const int arraySize) {
    const int moreThan = arraySize/k;

    int* result;
    int* value1;
    if (!makeArray(arraySize, &result, &value1)) return value1;

    for (int j = 0, i = 0; j < arraySize; j++) {
        const struct Nlist* a = lookup(xs[j]);//lookup integer at position j in hashmap
        if (a == NULL) {  //key not there so add a new entry to hashmap
            install(xs[j], 1);
        } else {  //found key
            if (a->value > moreThan) {
                result[i] = a->key;
                i++;
            }
            //update hashmap to increment number of appearances of int in xs array
            install(xs[j], a->value + 1);
        }
    }
    return result;
}


int testMoreThanK()
{
    const int xs[] = {3, 3,3,3,1,1,1,1};
    // Calculate the number of elements in the array
    int arraySize = sizeof(xs) / sizeof(xs[0]);

    int *result = moreThanK(xs, 3, arraySize);
    printf("result[0]:%d\n", result[0]);
    printf("result[1]:%d\n", result[1]);
    assert(result[0] == 3);
    assert(result[1] == 1);

    free(result);
}
