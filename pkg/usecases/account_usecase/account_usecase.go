package account_usecase

import (
	"Skyline/internal/utils"
	"Skyline/pkg/models/session_models"
	"Skyline/pkg/models/user-models"
	"Skyline/pkg/repository/session_repository"
	user_repository "Skyline/pkg/repository/user-repository"
	"Skyline/pkg/usecases/role_usecase"
	"Skyline/pkg/usecases/user_usecase"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type accountUsecase struct {
	userUsecase       user_usecase.UserUsecaseInterface
	sessionRepository session_repository.SessionRepositoryInterface
	userRepository    user_repository.UserRepositoryInterface
	roleUsecase       role_usecase.RoleUsecaseInterface
}

func NewAccountUsecase(userUsecase user_usecase.UserUsecaseInterface,
	sessionRepository session_repository.SessionRepositoryInterface,
	userRepository user_repository.UserRepositoryInterface,
	roleUsecase role_usecase.RoleUsecaseInterface) AccountUsecaseInterface {
	return &accountUsecase{
		userUsecase:       userUsecase,
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
		roleUsecase:       roleUsecase,
	}
}

func (usecase accountUsecase) Login(loginRequest *user_models.LoginRequest, ClientIp string, UserAgent string) (*user_models.LoginResponse, error) {
	user, err := usecase.userRepository.GetByEmail(loginRequest.Email)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(loginRequest.Password, user.Password)
	if err != nil {
		return nil, errors.New("Your email or password is incorrect!")
	}

	accessToken, err := usecase.createAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpiredAt, err := usecase.createRefreshToken(user)
	if err != nil {
		return nil, err
	}

	session := &session_models.Session{
		UserId:       user.UserId,
		RefreshToken: refreshToken,
		UserAgent:    UserAgent,
		ClientIp:     ClientIp,
		IsBlocked:    false,
		ExpiresAt:    refreshExpiredAt,
		CreatedAt:    time.Now(),
	}
	_, err = usecase.sessionRepository.Create(session)
	if err != nil {
		return nil, err
	}

	response := &user_models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return response, err
}

func (usecase accountUsecase) createAccessToken(user *user_models.User) (string, error) {
	role, err, _ := usecase.roleUsecase.Get(user.RoleId)
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["UserId"] = user.UserId
	claims["Email"] = user.Email
	claims["RoleName"] = role.Title
	claims["CreatedAt"] = time.Now()
	claims["ExpiredAt"] = time.Now().Add(time.Minute * time.Duration(utils.AppConfig.AccessTokenDuration))

	tokenString, err := token.SignedString([]byte(utils.AppConfig.AccessTokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (usecase accountUsecase) createRefreshToken(user *user_models.User) (string, time.Time, error) {
	role, err, _ := usecase.roleUsecase.Get(user.RoleId)
	if err != nil {
		return "", time.Time{}, err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["UserId"] = user.UserId
	claims["Email"] = user.Email
	claims["RoleName"] = role.Title
	claims["CreatedAt"] = time.Now()
	refreshExpiredAt := time.Now().Add(time.Hour * 24 * time.Duration(utils.AppConfig.RefreshTokenDuration))
	claims["ExpiredAt"] = refreshExpiredAt

	tokenString, err := token.SignedString([]byte(utils.AppConfig.AccessTokenSecret))
	if err != nil {
		return "", refreshExpiredAt, err
	}
	return tokenString, refreshExpiredAt, nil
}

func (usecase accountUsecase) ForgetPassword(email string) (bool, error) {
	user, err := usecase.userRepository.GetByEmail(email)
	if err != nil {
		return false, err
	}

	password := utils.RandomString(8)
	HashPassword, err := utils.HashPassword(password)
	if err != nil {
		return false, err
	}
	user.Password = HashPassword

	_, err = usecase.userRepository.Update(user)
	if err != nil {
		return false, err
	}

	err = utils.
		SendEmail("Skyline! forget password ",
			fmt.Sprintf("<h1> Dear user </h3>"+
				"<h3>Your password is changed to : <b> %s </b> </h5>", password),
			email,
			nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
