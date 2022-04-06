package apierror

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/status"
)

func AbortWithError(ctx *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.AbortWithError(runtime.HTTPStatusFromCode(st.Code()), errors.New(st.Message()))
}
