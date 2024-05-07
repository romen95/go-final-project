package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/romen95/go_final_project/app/model"
	"github.com/romen95/go_final_project/app/service"
)

func NextDateReadGET(w http.ResponseWriter, r *http.Request) {
	now, err := time.Parse(model.DatePattern, r.FormValue("now"))
	if err != nil {
		http.Error(w, fmt.Sprintf(""), http.StatusBadRequest)
		return
	}

	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	nextDate, err := service.NextDate(now, date, repeat)

	if err != nil {
		http.Error(w, fmt.Sprintf(""), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(nextDate))

	if err != nil {
		http.Error(w, fmt.Errorf("writing tasks data error: %w", err).Error(), http.StatusBadRequest)
	}
}
