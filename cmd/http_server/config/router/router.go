package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang_testovoe/cmd/http_server/app/controller"
	"golang_testovoe/cmd/http_server/app/controller/goodsController"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	appController := initializationController()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/good", func(r chi.Router) {
		r.Post("/create", appController.GoodsController.Create)
		r.Patch("/update", appController.GoodsController.Update)
		r.Delete("/remove", appController.GoodsController.Delete)
		r.Get("/list", appController.GoodsController.List)
		r.Patch("/rePrioritize", appController.GoodsController.RePrioritize)
	})
	return r
}

func initializationController() controller.Controller {
	return controller.Controller{
		GoodsController: goodsController.NewGoodsController(),
	}
}
