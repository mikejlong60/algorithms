#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "hashmap.h"


//TODO use utash hashtable implementation instead of ou own: https://troydhanson.github.io/uthash/userguide.html
//Question: You are given a pile of thousands of telephone bills and thousands of checks sent in to pay the bills. Find out who did not pay.
//Answer:
//    Steps:  Big O(n)  -- you iterate over each array once, first is to put cheques into map, second is to see if there is cheque for bill in hahmap.
//       1. //Question: You are given a pile of thousands of telephone bills and thousands of checks sent in to pay the bills. Find out who did not pay.
//Answer:
//    Steps:  you have two arrays bills and cheques
//       1. Add cheques array to hashmap.
//       2. unpaidBills = []UnpaidBills
//       3. for i := 0; i < len(bills); i++ {
//            lookup person from bills in checks hashmap.
//            if person is not in checks map {
//               unpaidBills = append(unpaidBills, bills[i])
//            }
//       }
//       return unpaidbills

int main() {
    // array of bills
    char *bills[] = {"dang", "mike","fred2", "dang", "mike","fred2", "dang", "mike","fred2", "dang", "mike","fred2", "dang", "mike","fred2", "dang", "mike","fred2", "dang", "mike","fred2", "dang", "mike","fred2", "dang", "mike","fred2","fred2", "dang", "mike","fred2","dang", "mike","fred2", "otis", "fred3", "fred3", "fred", "fred4", "fred5","fred0","fred2"};

    // array of checks
    char *cheques[] = {"fred2", "otis", "fred2","fred5", "fred3", "fred4"};

    // Calculate the number of elements in the array
    int na = sizeof(bills) / sizeof(bills[0]);
    int nb = sizeof(cheques) / sizeof(cheques[0]);
    printf("number of bills:%d\n", na);
    printf("number of cheques:%d\n", nb);

    // Dynamically allocate memory for the array of unpaid bills as pointers
    // It will not be larger than array of unpaid bills.
    char **unpaidbills = calloc(na, sizeof(char *));

    if (unpaidbills == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        return 1;
    }

   // Add cheques to hashmap
    for (int i = 0; i < nb; i++) {
        install(cheques[i], cheques[i]);
        printf("cheques[%d]:%s\n", i, cheques[i]);
    }

    //add unpaid bills to array
    for (int i = 0, j = 0; i < na; i++) {
        printf("bills[%d]:%s\n", i, bills[i]);
        struct PublisherNlist *np;
        if ((np = lookup(bills[i])) == NULL) {
            unpaidbills[j] = bills[i];
            j++;
        }
    }

   // Print the unpaid bills array
    printf("number of unpaid bills, some are NULL and we won't print those:%d\n", na);
    for (int j = 0; j < na; j++) {
        if (unpaidbills[j] != NULL)
            printf("unaidbills[%d]:%s\n", j, unpaidbills[j]);
    }
    free(unpaidbills);

    return 0;
}
