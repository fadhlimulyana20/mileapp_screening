package router

import (
	"net/http"

	"github.com/mileapp_screening/internal/handler"

	"github.com/go-chi/chi/v5"
)

func (rtr *router) helloRouter() http.Handler {
	helloHandler := handler.NewHelloHandler()
	uploadHandler := handler.NewUploadHandler()
	hello := chi.NewRouter()
	hello.Get("/", helloHandler.Hello)
	hello.Post("/upload", uploadHandler.Upload)
	return hello
}

func (rtr *router) nosqlRouter() http.Handler {
	nosqlHandler := handler.NewNoSQLHandler(rtr.cfg.Mongo)
	nosql := chi.NewRouter()
	nosql.Post("/", nosqlHandler.Insert)
	return nosql
}

func (rtr *router) transactionRouter() http.Handler {
	trxHandler := handler.NewTransactionHandler(rtr.cfg.Mongo)
	trxrouter := chi.NewRouter()

	trxrouter.Post("/", trxHandler.Insert)
	trxrouter.Get("/", trxHandler.Get)
	trxrouter.Get("/{id}", trxHandler.GetOne)
	trxrouter.Put("/{id}", trxHandler.Update)
	trxrouter.Patch("/{id}", trxHandler.Patch)
	trxrouter.Delete("/{id}", trxHandler.Delete)

	return trxrouter
}
