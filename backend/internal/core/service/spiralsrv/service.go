package spiralsrv

import (
	"martinlopez/spiral/internal/core/domain"
	"martinlopez/spiral/internal/pkg"
)

type service struct {
	//dependencies here
}

func New() *service {
	return &service{}
}

func (srv *service) Get(rows, cols int) domain.SpiralTable {
	var spiralTable domain.SpiralTable
	size := rows * cols
	fibSequence := pkg.FibonacciSequence(size)

	if cols == 1 {
		if rows == 1 {
			spiralTable = append(spiralTable, []int64{0})
			return spiralTable
		}
		for _, v := range fibSequence {
			spiralTable = append(spiralTable, []int64{v})
		}
		return spiralTable
	}

	spiralMatrix := pkg.SpiralMatrix(rows, cols, fibSequence)
	for r := 0; r < rows; r++ {
		spiralTable = append(spiralTable, spiralMatrix[r*cols:r*cols+cols])
	}

	return spiralTable
}
