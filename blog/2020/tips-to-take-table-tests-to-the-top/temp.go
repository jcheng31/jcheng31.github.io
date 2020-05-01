package main

import "time"

type User struct {
	Name         string
	EmailAddress string
	Birthday     time.Time // TODO: display on profile.

	Nickname *string
}

// CalculateAge calculates the user's age.
func (u User) CalculateAge() time.Duration {
	return time.Now().Sub(u.Birthday)
}

func IsUserActive(u User) bool {
	return true
}

func (u User) IsNicknameSet() bool {
	return u.Nickname != nil
}
