package test


//package httpresp

//import (
//	"encoding/json"
//	"net/http"
//)
//
//// Response v1
//type ResponseV1 struct {
//	ErrCode int         `json:"errCode"`
//	ErrMsg  string      `json:"errMsg"`
//	Payload interface{} `json:"payload"`
//}
//
//// Success ...
//func Success(w http.ResponseWriter, payload interface{}) {
//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
//	w.WriteHeader(http.StatusOK)
//
//	resp := ResponseV1{
//		Payload: payload,
//	}
//	err := json.NewEncoder(w).Encode(resp)
//	if err != nil {
//		panic(err)
//	}
//}
//
//// Fail ...
//func Fail(w http.ResponseWriter, code int, msg string) {
//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
//	w.WriteHeader(http.StatusOK)
//	resp := ResponseV1{
//		ErrCode: code,
//		ErrMsg:  msg,
//	}
//	if err := json.NewEncoder(w).Encode(resp); err != nil {
//		panic(err)
//	}
//}
