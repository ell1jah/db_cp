package delivery

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

type ItemService interface {
	Create(models.Item) (int, error)
	Get(int) (models.Item, error)
	GetAll(models.ItemsParams) ([]models.Item, error)
	Update(models.Item) (models.Item, error)
	Delete(int) error
}

type ItemHandler struct {
	ItemService ItemService
	Logger      logger.Logger
}

func (ih *ItemHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemIdString, ok := vars["ITEM_ID"]
	if !ok {
		ih.Logger.Errorw("no ITEM_ID var")
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	itemId, err := strconv.Atoi(itemIdString)
	if err != nil {
		ih.Logger.Errorw("fail to convert id to int",
			"err:", err.Error())
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	item, err := ih.ItemService.Get(itemId)
	if err != nil {
		ih.Logger.Infow("can`t get item",
			"err:", err.Error())
		http.Error(w, "can`t get item", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(item)

	if err != nil {
		ih.Logger.Errorw("can`t marshal item",
			"err:", err.Error())
		http.Error(w, "can`t make item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		ih.Logger.Errorw("can`t write response",
			"err:", err.Error())
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}

func (ih *ItemHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	itemsParams := &models.ItemsParams{}

	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		ih.Logger.Errorw("can`t read body of request",
			"err:", err.Error())
		http.Error(w, "bad data", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, itemsParams)
	if err != nil {
		ih.Logger.Infow("can`t unmarshal form",
			"err:", err.Error())
		http.Error(w, "bad  data", http.StatusBadRequest)
		return
	}

	_, err = govalidator.ValidateStruct(itemsParams)
	if err != nil {
		ih.Logger.Infow("can`t validate form",
			"err:", err.Error())
		http.Error(w, "bad data", http.StatusBadRequest)
		return
	}

	items, err := ih.ItemService.GetAll(*itemsParams)
	if err != nil {
		ih.Logger.Infow("can`t get items",
			"err:", err.Error())
		http.Error(w, "can`t get items", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(items)

	if err != nil {
		ih.Logger.Errorw("can`t marshal items",
			"err:", err.Error())
		http.Error(w, "can`t make items", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(resp)
	if err != nil {
		ih.Logger.Errorw("can`t write response",
			"err:", err.Error())
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}

func (ih *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {
	item := &models.Item{}

	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		ih.Logger.Errorw("can`t read body of request",
			"err:", err.Error())
		http.Error(w, "bad data", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, item)
	if err != nil {
		ih.Logger.Infow("can`t unmarshal form",
			"err:", err.Error())
		http.Error(w, "bad  data", http.StatusBadRequest)
		return
	}

	_, err = govalidator.ValidateStruct(item)
	if err != nil {
		ih.Logger.Infow("can`t validate form",
			"err:", err.Error())
		http.Error(w, "bad data", http.StatusBadRequest)
		return
	}

	item.ID, err = ih.ItemService.Create(*item)
	if err != nil {
		ih.Logger.Infow("can`t create item",
			"err:", err.Error())
		http.Error(w, "can`t create item", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(item)

	if err != nil {
		ih.Logger.Errorw("can`t marshal item",
			"err:", err.Error())
		http.Error(w, "can`t make item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(resp)
	if err != nil {
		ih.Logger.Errorw("can`t write response",
			"err:", err.Error())
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}

func (ih *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemIdString, ok := vars["ITEM_ID"]
	if !ok {
		ih.Logger.Errorw("no ITEM_ID var")
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	itemId, err := strconv.Atoi(itemIdString)
	if err != nil {
		ih.Logger.Errorw("fail to convert id to int",
			"err:", err.Error())
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	item := &models.Item{}
	body, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		ih.Logger.Errorw("can`t read body of request",
			"err:", err.Error())
		http.Error(w, "bad data", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, item)
	if err != nil {
		ih.Logger.Infow("can`t unmarshal form",
			"err:", err.Error())
		http.Error(w, "bad  data", http.StatusBadRequest)
		return
	}

	_, err = govalidator.ValidateStruct(item)
	if err != nil {
		ih.Logger.Infow("can`t validate form",
			"err:", err.Error())
		http.Error(w, "bad data", http.StatusBadRequest)
		return
	}

	item.ID = itemId
	*item, err = ih.ItemService.Update(*item)
	if err != nil {
		ih.Logger.Infow("can`t update item",
			"err:", err.Error())
		http.Error(w, "can`t update item", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(item)

	if err != nil {
		ih.Logger.Errorw("can`t marshal item",
			"err:", err.Error())
		http.Error(w, "can`t make item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		ih.Logger.Errorw("can`t write response",
			"err:", err.Error())
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}

func (ih *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemIdString, ok := vars["ITEM_ID"]
	if !ok {
		ih.Logger.Errorw("no ITEM_ID var")
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	itemId, err := strconv.Atoi(itemIdString)
	if err != nil {
		ih.Logger.Errorw("fail to convert id to int",
			"err:", err.Error())
		http.Error(w, "unknown error", http.StatusInternalServerError)
		return
	}

	err = ih.ItemService.Delete(itemId)
	if err != nil {
		ih.Logger.Infow("can`t delete item",
			"err:", err.Error())
		http.Error(w, "can`t delete item", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
