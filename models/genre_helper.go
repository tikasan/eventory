// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/tikasan/eventory/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": Model Helpers
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
	"golang.org/x/net/context"
)

// MediaType Retrieval Functions

// ListGenre returns an array of view: default.
func (m *GenreDB) ListGenre(ctx context.Context) []*app.Genre {
	defer goa.MeasureSince([]string{"goa", "db", "genre", "listgenre"}, time.Now())

	var native []*Genre
	var objs []*app.Genre
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Genre", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.GenreToGenre())
	}

	return objs
}

// GenreToGenre loads a Genre and builds the default view of media type Genre.
func (m *Genre) GenreToGenre() *app.Genre {
	genre := &app.Genre{}
	genre.Name = &m.Name

	return genre
}

// OneGenre loads a Genre and builds the default view of media type Genre.
func (m *GenreDB) OneGenre(ctx context.Context, id int) (*app.Genre, error) {
	defer goa.MeasureSince([]string{"goa", "db", "genre", "onegenre"}, time.Now())

	var native Genre
	err := m.Db.Scopes().Table(m.TableName()).Preload("EventGenres").Preload("UserFollowGenres").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Genre", "error", err.Error())
		return nil, err
	}

	view := *native.GenreToGenre()
	return &view, err
}