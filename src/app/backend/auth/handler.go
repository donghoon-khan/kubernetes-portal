package auth

import (
	"log"
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"

	authApi "github.com/donghoon-khan/kubeportal/src/app/backend/auth/api"
	"github.com/donghoon-khan/kubeportal/src/app/backend/docs"
)

type AuthHandler struct {
	manager authApi.AuthManager
}

func (authHandler AuthHandler) Install(ws *restful.WebService) {
	ws.Route(
		ws.POST("/login").
			To(authHandler.handleLogin).
			Reads(authApi.LoginSpec{}).
			Writes(authApi.AuthResponse{}).
			Doc("Get JWEToken by LoginSpec").
			Metadata(restfulspec.KeyOpenAPITags, docs.AuthenticationDocsTag).
			Returns(200, "OK", authApi.AuthResponse{}))
	ws.Route(
		ws.GET("/login/skippable").
			To(authHandler.handleLoginSkippable).
			Writes(authApi.LoginSkippableResponse{}).
			Doc("Is enable login skippable").
			Metadata(restfulspec.KeyOpenAPITags, docs.AuthenticationDocsTag).
			Returns(200, "OK", authApi.LoginSkippableResponse{}))
}

func (authHandler AuthHandler) handleLogin(request *restful.Request, resposne *restful.Response) {
	log.Println("Handle Login")
}

func (authHandler *AuthHandler) handleLoginSkippable(request *restful.Request, response *restful.Response) {
	response.WriteHeaderAndEntity(http.StatusOK,
		authApi.LoginSkippableResponse{Skippable: authHandler.manager.AuthenticationSkippable()})
}

func NewAuthHandler(manager authApi.AuthManager) AuthHandler {
	return AuthHandler{manager: manager}
}
