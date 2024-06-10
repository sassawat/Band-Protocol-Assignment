# Assignment_1

Title: Boss Baby's Revenge

Example script "main.js"

You can run with `node main.js`

Explanation

0. Define `shotsCount` as default value is 0. For counting shot.
1. Check numbers of input. If 1 > n,k > 1,000,000 return 0.
2. Define characters for receive from input. Only "S" and "R".
3. Check first characters, Boss Baby is trying to take revenge without any prior shots, so we immediately return "Bad boy".
4. Using loop for check each characters of S.
   - Check characters input only S, R.
   - Check value of index is equal "S".
     - If true. Increment the shotsCount because it represents a new shot that needs to be revenged.
   - Check value of index is equal R and shotsCount more than zero.
     - If true. Decrement shotsCount to indicate one shot has been revenged.
5. End of all, check if all shots have been avenged. If `shotsCount` equal zero, meaning is Boss Baby can avenged.

Test case:

- Add 2 case.
  1. Input contains character "Q".
  2. Length of input more than 1000000.
