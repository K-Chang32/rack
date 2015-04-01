package controllers

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/convox/kernel/Godeps/_workspace/src/github.com/ddollar/logger"
	"github.com/convox/kernel/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/convox/kernel/Godeps/_workspace/src/github.com/gorilla/websocket"

	"github.com/convox/kernel/models"
)

func init() {
	RegisterPartial("app", "builds")
	RegisterPartial("app", "changes")
	RegisterPartial("app", "logs")
	RegisterPartial("app", "releases")
	RegisterPartial("app", "resources")
	RegisterPartial("app", "services")

	RegisterTemplate("apps", "layout", "apps")
	RegisterTemplate("app", "layout", "app")

	log = logger.New("ns=kernel cn=app")
}

func AppList(rw http.ResponseWriter, r *http.Request) {
	log = log.At("list").Start()
	apps, err := models.ListApps()

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	sort.Sort(apps)

	RenderTemplate(rw, "apps", apps)
}

func AppShow(rw http.ResponseWriter, r *http.Request) {
	log = log.At("show").Start()
	name := mux.Vars(r)["app"]

	app, err := models.GetApp(name)

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	RenderTemplate(rw, "app", app)
}

func AppCreate(rw http.ResponseWriter, r *http.Request) {
	log = log.At("create").Start()
	name := GetForm(r, "name")
	repo := GetForm(r, "repo")

	app := &models.App{
		Name:       name,
		Repository: repo,
	}

	err := app.Create()

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	Redirect(rw, r, "/apps")
}

func AppDelete(rw http.ResponseWriter, r *http.Request) {
	log = log.At("delete").Start()
	vars := mux.Vars(r)
	name := vars["app"]

	app, err := models.GetApp(name)

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	log.Success("step=app.get app=%q", app.Name)

	err = app.Delete()

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	log.Success("step=app.delete app=%q", app.Name)

	RenderText(rw, "ok")
}

func AppPromote(rw http.ResponseWriter, r *http.Request) {
	log = log.At("promote").Start()
	app := mux.Vars(r)["app"]

	release, err := models.GetRelease(app, GetForm(r, "release"))

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}
	log.Success("step=release.get app=%q release=%q", release.App, release.Id)

	change := &models.Change{
		App:      app,
		Created:  time.Now(),
		Metadata: "{}",
		TargetId: release.Id,
		Type:     "PROMOTE",
		Status:   "changing",
		User:     "web",
	}

	change.Save()

	events, err := models.ListEvents(app)
	if err != nil {
		log.Error(err)
		change.Status = "failed"
		change.Metadata = err.Error()
		change.Save()

		RenderError(rw, err)
		return
	}
	log.Success("step=events.list app=%q release=%q", release.App, release.Id)

	err = release.Promote()

	if err != nil {
		log.Error(err)
		change.Status = "failed"
		change.Metadata = fmt.Sprintf("{\"error\": \"%s\"}", err.Error())
		change.Save()

		RenderError(rw, err)
		return
	}
	log.Success("step=release.promote app=%q release=%q", release.App, release.Id)

	Redirect(rw, r, fmt.Sprintf("/apps/%s", app))

	a, err := models.GetApp(app)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	log.Success("step=app.get app=%q", release.App)
	go a.WatchForCompletion(change, events)
}

func AppBuilds(rw http.ResponseWriter, r *http.Request) {
	log = log.At("builds").Start()
	app := mux.Vars(r)["app"]

	builds, err := models.ListBuilds(app)

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	RenderPartial(rw, "app", "builds", builds)
}

func AppChanges(rw http.ResponseWriter, r *http.Request) {
	log = log.At("changes").Start()
	app := mux.Vars(r)["app"]

	changes, err := models.ListChanges(app)

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	RenderPartial(rw, "app", "changes", changes)
}

func AppLogs(rw http.ResponseWriter, r *http.Request) {
	log = log.At("logs")
	app := mux.Vars(r)["app"]

	RenderPartial(rw, "app", "logs", app)
}

func AppLogStream(rw http.ResponseWriter, r *http.Request) {
	log = log.At("log stream").Start()
	app, err := models.GetApp(mux.Vars(r)["app"])

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}
	log.Success("step=app.get app=%q", app)

	logs := make(chan []byte)
	done := make(chan bool)

	app.SubscribeLogs(logs, done)

	ws, err := upgrader.Upgrade(rw, r, nil)

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}
	log.Success("step=upgrade app=%q", app)

	defer ws.Close()

	for data := range logs {
		ws.WriteMessage(websocket.TextMessage, data)
	}

	fmt.Println("ended")
}

func AppReleases(rw http.ResponseWriter, r *http.Request) {
	log = log.At("releases").Start()
	app := mux.Vars(r)["app"]

	releases, err := models.ListReleases(app)

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	RenderPartial(rw, "app", "releases", releases)
}

func AppResources(rw http.ResponseWriter, r *http.Request) {
	log = log.At("resources").Start()
	app := mux.Vars(r)["app"]

	resources, err := models.ListResources(app)

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	RenderPartial(rw, "app", "resources", resources)
}

func AppServices(rw http.ResponseWriter, r *http.Request) {
	log = log.At("services").Start()
	app := mux.Vars(r)["app"]

	services, err := models.ListServices(app)

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	RenderPartial(rw, "app", "services", services)
}

func AppStatus(rw http.ResponseWriter, r *http.Request) {
	log = log.At("status").Start()
	app, err := models.GetApp(mux.Vars(r)["app"])

	if err != nil {
		log.Error(err)
		RenderError(rw, err)
		return
	}

	RenderText(rw, app.Status)
}
