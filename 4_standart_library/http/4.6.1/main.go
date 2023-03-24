package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// StatusErr описывает ситуацию, когда на запрос
// пришел ответ с HTTP-статусом, отличным от 2xx.
type StatusErr struct {
	Code   int
	Status string
}

func (e StatusErr) Error() string {
	return "invalid response status: " + e.Status
}

// начало решения

// httpGet выполняет GET-запрос с указанными заголовками и параметрами,
// парсит ответ как JSON и возвращает получившуюся карту.
//
// Считает ошибкой любые ответы с HTTP-статусом, отличным от 2xx.
func httpGet(uri string, headers map[string]string, params map[string]string, timeout int) (map[string]any, error) {
	client := http.Client{Timeout: time.Duration(timeout) * time.Millisecond} //создаем http клиента с таймаутом

	req, err := http.NewRequest(http.MethodGet, uri, nil) // создаем  объект get запроса
	if err != nil {
		return nil, err
	}
	pr := url.Values{}               // создаем объект  для передачи параметров запроса
	for key, value := range params { // добавляем параметры в объект запроса
		pr.Add(key, value)
	}

	for key, value := range headers { // добавляем заголовки в объект запроса
		req.Header.Add(key, value)
	}

	req.URL.RawQuery = pr.Encode() // преобразуем переданные параметры в строку запроса

	resp, err := client.Do(req) //Отправляем запрос с помощью клиента client и объекта запроса req. получаем ответ и записываем в  resp типа http.Response
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close() // закрываем тело запроса

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, StatusErr{resp.StatusCode, resp.Status}
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result) // декодируем тело ответа из json в мапу
	if err != nil {
		return nil, err
	}

	return result, nil
}

// конец решения

func main() {
	{
		// GET-запрос
		const uri = "https://httpbingo.org/json"
		data, err := httpGet(uri, nil, nil, 3000)
		fmt.Printf("GET %v\n", uri)
		fmt.Println(data, err)
		fmt.Println()
		// GET https://httpbingo.org/json
		// map[slideshow:map[author:Yours Truly date:date of publication slides:[map[title:Wake up to WonderWidgets! type:all] map[items:[Why <em>WonderWidgets</em> are great Who <em>buys</em> WonderWidgets] title:Overview type:all]] title:Sample Slide Show]] <nil>
	}

	{
		// 404 Not Found
		const uri = "https://httpbingo.org/whatever"
		data, err := httpGet(uri, nil, nil, 3000)
		fmt.Printf("GET %v\n", uri)
		fmt.Println(data, err)
		fmt.Println()
		// GET https://httpbingo.org/whatever
		// map[] invalid response status: 404 Not Found
	}

	{
		// С заголовками
		const uri = "https://httpbingo.org/headers"
		headers := map[string]string{
			"accept": "application/xml",
			"host":   "httpbingo.org",
		}
		data, err := httpGet(uri, headers, nil, 3000)
		fmt.Printf("GET %v\n", uri)
		respHeaders := data["headers"].(map[string]any)
		fmt.Println(respHeaders["Accept"], respHeaders["Host"], err)
		fmt.Println()
		// GET https://httpbingo.org/headers
		// [application/xml] [httpbingo.org] <nil>
	}

	{
		// С URL-параметрами
		const uri = "https://httpbingo.org/get"
		params := map[string]string{"id": "42"}
		data, err := httpGet(uri, nil, params, 3000)
		fmt.Printf("GET %v\n", uri)
		fmt.Println(data["args"], err)
		fmt.Println()
		// GET https://httpbingo.org/get
		// map[id:[42]] <nil>
	}
}
