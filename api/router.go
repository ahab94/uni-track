package pool

import (
	"context"

	chi "github.com/go-chi/chi/v5"

	app "github.com/ahab94/uni-track"
	"github.com/ahab94/uni-track/models"
	"github.com/ahab94/uni-track/service"
)

type UniTrackService interface {
	GetDatapoint(ctx context.Context, poolID, blockNumber string) (*models.Datapoint, error)
	GetDatapointHistory(ctx context.Context, poolID string) ([]models.Datapoint, error)
}

type API struct {
	rt              *app.Runtime
	uniTrackService UniTrackService
}

func NewAPI(rt *app.Runtime) API {
	return API{
		rt:              rt,
		uniTrackService: service.NewUniTrackService(rt),
	}
}

func (a API) NewRouter(router chi.Router) chi.Router {
	router.Get("/v1/api/pool/{pool_id}", a.GetPoolDetails)
	router.Get("/v1/api/pool/{pool_id}/historic", a.GetPoolHistory)

	return router
}
