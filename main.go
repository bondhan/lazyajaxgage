package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/rs/cors"

	"github.com/labstack/echo"
)

// Gauge ...
type Gauge struct {
	Delay int `json:"delay`
	Value int `json:"value"`
}

// GetGaugeRand ...
func GetGaugeRand(delay int, name string) *Gauge {
	g := &Gauge{
		Delay: delay,
		Value: rand.Int() % 100,
	}
	fmt.Println("gauge:", name, "Delay..", delay)
	time.Sleep(time.Duration(delay) * time.Second)

	return g
}

func main() {

	rand.Seed(time.Hour.Nanoseconds())

	e := echo.New()
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/g1", func(c echo.Context) error {
		dur := rand.Int() % 10
		return c.JSON(http.StatusOK, GetGaugeRand(dur, "g1"))
	})

	e.GET("/g2", func(c echo.Context) error {
		dur := rand.Int() % 10
		return c.JSON(http.StatusOK, GetGaugeRand(dur, "g2"))
	})

	e.GET("/g3", func(c echo.Context) error {
		dur := rand.Int() % 10
		return c.JSON(http.StatusOK, GetGaugeRand(dur, "g3"))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
