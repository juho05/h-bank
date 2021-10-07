package services

import (
	"strings"
	"testing"

	"github.com/Bananenpro/hbank2-api/config"
	"github.com/stretchr/testify/assert"
)

func Test_IsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"Valid gmail email", "test@gmail.com", true},
		{"Valid gmx email", "test@gmx.de", true},
		{"Valid outlook email", "test@outlook.com", true},
		{"Valid protonmail email", "test@protonmail.com", true},
		{"Empty string", "", false},
		{"Missing @ sign", "test.gmail.com", false},
		{"Missing name", "@gmail.com", false},
		{"Missing domain", "test@com", false},
		{"Non-existant domain", "test@foomail.abc", false},
		{"Two @ signs", "test@foomail@abc", false},
		{"Too long", strings.Repeat("a", config.Data.UserMaxEmailLength) + "@gmail.com", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsValidEmail(tt.email))
		})
	}
}
