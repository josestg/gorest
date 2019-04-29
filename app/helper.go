package app

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// HELPER
func (A *App) RespondJSON (w http.ResponseWriter, status int, payload interface{}){
	res, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(res))
}

func (A *App) RespondError (w http.ResponseWriter, code int, msg string){
	A.RespondJSON(w,code,map[string]string{"error": msg})
}

func parserID(n string) (uint, error){
	id,err:= strconv.ParseInt(n,10,32)
	return uint(id), err
}

func parserEnable(n string) (bool, error){
	id,err:= strconv.ParseBool(n)
	return bool(id), err
}

