package service

import "context"

func (s securityService) CanManage(ctx context.Context) bool {
	panic("implement me")
}

func (s securityService) CanWrite(ctx context.Context) bool {
	panic("implement me")
}

func (s securityService) CanRead(ctx context.Context) bool {
	panic("implement me")
}
