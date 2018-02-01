package site

import (
	"context"
	"go-api/domain"
	"go-api/services/user"
	"github.com/jinzhu/gorm"
	"gopkg.in/unrolled/render.v1"
)

type Options struct {
	BasePath              string
	Database              *gorm.DB
	Renderer              *render.Render
	UserRepositoryFactory user.IUserRepositoryFactory
}

func NewResource(ctx context.Context, options *Options) *SiteResource {

	database := options.Database
	if &database == nil {
		panic("todos.Options.Database is required")
	}

	renderer := options.Renderer
	if renderer == nil {
		panic("todos.Options.Renderer is required")
	}

	userRepositoryFactory := options.UserRepositoryFactory
	if userRepositoryFactory == nil {
		userRepositoryFactory = user.NewUserRepositoryFactory()
	}

	u := &SiteResource{ctx, options, nil,
		database,
		renderer,
		userRepositoryFactory,
	}
	u.generateRoutes(options.BasePath)
	return u
}

// TodosResource implements IResource
type SiteResource struct {
	ctx                   context.Context
	options               *Options
	routes                *domain.Routes
	Database              *gorm.DB
	Renderer              *render.Render
	UserRepositoryFactory user.IUserRepositoryFactory
}

func (resource *SiteResource) Context() context.Context {
	return resource.ctx
}

func (resource *SiteResource) Routes() *domain.Routes {
	return resource.routes
}

func (resource *SiteResource) Render() *render.Render {
	return resource.Renderer
}

func (resource *SiteResource) UserRepository() user.IUserRepository {
	return resource.UserRepositoryFactory.New(resource.Database)
}

func (resource *SiteResource) UserRepositoryTransaction(tx *gorm.DB) user.IUserRepository {
	return resource.UserRepositoryFactory.New(tx)
}