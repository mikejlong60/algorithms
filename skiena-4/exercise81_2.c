#include <ctype.h>
#include <stdio.h>
#include <string.h>
#include "publisherhashmap.h"
#include <assert.h>

// Question: You are given a printed list containing the title, author, call number, and publisher of all the books in a school
//   library and another list of thirty publishers. Find out how many of the books in the library were published by each company.
// Answer:
//     Steps:  Big O(N)
//        1. booksPerPublisher = map[string]int{}
// 2. for i := 0; i < len(b); i++ {
//     lookup publisher in booksPerPublisher
//     if OK {
//         increment book count
//     } else {
//         Add new entry for publisher initialized to 1
//     }
// }


int main() {

    // array of books only listing their publisher for simplicity
    char *books[] = {"fred1","fred1","fred1","fred1","fred1","fred1",
        "fred1","fred2","fred1","fred1","fred1","fred1","fred2","fred5",
        "fred5","fred1","fred5","fred1","fred5","fred3","fred1","fred1","fred3",
        "fred5","fred3","fred1","fred1","fred4","fred3","fred1","fred4","fred1" };

    // Calculate the number of elements in the array
    int na = sizeof(books) / sizeof(books[0]);
    printf("number of books:%d\n", na);

    struct Publisher p1 = {.name = "fred1", .bookspublished = 0};
    struct Publisher p2 = {.name = "fred2", .bookspublished = 0};
    struct Publisher p3 = {.name = "fred3", .bookspublished = 0};
    struct Publisher p4 = {.name = "fred4", .bookspublished = 0};
    struct Publisher p5 = {.name = "fred5", .bookspublished = 0};

    install(p1.name, &p1);
    install(p2.name, &p2);
    install(p3.name, &p3);
    install(p4.name, &p4);
    install(p5.name, &p5);


    //increment book count per publisher
    for (int i = 0; i < na; i++) {
        struct PublisherNlist *b = lookup(books[i]);
        b->defn->bookspublished = b->defn->bookspublished + 1;
        install(b->name, b->defn);
        printf("books[%d]:%s\n", i, books[i]);
    }

    struct PublisherNlist *a = lookup("fred1");
    PublisherNlistprint(a);
    assert(strcmp("fred1", a->name) == 0);
    assert(a->defn->bookspublished == 19);

    a = lookup("fred2");
    PublisherNlistprint(a);
    assert(strcmp("fred2", a->name) == 0);
    assert(a->defn->bookspublished == 2);

    a = lookup("fred3");
    PublisherNlistprint(a);
    assert(strcmp("fred3", a->name) == 0);
    assert(a->defn->bookspublished == 4);
    a = lookup("fred4");
    PublisherNlistprint(a);
    assert(strcmp("fred4", a->name) == 0);
    assert(a->defn->bookspublished == 2);
    a = lookup("fred5");
    PublisherNlistprint(a);
    assert(strcmp("fred5", a->name) == 0);
    assert(a->defn->bookspublished == 5);
}