#include <ctype.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "publisherhashmap.h"
#include <assert.h>


int main() {
    struct Publisher p1 = {.name = "fred1", .bookspublished = 5};
    struct Publisher p2 = {.name = "fred2", .bookspublished = 6};
    struct Publisher p3 = {.name = "fred3", .bookspublished = 7};
    struct Publisher p4 = {.name = "fred4", .bookspublished = 8};
    struct Publisher p5 = {.name = "fred5", .bookspublished = 9};
    struct Publisher p6 = {.name = "fred5", .bookspublished = 10};
    install(p1.name, &p1);
    install(p2.name, &p2);
    install(p3.name, &p3);
    install(p4.name, &p4);
    install(p5.name, &p5);
    install(p6.name, &p6);
    PublisherNlistprint(lookup("fred1"));
    PublisherNlistprint(lookup("fred2"));
    PublisherNlistprint(lookup("fred3"));
    PublisherNlistprint(lookup("fred4"));
    PublisherNlistprint(lookup("fred5"));

    struct PublisherNlist *a = lookup("fred1");
    assert(strcmp("fred1", a->name) == 0);
    assert(a->defn->bookspublished == 5);

    struct PublisherNlist *b = lookup("fred2");
    assert(strcmp("fred2", b->name) == 0);
    assert(b->defn->bookspublished == 6);

    struct PublisherNlist *c = lookup("fred3");
    assert(strcmp("fred3", c->name) == 0);
    assert(c->defn->bookspublished == 7);

    struct PublisherNlist *d = lookup("fred4");
    assert(strcmp("fred4", d->name) == 0);
    assert(d->defn->bookspublished == 8);

    struct PublisherNlist *e = lookup("fred5");
    assert(strcmp("fred5", e->name) == 0);
    assert(e->defn->bookspublished == 10);
}
