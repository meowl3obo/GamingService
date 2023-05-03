package model

type Response[T any] struct {
	StatusCode int
	Response   T
}

type HttpResponse = Response[[]byte]

type ErrorResponse struct {
	Code    int
	Message string
}

type VersionResponse[T any] struct {
	Version string
	Data    T
}
