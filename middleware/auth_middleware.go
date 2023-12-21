package middleware

import (
	"net/http"

	"github.com/adityaqudaedah/go_restful_api/helpers"
	"github.com/adityaqudaedah/go_restful_api/model/web"
)

type AuthMiddleWare struct {
	Handler http.Handler
}

func NewAuthMiddleWare(handler http.Handler)  *AuthMiddleWare{
	return &AuthMiddleWare{
		Handler: handler,
	}
}

func (authMiddleware *AuthMiddleWare) ServeHTTP(w http.ResponseWriter,req *http.Request){
	

	if apiKey := req.Header.Get("MAMAT-API-KEY"); apiKey != "mamat-rahmat"{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code : http.StatusUnauthorized,
			Message: "UNAUTHORIZED",
			Status: "ERROR",
		}
		
		helpers.WriteToResponseBody(w,webResponse)
		
	}else{
		authMiddleware.Handler.ServeHTTP(w,req)
	}
}