package controller

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
	"github.com/tikasan/eventory/models"
)

// EventsController implements the events resource.
type EventsController struct {
	*goa.Controller
	db *gorm.DB
}

// NewEventsController creates a events controller.
func NewEventsController(service *goa.Service, db *gorm.DB) *EventsController {
	return &EventsController{
		Controller: service.NewController("EventsController"),
		db:         db,
	}
}

// KeepEvent runs the keep event action.
func (c *EventsController) KeepEvent(ctx *app.KeepEventEventsContext) error {
	// EventsController_KeepEvent: start_implement

	// Put your logic here
	//ufeDB := models.NewUserFollowEventDB(c.db)
	//ufe := models.UserFollowEvent{}
	//ufe.EventID = ctx.EventID

	// EventsController_KeepEvent: end_implement
	return nil
}

// List runs the list action.
func (c *EventsController) List(ctx *app.ListEventsContext) error {
	// EventsController_List: start_implement

	// Put your logic here
	eventDB := models.NewEventDB(c.db)
	events, err := eventDB.ListByQ(ctx.Context, ctx.Q, ctx.Sort, ctx.Page)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// EventsController_List: end_implement
	return ctx.OK(events)
}
