// Package auth ...
package auth

import (
	"context"
	"database/sql"
	"errors"
	"eticket/constant"
	"eticket/model/payload"
	"fmt"
	"time"

	"github.com/lukmanlukmin/go-lib/util"
)

// ValidateUserByCredential ...
func (s *Service) ValidateUserByCredential(ctx context.Context, data payload.LoginCredential) (*payload.Token, error) {
	userData, err := s.Repository.DB.Users.GetUserByUsername(ctx, data.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, constant.INVALID_USER_CREDENTIAL
		}
		return nil, err
	}

	if util.CheckPassword(userData.Password, data.Password) {
		jwtData := map[string]interface{}{
			constant.UserID: fmt.Sprintf("%d", userData.ID),
		}
		duration, err := time.ParseDuration(s.Conf.Security.JWTDuration)
		if err != nil {
			return nil, constant.FAIL_TO_GENERATE_TOKEN
		}
		additionalDuration, err := time.ParseDuration("24h")
		if err != nil {
			return nil, constant.FAIL_TO_GENERATE_TOKEN
		}
		jwt, err := util.GenerateJWT(s.Conf.Security.JWTSecret, duration, jwtData)
		if err != nil {
			return nil, constant.FAIL_TO_GENERATE_TOKEN
		}
		jwtRefresh, err := util.GenerateJWT(s.Conf.Security.JWTSecret, duration+additionalDuration, jwtData)
		if err != nil {
			return nil, constant.FAIL_TO_GENERATE_TOKEN
		}
		return &payload.Token{
			Token:        jwt,
			RefreshToken: jwtRefresh,
		}, nil
	}
	return nil, constant.INVALID_USER_CREDENTIAL
}
