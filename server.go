package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/jessevdk/go-flags"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
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

type TwimlGatherResponse struct {
	XMLName  xml.Name `xml:"Response"`
	Pause    TwimlPause
	Gather   TwimlGather
	Say      string
	Redirect string
}

type TwimlPause struct {
	XMLName xml.Name `xml:"Pause"`
	Length  int      `xml:"length,attr"`
}

type TwimlGather struct {
	Method    string `xml:"method,attr"`
	NumDigits int    `xml:"numDigits,attr"`
	Action    string `xml:"action,attr"`
	TimeOut   int    `xml:"timeout,attr"`
	Say       string
	Pause     TwimlPause
}

func debug(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello\n")
}

func answer(c *echo.Context) error {

	pause := TwimlPause{
		Length: 3,
	}

	gather := TwimlGather{
		Method:    "POST",
		NumDigits: 2,
		Action:    "/v1/verify",
		TimeOut:   60,
		Say:       "Press 1 to enter PIN, Press 2 to say password.",
		Pause:     pause,
	}

	response := TwimlGatherResponse{
		Pause:    pause,
		Gather:   gather,
		Say:      "I didn't quite get that. Please call again.",
		Redirect: "/v1/error",
	}

	// x, err := xml.MarshalIndent(response, "  ", "  ")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// var retval = xml.Header + string(x)

	fmt.Println(response)

	return c.XML(http.StatusOK, response)
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
