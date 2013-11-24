package tok3n

import (
	"appengine"
	"appengine/urlfetch"
)

func getAppEngineTok3n(c appengine.Context){
	client = urlfetch.Client(c)
}