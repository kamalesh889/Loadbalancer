package Router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Rout() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/urls/register", Register_url).Methods("POST")
	r.HandleFunc("/proxy", Proxy).Methods("GET")
	return r
}

//Register Url
func Register_url(res http.ResponseWriter, req *http.Request) {
	var seturl UrlRegister
	err := json.NewDecoder(req.Body).Decode(&seturl)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	seturl.Alive = true //Keeping Health Status true by default While registering
	Urllist = append(Urllist, seturl)
	fmt.Println(Urllist)
}

//Redirect request
func Proxy(res http.ResponseWriter, req *http.Request) {
	Url, err := Loadbalancer_url() //Get alive Url
	if err != nil {
		fmt.Println(err.Error())
		res.WriteHeader(http.StatusServiceUnavailable)
		res.Write([]byte(err.Error()))
	}
	fmt.Println("URl", Url)
	resp, err := http.Get(Url)
	if err != nil {
		fmt.Println(err)
		return
	}
	respdata, _ := ioutil.ReadAll(resp.Body)
	respstring := string(respdata)
	res.Write([]byte(respstring))
}
