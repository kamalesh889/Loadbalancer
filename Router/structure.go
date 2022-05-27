package Router

var Urllist []UrlRegister // All register Url Will be store here

type UrlRegister struct { //register url With this struct By default Alive is true
	Url   string //Url string
	Alive bool   //Health Status
}
