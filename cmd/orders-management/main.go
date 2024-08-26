package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-chi/chi"
	"github.com/rakyll/statik/fs"
	clayLog "github.com/utrack/clay/v3/log"
	"github.com/utrack/clay/v3/transport/middlewares/mwgrpc"
	"github.com/utrack/clay/v3/transport/server"
	"github.com/zhetkerbaevan/orders-management/internal/handler"
	"github.com/zhetkerbaevan/orders-management/internal/service"
	_ "github.com/zhetkerbaevan/orders-management/static/statik"
)

func main() {
	// Wire up our bundled Swagger UI
	staticFS, err := fs.New()
	if err != nil {
		logrus.Fatal(err)
	}
	hmux := chi.NewRouter()
	hmux.Mount("/", http.FileServer(staticFS))

	//orderService := service.NewOrderService()
	orderService := service.NewOrderService()
	impl := handler.NewOrderImplementation(orderService)

	srv := server.NewServer(
		12345,
		// Pass our mux with Swagger UI
		server.WithHTTPMux(hmux),
		// Recover from both HTTP and gRPC panics and use our own middleware
		server.WithGRPCUnaryMiddlewares(mwgrpc.UnaryPanicHandler(clayLog.Default)),
	)
	err = srv.Run(impl)
	if err != nil {
		logrus.Fatal(err)
	}
}
