package v1

import (
	"github.com/gin-gonic/gin"
	leshy_api "github.com/njxxdev/leshy/pkg/api"
	leshy_component "github.com/njxxdev/leshy/pkg/component"
)

func CreateApiV1() {
	leshy_component.GetComponentManager().Append(
		leshy_api.NewAPIServer("api_v1", gin.Default()).AddHandlers(handlers...),
	)
}
