# reverse of a positive number 
How loop is working is mentioned below:
Number = 12345
In first iteration, the condition in for loop is true,
    remainder = 5, reverseNumber = 5, originalNumber = 1234
In second iteration, the condition is still true,
    remainder = 4, reverseNumber = 5*10 + 4 = 54, originalNumber = 123
In third iteration, the condition is still true,
    remainder = 3, reverseNumber = 54*10 + 3 = 543, originalNumber = 12
In fourth iteration, the condition is still true
    remainder = 2, reverseNumber = 543*10 + 2 = 5432, originalNumber = 1
In fifth iteration, the condition is still true
    remainder = 1, reverseNumber = 5432*10 + 1 = 54321, originalNumber = 0
In next iteration, the condition is false and loops break

NOTE : This same program can be used for palindrome for an integer as well. Palindrome tells us whether a number is equal to when it is reversed. For example: 121, 111, 1551 