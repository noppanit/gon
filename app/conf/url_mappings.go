package mapping

type url_mappings struct{
	defaultMapping map[string]string
	wiredMapping map[string]string
}

var URL = new(url_mappings)

func init() {

	URL.defaultMapping = map[string]string{
	   	"url": "/",
	    "controller": "hello",
		"view": "index"}
	URL.wiredMapping = map[string]string{
		"url": "/test",
		"controller": "hello",
		"view": "index"}
	// "/", controller:"hello"

	// "/$controller/$action/$id?"

	// "/blog/$year/$month/$day"
    // Map{url:"/",controller:"hello"}
    // Map{url:"/$controller/$action/$id?"}
}