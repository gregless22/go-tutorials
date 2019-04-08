package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "test 1",
			args:    args{2},
			want:    4,
			wantErr: false,
		},
		{
			name:    "test 2",
			args:    args{5},
			want:    7,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.x)
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})
	}
}
