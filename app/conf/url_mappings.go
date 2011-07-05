package mapping

type url_mappings struct{
	defaultMapping map[string]string
	wiredMapping map[string]string
}

var url = new(url_mappings)

func init() {

	url.defaultMapping = map[string]string{
	   	"url": "/",
	    "controller": "hello",
		"view": "index"}
	url.wiredMapping = map[string]string{
		"url": "/hello",
		"controller": "hello",
		"view": "index"}
	// "/", controller:"hello"

	// "/$controller/$action/$id?"

	// "/blog/$year/$month/$day"
    // Map{url:"/",controller:"hello"}
    // Map{url:"/$controller/$action/$id?"}
}