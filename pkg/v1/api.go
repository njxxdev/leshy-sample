package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/njxxdev/leshy/pkg/api"
	"github.com/njxxdev/leshy/pkg/component"
)

func CreateApiV1() {
	component.GetComponentManager().Append(
		api.NewAPIServer("api_v1", gin.Default()).AddHandlers(handlers...),
	)
}
