// Code generated by goagen v1.1.0, DO NOT EDIT.
//
// API "eventory": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/tikasan/eventory/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty

package app

import (
	"golang.org/x/net/context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
	"unicode/utf8"
)

// AppendGenreCronContext provides the cron append genre action context.
type AppendGenreCronContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewAppendGenreCronContext parses the incoming request URL and body, performs validations and creates the
// context used by the cron controller append genre action.
func NewAppendGenreCronContext(ctx context.Context, r *http.Request, service *goa.Service) (*AppendGenreCronContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AppendGenreCronContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AppendGenreCronContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *AppendGenreCronContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *AppendGenreCronContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// FixUserFollowCronContext provides the cron fix user follow action context.
type FixUserFollowCronContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewFixUserFollowCronContext parses the incoming request URL and body, performs validations and creates the
// context used by the cron controller fix user follow action.
func NewFixUserFollowCronContext(ctx context.Context, r *http.Request, service *goa.Service) (*FixUserFollowCronContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := FixUserFollowCronContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *FixUserFollowCronContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *FixUserFollowCronContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *FixUserFollowCronContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NewEventFetchCronContext provides the cron new event fetch action context.
type NewEventFetchCronContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewNewEventFetchCronContext parses the incoming request URL and body, performs validations and creates the
// context used by the cron controller new event fetch action.
func NewNewEventFetchCronContext(ctx context.Context, r *http.Request, service *goa.Service) (*NewEventFetchCronContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := NewEventFetchCronContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *NewEventFetchCronContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *NewEventFetchCronContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *NewEventFetchCronContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// KeepEventsContext provides the events keep action context.
type KeepEventsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	EventID int
	IsKeep  bool
}

// NewKeepEventsContext parses the incoming request URL and body, performs validations and creates the
// context used by the events controller keep action.
func NewKeepEventsContext(ctx context.Context, r *http.Request, service *goa.Service) (*KeepEventsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := KeepEventsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramEventID := req.Params["eventID"]
	if len(paramEventID) > 0 {
		rawEventID := paramEventID[0]
		if eventID, err2 := strconv.Atoi(rawEventID); err2 == nil {
			rctx.EventID = eventID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("eventID", rawEventID, "integer"))
		}
	}
	paramIsKeep := req.Params["isKeep"]
	if len(paramIsKeep) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("isKeep"))
	} else {
		rawIsKeep := paramIsKeep[0]
		if isKeep, err2 := strconv.ParseBool(rawIsKeep); err2 == nil {
			rctx.IsKeep = isKeep
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("isKeep", rawIsKeep, "boolean"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *KeepEventsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *KeepEventsContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *KeepEventsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *KeepEventsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListEventsContext provides the events list action context.
type ListEventsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID   string
	Page int
	Q    string
	Sort string
}

// NewListEventsContext parses the incoming request URL and body, performs validations and creates the
// context used by the events controller list action.
func NewListEventsContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListEventsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListEventsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	paramPage := req.Params["page"]
	if len(paramPage) == 0 {
		rctx.Page = 0
	} else {
		rawPage := paramPage[0]
		if page, err2 := strconv.Atoi(rawPage); err2 == nil {
			rctx.Page = page
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("page", rawPage, "integer"))
		}
		if rctx.Page < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`page`, rctx.Page, 1, true))
		}
	}
	paramQ := req.Params["q"]
	if len(paramQ) == 0 {
		rctx.Q = ""
	} else {
		rawQ := paramQ[0]
		rctx.Q = rawQ
	}
	paramSort := req.Params["sort"]
	if len(paramSort) == 0 {
		rctx.Sort = ""
	} else {
		rawSort := paramSort[0]
		rctx.Sort = rawSort
		if !(rctx.Sort == "created_asc" || rctx.Sort == "created_desc" || rctx.Sort == "") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`sort`, rctx.Sort, []interface{}{"created_asc", "created_desc", ""}))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListEventsContext) OK(r EventCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.event+json; type=collection")
	if r == nil {
		r = EventCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListEventsContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListEventsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// CreateGenresContext provides the genres create action context.
type CreateGenresContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Name string
}

// NewCreateGenresContext parses the incoming request URL and body, performs validations and creates the
// context used by the genres controller create action.
func NewCreateGenresContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateGenresContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateGenresContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramName := req.Params["name"]
	if len(paramName) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("name"))
	} else {
		rawName := paramName[0]
		rctx.Name = rawName
		if utf8.RuneCountInString(rctx.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`name`, rctx.Name, utf8.RuneCountInString(rctx.Name), 1, true))
		}
		if utf8.RuneCountInString(rctx.Name) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`name`, rctx.Name, utf8.RuneCountInString(rctx.Name), 30, false))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateGenresContext) OK(r *Genre) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.genre+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *CreateGenresContext) OKTiny(r *GenreTiny) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.genre+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateGenresContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *CreateGenresContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// FollowGenresContext provides the genres follow action context.
type FollowGenresContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	GenreID int
}

// NewFollowGenresContext parses the incoming request URL and body, performs validations and creates the
// context used by the genres controller follow action.
func NewFollowGenresContext(ctx context.Context, r *http.Request, service *goa.Service) (*FollowGenresContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := FollowGenresContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramGenreID := req.Params["genreID"]
	if len(paramGenreID) > 0 {
		rawGenreID := paramGenreID[0]
		if genreID, err2 := strconv.Atoi(rawGenreID); err2 == nil {
			rctx.GenreID = genreID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("genreID", rawGenreID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *FollowGenresContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *FollowGenresContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *FollowGenresContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *FollowGenresContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListGenresContext provides the genres list action context.
type ListGenresContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Page int
	Q    string
	Sort string
}

// NewListGenresContext parses the incoming request URL and body, performs validations and creates the
// context used by the genres controller list action.
func NewListGenresContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListGenresContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListGenresContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPage := req.Params["page"]
	if len(paramPage) == 0 {
		rctx.Page = 0
	} else {
		rawPage := paramPage[0]
		if page, err2 := strconv.Atoi(rawPage); err2 == nil {
			rctx.Page = page
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("page", rawPage, "integer"))
		}
		if rctx.Page < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`page`, rctx.Page, 1, true))
		}
	}
	paramQ := req.Params["q"]
	if len(paramQ) == 0 {
		rctx.Q = ""
	} else {
		rawQ := paramQ[0]
		rctx.Q = rawQ
		if utf8.RuneCountInString(rctx.Q) < 0 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`q`, rctx.Q, utf8.RuneCountInString(rctx.Q), 0, true))
		}
		if utf8.RuneCountInString(rctx.Q) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`q`, rctx.Q, utf8.RuneCountInString(rctx.Q), 30, false))
		}
	}
	paramSort := req.Params["sort"]
	if len(paramSort) == 0 {
		rctx.Sort = ""
	} else {
		rawSort := paramSort[0]
		rctx.Sort = rawSort
		if !(rctx.Sort == "created_asc" || rctx.Sort == "created_desc" || rctx.Sort == "") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`sort`, rctx.Sort, []interface{}{"created_asc", "created_desc", ""}))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListGenresContext) OK(r GenreCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.genre+json; type=collection")
	if r == nil {
		r = GenreCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ListGenresContext) OKTiny(r GenreTinyCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.genre+json; type=collection")
	if r == nil {
		r = GenreTinyCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListGenresContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListGenresContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// FollowPrefsContext provides the prefs follow action context.
type FollowPrefsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PrefID int
}

// NewFollowPrefsContext parses the incoming request URL and body, performs validations and creates the
// context used by the prefs controller follow action.
func NewFollowPrefsContext(ctx context.Context, r *http.Request, service *goa.Service) (*FollowPrefsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := FollowPrefsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPrefID := req.Params["prefID"]
	if len(paramPrefID) > 0 {
		rawPrefID := paramPrefID[0]
		if prefID, err2 := strconv.Atoi(rawPrefID); err2 == nil {
			rctx.PrefID = prefID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("prefID", rawPrefID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *FollowPrefsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *FollowPrefsContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *FollowPrefsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *FollowPrefsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// LoginUsersContext provides the users login action context.
type LoginUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Email        string
	PasswordHash string
}

// NewLoginUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller login action.
func NewLoginUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*LoginUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := LoginUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramEmail := req.Params["email"]
	if len(paramEmail) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("email"))
	} else {
		rawEmail := paramEmail[0]
		rctx.Email = rawEmail
		if err2 := goa.ValidateFormat(goa.FormatEmail, rctx.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`email`, rctx.Email, goa.FormatEmail, err2))
		}
	}
	paramPasswordHash := req.Params["password_hash"]
	if len(paramPasswordHash) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("password_hash"))
	} else {
		rawPasswordHash := paramPasswordHash[0]
		rctx.PasswordHash = rawPasswordHash
		if ok := goa.ValidatePattern(`^[a-z0-9]{64}$`, rctx.PasswordHash); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`password_hash`, rctx.PasswordHash, `^[a-z0-9]{64}$`))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LoginUsersContext) OK(r *Message) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.message+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *LoginUsersContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// RegularCreateUsersContext provides the users regular create action context.
type RegularCreateUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Email        string
	PasswordHash string
}

// NewRegularCreateUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller regular create action.
func NewRegularCreateUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*RegularCreateUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := RegularCreateUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramEmail := req.Params["email"]
	if len(paramEmail) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("email"))
	} else {
		rawEmail := paramEmail[0]
		rctx.Email = rawEmail
		if err2 := goa.ValidateFormat(goa.FormatEmail, rctx.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`email`, rctx.Email, goa.FormatEmail, err2))
		}
	}
	paramPasswordHash := req.Params["password_hash"]
	if len(paramPasswordHash) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("password_hash"))
	} else {
		rawPasswordHash := paramPasswordHash[0]
		rctx.PasswordHash = rawPasswordHash
		if ok := goa.ValidatePattern(`^[a-z0-9]{64}$`, rctx.PasswordHash); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`password_hash`, rctx.PasswordHash, `^[a-z0-9]{64}$`))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *RegularCreateUsersContext) OK(r *Message) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.message+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *RegularCreateUsersContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *RegularCreateUsersContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// StatusUsersContext provides the users status action context.
type StatusUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ClientVersion string
	Platform      string
}

// NewStatusUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller status action.
func NewStatusUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*StatusUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := StatusUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramClientVersion := req.Params["client_version"]
	if len(paramClientVersion) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("client_version"))
	} else {
		rawClientVersion := paramClientVersion[0]
		rctx.ClientVersion = rawClientVersion
	}
	paramPlatform := req.Params["platform"]
	if len(paramPlatform) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("platform"))
	} else {
		rawPlatform := paramPlatform[0]
		rctx.Platform = rawPlatform
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *StatusUsersContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *StatusUsersContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// TmpCreateUsersContext provides the users tmp create action context.
type TmpCreateUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ClientVersion string
	Identifier    string
	Platform      string
}

// NewTmpCreateUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller tmp create action.
func NewTmpCreateUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*TmpCreateUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := TmpCreateUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramClientVersion := req.Params["client_version"]
	if len(paramClientVersion) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("client_version"))
	} else {
		rawClientVersion := paramClientVersion[0]
		rctx.ClientVersion = rawClientVersion
	}
	paramIdentifier := req.Params["identifier"]
	if len(paramIdentifier) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("identifier"))
	} else {
		rawIdentifier := paramIdentifier[0]
		rctx.Identifier = rawIdentifier
		if ok := goa.ValidatePattern(`(^[a-z0-9]{16}$|^[a-z0-9\-]{36}$)`, rctx.Identifier); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`identifier`, rctx.Identifier, `(^[a-z0-9]{16}$|^[a-z0-9\-]{36}$)`))
		}
	}
	paramPlatform := req.Params["platform"]
	if len(paramPlatform) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("platform"))
	} else {
		rawPlatform := paramPlatform[0]
		rctx.Platform = rawPlatform
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *TmpCreateUsersContext) OK(r *Token) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.token+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *TmpCreateUsersContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}
