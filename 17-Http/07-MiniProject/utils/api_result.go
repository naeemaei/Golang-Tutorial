package utils

import (
	"encoding/json"
	"net/http"
)

type ApiResult struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
}

func SetResult(w http.ResponseWriter, statusCode int, result interface{}, errorInfo interface{}) {

	w.Header().Add("Web-Server", "NGINX")
	w.Header().Set("Content-Type", "application/json")

	apiResult := &ApiResult{}
	if statusCode == http.StatusOK {
		apiResult.Success = true
	} else {
		apiResult.Success = false
	}

	apiResult.Result = result

	if errorInfo != nil {
		apiResult.Error = errorInfo
	}

	jsonResponse, err := json.Marshal(apiResult)
	if err != nil {
		apiResult.Success = false
		apiResult.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResponse)
	if err != nil {
		panic(err)
	}

}
