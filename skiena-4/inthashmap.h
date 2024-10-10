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
    for (struct Nlist* np = hashtab[hash(s)]; np != NULL; np = np->next)
    {
        if (s == np->key)
        {
            return np; //found
        }
    }
    return NULL; //not found
}

struct Nlist* install(const int key, const int value)
{
    struct Nlist* np;
    if ((np = lookup(key)) == NULL)
    {
        //not found
        np = (struct Nlist*)malloc(sizeof(*np));
        np->key = key;
        unsigned hashval = hash(key);
        np->next = hashtab[hashval];
        hashtab[hashval] = np;
    }
    np->value = value;
    return np;
}
