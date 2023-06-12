package restapi

func (r *controller) RegisterRouter() {
	resource := r.Router.Group("/api/v1", r.authentication())
	resource.GET("/todo", r.authorization(), r.getAllTodoHandler())
	resource.PUT("/todo/todo_id", r.authorization(), r.runTodoCheckHandler())
	resource.POST("/todo", r.authorization(), r.runTodoCreateHandler())
	//!

}
