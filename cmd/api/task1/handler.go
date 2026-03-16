package task1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type facility struct {
	id       string
	name     string
	location string
}

type booking struct {
	id          string
	facility_id string
	startTime   time.Time
	endTime     time.Time
	bookedBy    string
}

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func GetFacilities(w http.ResponseWriter, r *http.Request) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Entering GetFacilities")

	mockReturns := []facility{
		{
			id:       "1",
			name:     "Dummy1",
			location: "Norwich",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(gin.H{"facilities": mockReturns}); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		logger.Error(err.Error())
		return
	}

	logger.Info("Exiting GetFacilities")
}
