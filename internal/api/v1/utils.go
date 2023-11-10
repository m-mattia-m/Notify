package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"message-proxy/internal/model"
	"message-proxy/internal/service"
	"strings"
)

func getServiceAndUser(c *gin.Context, ifGetUser bool) (*service.Client, *model.OidcUser, error) {
	svc, err := getService(c)
	if err != nil {
		return nil, nil, err
	}

	if !ifGetUser {
		return svc, nil, nil
	}

	user, err := getUser(c)
	if err != nil {
		return nil, nil, err
	}

	return svc, user, nil
}

func getService(c *gin.Context) (*service.Client, error) {
	svc, found := c.Get("svc")
	if !found {
		return nil, fmt.Errorf("failed to get services from request-context")
	}
	svcObject := svc.(*service.Client)

	return svcObject, nil
}

func getUser(c *gin.Context) (*model.OidcUser, error) {
	bearer := getBearer(c)
	if bearer == "" {
		return nil, fmt.Errorf("failed to get bearer token")
	}

	claims, err := getClaims(bearer)
	if err != nil {
		return nil, err
	}

	roles, err := getRoles(claims.UrnZitadelIamOrgProjectRoles)
	if err != nil {
		return nil, fmt.Errorf("failed to get roles from bearer-claims %s", err.Error())
	}

	return &model.OidcUser{
		Email:             claims.Email,
		EmailVerified:     claims.EmailVerified,
		FamilyName:        claims.FamilyName,
		Gender:            claims.Gender,
		GivenName:         claims.GivenName,
		Locale:            claims.Locale,
		Name:              claims.Name,
		Nickname:          claims.Nickname,
		PreferredUsername: claims.PreferredUsername,
		Sub:               claims.Sub,
		//PrimaryDomain:     claims.UrnZitadelIamOrgDomainPrimary,
		UpdatedAt: claims.UpdatedAt,
		Roles:     roles,
	}, nil
}

// TODO: There are also possibilities to get the claims out without verifying the token,
// if this is guaranteed at an earlier point.
func getClaims(bearer string) (*model.BearerClaims, error) {
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

	var claims model.BearerClaims // map[string]interface{} // model.BearerClaims
	if err := IDToken.Claims(&claims); err != nil {
		return nil, fmt.Errorf("can't get custom claims -> %s", err)
	}

	return &claims, nil
}

func getBearer(c *gin.Context) string {
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

func getRoles(rolesInterface interface{}) ([]string, error) {
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
