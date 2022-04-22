package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/swiggy-ipp/cart-service/configs/database"
	"github.com/swiggy-ipp/cart-service/models"
	"github.com/swiggy-ipp/cart-service/repositories"
)

func main() {
	// Mock input
	userIDMock := "user1"

	// Create mock context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Make layered Architecture
	db := database.GetDynamoDBClient()                           // Database
	cartRepository := repositories.NewCartRepository(db, "cart") // Repository
	// configs.DeleteDynamoDBTable(db, "cart")                      // Delete Table
	// configs.CreateDynamoDBTable(db, "cart")                      // Create Table

	// Mock CRUD
	// Create
	cartRepository.Create(ctx, &models.Cart{
		UserID: userIDMock,
		Items: []models.CartItem{
			{
				ProductID: "product1",
				Quantity:  1,
			},
		},
	})
	// Read by UserID
	cart, _ := cartRepository.ReadByUserID(ctx, userIDMock)
	logrus.Info(cart)
	// Update
	cart.Items[0].ProductID = "updated"
	cartRepository.UpdateCartItems(ctx, cart)
	// Read
	cart, _ = cartRepository.Read(ctx, cart.Id)
	logrus.Info(cart)
	// Delete
	cartRepository.Delete(ctx, cart.Id)

}
