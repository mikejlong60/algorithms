#include <stdio.h>

#define T1 char*
#define T2 char*

T2 fold_right(T1 xs, T2 z, int arraylen, T2 (*f)(T1, T2)) { //}, f func(T1, T2) T2) T2 {
    printf("a:%s ---- b:%s arraylen:%d\n", xs, z, arraylen);
    if (arraylen < 15) { //Slice has a head and a tail.
		//char* h = xs*;
		  printf("inside 1\n", xs, z, arraylen);
		*xs++;

		return f(xs, fold_right(xs, z, ++arraylen, f));
//	} else if len(as) == 1 { //Slice has a head and an empty tail.
//		h := as[0]
//		return f(h, FoldRight(Zero[T1](), z, f))
    }
	return z;
}

char* f(char* a, char* b) {
    printf("a:%p ---- b:%p\n", a, b);
    return a;
}

void processString(T1 input, T2 output) {
    // Example implementation
    while (*input) {
        *output = *input;
        input++;
        output++;
    }
    *output = '\0';
}

//#include <stdio.h>

// Function pointer type for the operation
typedef int (*foldRightFunc)(int, int);

// foldRight function
int foldRight(int *arr, int size, int initial, foldRightFunc f) {
    if (size == 0) {
        return initial;
    } else {
        return f(arr[size - 1], foldRight(arr, size - 1, initial, f));
    }
}

// Example operation: sum function
int sumOperation(int element, int acc) {
    return element + acc;
}

// Example operation: product function
int productOperation(int element, int acc) {
    return element * acc;
}

// Define a function pointer type for the operation
typedef int (*foldLeftFunc)(int acc, int element);

// Recursive foldLeft function
int foldLeft(int *arr, int size, int index, int acc, foldLeftFunc f) {
    // Base case: when we've processed all elements
    if (index >= size) {
        return acc;
    }

    // Recursive case: apply the function and recurse
    return foldLeft(arr, size, index + 1, f(acc, arr[index]), f);
}

// Example operation: sum function
int lsumOperation(int acc, int element) {
    return acc + element;
}

// Example operation: product function
int lproductOperation(int acc, int element) {
    return acc * element;
}

int main() {
    int arr[] = {1, 2, 3, 4, 5};
    int size = sizeof(arr) / sizeof(arr[0]);

    // Folding with sum operation
    int sumResult = foldRight(arr, size, 0, sumOperation);
    printf("Sum: %d\n", sumResult); // Output: Sum: 15

    // Folding with product operation
    int productResult = foldRight(arr, size, 1, productOperation);
    printf("Product: %d\n", productResult); // Output: Product: 120

//////////////
    int arr2[] = {1, 2, 3, 4, 5, 6,7,8,9,0};
    int size2 = sizeof(arr) / sizeof(arr[0]);

    // Folding with sum operation
    int lsumResult = foldLeft(arr2, size2, 0, 0, lsumOperation);
    printf("Sum: %d\n", lsumResult); // Output: Sum: 15

    // Folding with product operation
    int lproductResult = foldLeft(arr2, size2, 0, 1, lproductOperation);
    printf("Product: %d\n", lproductResult); // Output: Product: 120

}
