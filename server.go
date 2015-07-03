package main

import (
	"fmt"
	"net/http"

	"github.com/jessevdk/go-flags"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

type doorman struct {
	twilio *string
	host   *string
	port   *string
}

func answer(c *echo.Context) error {
	return c.String(http.StatusOK, "answer POST\n")
}

func main() {

	fmt.Println("Starting Doorman")

	var opts struct {
		// Slice of bool will append 'true' each time the option
		// is encountered (can be set multiple times, like -vvv)
		Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information" default:"false"`

		Port string `short:"p" long:"Port" description:"Port to run on" required:"true"`

		TwilioSid string `short:"s" long:"sid" description:"Twilio SID" required:"true"`

		TwilioToken string `short:"t" long:"token" description:"Twilio Token" required:"true"`
	}

	_, err := flags.Parse(&opts)

	if err != nil {
		panic(err)
	}

	// Echo instance
	e := echo.New()

	// Debug mode
	e.SetDebug(opts.Verbose)

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Index("public/index.html")

	// Deployment Routes
	e.Post("/v1/answer", answer)

	// Start server
	e.Run(":" + opts.Port)

}
