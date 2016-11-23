package models

import (
	"net/http"
)

type HSetRequest struct {
	Key		interface{}	`json:"key"`
	Field	interface{}	`json:"field"`
	Value	interface{}	`json:"value"`
}

func HSet(key,field,value string)(int,string,error){
	hset := client.HSet(key,field,value)
	if err := hset.Err(); err != nil {
		return http.StatusInternalServerError,"",err
	}

	if hset.Val() {
		return http.StatusCreated,"Success insert.",nil
	} else {
		return http.StatusOK,"Success updated.",nil
	}
}