package delivery

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

type OrderService interface {
	Get(int) (models.Order, error)
	GetUsersAll(int) ([]models.Order, error)
	GetAll() ([]models.Order, error)
	Update(models.Order) (models.Order, error)
}

type ContextManager interface {
	UserIDFromContext(ctx context.Context) (int, error)
}

type OrderHandler struct {
	OrderService   OrderService
	ContextManager ContextManager
	Logger         logger.Logger
}

func (oh *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderIdString, ok := vars["ORDER_ID"]
	if !ok {
		oh.Logger.Errorw("no ORDER_ID var")
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	orderId, err := strconv.Atoi(orderIdString)
	if err != nil {
		oh.Logger.Errorw("fail to convert id to int",
			"err:", err.Error())
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	order, err := oh.OrderService.Get(orderId)
	if err != nil {
		oh.Logger.Infow("can`t get order",
			"err:", err.Error())
		http.Error(w, "can`t get order", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(order)

	if err != nil {
		oh.Logger.Errorw("can`t marshal order",
			"err:", err.Error())
		http.Error(w, "can`t make order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		oh.Logger.Errorw("can`t write response",
			"err:", err.Error())
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}

func (oh *OrderHandler) GetAllMy(w http.ResponseWriter, r *http.Request) {
	userID, err := oh.ContextManager.UserIDFromContext(r.Context())
	if err != nil {
		oh.Logger.Errorw("fail to get id from context",
			"err:", err.Error())
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	orders, err := oh.OrderService.GetUsersAll(userID)
	if err != nil {
		oh.Logger.Infow("can`t get order",
			"err:", err.Error())
		http.Error(w, "can`t get order", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(orders)

	if err != nil {
		oh.Logger.Errorw("can`t marshal orders",
			"err:", err.Error())
		http.Error(w, "can`t get orders", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		oh.Logger.Errorw("can`t write response",
			"err:", err.Error())
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}

func (oh *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	orders, err := oh.OrderService.GetAll()
	if err != nil {
		oh.Logger.Infow("can`t get order",
			"err:", err.Error())
		http.Error(w, "can`t get order", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(orders)

	if err != nil {
		oh.Logger.Errorw("can`t marshal orders",
			"err:", err.Error())
		http.Error(w, "can`t get orders", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		oh.Logger.Errorw("can`t write response",
			"err:", err.Error())
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}

func (oh *OrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderIdString, ok := vars["ORDER_ID"]
	if !ok {
		oh.Logger.Errorw("no ORDER_ID var")
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	orderId, err := strconv.Atoi(orderIdString)
	if err != nil {
		oh.Logger.Errorw("fail to convert id to int",
			"err:", err.Error())
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	order := &models.Order{}
	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		oh.Logger.Errorw("can`t read body of request",
			"err:", err.Error())
		http.Error(w, "bad data", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, order)
	if err != nil {
		oh.Logger.Infow("can`t unmarshal form",
			"err:", err.Error())
		http.Error(w, "bad  data", http.StatusBadRequest)
		return
	}

	_, err = govalidator.ValidateStruct(order)
	if err != nil {
		oh.Logger.Infow("can`t validate form",
			"err:", err.Error())
		http.Error(w, "bad data", http.StatusBadRequest)
		return
	}

	order.ID = orderId
	*order, err = oh.OrderService.Update(*order)
	if err != nil {
		oh.Logger.Infow("can`t update order",
			"err:", err.Error())
		http.Error(w, "can`t update order", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(order)

	if err != nil {
		oh.Logger.Errorw("can`t marshal order",
			"err:", err.Error())
		http.Error(w, "can`t make order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		oh.Logger.Errorw("can`t write response",
			"err:", err.Error())
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}
