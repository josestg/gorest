package helper

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func ParseID(n string) int {
	id,_ := strconv.ParseInt(n,10,32)
	return int(id)
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}){
	res, err := json.Marshal(payload)
	if err !=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(res))
	return
}

func RespondError(w http.ResponseWriter, status int, msg string)  {
	RespondJSON(w,status,map[string]string{"error": msg})
}
