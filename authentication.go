package pagarmeapisdk

import (
	"net/http"

	"github.com/apimatic/go-core-runtime/https"
)

func BasicAuthentication(config Configuration) https.Authenticator {
	return func(requiresAuth bool) https.HttpInterceptor {
		if !requiresAuth {
			return https.PassThroughInterceptor
		}
		return func(req *http.Request,
			next https.HttpCallExecutor,
		) https.HttpContext {
			req.SetBasicAuth(config.BasicAuthUserName(), config.BasicAuthPassword())
			return next(req)
		}
	}
}
