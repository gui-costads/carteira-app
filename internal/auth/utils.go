package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gui-costads/carteira-app/internal/config"
)

type AuthService struct {
	cfg *config.AppConfig
}

func NewAuthService(cfg *config.AppConfig) *AuthService {
	return &AuthService{cfg: cfg}
}

var (
	ErrInvalidToken     = errors.New("token inválido")
	ErrTokenExpired     = errors.New("token expirado")
	ErrInvalidClaims    = errors.New("claims inválidos")
	ErrMissingSecretKey = errors.New("chave secreta não configurada")
	ErrInvalidUserID    = errors.New("ID de usuário inválido")
	ErrInvalidUserName  = errors.New("nome de usuário inválido")
)

const (
	ClaimKeyUserID   = "user_id"
	ClaimKeyUserName = "user_name"
	ClaimKeyExp      = "exp"
)

func (s *AuthService) getSecretKey() ([]byte, error) {
	key := s.cfg.JWTSecretKey
	if len(key) == 0 {
		return nil, errors.New("chave secreta não configurada")
	}
	return key, nil

}

func (s *AuthService) GenerateToken(id uint, nome string) (string, error) {
	secretKey, err := s.getSecretKey()
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		ClaimKeyUserID:   id,
		ClaimKeyUserName: nome,
		ClaimKeyExp:      time.Now().Add(s.cfg.JWTExpiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func (s *AuthService) ExtractToken(tokenString string) (uint, string, error) {
	secretKey, err := s.getSecretKey()
	if err != nil {
		return 0, "", err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, "", ErrTokenExpired
		}
		return 0, "", fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	if !token.Valid {
		return 0, "", ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", ErrInvalidClaims
	}

	idFloat, ok := claims[ClaimKeyUserID].(float64)
	if !ok {
		return 0, "", ErrInvalidUserID
	}
	id := uint(idFloat)

	nome, ok := claims[ClaimKeyUserName].(string)
	if !ok {
		return 0, "", ErrInvalidUserName
	}

	return id, nome, nil
}

func (s *AuthService) ValidateToken(tokenString string) error {
	secretKey, err := s.getSecretKey()
	if err != nil {
		return err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return ErrTokenExpired
		}
		return fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	if !token.Valid {
		return ErrInvalidToken
	}

	return nil
}
