package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"notify/internal/helper"
	"notify/internal/model"
	"notify/internal/service"
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
	bearer := helper.GetBearer(c)
	if bearer == "" {
		return nil, fmt.Errorf("failed to get bearer token")
	}

	claims, err := helper.GetClaims(bearer)
	if err != nil {
		return nil, err
	}

	roles, err := helper.GetRoles(claims.UrnZitadelIamOrgProjectRoles)
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
