#include <ctype.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct PublisherNlist { //table entry
    struct PublisherNlist *next; //nest entry in liked list chain
    char *name; //defined name
    char *defn; //replacement text
};

void PublisherNlistprint(struct PublisherNlist *);
void PublisherNlistprint(struct PublisherNlist *p) {
    if (p != NULL) {
        PublisherNlistprint(p->next);
    }
}

#define HASHSIZE 3000
#define NUMELMTS 200000


static struct PublisherNlist *hashtab[HASHSIZE]; //pointer table

// hash - form hash value for string s
unsigned hash(char *s) {
    unsigned hashval;

    for (hashval = 0; *s != '\0'; s++)
        hashval = *s + 31 * hashval;
    
    return hashval % HASHSIZE;
}

//lookup
struct PublisherNlist *lookup(char *s) {
    struct PublisherNlist *np;
    for (np = hashtab[hash(s)]; np != NULL; np = np->next) {
        if (strcmp(s, np->name) == 0) {
            return np;  //found
        }
    }
    return NULL;  //not found
}

// install - put(name defn) in hashtab
struct PublisherNlist *install(char *name, char *defn) {
    struct PublisherNlist *np;
    unsigned hashval;
    if ((np = lookup(name)) == NULL) { //not found
        np = (struct PublisherNlist *) malloc(sizeof(*np));
        if (np == NULL || (np->name = strdup(name)) == NULL)
            return NULL;
        hashval = hash(name);
        np->next = hashtab[hashval];
        hashtab[hashval] = np;
    } else { //already there
        free((void *) np->defn); //free previous defn
    }
    if ((np->defn = strdup(defn)) == NULL) {
        return NULL;
    }
    return np;
}
