package auth

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"notify/internal/helper"
	"notify/internal/model"
)

func Authenticate(c *gin.Context) {

	oidcClientId := viper.GetString("authentication.oidc.clientId")
	oidcAuthority := viper.GetString("authentication.oidc.issuer")
	runMode := viper.GetString("app.env")

	bearer := helper.GetBearer(c)
	if bearer == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.HttpError{Message: "Unauthorized"})
		return
	}

	_, err := validateToken(c.Request.Context(), bearer, oidcClientId, oidcAuthority, runMode)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.HttpError{Message: "Unauthorized"})
		return
	}

}

func validateToken(ctx context.Context, bearer, oidcClientId, oidcAuthority, runMode string) (*oidc.Provider, error) {
	provider, err := oidc.NewProvider(ctx, oidcAuthority)
	if err != nil {
		return nil, fmt.Errorf("can't create new provider -> %s", err)
	}

	insecureSkipSignatureCheck := runMode == "DEV"
	var verifier = provider.Verifier(&oidc.Config{ClientID: oidcClientId, InsecureSkipSignatureCheck: insecureSkipSignatureCheck})

	_, err = verifier.Verify(ctx, bearer)
	if err != nil {
		return nil, fmt.Errorf("can't verify bearer -> %s", err)
	}

	return provider, nil
}
