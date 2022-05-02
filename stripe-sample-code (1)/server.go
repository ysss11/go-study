package main

import (
  "bytes"
  "encoding/json"
  "io"
  "log"
  "net/http"

  "github.com/stripe/stripe-go/v72"
  "github.com/stripe/stripe-go/v72/paymentintent"
  "github.com/stripe/stripe-go/v72/customer"
  "github.com/stripe/stripe-go/v72/paymentmethod"
)

func main() {
  // This is a public sample test API key.
  // Don’t submit any personally identifiable information in requests made with this key.
  // Sign in to see your own test API key embedded in code samples.
  stripe.Key = "sk_test_09l3shTSTKHYCzzZZsiLl2vA"

  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)
  http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

  addr := "localhost:4242"
  log.Printf("Listening on %s ...", addr)
  log.Fatal(http.ListenAndServe(addr, nil))
}

type item struct {
  id string
}

func calculateOrderAmount(items []item) int64 {
  // Replace this constant with a calculation of the order's amount
  // Calculate the order total on the server to prevent
  // people from directly manipulating the amount on the client
  return 1400
}

func chargeCustomer(CustomerID string) {
  // Lookup the payment methods available for the customer
  params := &stripe.PaymentMethodListParams{
    Customer: stripe.String(CustomerID),
    Type:     stripe.String(string(stripe.PaymentMethodTypeCard)),
  }
  i := paymentmethod.List(params)
  for i.Next() {
    pm := i.PaymentMethod()
  }

  piparams := &stripe.PaymentIntentParams{
    Amount:        stripe.Int64(1099),
    Currency:      stripe.String(string(stripe.CurrencyEUR)),
    Customer:      stripe.String(CustomerID),
    PaymentMethod: stripe.String(pm.ID),
    Confirm:       stripe.Bool(true),
    OffSession:    stripe.Bool(true),
  }

  // Charge the customer and payment method immediately
  _, err := paymentintent.New(piparams)

  if err != nil {
    if stripeErr, ok := err.(*stripe.Error); ok {
      // Error code will be authentication_required if authentication is needed
      fmt.Printf("Error code: %v", stripeErr.Code)

      paymentIntentID := stripeErr.PaymentIntent.ID
      paymentIntent, _ := paymentintent.Get(paymentIntentID, nil)

      fmt.Printf("PI: %v", paymentIntent.ID)
    }
  }
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
    return
  }

  var req struct {
    Items []item `json:"items"`
  }

  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Printf("json.NewDecoder.Decode: %v", err)
    return
  }
  // Alternatively, set up a webhook to listen for the payment_intent.succeeded event
  // and attach the PaymentMethod to a new Customer
  cparams := &stripe.CustomerParams{}
  c, _ := customer.New(cparams)

  // Create a PaymentIntent with amount and currency
  params := &stripe.PaymentIntentParams{
    Customer: stripe.String(c.ID),
    SetupFutureUsage: stripe.String("off_session"),
    Amount:   stripe.Int64(calculateOrderAmount(req.Items)),
    Currency: stripe.String(string(stripe.CurrencyEUR)),
    AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
      Enabled: stripe.Bool(true),
    },
  }

  pi, err := paymentintent.New(params)
  log.Printf("pi.New: %v", pi.ClientSecret)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Printf("pi.New: %v", err)
    return
  }

  writeJSON(w, struct {
    ClientSecret string `json:"clientSecret"`
  }{
    ClientSecret: pi.ClientSecret,
  })
}

func writeJSON(w http.ResponseWriter, v interface{}) {
  var buf bytes.Buffer
  if err := json.NewEncoder(&buf).Encode(v); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    log.Printf("json.NewEncoder.Encode: %v", err)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  if _, err := io.Copy(w, &buf); err != nil {
    log.Printf("io.Copy: %v", err)
    return
  }
}