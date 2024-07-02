namespace go api

struct Request {
	1: string message(api.body = 'message')
}

struct Response {
	1: string message(api.body = 'message')
}

service Hello {
	Response echo(1: Request req)(
		api.post = '/gateway/hello/',
		api.baseurl = 'localhost:8888',
		api.param = 'true',
		// api.serializer = 'form',
		// api.category = 'demo'
	)
}
