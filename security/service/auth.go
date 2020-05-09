package service

import (

	"context"
	"log"
	"os"

	appContext "github.com/monchemin/C-19-API/context"
	"github.com/monchemin/C-19-API/security/jwt"
	"github.com/monchemin/C-19-API/security/repository"
)

const (
	MANAGE = "M"
	WRITE  = "W"
	READ   = "R"
)

func (ss securityService) CanManage(ctx context.Context) bool {
	claims := contextToClaims(ctx)
	if claims == nil  {
		return false
	}
	if claims.UserName == os.Getenv("SU") {
		return true
	}
	if !ss.isActive(claims.UserID) {
		return false
	}
	privileges := ss.privileges(ctx)
	for _, role := range privileges {
		if role.Code == MANAGE {
			return true
		}
	}
	return false
}

func (ss securityService) CanWrite(ctx context.Context) bool {
	claims := contextToClaims(ctx)
	if !ss.isActive(claims.UserID) {
		return false
	}
	privileges := ss.privileges(ctx)
	for _, role := range privileges {
		if role.Code == MANAGE ||  role.Code == WRITE {
			return true
		}
	}
	return false
}

func (ss securityService) CanRead(ctx context.Context) bool {
	claims := contextToClaims(ctx)
	if claims.UserID == os.Getenv("SU") {
		return true
	}
	if !ss.isActive(claims.UserID) {
		return false
	}
	privileges := ss.privileges(ctx)
	for _, role := range privileges {
		if role.Code == MANAGE ||  role.Code == WRITE ||  role.Code == READ {
			return true
		}
	}
	return false
}

func contextToClaims(ctx context.Context) *jwt.AppClaims {
	claims, err := parseJwt(ctx)
	if err != nil {
		log.Println(err)
	}
	return claims
}

func (ss securityService) isActive(userID string) bool {
	user := ss.repository.UserByID(userID)
	return user.IsActive
}

func (ss securityService) privileges(ctx context.Context) []repository.PrivilegeResult {
	claims := contextToClaims(ctx)
	ctxValues := appContext.ContextKeys(ctx)
	return ss.repository.UserPrivileges(claims.UserID, ctxValues.ResourceID)
}
