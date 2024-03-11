package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



type Dados struct{
	 Name string `json:"firstname"`
	 Email string `json:"email"`
	 CardNumber string `json:"cardnumber"`
	 CardName string  `json:"cardname"`
	 ExpireMonth string `json:"expmonth"`
	 ExpireYear string `json:"expyear"`
	 CVV string `json:"cvv"`
}
func main() {

	var listOfCards = []Dados{}

	
	router := gin.Default()
	router.Use(corsMiddleware())
	router.POST("/", func(c *gin.Context) {
		var requestData = Dados{}
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Request data received",
		})
		listOfCards = append(listOfCards, requestData)
	})

	router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
			"data": listOfCards,
		})
	})

	router.GET("delete/:id", func(c *gin.Context) {
		cardid, _:= c.Params.Get("id")
		for index, card := range listOfCards {
			if card.CardNumber == cardid{
				listOfCards = append(listOfCards[:index], listOfCards[index+1:]... )
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"data": listOfCards,
		})
	})

	router.Run(":9000")
}


func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}