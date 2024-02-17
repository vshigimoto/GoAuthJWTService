package services

type Services struct {
	TokenServiceI
}

type TokenServiceI struct {
	GenerateAccessToken()
	GenerateRefreshToken()
}

func New() *Services {
	return &Services{

	}
}