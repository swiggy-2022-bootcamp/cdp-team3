package grpcs

import (
	"github.com/swiggy-ipp/checkout-service/grpcs/cart_checkout"
	"github.com/swiggy-ipp/checkout-service/grpcs/shipping_checkout"
)

var CartCheckoutGRPCClient cart_checkout.CheckoutServiceClient

var ShippingCheckoutGRPCClient shipping_checkout.ShippingClient

