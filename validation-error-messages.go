package main

var validationErrorMessages = map[validationError]string{
	{
		"Username",
		"required",
	}: "User name is required",
	{
		"Password",
		"required",
	}: "Password is required",
	{
		"Password",
		"alphanum",
	}: "Password must contain letters and numbers only",
}
