package main

var validationErrorMessages = map[validationError]string{
	{
		"Username",
		"required",
	}: "User name is required",
	{
		"Username",
		"alpha",
	}: "Username must contain letters only",
	{
		"Password",
		"required",
	}: "Password is required",
	{
		"Password",
		"alphanum",
	}: "Password must contain letters and numbers only",
	{
		"Password",
		"min",
	}: "Password must contain at least 5 characters",
	{
		"NewUsername",
		"required",
	}: "New user name is required",
	{
		"NewUsername",
		"alpha",
	}: "New user name must contain letters only",
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
		"NewPassword",
		"min",
	}: "NewPassword must contain at least 5 characters",
	{
		"ConfirmNewPassword",
		"required",
	}: "Confirm new password is required",
	{
		"ConfirmNewPassword",
		"alphanum",
	}: "Confirm new password must contain letters and numbers only",
}
