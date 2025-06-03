package controllers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type CreateTransactionRequest struct {
	OrderID      string `json:"orderId"`
	Amount       int64  `json:"amount"`
	CustomerName string `json:"customerName"`
}

func CreateTransaction(c *fiber.Ctx) error {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Ambil key dari env
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Midtrans Server Key is not set"})
	}

	var reqBody CreateTransactionRequest

	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if reqBody.OrderID == "" || reqBody.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID or amount"})
	}

	snapClient := snap.Client{}
	snapClient.New(serverKey, midtrans.Production)

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  reqBody.OrderID,
			GrossAmt: reqBody.Amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: reqBody.CustomerName,
			Email: "pedromhutagaol@gmail.com",
		},
		Expiry: &snap.ExpiryDetails{
			StartTime: time.Now().Format("2006-01-02 15:04:05 -0700"),
			Unit:      "minutes",
			Duration:  30,
		},
	}

	resp, err := snapClient.CreateTransaction(snapReq)
	if err != nil || resp == nil || resp.Token == "" {
		log.Printf("Midtrans Error: %+v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create transaction"})
	}

	return c.JSON(fiber.Map{
		"snapToken": resp.Token,
	})
}
