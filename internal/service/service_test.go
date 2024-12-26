package service

import (
	"testing"
)

func TestHelloService_Hello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"success",
			"Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := HelloService{}
			if got := s.Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
