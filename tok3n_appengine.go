package tok3nsdkgo

import (
	"appengine"
	"appengine/urlfetch"
)

func GetAppEngineTok3nInstance(c appengine.Context, conf Tok3nConfig) *Tok3nInstance{
	tok3n := new(Tok3nInstance)
	tok3n.Client = urlfetch.Client(c)
	tok3n.Config = conf
	return tok3n
}