package services

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/thompsch/app-tester/configs"
	"net/http"
)

func SetupWebServerAndListen() {
	configs.LoadEnvironment()
	InitializeGoogleLogger()
	http.HandleFunc(configs.WebserverPath, ParseWebhookData)
	port := configs.Port
	if port != "" {
		port = ":" + port
	}

	LogInfo(fmt.Sprintf("Starting web server on port %s", port))

	e := http.ListenAndServe(port, nil)
	if e != nil && !errors.Is(e, http.ErrServerClosed) {
		LogCritical(fmt.Sprintf("Error starting server:", e))
	} else {
		LogInfo(fmt.Sprintf("Web server listening on " + configs.WebserverPath))
	}
}
