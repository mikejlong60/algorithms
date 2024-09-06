Question: Given a set S of n positive(simplifies fo now) integers and an integer T, give an O(n(k-1))
algorithm to test whether k of three distinct integers in S add up to T

Answer:
    1. Sort the set of integers
    2. Find the closest spot j in S using a binary search where s[j] < T. Example
        Example data set: S = {1,2,3,4,5}. k = 3, T = 8
        2.1 Find closest spot in array with binary search less than 8(T).  It's 5, position 4.
        2.2 Work backwards 2 spots from 5 since you have three numbers you have to add and you cannot repeat usage 
        of the same number.
        2.3 recursively function f- 
            2.3.1 start at position 3.  Add position 2 to working sum(ws). Does ws exceed T?
                Yes - stop algorithm - no such set of k integers exits
                No - call f with 
                    

func f(int[] S, int T, int k, int workingIdx, int workingSum) bool {
    

}