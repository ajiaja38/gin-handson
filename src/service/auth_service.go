package service

import "res-gin/src/dto"

type AuthService interface {
	LoginUser(loginDto *dto.LoginDTO) (*dto.LoginResponseDTO, error)
	RefreshToken(refreshTokenDto *dto.RefreshTokenDTO) (*dto.RefreshTokenResponseDTO, error)
}
