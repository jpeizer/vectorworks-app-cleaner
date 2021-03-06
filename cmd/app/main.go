package main

import (
	"github.com/jpeizer/vectorworks-app-cleaner/internal/software"
	"github.com/webview/webview"
	"net/http"
)

// TODO: https://github.com/electron/windows-installer

// Initialize by generating templates
func init() {
	software.GenerateTemplates()
}

// Set up and run the webview.
func main() {
	go runWebserver()

	// Set up Webview
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Vectorworks, Inc. - Application Utility Tool")
	w.SetSize(1000, 700, webview.HintFixed)
	w.Navigate("http://127.0.0.1:12346")
	w.Run()
}

// Creates all Mux objects, Handlers, configures the web server and
// deploys the http server on http://127.0.0.1:12346
func runWebserver() {
	mux := http.NewServeMux()
	// Routes
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(software.GetWD() + "/static"))))
	mux.HandleFunc("/", software.HomePageHandler) // Handle home page, also catch all
	mux.HandleFunc("/editSerial", software.EditSerialHandler)
	mux.HandleFunc("/updateSerial", software.UpdateSerialHandler)

	// Configure the webserver
	webServer := &http.Server{
		Addr:    "127.0.0.1:12346",
		Handler: mux,
	}
	//Start the web server
	err := webServer.ListenAndServe()
	software.Check(err)
}
