package controller

import (
	"encoding/json"
	"github.com/tealeg/xlsx"
	"net/http"
	"test/database"
	"test/model"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var input model.FileUploadInput
	res := make(map[string]string)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		res["Message"] = "Input error"
		js, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
	}
	file, err := xlsx.OpenFile(input.Path)
	if err != nil {
		res["Message"] = "File not found"
		js, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
	}
	for _, sheet := range file.Sheets {
		for i, row := range sheet.Rows {
			if i == 0 {
				continue
			}
			var prod []string
			for _, cell := range row.Cells {
				if cell.Value == "" {
					break
				}
				prod = append(prod, cell.Value)
			}
			database.CreateProduct(prod)
		}
	}
	res["Message"] = "Success"
	js, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
