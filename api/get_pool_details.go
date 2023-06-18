package pool

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func (a API) GetPoolDetails(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		poolID := chi.URLParam(req, "pool_id")
		if poolID == "" {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		blockParam := req.URL.Query().Get("block")

		var blockNumber string

		if strings.ToLower(blockParam) == "latest" {
		} else if _, err := strconv.Atoi(blockParam); err == nil {
			blockNumber = blockParam
		} else {
			logrus.Errorf("failed to process blockNumber err: %+v", err)

			w.WriteHeader(http.StatusBadRequest)

			return
		}

		poolDetails, err := a.uniTrackService.GetDatapoint(req.Context(), poolID, blockNumber)
		if err != nil {
			logrus.Errorf("failed to get datapoint err: %+v", err)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		if poolDetails == nil {
			w.WriteHeader(http.StatusNoContent)

			return
		}

		response := map[string]interface{}{
			"token0Balance": poolDetails.Token0Balance,
			"token1Balance": poolDetails.Token1Balance,
			"blockNumber":   poolDetails.BlockNumber,
			"tick":          poolDetails.Tick,
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		if _, err = w.Write(responseJSON); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
