
#include <stdbool.h>
#include <stdlib.h>

bool makeArray(const int arraySize, int** result, int** value1)
{
    *result = malloc(arraySize * sizeof(int));
    if (*result == NULL) {
        // Handle allocation failure
        *value1 = NULL;
        return false;
    }
    return true;
}
