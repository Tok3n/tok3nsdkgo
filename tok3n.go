package tok3nsdkgo

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"errors"
	"net/url"
	"encoding/json"
)

type Tok3n_OTP_Response struct {
	Error string
	Result string
}

type Tok3nConfig struct{
	Domain string

	SecretKey string
	PublicKey string
}

func GetTok3nConfigWithSecretPublic(s,p string)Tok3nConfig{
	var conf Tok3nConfig
	conf.Domain = "secure.tok3n.com"
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

func (t Tok3nInstance) _addDomain(path string)string{
	return fmt.Sprintf("http://%s%s",t.Config.Domain,path)
}

func (t Tok3nInstance) _getRemote(path string) (string,error ){
	url := t._addDomain(path)
	res, err := t.Client.Get(url)
	if err != nil{
		return "", err
	}
	response, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	resp := string(response)
	if strings.HasPrefix(resp, "ERROR"){
		return "",errors.New(resp)
	}

	return resp,nil
}

func (t Tok3nInstance) GetActiveSession(kind string) (string, error){
	url := fmt.Sprintf("/api/v1/getSession?kind=%s&secretKey=%s",kind,t.Config.SecretKey)
	return t._getRemote(url)
	//$url = $this->tok3nURL."/api/v1/getSession?kind=$kind&secretKey=".$this->tok3nSecretKey;
}

func (t Tok3nInstance) GetAccessUrl(callback, callbackdata string) (string,error){
	session, err := t.GetActiveSession("access")
	if err!= nil{
		return "", err
	}

	u, _ := url.Parse(t._addDomain("/login.do?"))
	q := u.Query()
	q.Set("publicKey", t.Config.PublicKey)
	q.Set("session", session)
	q.Set("callbackurl", callback)
	q.Set("callbackdata", callbackdata)
	u.RawQuery = q.Encode()
	return u.String(),nil

}

func (t Tok3nInstance) GetJsClientUrl(action,userkey string) string {
	u, _ := url.Parse(t._addDomain("/api/v1/client.js?"))
	q := u.Query()
	q.Set("publicKey", t.Config.PublicKey)
	q.Set("actionName", action)
	q.Set("userkey", userkey)
	u.RawQuery = q.Encode()
	return u.String();
}

func (t Tok3nInstance) ValidateOTP(userkey, otp, sesion string) (string,error) {
	u, _ := url.Parse("/api/v1/otpValid?")
	q := u.Query()
	q.Set("SecretKey", t.Config.SecretKey)
	q.Set("otp", otp)
	q.Set("UserKey", userkey)
	q.Set("secion",sesion)
	u.RawQuery = q.Encode()
	u.String()
	response,err := t._getRemote(u.String())
	if err!= nil{
		return "", err
	}

	var responseStruct Tok3n_OTP_Response
	err = json.Unmarshal([]byte(response),&responseStruct)
	if err!= nil{
		return "", err
	}
	if responseStruct.Error != ""{
		return "",errors.New("Error: with the channel")
	}else{
		return responseStruct.Result, nil
	}


}