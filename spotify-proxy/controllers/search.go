package controllers

import (
	"errors"
	"net/http"

	"samples/spotify-proxy/models"
	"samples/spotify-proxy/services"

	"github.com/gin-gonic/gin"
)

type SearchController interface {
	Search(ctx *gin.Context)
}

type searchController struct {
	client        *http.Client
	authService   services.AuthService
	searchService services.SearchService
	config        models.Config
	logger        services.Logger
}

func NewSearchController(client *http.Client, authSvc services.AuthService, searchSvc services.SearchService,
	config models.Config, logger services.Logger) SearchController {
	return &searchController{
		client:        client,
		authService:   authSvc,
		searchService: searchSvc,
		config:        config,
		logger:        logger,
	}
}

func (ctrl *searchController) Search(ctx *gin.Context) {
	term := ctx.Params.ByName("term")
	ctrl.logger.Log("Start searching for term %s", term)
	if len(term) == 0 {
		err := errors.New("Search term cannot be empty")
		ctrl.logger.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	context := services.NewContext(ctrl.config)

	clientID := ctx.GetHeader("X-API-CLIENT_ID")
	if len(clientID) > 0 {
		context.Config.ClientID = clientID
	}
	if len(context.Config.ClientID) == 0 {
		err := errors.New("ClientID is required in X-API-CLIENT_ID header")
		ctrl.logger.Error(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
	}

	secret := ctx.GetHeader("X-API-SECRET")
	if len(secret) > 0 {
		context.Config.Secret = secret
	}
	if len(context.Config.Secret) == 0 {
		err := errors.New("Secret is required in X-API-SECRET header")
		ctrl.logger.Error(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
	}

	spotifyCtx, err := ctrl.authService.AuthenticateClient(context)
	if err != nil {
		ctrl.logger.Error(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	context.SpotifyContext = spotifyCtx

	results, err := ctrl.searchService.Search(context, term)
	if err != nil {
		ctrl.logger.Error(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	if len(results.Tracks) == 0 {
		err := errors.New("No Results Found")
		ctrl.logger.Error(err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, results)
}
