package internal

import (
	"b2b-api/internal/config"
	"b2b-api/internal/handlers/authentication"
	merchHandler "b2b-api/internal/handlers/merchant"
	"b2b-api/internal/middleware"
	"b2b-api/internal/pkg/db/oracle"
	"b2b-api/internal/pkg/db/redis"
	"b2b-api/internal/pkg/logger"
	authenticationRepo "b2b-api/internal/pkg/repository/authentication"
	"b2b-api/internal/pkg/repository/merchant"
	"b2b-api/internal/pkg/service"
	"b2b-api/internal/pkg/totp"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	service.NewService,
	config.NewConfig,
	middleware.NewMiddleware,
	authentication.NewAuthHandler,
	totp.NewTOTP,
	oracle.NewOracle, //Added Usmon
	redis.NewRedis,
	logger.NewLogger,
	authenticationRepo.NewAuthRepository,
	merchant.NewRepositoryMerchant,
	merchHandler.NewMerchHandler,
)
