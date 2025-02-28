package coordinator

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/suborbital/reactr/rt"
	"github.com/suborbital/vektor/vk"
)

type MetricsResponse struct {
	Scheduler rt.ScalerMetrics `json:"scheduler"`
}

func (c *Coordinator) metricsHandler() vk.HandlerFunc {
	return func(r *http.Request, ctx *vk.Ctx) (interface{}, error) {
		metrics, err := c.exec.Metrics()
		if err != nil {
			ctx.Log.Error(errors.Wrap(err, "failed to exec.Metrics"))
			return nil, vk.E(http.StatusInternalServerError, "unknown error")
		}

		resp := &MetricsResponse{
			Scheduler: *metrics,
		}

		return resp, nil
	}
}
