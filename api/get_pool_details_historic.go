package pool

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func (a API) GetPoolHistory(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		poolID := chi.URLParam(req, "pool_id")
		if poolID == "" {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		blocks, err := a.uniTrackService.GetDatapointHistory(req.Context(), poolID)
		if err != nil {
			logrus.Errorf("failed to get datapoint history err: %+v", err)

			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		logrus.Infof("got %d blocks for pool id %s", len(blocks), poolID)

		response := make([]map[string]interface{}, len(blocks))
		previousToken0Balance := 0
		previousToken1Balance := 0

		for i, poolDetails := range blocks {
			// Convert string values to integers
			token0Balance, _ := strconv.Atoi(poolDetails.Token0Balance)
			token1Balance, _ := strconv.Atoi(poolDetails.Token1Balance)

			token0Delta := token0Balance - previousToken0Balance
			token1Delta := token1Balance - previousToken1Balance

			if previousToken0Balance == 0 && previousToken1Balance == 0 {
				token0Delta = 0
				token1Delta = 0
			}

			response[i] = map[string]interface{}{
				"blockNumber":   poolDetails.BlockNumber,
				"token0Delta":   token0Delta,
				"token1Delta":   token1Delta,
				"token0Balance": poolDetails.Token0Balance,
				"token1Balance": poolDetails.Token1Balance,
			}

			previousToken0Balance = token0Balance
			previousToken1Balance = token1Balance
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(responseJSON)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
