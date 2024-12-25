package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gitlab.com/idoko/bucketeer/db"
	"gitlab.com/idoko/bucketeer/models"
)

var orderIdKey = "orderId"

func orders(router chi.Router) {
	router.Get("/", getAllOrders)
	router.Post("/", createOrder)

	router.Route("/{orderId}", func(router chi.Router) {
		router.Use(OrderContext)
		router.Get("/", getOrder)
		router.Delete("/", deleteOrder)
	})
}

func OrderContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		strOrderId := chi.URLParam(r, "orderId")
		if strOrderId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("orderId is required")))
			return
		}

		orderId, err := strconv.ParseUint(strOrderId, 10, 64)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid orderId")))
		}

		ctx := context.WithValue(r.Context(), orderIdKey, orderId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := dbInstance.GetAllOrders()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, orders); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	order := &models.Order{}

	if err := render.Bind(r, order); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := dbInstance.AddOrder(order); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, order); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

// Добавить получение заказа по userId, такой же метод, но getOrdersByUserId
func getOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.Context().Value(orderIdKey).(uint64)

	order, err := dbInstance.GetOrderById(orderId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &order); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.Context().Value(orderIdKey).(uint64)

	err := dbInstance.DeleteOrder(orderId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
