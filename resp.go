package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SwapToStruct (req, target interface{}) (err error) {
	dataByte, err := json.Marshal(req)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataByte, target)
	return
}

type H struct {
	Code                   string
	Message                string
	TraceId                string
	Data                   interface{}
	Rows                   interface{}
	Total                  interface{}
	SkyWalkingDynamicField string
}

func Resp(w http.ResponseWriter, code string, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:    code,
		Data:    data,
		Message: message,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

func RespList(w http.ResponseWriter, code string, data interface{}, message string, rows interface{}, total interface{}, skyWalkingDynamicField string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:                   code,
		Data:                   data,
		Message:                message,
		Rows:                   rows,
		Total:                  total,
		SkyWalkingDynamicField: skyWalkingDynamicField,
	}
	ret, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

/**
200 OKLoginSuccessVO
201 Created
401 Unauthorized
403 Forbidden
404 Not Found
**/
func RespOK(w http.ResponseWriter, data interface{}, message string) {
	Resp(w, "SUCCESS", data, message)
}
func RespFail(w http.ResponseWriter, data interface{}, message string) {
	Resp(w, "TOKEN_FAIL", data, message)
}

//writer  data  message  row  total  field
func RespListOK(w http.ResponseWriter, data interface{}, message string, rows interface{}, total interface{}, skyWalkingDynamicField string) {
	RespList(w, "SUCCESS", data, message, rows, total, skyWalkingDynamicField)
}
func RespListFail(w http.ResponseWriter, data interface{}, message string, rows interface{}, total interface{}, skyWalkingDynamicField string) {
	RespList(w, "TOKEN_FAIL", data, message, rows, total, skyWalkingDynamicField)
}
