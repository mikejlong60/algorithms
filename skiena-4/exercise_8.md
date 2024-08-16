Question: You are given a pile of thousands of telephone bills and thousands of checks sent in to pay the bills. Find out who did not pay.
Answer:
    Steps:  Big O(n log n) due to the sorting
       1. Sort both arrays by name a - checks, b - bills.
       2. unpaidBills = []UnpaidBills
       3. for i := 0; i < len(a); i++ {
            if array a[i] is > b[i] {// There is not a matching check so add it to if a[i].firstName < b[i].firstName && a[i].lastName == b[i].lastName {
               unpaidBills = append(unpaidBills, b[i])
            }
       }


Question: You are given a printed list containing the title, author, call number, and publisher of all the books in a school
  library and another list of thirty publishers. Find out how many of the books in the library were published by each company.
Answer:
    Steps:  Big O(N)
       1. booksPerPublisher = map[string]int{}
       2. for i := 0; i < len(b); i++ {
            lookup publisher in booksPerPublisher
            if OK {
                increment book count
            } else {
                Add new entry for publisher initialized to 1
            }
       }


Question: You are given all the book checkout cards used in the campus library during the past year, each of which contains
   the name of the person who took out the book. Determine how many distinct people checked out at least one book.
   Steps:  Big O(N)
       1. distinctPeople = map[string]interface{}{}
       2. for i := 0; i < len(cards); i++ {
            lookup person in distinctPeople
            if !OK {
                Add new entry for person
            }
          }
          return len(distinctPeople)


