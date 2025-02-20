package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/gijan/templui/assets"
	"github.com/gijan/templui/internal/config"
	"github.com/gijan/templui/internal/middleware"
	"github.com/gijan/templui/internal/ui/pages"
	"github.com/gijan/templui/pkg/components"
	mw "github.com/gijan/templui/pkg/middleware"
)

func toastDemoHandler(w http.ResponseWriter, r *http.Request) {
	duration, err := strconv.Atoi(r.FormValue("duration"))
	if err != nil {
		duration = 0
	}

	toastProps := components.ToastProps{
		Message:     r.FormValue("message"),
		Type:        r.FormValue("type"),
		Position:    r.FormValue("position"),
		Duration:    duration,
		Size:        r.FormValue("size"),
		Dismissible: r.FormValue("dismissible") == "on",
		Icon:        r.FormValue("icon") == "on",
	}

	components.Toast(toastProps).Render(r.Context(), w)
}

func main() {
	mux := http.NewServeMux()
	config.LoadConfig()
	SetupAssetsRoutes(mux)

	cspConfig := mw.CSPConfig{
		ScriptSrc: []string{
			"cdn.jsdelivr.net",     // HTMX
			"cdnjs.cloudflare.com", // highlight.js
		},
	}

	wrappedMux := middleware.WithURLPathValue(
		middleware.CacheControlMiddleware(
			mw.WithCSP(cspConfig)(mux),
		),
	)

	mux.Handle("GET /", templ.Handler(pages.Landing()))
	mux.Handle("GET /docs/getting-started", http.RedirectHandler("/docs/introduction", http.StatusSeeOther))
	mux.Handle("GET /docs/components", http.RedirectHandler("/docs/components/accordion", http.StatusSeeOther))
	mux.Handle("GET /docs/introduction", templ.Handler(pages.Introduction()))
	mux.Handle("GET /docs/how-to-use", templ.Handler(pages.HowToUse()))
	mux.Handle("GET /docs/themes", templ.Handler(pages.Themes()))
	// Components
	mux.Handle("GET /docs/components/accordion", templ.Handler(pages.Accordion()))
	mux.Handle("GET /docs/components/alert", templ.Handler(pages.Alert()))
	mux.Handle("GET /docs/components/avatar", templ.Handler(pages.Avatar()))
	mux.Handle("GET /docs/components/badge", templ.Handler(pages.Badge()))
	mux.Handle("GET /docs/components/button", templ.Handler(pages.Button()))
	mux.Handle("GET /docs/components/card", templ.Handler(pages.Card()))
	mux.Handle("GET /docs/components/code", templ.Handler(pages.Code()))
	mux.Handle("GET /docs/components/checkbox", templ.Handler(pages.Checkbox()))
	mux.Handle("GET /docs/components/datepicker", templ.Handler(pages.Datepicker()))
	mux.Handle("GET /docs/components/dropdown-menu", templ.Handler(pages.DropdownMenu()))
	mux.Handle("GET /docs/components/form", templ.Handler(pages.Form()))
	mux.Handle("GET /docs/components/icon", templ.Handler(pages.Icon()))
	mux.Handle("GET /docs/components/input", templ.Handler(pages.Input()))
	mux.Handle("GET /docs/components/label", templ.Handler(pages.Label()))
	mux.Handle("GET /docs/components/modal", templ.Handler(pages.Modal()))
	mux.Handle("GET /docs/components/radio", templ.Handler(pages.Radio()))
	mux.Handle("GET /docs/components/select", templ.Handler(pages.Select()))
	mux.Handle("GET /docs/components/sheet", templ.Handler(pages.Sheet()))
	mux.Handle("GET /docs/components/slider", templ.Handler(pages.Slider()))
	mux.Handle("GET /docs/components/tabs", templ.Handler(pages.Tabs()))
	mux.Handle("GET /docs/components/textarea", templ.Handler(pages.Textarea()))
	mux.Handle("GET /docs/components/timepicker", templ.Handler(pages.Timepicker()))
	mux.Handle("GET /docs/components/toast", templ.Handler(pages.Toast()))
	mux.Handle("GET /docs/components/toggle", templ.Handler(pages.Toggle()))
	mux.Handle("GET /docs/components/tooltip", templ.Handler(pages.Tooltip()))
	// Showcase API
	mux.Handle("POST /docs/toast/demo", http.HandlerFunc(toastDemoHandler))

	fmt.Println("Server is running on http://localhost:8090")
	http.ListenAndServe(":8090", wrappedMux)
}

func SetupAssetsRoutes(mux *http.ServeMux) {
	var isDevelopment = config.AppConfig.GoEnv != "production"

	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}

		var fs http.Handler
		if isDevelopment {
			fs = http.FileServer(http.Dir("./assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}

		fs.ServeHTTP(w, r)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))
}
