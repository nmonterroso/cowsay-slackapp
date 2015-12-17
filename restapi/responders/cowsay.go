package responders

import (
	"github.com/nmonterroso/pacowsay/restapi/operations"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/nmonterroso/pacowsay/models"
	"github.com/nmonterroso/cowsay/lib"
	"net/http"
	"fmt"
)

func CowsayResponder(params operations.CowsayParams) middleware.Responder {
	cowsay, err := lib.Cowsay(params.Text)

	if err != nil {
		return &operations.CowsayDefault{
			&models.Error{
				Code: http.StatusInternalServerError,
				Message: err.Error(),
			},
		}
	}

	return &operations.CowsayOK{
		Payload: &models.SLACKResponse{
			Text: fmt.Sprintf("```%s```", cowsay),
			ResponseType: "ephemeral",
		},
	}
}
