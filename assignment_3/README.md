# Assignment_3

Title: Transaction Broadcasting and Monitoring Client

Example script "main.go"

You can run with `go run main.go`

Explanation

Methods

1. BroadcastTransaction(payload BroadcastPayload) (string, error)

   Sends a transaction broadcast request.
   Parameters:
   payload (BroadcastPayload): The transaction details including symbol, price, and timestamp.
   Returns:
   A string containing the transaction hash.
   An error if the request fails.

2. CheckTransactionStatus(txHash string) (string, error)

   Checks the status of a transaction using its hash.
   Parameters:
   txHash (string): The transaction hash.
   Returns:
   A string containing the transaction status (CONFIRMED, FAILED, PENDING, or DNE).
   An error if the request fails.

3. MonitorTransaction(txHash string, interval time.Duration) error

   Monitors the transaction status periodically until it is finalized.
   Parameters:
   txHash (string): The transaction hash.
   interval (time.Duration): The interval between status checks.
   Returns:
   An error if monitoring fails.

In the main function
I used `NewTransactionClient` for create new object and setup payload with `BroadcastPayload` type. Next, I used `BroadcastTransaction` function from object with payload to sends a transaction broadcast request. After its return `txHash`. Finally I used `MonitorTransaction` with `txHash` and `interval` (this params I make for wait respone from server for test) parameters to monitors the transaction status.

Conclusion
This example module can be demonstrates how to send a transaction and continuously check its status until it is finalized.
