#include <ctype.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct Publisher {
    char name[50];
    int bookspublished;
};

struct PublisherNlist { //table entry
    struct PublisherNlist *next; //nest entry in liked list chain
    char *name; //defined name
    struct Publisher *defn;
};

void PublisherNlistprint(const struct PublisherNlist *p) {
    if (p != NULL) {
        PublisherNlistprint(p->next);
        printf("Next guy in linked list - %s : %d\n", p->name, p->defn->bookspublished);
    } else {
        printf("p is nil\n");
    }
    printf("leaving Nlistprint\n");
}

#define HASHSIZE 6
#define NUMELMTS 5


static struct PublisherNlist *hashtab[HASHSIZE]; //pointer table

// hash - form hash value for string s
unsigned hash(const char *s) {
    unsigned hashval;

    for (hashval = 0; *s != '\0'; s++)
        hashval = *s + 31 * hashval;

    return hashval % HASHSIZE;
}

//lookup
struct PublisherNlist *lookup(const char *s) {
    struct PublisherNlist *np;
    printf("lookup1\n");
    for (np = hashtab[hash(s)]; np != NULL; np = np->next) {
        printf("searching linked list in cell[%d]\n", hash(s));
        if (strcmp(s, np->name) == 0) {
            printf("found[%s] in linked list\n", np->name);
            return np;  //found
        }
    }
    printf("lookup2\n");
    return NULL;  //not found
}

// install - put(name defn) in hashtab
struct PublisherNlist *install(const char *name, struct Publisher *defn) {
    //printf("install1:name[%s] --- defn[%s]\n", name, defn);
    struct PublisherNlist *np;
    printf("install111 name[%s]\n", name);
    if ((np = lookup(name)) == NULL) { //not found
        printf("install2\n");
        np = (struct PublisherNlist *) malloc(sizeof(*np));
        if (np == NULL || (np->name = strdup(name)) == NULL)
            return NULL;
        unsigned hashval = hash(name);
        np->next = hashtab[hashval];
        printf("hashval[%d]\n", hashval);
        hashtab[hashval] = np;
    } else { //already there
        printf("install3 publisher is already there \n");
        np->defn = NULL;
        free((void *) np->defn); //free previous defn
        printf("install4\n");
    }
    printf("install5\n");
    if ((np->defn = defn) == NULL) {
        printf("install6\n");
        return NULL;
    }
    return np;
}


main() {
   //*printf("1\n");
    struct Publisher p1 = {.name = "fred1", .bookspublished = 0};
    struct Publisher p2 = {.name = "fred2", .bookspublished = 0};
    struct Publisher p3 = {.name = "fred3", .bookspublished = 0};
    struct Publisher p4 = {.name = "fred4", .bookspublished = 0};
    struct Publisher p5 = {.name = "fred5", .bookspublished = 0};
    struct Publisher p6 = {.name = "fred5", .bookspublished = 10};
    install(p1.name, &p1);
    install(p2.name, &p2);
    install(p3.name, &p3);
    install(p4.name, &p4);
    install(p5.name, &p5);
    install(p6.name, &p6);
    //for (int i = 0; i < NUMELMTS; i++) {
    //    char key[20];
    //    sprintf(key, "MAX%d", i);
    //    char val[30];
    //    sprintf(val, "fred%d", i);
    //printf("2\n");
    //    struct Nlist *x = install(key, val);
    //    printf("3\n");
    //}
    for (int i = 0; i < NUMELMTS; i++) {
        struct PublisherNlist *n = hashtab[i];
       printf("in loop name[%s] number[%d]\n", n->name, n->defn->bookspublished);
        PublisherNlistprint(n);
    }
}
