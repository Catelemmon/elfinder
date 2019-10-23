package elfinder

type method string

//type MethodHandler func(con Connector, req *http.Request, rw http.ResponseWriter)
//
//var handlers = map[method]MethodHandler{
//	"open": Open,
//}
//
//func Open(con Connector, req *http.Request, rw http.ResponseWriter) {
//
//	var (
//		init   bool
//		target string
//		tree   bool
//
//		response map[string]interface{}
//
//		err error
//	)
//	init = ParseStringToBool(req.Form.Get("init"))
//	target = req.Form.Get("target")
//	tree = ParseStringToBool(req.Form.Get("tree"))
//
//	response = make(map[string]interface{})
//
//	if init {
//		response["api"] = APIVERSION
//	}
//
//	response["cwd"] = con.getVolume()
//}


