package main

import (
	"reflect"
	"testing"
	"time"
)

type User struct {
	Name         string
	EmailAddress string
	Birthday     time.Time
}

func TestIsUserActive(t *testing.T) {
	type args struct {
		u User
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUserActive(tt.args.u); got != tt.want {
				t.Errorf("IsUserActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_CalculateAge(t *testing.T) {
	type fields struct {
		Name         string
		EmailAddress string
		Birthday     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Name:         tt.fields.Name,
				EmailAddress: tt.fields.EmailAddress,
				Birthday:     tt.fields.Birthday,
			}
			if got := u.CalculateAge(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.CalculateAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_IsNicknameSet(t *testing.T) {
	initDB := func() User {
		return User{}
	}

	type fields struct {
		Nickname *string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"returns true if a nickname is set",
			fields{
				Nickname: &"Hello", // This isn't valid and won't compile!
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Nickname: tt.fields.Nickname,
			}
			if got := u.IsNicknameSet(); got != tt.want {
				t.Errorf("User.IsNicknameSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
