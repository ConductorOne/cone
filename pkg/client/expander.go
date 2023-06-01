package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/conductorone/cone/internal/c1api"
	"golang.org/x/sync/errgroup"
)

const (
	expanderConcurrency = 10
)

type Expander struct {
	keys    sync.Map
	fetched sync.Map
}

type appExpandKey struct {
	appID string
}

type hasAppID interface {
	GetAppId() string
}

func (ee *Expander) ExpandApp(v hasAppID) {
	ee.keys.Store(appExpandKey{v.GetAppId()}, true)
}

type resourceTypeExpandKey struct {
	appID          string
	resourceTypeID string
}

type hasResourceTypeID interface {
	GetAppId() string
	GetAppResourceTypeId() string
}

func (ee *Expander) ExpandResourceType(v hasResourceTypeID) {
	ee.keys.Store(resourceTypeExpandKey{v.GetAppId(), v.GetAppResourceTypeId()}, true)
}

type resourceExpandKey struct {
	appID          string
	resourceTypeID string
	resourceID     string
}

type hasResourceID interface {
	GetAppId() string
	GetAppResourceTypeId() string
	GetAppResourceId() string
}

func (ee *Expander) ExpandResource(v hasResourceID) {
	ee.keys.Store(resourceExpandKey{v.GetAppId(), v.GetAppResourceTypeId(), v.GetAppResourceId()}, true)
}

func (ee *Expander) GetApp(appID string) (*c1api.C1ApiAppV1App, bool) {
	if app, ok := ee.fetched.Load(appExpandKey{appID}); ok {
		return app.(*c1api.C1ApiAppV1App), true
	}
	return &c1api.C1ApiAppV1App{}, false
}

func (ee *Expander) GetResourceType(appID string, resourceTypeID string) (*c1api.C1ApiAppV1AppResourceType, bool) {
	if rt, ok := ee.fetched.Load(resourceTypeExpandKey{appID, resourceTypeID}); ok {
		return rt.(*c1api.C1ApiAppV1AppResourceType), true
	}
	return &c1api.C1ApiAppV1AppResourceType{}, false
}

func (ee *Expander) GetResource(appID string, resourceTypeID string, resourceID string) (*c1api.C1ApiAppV1AppResource, bool) {
	if r, ok := ee.fetched.Load(resourceExpandKey{appID, resourceTypeID, resourceID}); ok {
		return r.(*c1api.C1ApiAppV1AppResource), true
	}
	return &c1api.C1ApiAppV1AppResource{}, false
}

func (ee *Expander) Run(ctx context.Context, c C1Client) error {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(expanderConcurrency)

	ee.keys.Range(func(k, v interface{}) bool {
		switch v := k.(type) {
		case appExpandKey:
			g.Go(func() error {
				app, err := c.GetApp(ctx, v.appID)
				if err != nil {
					return err
				}
				ee.fetched.Store(v, app)
				return nil
			})
		case resourceTypeExpandKey:
			g.Go(func() error {
				rt, err := c.GetResourceType(ctx, v.appID, v.resourceTypeID)
				if err != nil {
					return err
				}
				ee.fetched.Store(v, rt)
				return nil
			})
		case resourceExpandKey:
			g.Go(func() error {
				r, err := c.GetResource(ctx, v.appID, v.resourceTypeID, v.resourceID)
				if err != nil {
					return err
				}
				ee.fetched.Store(v, r)
				return nil
			})
		default:
			g.Go(func() error {
				return fmt.Errorf("unknown expand key: %v", k)
			})
			return false
		}
		return true
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	return nil
}
