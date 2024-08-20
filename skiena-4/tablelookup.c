#include <ctype.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct Nlist { //table entry
    struct Nlist *next; //nest entry in liked list chain
    char *name; //defined name
    char *defn; //replacement text
};

void Nlistprint(struct Nlist *);
void Nlistprint(struct Nlist *p) {
    if (p != NULL) {
        Nlistprint(p->next);
        printf("Next guy in linked list - %s:%s\n", p->name, p->defn);
    }
    printf("leaving Nlistprint\n");
}

#define HASHSIZE 3000
#define NUMELMTS 200000


static struct Nlist *hashtab[HASHSIZE]; //pointer table

// hash - form hash value for string s
unsigned hash(char *s) {
    unsigned hashval;

    for (hashval = 0; *s != '\0'; s++)
        hashval = *s + 31 * hashval;
    
    return hashval % HASHSIZE;
}

//lookup
struct Nlist *lookup(char *s) {
    struct Nlist *np;
    printf("lookup1\n");
    for (np = hashtab[hash(s)]; np != NULL; np = np->next) {
        printf("searching linked list in cell[%d]\n", hash(s));
        if (strcmp(s, np->name) == 0) {
            printf("found in linked list\n");
            return np;  //found
        }
    }
     printf("lookup2\n");
    return NULL;  //not found
}

// install - put(name defn) in hashtab
struct Nlist *install(char *name, char *defn) {
    printf("install1:name[%s] --- defn[%s]\n", name, defn);
    struct Nlist *np;
    unsigned hashval;
    printf("install111\n");
    if ((np = lookup(name)) == NULL) { //not found
        printf("install2\n");
        np = (struct Nlist *) malloc(sizeof(*np));
        if (np == NULL || (np->name = strdup(name)) == NULL)
            return NULL;
        hashval = hash(name);
        np->next = hashtab[hashval];
        printf("hashval[%d]\n", hashval);
        hashtab[hashval] = np;
    } else { //already there 
        printf("install3\n");
        free((void *) np->defn); //free previous defn
        printf("install4\n");
    }
    printf("install5\n");
    if ((np->defn = strdup(defn)) == NULL) {
        printf("install6\n");
        return NULL;
    }
    return np;
}
    

main() {
    printf("1\n");
    for (int i = 0; i < NUMELMTS; i++) {
        char key[20];
        sprintf(key, "MAX%d", i);    
        char val[30];
        sprintf(val, "fred%d", i);  
    printf("2\n");
        struct Nlist *x = install(key, val);
        printf("3\n");
    }
    for (int i = 0; i < NUMELMTS; i++) {
        printf("in loop\n");
        Nlistprint(hashtab[i]);
    }
}