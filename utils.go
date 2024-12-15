package httpresponse

import (
	"encoding/json"
	"net/http"

	"github.com/IvanSkripnikov/logger"
)

// InitRoutes Инициализация маршрутов для Rest API.
func InitRoutes(routes map[string]func(http.ResponseWriter, *http.Request)) {
	// Инициализация эндпоинов для Rest API
	for route, handler := range routes {
		http.HandleFunc(route, handler)
	}

	// Вешаем обработчик для главной страницы и логирования ошибочных запросов
	http.HandleFunc("/", homeHandler)
}

// SendResponse Отправить ответ клиенту.
func SendResponse(w http.ResponseWriter, data ResponseData, caption string) {
	response, errEncode := json.Marshal(data)
	if errEncode != nil {
		logger.Errorf("Failed to serialize data to get %s. Error: %v", caption, errEncode)
		http.Error(w, errEncode.Error(), http.StatusInternalServerError)

		return
	} else {
		logger.Debugf("Data for receiving %s has been successfully serialized.", caption)
	}

	w.Header().Set("Content-Type", "application/json")
	_, errWrite := w.Write(response)
	if errWrite != nil {
		logger.Errorf("Failed to send %s data. Error: %v", caption, errWrite)
		http.Error(w, errWrite.Error(), http.StatusInternalServerError)

		return
	} else {
		logger.Debugf("Data with %s sent successfully.", caption)
	}
}
