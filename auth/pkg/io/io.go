package io

type UserDetails struct {
	Username string
	Roles []string
	Others map[string]string
}

type Session struct {
	User UserDetails
}
