# Assignment_2

Title: Superman's Chicken Rescue

Example script "main.js"

You can run with `node main.js`

Explanation

1. Initialize `maxProtected` to keep track of the maximum number of chickens protected
2. Initialize `start` to keep the starting index of the current window.
3. Check numbers of input. If 1 > n,k > 1,000,000 return 0.
4. iterate through the positions array using the end pointer.
5. For each position end, I check if the difference between positions[end] and positions[start] is greater than or equal to k. If it is, I increment the start pointer to reduce the scope from the left until the condition is false.
6. I update the maxProtected with the maximum number of chickens that can be covered within the current scope (end - start + 1).
7. After iterating through all positions, this function return the maxProtected which holds the maximum number of chickens that can be protected.
