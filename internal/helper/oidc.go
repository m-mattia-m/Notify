package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"message-proxy/internal/model"
	"strings"
)

// GetClaims
// TODO: There are also possibilities to get the claims out without verifying the token,
//   - if this is guaranteed at an earlier point.
func GetClaims(bearer string) (*model.BearerClaims, error) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, viper.GetString("authentication.oidc.issuer"))
	if err != nil {
		return nil, fmt.Errorf("can't create new provider -> %s", err)
	}

	insecureSkipSignatureCheck := viper.GetString("app.env") == "DEV"
	var verifier = provider.Verifier(&oidc.Config{
		ClientID:                   viper.GetString("authentication.oidc.clientId"),
		InsecureSkipSignatureCheck: insecureSkipSignatureCheck,
	})

	IDToken, err := verifier.Verify(ctx, bearer)
	if err != nil {
		return nil, err
	}

	var claims model.BearerClaims
	if err := IDToken.Claims(&claims); err != nil {
		return nil, fmt.Errorf("can't get custom claims -> %s", err)
	}

	return &claims, nil
}

func GetBearer(c *gin.Context) string {
	authToken := c.Request.Header.Get("Authorization")
	if authToken == "" {
		return ""
	}

	authTokenSections := strings.Split(authToken, " ")
	if len(authTokenSections) != 2 {
		return ""
	}

	return authTokenSections[1]
}

func GetRoles(rolesInterface interface{}) ([]string, error) {
	var rolesMap map[string]interface{}
	roleBytes, err := json.Marshal(rolesInterface)
	if err != nil {
		return nil, fmt.Errorf("can't marshal roles to []bytes: %s", err)
	}

	err = json.Unmarshal(roleBytes, &rolesMap)
	if err != nil {
		return nil, err
	}

	var roleNames []string
	for rolesMapName, _ := range rolesMap {
		roleNames = append(roleNames, rolesMapName)
	}
	return roleNames, nil
}
