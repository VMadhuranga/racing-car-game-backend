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
	{
		"NewUsername",
		"required",
	}: "New user name is required",
	{
		"OldPassword",
		"required",
	}: "Old password is required",
	{
		"OldPassword",
		"alphanum",
	}: "Old password must contain letters and numbers only",
	{
		"NewPassword",
		"required",
	}: "New password is required",
	{
		"NewPassword",
		"alphanum",
	}: "New password must contain letters and numbers only",
	{
		"ConfirmNewPassword",
		"required",
	}: "Confirm new password is required",
	{
		"ConfirmNewPassword",
		"alphanum",
	}: "Confirm new password must contain letters and numbers only",
}
