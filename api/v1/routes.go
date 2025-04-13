package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ooqls/go-log/api/v1/gen"
	"github.com/ooqls/go-log/api/v1/impl"
)

func AddRoutes(e *gin.Engine) {	
	gen.RegisterHandlers(e, &impl.LoggingServerImpl{})
}
