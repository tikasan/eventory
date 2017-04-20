// Code generated by goagen v1.1.0, DO NOT EDIT.
//
// API "eventory": cron Resource Client
//
// Command:
// $ goagen
// --design=github.com/tikasan/eventory/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty

package client

import (
	"golang.org/x/net/context"
	"fmt"
	"net/http"
	"net/url"
)

// AppendGenreCronPath computes a request path to the append genre action of cron.
func AppendGenreCronPath() string {

	return fmt.Sprintf("/api/v2/cron/events/appendgenre")
}

// <b>イベントにジャンルを付加する<b>
func (c *Client) AppendGenreCron(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAppendGenreCronRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAppendGenreCronRequest create the request corresponding to the append genre action endpoint of the cron resource.
func (c *Client) NewAppendGenreCronRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.CronTokenSigner != nil {
		c.CronTokenSigner.Sign(req)
	}
	return req, nil
}

// FixUserKeepCronPath computes a request path to the fix user keep action of cron.
func FixUserKeepCronPath() string {

	return fmt.Sprintf("/api/v2/cron/user/events/fixkeep")
}

// <b>ユーザーのイベントのキープ操作の確定</b><br>
// user_keep_statusesテーブルのbatch_processedをtrueに変更する
func (c *Client) FixUserKeepCron(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFixUserKeepCronRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFixUserKeepCronRequest create the request corresponding to the fix user keep action endpoint of the cron resource.
func (c *Client) NewFixUserKeepCronRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.CronTokenSigner != nil {
		c.CronTokenSigner.Sign(req)
	}
	return req, nil
}

// NewEventFetchCronPath computes a request path to the new event fetch action of cron.
func NewEventFetchCronPath() string {

	return fmt.Sprintf("/api/v2/cron/events/fetch")
}

// <b>最新イベント情報の取得<b>
func (c *Client) NewEventFetchCron(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewNewEventFetchCronRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewNewEventFetchCronRequest create the request corresponding to the new event fetch action endpoint of the cron resource.
func (c *Client) NewNewEventFetchCronRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.CronTokenSigner != nil {
		c.CronTokenSigner.Sign(req)
	}
	return req, nil
}