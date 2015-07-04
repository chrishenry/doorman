package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/jessevdk/go-flags"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	// "github.com/chrishenry/twilio"
)

// type doorman struct {
// 	twilio *twilio.Client
// }

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information" default:"false"`

	Port string `short:"p" long:"Port" description:"Port to run on" required:"true"`

	TwilioSID string `short:"s" long:"sid" description:"Twilio SID" env:"TWILIO_SID"`

	TwilioToken string `short:"t" long:"token" description:"Twilio Token" env:"TWILIO_TOKEN"`
}

type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Say     string   `xml:",omitempty"`
}

func debug(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello\n")
}

func answer(c *echo.Context) error {
	twiml := TwiML{Say: "Hello World!"}
	x, err := xml.MarshalIndent(twiml, "  ", "  ")
	if err != nil {
		panic(err.Error())
	}

	var retval = string(x)

	fmt.Println(retval)

	return c.String(http.StatusOK, retval)
}

func main() {

	fmt.Println("Starting Doorman")

	_, err := flags.Parse(&opts)

	if err != nil {
		panic(err)
	}

	if len(opts.TwilioSID) == 0 {
		panic("Twilio SID is required")
	}

	if len(opts.TwilioToken) == 0 {
		panic("Twilio Token is required")
	}

	// Echo instance
	e := echo.New()

	// Debug mode
	e.SetDebug(opts.Verbose)

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Get("/", debug)

	// Deployment Routes
	e.Post("/v1/answer", answer)

	// Start server
	e.Run(":" + opts.Port)

}
