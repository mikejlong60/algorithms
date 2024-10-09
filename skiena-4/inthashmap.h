#include <ctype.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct Nlist
{
    //table entry
    struct Nlist* next; //nest entry in liked list chain
    int key;
    int value; //replacement text
};

void Nlistprint(const struct Nlist*);

inline void Nlistprint(const struct Nlist* p)
{
    if (p != NULL)
    {
        Nlistprint(p->next);
        printf("Next guy in linked list - %n:%n\n", p->key, p->value);
    }
    printf("leaving Nlistprint\n");
}

#define HASHSIZE 3000
#define NUMELMTS 200000


static struct Nlist* hashtab[HASHSIZE]; //pointer table

// hash - form hash value for string s
unsigned hash(const int s)
{
    return s % HASHSIZE;
}

//lookup
struct Nlist* lookup(const int s)
{
    printf("lookup1\n");
    for (struct Nlist* np = hashtab[hash(s)]; np != NULL; np = np->next)
    {
        printf("searching linked list in cell[%d]\n", hash(s));
        if (s == np->key)
        {
            printf("found in linked list\n");
            return np; //found
        }
    }
    printf("did not find in lookup in cell[%d]\n", hash(s));
    return NULL; //not found
}

struct Nlist* install(const int key, const int value)
{
    printf("install1:key[%d] --- value[%d]\n", key, value);
    struct Nlist* np;
    printf("install111\n");
    if ((np = lookup(key)) == NULL)
    {
        //not found
        printf("install2\n");
        np = (struct Nlist*)malloc(sizeof(*np));
        if (np == NULL || (np->key == key))
            return NULL;
        unsigned hashval = hash(key);
        np->next = hashtab[hashval];
        printf("hashval[%d]\n", hashval);
        hashtab[hashval] = np;
    }
    else
    {
        //already there
        printf("install3\n");
        free((void*)np->value); //free previous defn
        printf("install4\n");
    }
    printf("install5\n");
    if (np->value == value)
    {
        printf("install6\n");
        return NULL;
    }
    return np;
}


// inline int main() {
//     printf("1\n");
//     for (int i = 0; i < NUMELMTS; i++) {
//         //char key[20];
//         //sprintf(key, "MAX%d", i);
//         //char val[30];
//         //sprintf(val, "fred%d", i);
//     //printf("2\n");
//         struct Nlist *x = install(&i, &i);
//         printf("3\n");
//     }
//     for (int i = 0; i < NUMELMTS; i++) {
//         printf("in loop\n");
//         Nlistprint(hashtab[i]);
//     }
// }
