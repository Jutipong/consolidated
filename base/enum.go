package base

//## Http
const (
	Successfully float32 = 0000
	DataNotFound float32 = 1111

	Error        float32 = 500
	BadRequest   float32 = 400
	Unauthorized float32 = 401
)

//## Jwt Key
const (
	Authorization string = "Authorization"
	Bearer        string = "Bearer "
	UserRequest   string = "UserRequest"
)
