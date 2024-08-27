#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "hashmap.h"


//Question: You are given a printed list containing the title, author, call number, and publisher of all the books in a school
//  library and another list of thirty publishers. Find out how many of the books in the library were published by each company.
//Answer:
//    Steps:  Big O(N)
//       1. booksPerPublisher = map[string]int{}
//             Populate
//       2. for i := 0; i < len(b); i++ {
//            lookup publisher in booksPerPublisher
//            if OK {
//                increment book count
//            } else {
//                Add new entry for publisher initialized to 1
//            }
//       }

struct book {
    char title[50];
    char publisher[50];
};


int main() {
    // array of publishers
    char *publishers[] = {"dang0", "dang1","dang2","dang3","dang4","dang5","dang6","dang7","dang8","dang9","dang10","dang11",
        "dang12","dang13","dang14","dang15","dang16","dang17","dang18","dang19","dang20","dang21","dang22","dang23",
        "dang24","dang25","dang26","dang27","dang28","dang29"};
    struct book books[10];
    strcpy(books[0].title, "Alice");
    strcpy(books[0].publisher, "dang0");
    strcpy(books[1].title, "Alice1");
    strcpy(books[1].publisher, "dang3");
    strcpy(books[2].title, "Alice3");
    strcpy(books[2].publisher, "dang0");
    strcpy(books[3].title, "Alice3");
    strcpy(books[3].publisher, "dang0");
    strcpy(books[4].title, "Alice4");
    strcpy(books[4].publisher, "dang5");
    strcpy(books[5].title, "Alice5");
    strcpy(books[5].publisher, "dang0");
    strcpy(books[6].title, "Alice6");
    strcpy(books[6].publisher, "dang0");
    strcpy(books[7].title, "Alice7");
    strcpy(books[7].publisher, "dang9");
    strcpy(books[8].title, "Alice8");
    strcpy(books[8].publisher, "dang9");
    strcpy(books[9].title, "Alice9");
    strcpy(books[9].publisher, "dang0");


    // array of checks
    char *cheques[] = {"fred2", "otis", "fred2","fred5", "fred3", "fred4"};

    // Calculate the number of elements in the array
    int na = sizeof(books) / sizeof(books[0]);
    int nb = sizeof(cheques) / sizeof(cheques[0]);
    printf("number of bills:%d\n", na);
    printf("number of cheques:%d\n", nb);

    // Dynamically allocate memory for the array of unpaid bills as pointers
    // It will not be larger than array of unpaid bills.
    char **booksperpublisher = calloc(na, sizeof(howmany *));

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
