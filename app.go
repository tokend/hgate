package hgate

import (
	"gitlab.com/tokend/hgate/config"
	"gitlab.com/tokend/hgate/server"
	"gitlab.com/tokend/go/signcontrol"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"
	"log"
	"net/http"
	"net/http/httputil"
)

type App struct {
	Conf *config.GateConfig
	Log  *logan.Entry
}

func NewApp(configPath string) (app *App, err error) {
	app = new(App)

	app.Conf, err = config.InitConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("can not initialize config: %s", err.Error())
	}

	initLog(app)

	return
}

func (app *App) Serve() {
	mux := server.NewMux()

	mux.HandleFunc("/", app.RedirectHandler)

	fmt.Println("hgate server listening at:" + app.Conf.Port)
	err := http.ListenAndServe("localhost:"+app.Conf.Port, mux)
	log.Fatal(err)
}

/*
lEntry := app.Log.WithField("service", "RedirectHandler")
	lEntry.WithField("path", r.URL.Path).Info("Started request")

	request, err := http.NewRequest("GET", "https://api.swarm.fund/users", nil)
	if err != nil {
		panic(err)
	}

	err = signcontrol.SignRequest(request, app.Conf.KP)
	if err != nil {
		lEntry.WithError(err).Error("SignRequest failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	lEntry.WithField("data", string(data)).Error("AAAAAAAAAA")

	w.Write(data)
	lEntry.WithField("path", r.URL.Path).Info("Finished request")
 */

func (app *App) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	lEntry := app.Log.WithField("service", "RedirectHandler")
	lEntry.WithField("path", r.URL.Path).Info("Started request")

	err := signcontrol.SignRequest(r, app.Conf.KP)
	if err != nil {
		lEntry.WithError(err).Error("SignRequest failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lEntry.Error(app.Conf.HUrl.String())
	proxy := httputil.NewSingleHostReverseProxy(app.Conf.HUrl)
	proxy.ServeHTTP(w, r)
	lEntry.WithField("path", r.URL.Path).Info("Finished request")
}

func initLog(app *App) {
	app.Log = logan.New().Level(app.Conf.LL).WithField("application", "hgate")
}
