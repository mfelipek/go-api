package main

import (
	"fmt"
	"time"
	"context"
	"errors"	
	"io/ioutil"
	"go-api/server"
	"go-api/domain"
	"gopkg.in/unrolled/render.v1"
	"go-api/services/site"
	"go-api/services/todo"
	"go-api/services/session"
	"github.com/dgrijalva/jwt-go"
)

func main() {

	staticContentPath := "/root/go-api"

	// try to load signing keys for token authority
	// NOTE: DO NOT USE THESE KEYS FOR PRODUCTION! FOR DEMO ONLY
	privateSigningKeyBytes, err := ioutil.ReadFile(staticContentPath + "/keys/jwtRS256.key")
	if err != nil {
		panic(errors.New(fmt.Sprintf("Error loading private signing key: %v", err.Error())))
	}	
	privateSigningKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateSigningKeyBytes)
    if err != nil {
       panic(errors.New(fmt.Sprintf("Error parsing rsa priv key: %v", err.Error())))
    }	    
	
	publicSigningKeyBytes, err := ioutil.ReadFile(staticContentPath + "/keys/jwtRS256.key.pub")
	if err != nil {
		panic(errors.New(fmt.Sprintf("Error loading public signing key: %v", err.Error())))
	}
	publicSigningKey, err := jwt.ParseRSAPublicKeyFromPEM(publicSigningKeyBytes)
    if err != nil {
       panic(errors.New(fmt.Sprintf("Error parsing rsa pub key: %v", err.Error())))
    }

    // create new empty context
	ctx := context.Background()
	
	// create mysql conn
	dbConn, err := domain.NewMysqlConn(&domain.DbConfig{
		Username: "root",
		Password: "password",
		Host: "mysqldb",
		Port: 3306,
	})	
	defer dbConn.Close()
	
	if err != nil {
		panic(errors.New(fmt.Sprintf("Error opening mysql connection: %v", err.Error())))
	}

	// set up Renderer
	renderer := render.New(render.Options{
		Directory: staticContentPath + "/templates",
		Layout: "layouts/main",
	})

	// set up site resource
	siteResource := site.NewResource(ctx, &site.Options{
		Renderer: renderer,
		Database: dbConn,
	})
	
	// set up todo resource
	todosResource := todo.NewResource(ctx, &todo.Options{
		Renderer: renderer,
		Database: dbConn,
	})
	
	// set up session resource
	sessionResource := session.NewResource(ctx, &session.Options{
		Renderer: 			renderer,
		Database: 			dbConn,
		PrivateSigningKey:  privateSigningKey,
		PublicSigningKey:   publicSigningKey,
	})
	
	// init server
	s := server.NewServer()

	// set up router
	ac := server.NewAccessController(ctx, renderer)
	router := server.NewRouter(ctx, ac, staticContentPath + "/static")

	// add REST resources to router
	router.AddResources(siteResource, todosResource, sessionResource)
	
	// add middlewares
	s.UseMiddleware(session.NewAuthenticator(sessionResource))

	// setup router
	s.UseRouter(router)

	// bam!
	s.Run(":8080", server.Options{
		Timeout: 10*time.Second,
	})	
}