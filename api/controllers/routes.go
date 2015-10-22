package controllers

import (
	"net/http"

	"github.com/convox/rack/api/Godeps/_workspace/src/github.com/gorilla/mux"
)

func HandlerFunc(w http.ResponseWriter, req *http.Request) {
	router := NewRouter()
	router.ServeHTTP(w, req)
}

func NewRouter() (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/sns", SNSHandler)
	router.HandleFunc("/apps", api("app.list", AppList)).Methods("GET")
	router.HandleFunc("/apps", api("app.create", AppCreate)).Methods("POST")
	router.HandleFunc("/apps/{app}", api("app.get", AppShow)).Methods("GET")
	router.HandleFunc("/apps/{app}", api("app.delete", AppDelete)).Methods("DELETE")
	router.HandleFunc("/apps/{app}/builds", api("build.list", BuildList)).Methods("GET")
	router.HandleFunc("/apps/{app}/builds", api("build.create", BuildCreate)).Methods("POST")
	router.HandleFunc("/apps/{app}/builds/{build}", api("build.get", BuildGet)).Methods("GET")
	router.HandleFunc("/apps/{app}/environment", api("environment.list", EnvironmentList)).Methods("GET")
	router.HandleFunc("/apps/{app}/environment", api("environment.set", EnvironmentSet)).Methods("POST")
	router.HandleFunc("/apps/{app}/environment/{name}", api("environment.delete", EnvironmentDelete)).Methods("DELETE")
	router.HandleFunc("/apps/{app}/formation", api("formation.list", FormationList)).Methods("GET")
	router.HandleFunc("/apps/{app}/formation/{process}", api("formation.set", FormationSet)).Methods("POST")
	router.HandleFunc("/apps/{app}/processes", api("process.list", ProcessList)).Methods("GET")
	router.HandleFunc("/apps/{app}/processes/{process}", api("process.get", ProcessShow)).Methods("GET")
	router.HandleFunc("/apps/{app}/processes/{process}", api("process.stop", ProcessStop)).Methods("DELETE")
	router.HandleFunc("/apps/{app}/processes/{process}/run", api("process.run.detach", ProcessRunDetached)).Methods("POST")
	router.HandleFunc("/apps/{app}/releases", api("release.list", ReleaseList)).Methods("GET")
	router.HandleFunc("/apps/{app}/releases/{release}", api("release.get", ReleaseShow)).Methods("GET")
	router.HandleFunc("/apps/{app}/releases/{release}/promote", api("release.promote", ReleasePromote)).Methods("POST")
	router.HandleFunc("/apps/{app}/ssl", api("ssl.list", SSLList)).Methods("GET")
	router.HandleFunc("/apps/{app}/ssl", api("ssl.create", SSLCreate)).Methods("POST")
	router.HandleFunc("/apps/{app}/ssl/{process}/{port}", api("ssl.delete", SSLDelete)).Methods("DELETE")
	router.HandleFunc("/services", api("service.list", ServiceList)).Methods("GET")
	router.HandleFunc("/services", api("service.create", ServiceCreate)).Methods("POST")
	router.HandleFunc("/services/{service}", api("service.show", ServiceShow)).Methods("GET")
	router.HandleFunc("/services/{service}", api("service.delete", ServiceDelete)).Methods("DELETE")
	router.HandleFunc("/services/{service}/links", api("link.create", LinkCreate)).Methods("POST")
	router.HandleFunc("/services/{service}/links/{app}", api("link.delete", LinkDelete)).Methods("DELETE")
	router.HandleFunc("/system", api("system.show", SystemShow)).Methods("GET")
	router.HandleFunc("/system", api("system.update", SystemUpdate)).Methods("PUT")
	router.HandleFunc("/switch", api("switch", Switch)).Methods("POST")

	// websockets
	router.Handle("/apps/{app}/logs", ws("app.logs", AppLogs)).Methods("GET")
	router.Handle("/apps/{app}/builds/{build}/logs", ws("build.logs", BuildLogs)).Methods("GET")
	router.Handle("/apps/{app}/processes/{pid}/exec", ws("process.exec.attach", ProcessExecAttached)).Methods("GET")
	router.Handle("/apps/{app}/processes/{process}/run", ws("process.run.attach", ProcessRunAttached)).Methods("GET")
	router.Handle("/services/{service}/logs", ws("service.logs", ServiceLogs)).Methods("GET")

	// utility
	router.HandleFunc("/boom", UtilityBoom).Methods("GET")
	router.HandleFunc("/check", UtilityCheck).Methods("GET")

	// limbo
	// auth.HandleFunc("/apps/{app}/debug", controllers.AppDebug).Methods("GET")
	// auth.HandleFunc("/apps/{app}/processes/{id}", controllers.ProcessStop).Methods("DELETE")
	// auth.HandleFunc("/apps/{app}/processes/{id}/top", controllers.ProcessTop).Methods("GET")
	// auth.HandleFunc("/top/{metric}", controllers.ClusterTop).Methods("GET")

	return
}
