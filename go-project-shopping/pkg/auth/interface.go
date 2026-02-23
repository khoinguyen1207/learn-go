package auth

type JWTService interface {
	GenerateAccessToken(uuid, email string, role int32) (string, error)
	GenerateRefreshToken(uuid, email string, role int32) (string, error)
}
