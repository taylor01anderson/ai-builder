package routes

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/stripe/stripe-go/v82"
    "github.com/stripe/stripe-go/v82/checkout/session"
)

func StripeRoutes(rg *gin.RouterGroup) {

    rg.POST("/create-checkout", CreateCheckout)
}

func CreateCheckout(c *gin.Context) {

    stripe.Key = "YOUR_STRIPE_SECRET"

    params := &stripe.CheckoutSessionParams{
        SuccessURL: stripe.String("http://localhost:3000/success"),
        CancelURL: stripe.String("http://localhost:3000/cancel"),
        Mode: stripe.String(string(stripe.CheckoutSessionModeSubscription)),
        LineItems: []*stripe.CheckoutSessionLineItemParams{
            {
                Price: stripe.String("price_123"),
                Quantity: stripe.Int64(1),
            },
        },
    }

    s, err := session.New(params)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })

        return
    }

    c.JSON(http.StatusOK, gin.H{
        "url": s.URL,
    })
}