package main

import "http-examples/examples"

// HTTP Server
// Routing and handlers
// HTTP request segments // url, method, path, query, params, headers, body
// HTTP response segments , statusCode, body, headers
// Context
// Mini Project => Simple CRUD store without database

func main() {
	// 02- Create server
	//examples.CreateServer()

	// 03- Create server and servemux
	examples.CreateServerWithMux()

}
