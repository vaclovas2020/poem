/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package adminserver

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"webimizer.dev/poem/admin"
)

func (srv *adminServer) AddCategory(context.Context, *admin.AdminCategory) (*admin.CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
