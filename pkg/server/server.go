package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/linkc0829/go-icsharing/pkg/server/routes"
	"github.com/linkc0829/go-icsharing/pkg/utils"
	"github.com/linkc0829/go-icsharing/pkg/utils/datasource"
	"github.com/unrolled/secure"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func RegisterRoutes(cfg *utils.ServerConfig, r *gin.Engine, db *datasource.DB) {
	routes.Auth(cfg, r, db)
	routes.Graph(cfg, r, db.Mongo)
	routes.FreeTrial(cfg, r, db.Sqlite)
	routes.RestAPI(cfg, r, db.Mongo)

}

//Run will steup the routes and start the server
func SetupServer(serverconf *utils.ServerConfig, db *datasource.DB) *gin.Engine {

	r := gin.Default()

	InitalizeAuthProviders(serverconf)
	RegisterRoutes(serverconf, r, db)

	// Inform the user where the server is listening
	r.LoadHTMLGlob("views/*")
	r.Static(serverconf.StaticPath, "./public")
	r.StaticFile("/favicon.ico", "./favicon.ico")

	// HTTPS
	// To generate a development cert and key, run the following from your *nix terminal:
	// go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
	if serverconf.URISchema == "https" {
		r.Use(TlsHandler(serverconf))
	}

	log.Println("Running @ " + serverconf.ListenEndpoint())

	// Run the server
	// Print out and exit(1) to the OS if the server cannot run
	//log.Fatal(r.Run(serverconf.ListenEndpoint()))
	return r
}

func TlsHandler(cfg *utils.ServerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     cfg.CookiesDomain,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
