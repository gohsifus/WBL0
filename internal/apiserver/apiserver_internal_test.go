package apiserver

import (
	"github.com/sirupsen/logrus"
	"natTest/internal/cache"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig(), logrus.New(), cache.New()) //Создаем обект apiServer
	rec := httptest.NewRecorder() //Замена для http.ResponseWriter
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil) //Создаем запрос - замена http.Request
	s.handleHello().ServeHTTP(rec, req) //Вызываем метод обработки url
}
