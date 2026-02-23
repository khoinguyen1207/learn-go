package auth

type JWTService interface {
	GenerateAccessToken(userID string, role int32) (string, error)
	GenerateRefreshToken(userID string, role int32) (string, error)
}
