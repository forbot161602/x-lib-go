package xbgin

func NoRouteHandler(ctx *Context) {
	flow := &RESTFlow{}
	flow.Initiate(ctx)
	flow.SetNotFoundError()
}
