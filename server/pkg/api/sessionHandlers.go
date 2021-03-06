package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/darwinfroese/hacksite/server/pkg/auth"
)

var sessionHandlersMap = map[string]handler{
	"GET":     sessionHandler,
	"OPTIONS": optionsHandler,
}

func sessionRoute(ctx apiContext, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return callHandler(ctx, w, r, sessionHandlersMap)
}

func sessionHandler(ctx apiContext, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	session, err := auth.GetCurrentSession(r, ctx.db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if time.Now().After(session.Expiration) {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	session.Expiration = time.Now().Add(time.Second * auth.SessionMaxAge)
	err = ctx.db.StoreSession(session)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	auth.SetCookie(w, auth.SessionCookieName, session.Token)

	w.WriteHeader(http.StatusOK)
}
