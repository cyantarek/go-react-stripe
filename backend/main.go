package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"log"
	"os"
)

func main() {
	app := echo.New()
	app.Use(middleware.CORS())
	
	if os.Getenv("STRIPE_KEY") != "" {
		stripe.Key = os.Getenv("STRIPE_KEY")
	} else {
		log.Fatal("STRIPE_KEY env var not set")
	}
	
	app.POST("/api/payment", func(c echo.Context) error {
		var in = make(map[string]interface{})
		
		err := c.Bind(&in)
		if err != nil {
			return err
		}
		
		amount := int64(in["amount"].(float64))
		token := in["token"].(map[string]interface{})["id"].(string)
		currency := "USD"
		
		params := stripe.ChargeParams{
			Amount:   &amount,
			Currency: &currency,
		}
		
		params.SetSource(token)
		
		ch, err := charge.New(&params)
		if err != nil {
			return err
		}
		
		fmt.Println(ch)
		
		return nil
	})
	
	log.Fatal(app.Start(":5300"))
}
