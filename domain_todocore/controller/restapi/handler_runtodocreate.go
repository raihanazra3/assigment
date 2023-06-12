package restapi

import (
	"assigment/domain_todocore/model/entity"
	"assigment/domain_todocore/usecase/runtodocreate"
	"assigment/shared/gogen"
	"assigment/shared/infrastructure/logger"
	"assigment/shared/model/payload"
	"assigment/shared/util"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *controller) runTodoCreateHandler() gin.HandlerFunc {

	type InportRequest = runtodocreate.InportRequest
	type InportResponse = runtodocreate.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
		Todo *entity.Todo `json:"todo"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		err := c.BindJSON(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.Message = jsonReq.Message
		req.Now = time.Now()
		req.RandomString = util.GenerateID(5)

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Todo = res.Todo

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
