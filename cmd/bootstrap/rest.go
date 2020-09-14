package bootstrap

import (
	"context"
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"net/http"
	"service/internal/middleware/rest"
	"service/pkg/echoserver"
	"service/pkg/logger"
	"service/pkg/reconf"

	// products module
	pdt "service/internal/module/products/transport/rest"
	pdu "service/internal/module/products/usecase"
)

var (
	restCommand = &cobra.Command{
		Use:   "rest",
		Short: "Starting REST API service",
		Run: func(cmd *cobra.Command, args []string) {
			// init root context
			ctx := context.Background()

			// init database
			dbread := InitPostgresRead()
			dbwrite := InitPostgresWrite()

			// init echo server
			echoserver.InitServer(reconf.Config().GetBool("debug"))

			// set default middleware
			echoserver.GetServer().Use(
				rest.CORS(),
				rest.Headers(),
				rest.Logger(logger.GetLogger()),
			)

			// versioning
			v1 := echoserver.GetServer().Group("/v1")

			// set healthcheck endpoint
			v1.GET("/healthcheck", func(c echo.Context) error {
				return c.JSON(http.StatusOK, "OK")
			})

			// init products
			pdt.NewHandler(v1, pdu.Initialize(dbread, dbwrite))

			// start server
			go echoserver.StartServer(
				reconf.Config().GetInt("rest.port"),
				reconf.Config().GetInt("rest.write_timeout"),
				reconf.Config().GetInt("rest.read_timeout"),
				reconf.Config().GetInt("rest.idle_timeout"),
			)

			// watching config changes
			go reconf.Watch(ctx, reconf.Config().GetInt("sync_config"))

			// register service
			RegistrationService()

			// deregister service
			defer DeregistrationService()

			// shutdown server gracefully
			echoserver.Shutdown(ctx)
		},
	}
)
