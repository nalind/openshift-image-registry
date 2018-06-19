package server

import (
	"net/http"

	"github.com/docker/distribution/registry/auth"
	"github.com/docker/distribution/registry/handlers"

	"github.com/openshift/image-registry/pkg/dockerregistry/server/api"
	"github.com/openshift/image-registry/pkg/dockerregistry/server/metrics"
)

func RegisterMetricHandler(app *handlers.App) {
	getMetricsAccess := func(r *http.Request) []auth.Access {
		return []auth.Access{
			{
				Resource: auth.Resource{
					Type: "metrics",
				},
				Action: "get",
			},
		}
	}
	extensionsRouter := app.NewRoute().PathPrefix(api.ExtensionsPrefix).Subrouter()
	app.RegisterRoute(
		extensionsRouter.Path(api.MetricsPath).Methods("GET"),
		metrics.Dispatcher,
		handlers.NameNotRequired,
		getMetricsAccess,
	)
}
