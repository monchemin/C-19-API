package service

import "context"

func (ss securityService) CanManage(ctx context.Context) bool {
	return false
}

func (ss securityService) CanWrite(ctx context.Context) bool {
	return false
}

func (ss securityService) CanRead(ctx context.Context) bool {
	return false
}
