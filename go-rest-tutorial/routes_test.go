package main

import (
	"net/http"
	"testing"
)

func Test_homePage(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "return"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			homePage(tt.args.w, tt.args.r)
		})
	}
}
