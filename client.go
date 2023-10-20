/*
Package pagarmeapisdk

This file was automatically generated by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package pagarmeapisdk

import (
	"net/http"

	"github.com/nanocfw/pagarme-golang-sdk/controllers"

	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
)

type Client struct {
	callBuilderFactory          https.CallBuilderFactory
	config                      Configuration
	userAgent                   string
	PlansController             controllers.PlansController
	SubscriptionsController     controllers.SubscriptionsController
	InvoicesController          controllers.InvoicesController
	OrdersController            controllers.OrdersController
	CustomersController         controllers.CustomersController
	RecipientsController        controllers.RecipientsController
	ChargesController           controllers.ChargesController
	TransfersController         controllers.TransfersController
	TokensController            controllers.TokensController
	TransactionsController      controllers.TransactionsController
	PayablesController          controllers.PayablesController
	BalanceOperationsController controllers.BalanceOperationsController
}

// Constructor for client.
func NewClient(config Configuration) *Client {
	client := &Client{
		config: config,
	}

	client.userAgent = utilities.UpdateUserAgent("PagarmeApiSDK - Go 6.8.0")
	client.callBuilderFactory = callBuilderHandler(
		func(server string) string {
			if server == "" {
				server = "default"
			}
			return getBaseUri(Server(server), client.config)
		},
		BasicAuthentication(config),
		https.NewHttpClient(config.HttpConfiguration()),
		config.httpConfiguration.RetryConfiguration(),
		withServiceRefererName(config.ServiceRefererName()),
		withUserAgent(client.userAgent),
	)

	baseController := controllers.NewBaseController(client)
	client.PlansController = *controllers.NewPlansController(*baseController)
	client.SubscriptionsController = *controllers.NewSubscriptionsController(*baseController)
	client.InvoicesController = *controllers.NewInvoicesController(*baseController)
	client.OrdersController = *controllers.NewOrdersController(*baseController)
	client.CustomersController = *controllers.NewCustomersController(*baseController)
	client.RecipientsController = *controllers.NewRecipientsController(*baseController)
	client.ChargesController = *controllers.NewChargesController(*baseController)
	client.TransfersController = *controllers.NewTransfersController(*baseController)
	client.TokensController = *controllers.NewTokensController(*baseController)
	client.TransactionsController = *controllers.NewTransactionsController(*baseController)
	client.PayablesController = *controllers.NewPayablesController(*baseController)
	client.BalanceOperationsController = *controllers.NewBalanceOperationsController(*baseController)
	return client
}

func (c *Client) GetCallBuilder() https.CallBuilderFactory {
	return c.callBuilderFactory
}

func getBaseUri(
	server Server,
	config Configuration) string {
	if config.Environment() == Environment(PRODUCTION) {
		if server == Server(ENUMDEFAULT) {
			return "https://api.pagar.me/core/v5"
		}
	}
	return "TODO: Select a valid server."
}

type clientOptions func(cb https.CallBuilder)

func callBuilderHandler(
	baseUrlProvider func(server string) string,
	auth https.Authenticator,
	httpClient https.HttpClient,
	retryConfig https.RetryConfiguration,
	opts ...clientOptions) https.CallBuilderFactory {
	callBuilderFactory := https.CreateCallBuilderFactory(baseUrlProvider, auth, httpClient, retryConfig)
	return tap(callBuilderFactory, opts...)
}

func tap(
	callBuilderFactory https.CallBuilderFactory,
	opts ...clientOptions) https.CallBuilderFactory {
	return func(httpMethod, path string) https.CallBuilder {
		callBuilder := callBuilderFactory(httpMethod, path)
		for _, opt := range opts {
			opt(callBuilder)
		}
		return callBuilder
	}
}

func withServiceRefererName(serviceRefererName string) clientOptions {
	f := func(request *http.Request) *http.Request {
		request.Header.Add("ServiceRefererName", serviceRefererName)
		return request
	}
	return func(cb https.CallBuilder) {
		cb.InterceptRequest(f)
	}
}

func withUserAgent(userAgent string) clientOptions {
	f := func(request *http.Request) *http.Request {
		request.Header.Add("user-agent", userAgent)
		return request
	}
	return func(cb https.CallBuilder) {
		cb.InterceptRequest(f)
	}
}
