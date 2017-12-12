package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	cf "github.com/jasonmichels/go-journey-server-utils/config"
	md "github.com/jasonmichels/go-journey-server-utils/middleware"
	"github.com/jasonmichels/journey-registry/journey"
)

var templates = template.Must(template.ParseFiles(cf.INDEX))

// Handle serving assets from a journey registry
func handler(w http.ResponseWriter, r *http.Request, assets *journey.DependencyAssets) {

	if err := templates.ExecuteTemplate(w, "index.html", assets); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	nrName := os.Getenv("NEW_RELIC_NAME")
	nrKey := os.Getenv("NEW_RELIC_KEY")
	registryURL := cf.Getenv("JOURNEY_REGISTRY_URL", "localhost:9011")
	port := cf.Getenv("PORT", ":9010")
	pathPrefix := os.Getenv("PATH_PREFIX")

	j, err := cf.LoadJourneyConfig("journey.json")
	if err != nil {
		log.Panic(err)
	}

	http.Handle(md.NewRelicMiddleware(nrName, nrKey, "/", md.LoggingMiddleware(md.LocalAssetMiddleware(pathPrefix, md.JourneyAssetMiddleware(j, registryURL, handler)))))
	http.ListenAndServe(port, nil)
}
