function maxChickensProtected(
  n: number,
  k: number,
  positions: number[]
): number {
  let maxProtected = 0;
  let start = 0;

  // Check numbers of input
  if (!(1 <= n && n <= 1000000) || !(1 <= k && k <= 1000000)) {
    return 0;
  }

  // Using loop for get end position
  for (let end = 0; end < n; end++) {
    // Check positions difference between end and start position.
    if (positions[end] - positions[start] >= k) {
      // Increment start for check next start position value
      start++;
    }
    // Set max protected value by compare value between previously value and index of end minus index of start with add 1.
    // numbers of protected will be update when start not increment
    maxProtected = Math.max(maxProtected, end - start + 1);
  }

  return maxProtected;
}

// Example test case:
const n1 = 5;
const k1 = 5;
const positions1 = [2, 5, 10, 12, 15];
console.log(maxChickensProtected(n1, k1, positions1)); // Output: 2

const n2 = 6;
const k2 = 10;
const positions2 = [1, 11, 30, 34, 35, 37];
console.log(maxChickensProtected(n2, k2, positions2)); // Output: 4
