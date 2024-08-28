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
        printf("Next guy in linked list - %s : %d\n", p->name, p->defn->bookspublished);
        PublisherNlistprint(p->next);
    }
}

#define HASHSIZE 10000

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
    for (np = hashtab[hash(s)]; np != NULL; np = np->next) {
        if (strcmp(s, np->name) == 0) {
            return np;  //found
        }
    }
    return NULL;  //not found
}

// install - put(name defn) in hashtab
struct PublisherNlist *install(const char *name, struct Publisher *defn) {
    //printf("install1:name[%s] --- defn[%s]\n", name, defn);
    struct PublisherNlist *np;
    if ((np = lookup(name)) == NULL) { //not found
        np = (struct PublisherNlist *) malloc(sizeof(*np));
        if (np == NULL || (np->name = strdup(name)) == NULL)
            return NULL;
        unsigned hashval = hash(name);
        np->next = hashtab[hashval];
        hashtab[hashval] = np;
    } else { //already there
        np->defn = NULL;
        free((void *) np->defn); //free previous defn
    }
    if ((np->defn = defn) == NULL) {
        return NULL;
    }
    return np;
}
