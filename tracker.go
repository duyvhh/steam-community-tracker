package tracker

import (
	"appengine"
	"appengine/mail"
	"appengine/urlfetch"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

const ITEM_URL string = "<ITEM_URL_HERE>"
const SENDER string = "<SENDER_EMAIL_ADDRESS>"
const THRESHOLD float64 = 15.00 //THRESHOLD
const MESSAGE string = `
Your favorite item has its price dropped below your threshold!

Item URL: %s
Threshold Set: %f
Current Price: %f

Good luck!
`

func init() {
	http.HandleFunc("/trigger", handleTrigger)
}

func handleTrigger(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	client := urlfetch.Client(c)

	resp, err := client.Get(ITEM_URL)

	if err != nil {
		c.Infof(err.Error())
		return
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		c.Infof(err.Error())
		return
	}

	var price float64 = 0.0

	doc.Find("span.market_listing_price_with_fee").Each(func(i int, s *goquery.Selection) {

		priceVal := strings.TrimSpace(s.Text())

		if price != 0 {
			return
		}

		if strings.HasSuffix(priceVal, "USD") || strings.HasPrefix(priceVal, "$") {
			price, _ = strconv.ParseFloat(strings.TrimPrefix(strings.TrimSuffix(priceVal, " USD"), "$"), 64)
		}
	})

	if price < THRESHOLD {
		sendEmail(price, c)
	}

	fmt.Fprintf(rw, "Triggered: %f", price)
}

func sendEmail(price float64, c appengine.Context) {
	msg := &mail.Message{
		Sender:  SENDER,
		To:      []string{"<RECIPIENT_EMAIL_ADDRESSES"}, //comma-delimited email addresses
		Subject: "Steam Item price dropped below threshold",
		Body:    fmt.Sprintf(MESSAGE, ITEM_URL, THRESHOLD, price),
	}

	if err := mail.Send(c, msg); err != nil {
		c.Errorf("Couldn't send email: %v", err)
	}
}
