function checkBossBabyBehavior(S: string): string {
  let shotsCount = 0;

  // Check numbers of input.
  if (!(1 <= S.length && S.length <= 1000000)) {
    return "Input is exceeded or a few";
  }

  // Define characters for receive from input.
  const allowChar = ["S", "R"];

  // Check first characters, Boss Baby is trying to take revenge without any prior shots, so we immediately return "Bad boy".
  if (S[0] === "R") return "Bad boy";

  // Using loop for check each characters of S.
  for (let charIndex = 0; charIndex < S.length; charIndex++) {
    // Check characters input only S, R.
    if (!allowChar.includes(S[charIndex])) {
      return "Input only required 'S' or 'R'";
    }

    // Check value of index is equal "S".
    if (S[charIndex] === "S") {
      // Increment the shotsCount because it represents a new shot that needs to be revenged.
      shotsCount++;
    }

    // Check value of index is equal R and shotsCount more than zero.
    if (S[charIndex] === "R" && shotsCount > 0) {
      // Decrement shotsCount to indicate one shot has been revenged.
      shotsCount--;
    }
  }

  // Check if all shots have been revenged.
  return shotsCount === 0 ? "Good boy" : "Bad boy";
}

// Example test case:
console.log(checkBossBabyBehavior("SRSSRRR")); // Good boy
console.log(checkBossBabyBehavior("RSSRR")); // Bad boy
console.log(checkBossBabyBehavior("SSSRRRRS")); // Bad boy
console.log(checkBossBabyBehavior("SRRSSR")); // Bad boy
console.log(checkBossBabyBehavior("SSRSRRR")); // Good boy

console.log(checkBossBabyBehavior("SSRSRRQ")); // Input only required 'S' or 'R'

const randomMixedSequence = Array(1000001)
  .fill(0)
  .map((_, i) => (i % 2 === 0 ? "S" : "R"))
  .join("");
console.log(checkBossBabyBehavior(randomMixedSequence)); // Input is exceeded or a few
