package resource

import (
	"golnfuturecapacities/api/service"
	"net/http"
)

type WelcomeResource struct {
	WelcomeService service.WelcomeService
}

func WelcomeController(welcomeService service.WelcomeService) *WelcomeResource {
	return &WelcomeResource{
		WelcomeService: welcomeService,
	}
}

func (app *WelcomeResource) WelcomeGetHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Welcome Resource"))
	if err != nil {
		return
	}
}
