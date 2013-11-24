package tok3nsdkgo

import (
	"net/http"
	"io/ioutil"
	"fmt"
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
	conf := GetTok3nConfigWithSecretPublic(s,p)
	conf.Domain = d
	return conf
}

type Tok3nInstance struct {
	Client *http.Client
	Config Tok3nConfig
}

func (t Tok3nInstance) _getRemote(path string) (string,error ){
	url := fmt.Sprintf("http://%s%s",t.Config.Domain,path)
	res, err := Client.Get(url)
	if err != nil{
		return "", err
	}
	response, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(response),nil
}

func (t Tok3nInstance) getActiveSession(kind string) (string, error){
	url := fmt.Sprintf("/api/v1/getSession?kind=%s&secretKey=%s",kind,t.Config.SecretKey)
	return t._getRemote(url)
	//$url = $this->tok3nURL."/api/v1/getSession?kind=$kind&secretKey=".$this->tok3nSecretKey;
}

func (t Tok3nInstance) getAccessUrl(callback, callbackdata string) (string,error){
	session, err := t.getActiveSession("access")
	if err!= nil{
		return "", err
	}
	return fmt.Sprintf("/login.do?publicKey=%s&session=%s&callbackurl=%s&callbackdata=%s",t.Config.PublicKey,session,callback,callbackdata),nil

}