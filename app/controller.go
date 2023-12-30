package app

// BaseController is a struct that all controllers must embed, you can add anything that every controller needs to be able to access here
type BaseController struct {
	App *App
}
