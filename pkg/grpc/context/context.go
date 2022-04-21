package context

import (
	ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return ctxtags.UnaryServerInterceptor(
		ctxtags.WithFieldExtractor(ctxtags.CodeGenRequestFieldExtractor),
	)
}
