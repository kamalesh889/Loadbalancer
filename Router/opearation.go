package Router

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var Counter = 0

//It will Provide the Healthy Url In roundrobin method
func Loadbalancer_url() (string, error) {

	for i := 0; i < len(Urllist); i++ {
		if !Urllist[Counter].Alive {
			Counter = (Counter + 1) % len(Urllist)
		}
		if Urllist[Counter].Alive {
			Action_url := Urllist[Counter].Url
			Counter = (Counter + 1) % len(Urllist)
			return Action_url, nil
		}
	}
	return "", errors.New("no active server is persent")

}

//It will check Url is alive or not
func Healthchkurl() {
	fmt.Println("Enterin")
	for i := range Urllist {
		_, err := http.Head(Urllist[i].Url)
		fmt.Println("server", Urllist[i].Url)
		if err != nil {
			fmt.Println("error is", err)
			Urllist[i].Alive = false
		} else {
			Urllist[i].Alive = true
		}
	}
	fmt.Println("health chk done..", "All Servers Status are", Urllist)
}

//It will run (Healthchkurl) func in a perticular time interval to get the Status of Urls
func Healthchk() {
	for {
		time.Sleep(5 * time.Second)
		go Healthchkurl()
	}
}
