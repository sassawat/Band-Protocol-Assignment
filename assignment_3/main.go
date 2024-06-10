package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BroadcastPayload struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type BroadcastResponse struct {
	TxHash string `json:"tx_hash"`
}

type TransactionClient struct {
	BroadcastURL string
	CheckURL     string
}

func NewTransactionClient() *TransactionClient {
	return &TransactionClient{
		BroadcastURL: "https://mock-node-wgqbnxruha-as.a.run.app/broadcast",
		CheckURL:     "https://mock-node-wgqbnxruha-as.a.run.app/check",
	}
}

func (c *TransactionClient) BroadcastTransaction(payload BroadcastPayload) (string, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp, err := http.Post(c.BroadcastURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("failed to broadcast transaction: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var broadcastResponse BroadcastResponse
	if err := json.NewDecoder(resp.Body).Decode(&broadcastResponse); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return broadcastResponse.TxHash, nil
}

func (c *TransactionClient) CheckTransactionStatus(txHash string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", c.CheckURL, txHash))
	if err != nil {
		return "", fmt.Errorf("failed to check transaction status: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}

func (c *TransactionClient) MonitorTransaction(txHash string, interval time.Duration) error {
	for {
		status, err := c.CheckTransactionStatus(txHash)
		if err != nil {
			return err
		}

		fmt.Printf("Transaction status: %s\n", status)

		if status == "CONFIRMED" {
			fmt.Println("Transaction confirmed successfully.")
			break
		} else if status == "FAILED" {
			fmt.Println("Transaction failed.")
			break
		} else if status == "DNE" {
			fmt.Println("Transaction does not exist.")
			break
		}

		time.Sleep(interval)
	}

	return nil
}

func main() {
	client := NewTransactionClient()

	payload := BroadcastPayload{
		Symbol:    "ETH",
		Price:     4500,
		Timestamp: 1678912345,
	}

	txHash, err := client.BroadcastTransaction(payload)
	if err != nil {
		fmt.Printf("Error broadcasting transaction: %v\n", err)
		return
	}

	fmt.Printf("Transaction broadcasted with hash: %s\n", txHash)

	err = client.MonitorTransaction(txHash, 5*time.Second)
	if err != nil {
		fmt.Printf("Error during transaction monitoring: %v\n", err)
	}
}
