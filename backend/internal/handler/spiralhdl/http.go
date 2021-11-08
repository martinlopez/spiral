package spiralhdl

import (
	"martinlopez/spiral/internal/core/port"
	"martinlopez/spiral/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SpiralParams struct {
	Rows int `form:"rows"`
	Cols int `form:"cols"`
}

type SpiralResponse struct {
	Rows [][]int64 `json:"rows"`
	Ts   int64     `json:"ts"`
}

type HTTPHandler struct {
	spiralService port.SpiralService
}

func NewHTTPHandler(spiralService port.SpiralService) *HTTPHandler {
	return &HTTPHandler{
		spiralService: spiralService,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	var spiralParams SpiralParams
	if c.ShouldBindQuery(&spiralParams) == nil {
		cols := spiralParams.Cols
		rows := spiralParams.Rows
		if cols == 0 || rows == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request - cols and rows query params must be greater than 0",
			})
			return
		}
		spiralTable := hdl.spiralService.Get(rows, cols)
		response := SpiralResponse{
			Rows: spiralTable,
			Ts:   pkg.MakeTimestamp(),
		}

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

}
