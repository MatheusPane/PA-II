package controllers

import (
	"api/database" // ✅ import database
	"api/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

    func CreateOrder(c *fiber.Ctx) error {
        // Struktur input dari client
        var input struct {
            UserID  uint `json:"user_id"`
            Items   []struct {
                ProductID uint `json:"product_id"`
                Quantity  int  `json:"quantity"`
            } `json:"items"`
        }

        // Parse input dari body request
        if err := c.BodyParser(&input); err != nil {
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{
                "message": "Invalid input",
                "error":   err.Error(),
            })
        }

        var total float64
        var orderItems []models.OrderItem

        // Hitung total dan siapkan order items
        for _, item := range input.Items {
            var product models.Product
            if err := database.DB.First(&product, item.ProductID).Error; err != nil {
                return c.Status(http.StatusBadRequest).JSON(fiber.Map{
                    "message": "Product not found",
                    "product_id": item.ProductID,
                })
            }

            subtotal := float64(item.Quantity) * product.Price
            total += subtotal

            orderItems = append(orderItems, models.OrderItem{
                ProductID: item.ProductID,
                Quantity:  item.Quantity,
                Price:     product.Price,
            })
        }

        // Buat order baru
        order := models.Order{
            UserID: input.UserID,
            TotalPrice: total,
            Status: "pending", // default status
            Items:  orderItems,
        }

        if err := database.DB.Create(&order).Error; err != nil {
            return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
                "message": "Failed to create order",
                "error":   err.Error(),
            })
        }

        return c.Status(http.StatusCreated).JSON(fiber.Map{
            "message": "Order created successfully",
            "order_id": order.ID,
        })
    }


func GetAllOrders(c *fiber.Ctx) error {
    var orders []models.Order

    
    result := database.DB.
        Preload("User").
        Preload("Items.Product").
        Find(&orders)

 
    if result.Error != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to fetch orders",
        })
    }
    
    // Menyiapkan response yang lebih terstruktur
    var formattedOrders []fiber.Map
    for _, order := range orders {
        var products []fiber.Map
        for _, item := range order.Items {
            var productDetail models.Product
            err := database.DB.First(&productDetail, item.ProductID).Error
            if err != nil {
                productDetail.Title = "Unknown"
            }
            products = append(products, fiber.Map{
                "product_id":   item.ProductID,
                "product_name": productDetail.Title,
                "quantity":     item.Quantity,
                "price":        item.Price,
            })

            fmt.Println("Product Name:", item.Product.Title)
        }
    
        formattedOrders = append(formattedOrders, fiber.Map{
            "order_id":     order.ID,
            "customer_name": order.User.Name,
            "status":       order.Status,
            "products":     products,
        })
    }
    

    return c.JSON(fiber.Map{
        "orders": formattedOrders,        
    })
    
}
