package model

type OidcUser struct {
	Email             string `json:"email"`
	EmailVerified     bool   `json:"email_verified"`
	FamilyName        string `json:"family_name"`
	Gender            string `json:"gender"`
	GivenName         string `json:"given_name"`
	Locale            string `json:"locale"`
	Name              string `json:"name"`
	Nickname          string `json:"nickname"`
	PreferredUsername string `json:"preferred_username"`
	Sub               string `json:"sub"`
	UpdatedAt         int    `json:"updated_at"`
	//PrimaryDomain     string   `json:"primary_domain"`
	//Roles             []string `json:"roles"`
}

type BearerClaims struct {
	Iss               string   `json:"iss"`
	Aud               []string `json:"aud"`
	Azp               string   `json:"azp"`
	AtHash            string   `json:"at_hash"`
	CHash             string   `json:"c_hash"`
	Amr               []string `json:"amr"`
	Exp               int      `json:"exp"`
	Iat               int      `json:"iat"`
	AuthTime          int      `json:"auth_time"`
	Email             string   `json:"email"`
	EmailVerified     bool     `json:"email_verified"`
	FamilyName        string   `json:"family_name"`
	Gender            string   `json:"gender"`
	GivenName         string   `json:"given_name"`
	Locale            string   `json:"locale"`
	Name              string   `json:"name"`
	Nickname          string   `json:"nickname"`
	PreferredUsername string   `json:"preferred_username"`
	Sub               string   `json:"sub"`
	UpdatedAt         int      `json:"updated_at"`
	//UrnZitadelIamOrgDomainPrimary interface{} `json:"urn:zitadel:iam:org:domain:primary"`
	//UrnZitadelIamOrgProjectRoles  interface{} `json:"urn:zitadel:iam:org:project:roles"`
}
