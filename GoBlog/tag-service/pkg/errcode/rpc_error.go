package errcode

import (
	pb "github.com/darrenli6/GolangStudy/GoBlog/tag-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TogRPCError(err *Error) error {
	s, _ := status.New(ToRPCError(err.Code()), err.Msg()).WithDetails(&pb.Error{Code: int32(err.Code()), Message: err.Msg()})
	return s.Err()
}

func ToRPCStatus(err *Error) *Status {
	s, _ := status.New(ToRPCError(err.Code()), err.Msg()).WithDetails(&pb.Error{Code: int32(err.Code()), Message: err.Msg()})
	return &Status{s}
}

func ToRPCError(code int) codes.Code {

	var statusCode codes.Code

	switch code {
	case Fail.Code():
		statusCode = codes.Internal
	default:
		statusCode = codes.Unknown
	}

	return statusCode

}

type Status struct {
	*status.Status
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}
