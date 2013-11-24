package tok3n

import (
	"net/http"
)

type Tok3nConfig struct{
	Domain string

	SecretKey string
	PublicKey string
}

func GetTok3nConfigWithSecretPublic(s,p string)Tok3nConfig{
	var conf Tok3nConfig
	conf.Domain = "secret.tok3n.com"
	conf.SecretKey = s
	conf.PublicKey = p
	return conf
}
func GetTok3nConfigWithDomainSecretPublic(d,s,p string)Tok3nConfig{
	conf := GetTok3nConfigWithSecretPublic(sp)
	conf.Domain = d
	return conf
}

type Tok3nInstance struct {
	Client *http.Client
	Config Tok3nConfig
}

