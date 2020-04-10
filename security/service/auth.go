package service

import "context"

func (ss securityService) CanManage(ctx context.Context) bool {
	panic("implement me")
}

func (ss securityService) CanWrite(ctx context.Context) bool {
	panic("implement me")
}

func (ss securityService) CanRead(ctx context.Context) bool {
	panic("implement me")
}
