package session

import (
	"context"
	"crypto/rsa"
	"go-api/domain"	
	"go-api/services/user"
	"github.com/jinzhu/gorm"
	"gopkg.in/unrolled/render.v1"
)

type Options struct {
	BasePath                      string
	TokenAuthority                ITokenAuthority
	PrivateSigningKey             *rsa.PrivateKey
	PublicSigningKey              *rsa.PublicKey
	Database                      *gorm.DB
	Renderer                      *render.Render
	TokenRepositoryFactory 		  ITokenRepositoryFactory
	UserRepositoryFactory 		  user.IUserRepositoryFactory
}

func NewResource(ctx context.Context, options *Options) *SessionResource {

	database := options.Database
	if database == nil {
		panic("sessions.Options.Database is required")
	}
	
	renderer := options.Renderer
	if renderer == nil {
		panic("sessions.Options.Renderer is required")
	}
	
	userRepositoryFactory := options.UserRepositoryFactory
	if userRepositoryFactory == nil {
		userRepositoryFactory = user.NewUserRepositoryFactory()
	}
	
	tokenRepositoryFactory := options.TokenRepositoryFactory
	if tokenRepositoryFactory == nil {
		// init default TokenRepositoryFactory
		tokenRepositoryFactory = NewTokenRepositoryFactory()
	}

	tokenAuthority := options.TokenAuthority
	if tokenAuthority == nil {
		if options.PrivateSigningKey == nil {
			panic("sessions.options.PrivateSigningKey is required")
		}
		if options.PublicSigningKey == nil {
			panic("sessions.options.PublicSigningKey is required")
		}
    
		tokenAuthority = NewTokenAuthority(&TokenAuthorityOptions{
			PrivateSigningKey: options.PrivateSigningKey,
			PublicSigningKey:  options.PublicSigningKey,
			Database: database,
			UserRepositoryFactory: userRepositoryFactory,
			TokenRepositoryFactory: tokenRepositoryFactory,
		})
	}
	
	resource := &SessionResource{ctx, options, nil,
		database,
		renderer,
		tokenAuthority,
		tokenRepositoryFactory,
		userRepositoryFactory,
	}
	resource.generateRoutes(options.BasePath)
	return resource
}

// SessionResource implements IResource
type SessionResource struct {
	ctx                           context.Context
	options                       *Options
	routes                        *domain.Routes
	Database                      *gorm.DB
	Renderer                      *render.Render
	TokenAuthority                ITokenAuthority
	TokenRepositoryFactory 		  ITokenRepositoryFactory
	UserRepositoryFactory         user.IUserRepositoryFactory
}

func (resource *SessionResource) Context() context.Context {
	return resource.ctx
}

func (resource *SessionResource) Routes() *domain.Routes {
	return resource.routes
}

func (resource *SessionResource) TokenRepository() ITokenRepository {
	return resource.TokenRepositoryFactory.New(resource.Database)
}

func (resource *SessionResource) TokenRepositoryTransaction(tx *gorm.DB) ITokenRepository {
	return resource.TokenRepositoryFactory.New(tx)
}

func (resource *SessionResource) UserRepository() user.IUserRepository {
	return resource.UserRepositoryFactory.New(resource.Database)
}

func (resource *SessionResource) UserRepositoryTransaction(tx *gorm.DB) user.IUserRepository {
	return resource.UserRepositoryFactory.New(tx)
}

func (resource *SessionResource) Render() *render.Render {
	return resource.Renderer
}