package spiralsrv

import (
	"martinlopez/spiral/internal/core/domain"
	"reflect"
	"testing"
)

func Test_service_Get(t *testing.T) {
	type args struct {
		rows int
		cols int
	}
	tests := []struct {
		name string
		srv  *service
		args args
		want domain.SpiralTable
	}{
		{
			"given rows=1 and cols=1, should return [0]",
			New(),
			args{
				rows: 1,
				cols: 1,
			},
			domain.SpiralTable{
				0: []int64{0},
			},
		},
		{
			"given rows=1 and cols=2, should return the spiral table",
			New(),
			args{
				rows: 1,
				cols: 2,
			},
			domain.SpiralTable{
				[]int64{0, 1},
			},
		},
		{
			"given rows=2 and cols=1, should return the spiral table",
			New(),
			args{
				rows: 2,
				cols: 1,
			},
			domain.SpiralTable{
				[]int64{0},
				[]int64{1},
			},
		},
		{
			"given rows=2 and cols=2, should return the spiral table",
			New(),
			args{
				rows: 2,
				cols: 2,
			},
			domain.SpiralTable{
				[]int64{0, 1},
				[]int64{2, 1},
			},
		},
		{
			"given rows=3 and cols=5 (cols%2=1), should return the spiral table",
			New(),
			args{
				cols: 5,
				rows: 3,
			},
			domain.SpiralTable{
				[]int64{0, 1, 1, 2, 3},
				[]int64{89, 144, 233, 377, 5},
				[]int64{55, 34, 21, 13, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &service{}
			if got := srv.Get(tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Get(%v) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
