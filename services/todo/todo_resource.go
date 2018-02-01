package todo

import (
	"context"
	"go-api/domain"
	"github.com/jinzhu/gorm"
	"gopkg.in/unrolled/render.v1"
)

type Options struct {
	BasePath              string
	Database              *gorm.DB
	Renderer              *render.Render
	TodoRepositoryFactory ITodoRepositoryFactory
}

func NewResource(ctx context.Context, options *Options) *TodoResource {

	database := options.Database
	if &database == nil {
		panic("todos.Options.Database is required")
	}
	
	renderer := options.Renderer
	if renderer == nil {
		panic("todos.Options.Renderer is required")
	}

	todoRepositoryFactory := options.TodoRepositoryFactory
	if todoRepositoryFactory == nil {
		// init default TodoRepositoryFactory
		todoRepositoryFactory = NewTodoRepositoryFactory()
	}

	u := &TodoResource{ctx, options, nil,
		database,
		renderer,		
		todoRepositoryFactory,
	}
	u.generateRoutes(options.BasePath)
	return u
}

// TodosResource implements IResource
type TodoResource struct {
	ctx                   context.Context
	options               *Options
	routes                *domain.Routes
	Database              *gorm.DB
	Renderer              *render.Render
	TodoRepositoryFactory ITodoRepositoryFactory
}

func (resource *TodoResource) Context() context.Context {
	return resource.ctx
}

func (resource *TodoResource) Routes() *domain.Routes {
	return resource.routes
}

func (resource *TodoResource) Render() *render.Render {
	return resource.Renderer
}

func (resource *TodoResource) TodoRepository() ITodoRepository {
	return resource.TodoRepositoryFactory.New(resource.Database)
}

func (resource *TodoResource) TodoRepositoryTransaction(tx *gorm.DB) ITodoRepository {
	return resource.TodoRepositoryFactory.New(tx)
}