package responders

import (
	"github.com/nmonterroso/pacowsay/restapi/operations"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/nmonterroso/pacowsay/models"
)

func CowsayResponder(params operations.CowsayParams) middleware.Responder {
	response := &operations.CowsayOK{
		Payload: &models.SLACKResponse{
			Text: "`moo`",
			ResponseType: "in_channel",
		},
	}


	return response
}
