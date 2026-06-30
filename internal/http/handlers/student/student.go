package student

import (
	"encoding/json"
	"net/http"

	"github.com/anamika8076/go-api/internal/types"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var student types.Student
		json.NewDecoder(r.Body).Decode(&student)

		slog.Info("creating a student ")
		w.Write([]byte("welcome to student api"))
	}
}
 