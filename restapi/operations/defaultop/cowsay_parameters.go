package defaultop

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewCowsayParams creates a new CowsayParams object
// with the default values initialized.
func NewCowsayParams() CowsayParams {
	return CowsayParams{}
}

// CowsayParams contains all the bound params for the cowsay operation
// typically these are obtained from a http.Request
//
// swagger:parameters cowsay
type CowsayParams struct {
	/*the channel id the request originated from
	  Required: true
	  In: formData
	*/
	ChannelID string
	/*the name of the corresponding `channel_id`
	  Required: true
	  In: formData
	*/
	ChannelName string
	/*the command sent by the user
	  Required: true
	  In: formData
	*/
	Command string
	/*the response url of the Slack command
	  Required: true
	  In: formData
	*/
	ResponseURL string
	/*domain for the team `team_id`
	  Required: true
	  In: formData
	*/
	TeamDomain string
	/*Team ID the request originated from
	  Required: true
	  In: formData
	*/
	TeamID string
	/*the arguments to the command
	  Required: true
	  In: formData
	*/
	Text string
	/*Team command token
	  Required: true
	  In: formData
	*/
	Token string
	/*id of the user who sent the command
	  Required: true
	  In: formData
	*/
	UserID string
	/*name of the corresponding `user_id`
	  Required: true
	  In: formData
	*/
	UserName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CowsayParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return err
	}
	fds := httpkit.Values(r.Form)

	fdChannelID, fdhkChannelID, _ := fds.GetOK("channel_id")
	if err := o.bindChannelID(fdChannelID, fdhkChannelID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdChannelName, fdhkChannelName, _ := fds.GetOK("channel_name")
	if err := o.bindChannelName(fdChannelName, fdhkChannelName, route.Formats); err != nil {
		res = append(res, err)
	}

	fdCommand, fdhkCommand, _ := fds.GetOK("command")
	if err := o.bindCommand(fdCommand, fdhkCommand, route.Formats); err != nil {
		res = append(res, err)
	}

	fdResponseURL, fdhkResponseURL, _ := fds.GetOK("response_url")
	if err := o.bindResponseURL(fdResponseURL, fdhkResponseURL, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTeamDomain, fdhkTeamDomain, _ := fds.GetOK("team_domain")
	if err := o.bindTeamDomain(fdTeamDomain, fdhkTeamDomain, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTeamID, fdhkTeamID, _ := fds.GetOK("team_id")
	if err := o.bindTeamID(fdTeamID, fdhkTeamID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdText, fdhkText, _ := fds.GetOK("text")
	if err := o.bindText(fdText, fdhkText, route.Formats); err != nil {
		res = append(res, err)
	}

	fdToken, fdhkToken, _ := fds.GetOK("token")
	if err := o.bindToken(fdToken, fdhkToken, route.Formats); err != nil {
		res = append(res, err)
	}

	fdUserID, fdhkUserID, _ := fds.GetOK("user_id")
	if err := o.bindUserID(fdUserID, fdhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdUserName, fdhkUserName, _ := fds.GetOK("user_name")
	if err := o.bindUserName(fdUserName, fdhkUserName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CowsayParams) bindChannelID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("channel_id", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("channel_id", "formData", raw); err != nil {
		return err
	}

	o.ChannelID = raw

	return nil
}

func (o *CowsayParams) bindChannelName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("channel_name", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("channel_name", "formData", raw); err != nil {
		return err
	}

	o.ChannelName = raw

	return nil
}

func (o *CowsayParams) bindCommand(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("command", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("command", "formData", raw); err != nil {
		return err
	}

	o.Command = raw

	return nil
}

func (o *CowsayParams) bindResponseURL(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("response_url", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("response_url", "formData", raw); err != nil {
		return err
	}

	o.ResponseURL = raw

	return nil
}

func (o *CowsayParams) bindTeamDomain(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("team_domain", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("team_domain", "formData", raw); err != nil {
		return err
	}

	o.TeamDomain = raw

	return nil
}

func (o *CowsayParams) bindTeamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("team_id", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("team_id", "formData", raw); err != nil {
		return err
	}

	o.TeamID = raw

	return nil
}

func (o *CowsayParams) bindText(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("text", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("text", "formData", raw); err != nil {
		return err
	}

	o.Text = raw

	return nil
}

func (o *CowsayParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("token", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("token", "formData", raw); err != nil {
		return err
	}

	o.Token = raw

	return nil
}

func (o *CowsayParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("user_id", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("user_id", "formData", raw); err != nil {
		return err
	}

	o.UserID = raw

	return nil
}

func (o *CowsayParams) bindUserName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("user_name", "formData")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("user_name", "formData", raw); err != nil {
		return err
	}

	o.UserName = raw

	return nil
}