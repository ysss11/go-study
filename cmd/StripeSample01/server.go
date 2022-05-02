package main

import (
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func main() {
	// This is a public sample test API key.
	// Don’t submit any personally identifiable information in requests made with this key.
	// Sign in to see your own test API key embedded in code samples.
	stripe.Key = "sk_test_51J7JAuJgIlSK0iakkSNk2SHR7e7ZoX6XA7lOfQbAkkMnjMtZUCzlMvcKogo3XBgTbPAuok05UKN5UhSSaCatguX4002Q5n7izR"

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/create-checkout-session", createCheckoutSession)
	addr := "127.0.0.1:4242"
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func createCheckoutSession(w http.ResponseWriter, r *http.Request) {
	domain := "http://127.0.0.1:4242"
	params := &stripe.CheckoutSessionParams{
		AutomaticTax:             &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
		PaymentMethodTypes:       stripe.StringSlice([]string{"card"}),
		CustomerEmail:            stripe.String("customer@example.com"),
		SubmitType:               stripe.String("donate"),
		BillingAddressCollection: stripe.String("auto"),
		ShippingAddressCollection: &stripe.CheckoutSessionShippingAddressCollectionParams{
			AllowedCountries: stripe.StringSlice([]string{
				// "US",
				// "CA",
				"JP",
			}),
		},
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String("price_1KttjiJgIlSK0iak9fqMrmr6"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "?success=true"),
		CancelURL:  stripe.String(domain + "?canceled=true"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
