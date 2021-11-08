package port

import "martinlopez/spiral/internal/core/domain"

type SpiralService interface {
	Get(rows, cols int) domain.SpiralTable
}
