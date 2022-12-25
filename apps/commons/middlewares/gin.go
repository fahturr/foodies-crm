package middlewares

type GinMiddleware struct{}

func NewGinMiddleware() *GinMiddleware {
	return &GinMiddleware{}
}
