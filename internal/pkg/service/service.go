package service

import (
	"b2b-api/internal/config"
	"b2b-api/internal/pkg/db/redis"
	"b2b-api/internal/pkg/logger"
	"b2b-api/internal/pkg/repository/authentication"
	"b2b-api/internal/pkg/repository/merchant"
	"b2b-api/internal/pkg/totp"
	"go.uber.org/fx"
)

var NewService = fx.Provide(newService)

type IService interface {
	TwoFaInstance() totp.Iotp
	ConfigInstance() config.IConfig
	RedisInstance() redis.IRedis
	LoggerInstance() logger.ILogger
	AuthRepositoryInstance() authentication.IAuthRepository
	MerchRepositoryInstance() merchant.IMerchantRepository
}

type dependencies struct {
	fx.In
	TwoFA           totp.Iotp
	Config          config.IConfig
	Redis           redis.IRedis
	AuthRepository  authentication.IAuthRepository
	MerchRepository merchant.IMerchantRepository
	Logger          logger.ILogger
}

type service struct {
	TwoFA           totp.Iotp
	Config          config.IConfig
	Redis           redis.IRedis
	AuthRepository  authentication.IAuthRepository
	MerchRepository merchant.IMerchantRepository
	Logger          logger.ILogger
}

func newService(d dependencies) IService {
	return &service{
		d.TwoFA,
		d.Config,
		d.Redis,
		d.AuthRepository,
		d.MerchRepository,
		d.Logger,
	}
}

func (s service) TwoFaInstance() totp.Iotp {
	return s.TwoFA
}

func (s service) ConfigInstance() config.IConfig {
	return s.Config
}

func (s service) RedisInstance() redis.IRedis {
	return s.Redis
}

func (s service) AuthRepositoryInstance() authentication.IAuthRepository {
	return s.AuthRepository
}

func (s service) MerchRepositoryInstance() merchant.IMerchantRepository {
	return s.MerchRepository
}

func (s service) LoggerInstance() logger.ILogger {
	return s.Logger
}
