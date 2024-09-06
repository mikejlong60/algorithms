#include <ctype.h>
#include <stdio.h>
#include <string.h>
#include "hashmap.h"
#include <assert.h>

//Question: You are given all the book checkout cards used in the campus library during the past year, each of which contains
//   the name of the person who took out the book. Determine how many distinct people checked out at least one book.
//   Steps:  Big O(N)
//       1. distinctPeople = map[string]interface{}{}
//2. for i := 0; i < len(cards); i++ {
//    lookup person in distinctPeople
//    if !OK {
//        Add new entry for person
//    }
//}
//return len(distinctPeople)

int main() {
    // array of checkout cards with persons name who checked the book out.
    char *books[] = {"fred1","fred1","fred1","fred1","fred1","fred1",
        "fred1","fred2","fred1","fred1","fred1","fred1","fred2","fred5",
        "fred5","fred1","fred5","fred1","fred5","fred3","fred1","fred1","fred3",
        "fred5","fred3","fred1","fred1","fred4","fred3","fred1","fred4","fred1" };

    // Calculate the number of elements in the array
    int na = sizeof(books) / sizeof(books[0]);//TODO
    printf("number of books:%d\n", na);

    int x = 0;
    // Add books to hashmap
    for (int i = 0; i < na; i++) {
        struct PublisherNlist *np;
        if ((np = lookup(books[i])) == NULL) {
            install(books[i], books[i]);
            x++;
        }
    }
    printf("number of distinct people who checked out books:%d\n", x);
    assert(x == 5);
}