package auth

import (
	"log"
	"net/http"

	authApi "github.com/donghoon-khan/kubeportal/src/app/backend/auth/api"
	"github.com/emicklei/go-restful"
)

type AuthHandler struct {
	manager authApi.AuthManager
}

func (self AuthHandler) Install(ws *restful.WebService) {
	ws.Route(
		ws.POST("/login").
			To(self.handleLogin).
			Reads(authApi.LoginSpec{}).
			Writes(authApi.AuthResponse{}))
	ws.Route(
		ws.GET("/login/skippable").
			To(self.handleLoginSkippable).
			Writes(authApi.LoginSkippableResponse{}))
}

// handleLogin godoc
// @Tags Authentication
// @Summary Get JWEToken
// @Accept  json
// @Produce  json
// @Router /login [post]
// @Param LoginSpec body authApi.LoginSpec true "Information required to authenticate user"
// @Success 200 {object} authApi.AuthResponse
func (self AuthHandler) handleLogin(request *restful.Request, resposne *restful.Response) {
	log.Println("Handle Login")
}

func (self *AuthHandler) handleLoginSkippable(request *restful.Request, response *restful.Response) {
	response.WriteHeaderAndEntity(http.StatusOK, authApi.LoginSkippableResponse{Skippable: self.manager.AuthenticationSkippable()})
}

func NewAuthHandler(manager authApi.AuthManager) AuthHandler {
	return AuthHandler{manager: manager}
}