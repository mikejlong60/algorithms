#include <stdio.h>
#include <stdlib.h>
#include <string.h>

//Question: You are given a pile of thousands of telephone bills and thousands of checks sent in to pay the bills. Find out who did not pay.
//Answer:
//    Steps:  Big O(n log n) due to the sorting
//       1. //Question: You are given a pile of thousands of telephone bills and thousands of checks sent in to pay the bills. Find out who did not pay.
//Answer:
//    Steps:  Big O(n log n) due to the sorting
//       1. Sort both arrays by name a - checks, b - bills.
//       2. unpaidBills = []UnpaidBills
//       3. for i := 0; i < len(a); i++ {
//            if array a[i] is > b[i] {// There is not a matching check so add it to if a[i].firstName < b[i].firstName && a[i].lastName == b[i].lastName {
//               unpaidBills = append(unpaidBills, b[i])
//            }
//       }


// Comparison function for qsort
int compareStrings(const void *a, const void *b) {
    // Convert the pointers to the correct type
    const char **strA = (const char **)a;
    const char **strB = (const char **)b;
    // Use strcmp to compare the strings
    return strcmp(*strA, *strB);
}

int main() {
    // array of bills
    char *a[] = {"fred2", "fred3", "fred", "fred4", "fred5"};

    // array of checks
    const char *b[] = {"fred2", "fred2","fred5", "fred3", "fred4"};

    // Dynamically allocate memory for the array of unpaid bills as pointers
    char **unpaidbills = calloc(7, sizeof(char *));

    if (unpaidbills == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        return 1;
    }

    // Calculate the number of elements in the array
    int na = sizeof(a) / sizeof(a[0]);
    int nb = sizeof(b) / sizeof(b[0]);

    // Sort the array using qsort
    qsort(a, na, sizeof(char *), compareStrings);
    qsort(b, nb, sizeof(char *), compareStrings);

    // Print the sorted array
    for (int i, j = 0; i < na; i++) {
        int result = strcmp(a[i], b[i]);
        printf("result:a[%d] --- b[%d] --- result:%d\n", a[i], b[i], result);
        if (result != 0) { // There is not a matching check so add it to list of unpaid bills
            unpaidbills[j] = a[i];
            j++;
        }
    }

   // Print the unpaid bills array
    for (int i = 0; i < 7; i++) {
        printf("unpaidbills[%d]:%s\n", i, unpaidbills[i]);
    }
    free(unpaidbills);
    return 0;
}
