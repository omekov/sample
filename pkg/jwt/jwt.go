package jwt

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/omekov/sample/internal/config"
	"github.com/omekov/sample/internal/model"
	"github.com/omekov/sample/pkg/contant"
	"github.com/omekov/sample/pkg/domain"
	"golang.org/x/crypto/bcrypt"
)

// Claims ...
type Claims struct {
	User domain.User
	jwtgo.StandardClaims
}

// Config ...
type Config struct {
	refreshTokenSecret []byte
	accessTokenSecret  []byte
	lifeTimeAccess     int
	lifeTimeRefresh    int
}

func NewJWT(cfg config.JWT) *Config {
	return &Config{
		refreshTokenSecret: []byte(cfg.AccessSecret),
		accessTokenSecret:  []byte(cfg.RefreshSecret),
		lifeTimeAccess:     cfg.LifeTimeAccess,
		lifeTimeRefresh:    cfg.LifeTimeRefresh,
	}
}

// NewRefreshJWT ...
func (c *Config) generateRefresh(user *domain.User) (string, error) {
	return c.getClaims(user, c.lifeTimeRefresh, c.refreshTokenSecret)
}

// NewAccessJWT ...
func (c *Config) generateAccess(user *domain.User) (string, error) {
	return c.getClaims(user, c.lifeTimeAccess, c.accessTokenSecret)
}

func (c *Config) getClaims(user *domain.User, lifeTime int, secret []byte) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, Claims{
		User: *user,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(lifeTime) * time.Minute).Unix(),
		},
	})
	return token.SignedString(secret)
}

// GetClaims ...
func (c *Config) GetParseClaims(token string) (*Claims, error) {
	var claims Claims
	t, err := c.parseJWT(token, c.accessTokenSecret, &claims)
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, contant.ErrInvalidAccessToken
	}
	return &claims, nil
}

// GetToken ...
func (c *Config) GetToken(user *domain.User) (model.Token, error) {
	token := model.Token{}
	var err error
	token.Refreshtoken, err = c.generateRefresh(user)
	if err != nil {
		return token, err
	}

	token.AccessToken, err = c.generateAccess(user)
	if err != nil {
		return token, err
	}

	return token, nil
}

//// GetRefreshJWT ...
//func (c *Config) GetRefreshJWT(refToken string) (*http.Token, error) {
//	var claims Claims
//	t, err := c.parseJWT(refToken, c.refreshTokenSecret, &claims)
//	if err != nil {
//		return nil, err
//	}
//	if !t.Valid {
//		return nil, contant.ErrInvalidAccessToken
//	}
//	newRefToken, err := c.NewRefreshJWT(&claims.Customer)
//	if err != nil {
//		return nil, err
//	}
//	newAccToken, err := c.NewAccessJWT(&claims.Customer)
//	if err != nil {
//		return nil, err
//	}
//	return &http.Token{
//		AccessToken:  newAccToken,
//		Refreshtoken: newRefToken,
//	}, nil
//}

func (c *Config) parseJWT(tokenString string, key []byte, claims *Claims) (*jwtgo.Token, error) {
	return jwtgo.ParseWithClaims(tokenString, claims, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, contant.ErrInvalidAccessToken
		}
		return key, nil
	})
}

// Validate ...
func (c *Config) Validate(credential model.Credential) error {
	return validation.ValidateStruct(
		validation.Field(&credential.Username, validation.Required, is.Email),
		validation.Field(&credential.Password, validation.Length(8, 32)),
	)
}

// BeforeCreate ...
//func (c *Config) BeforeCreate(password string) (enc string, err error) {
//	if len(password) > 0 {
//		enc, err = encryptString(password)
//		if err != nil {
//			return enc, err
//		}
//	}
//	return enc, nil
//}

//func encryptString(p string) (string, error) {
//	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
//	if err != nil {
//		return "", err
//	}
//	return string(b), nil
//}

// ComparePassword ...
func (c *Config) ComparePassword(encPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(encPassword), []byte(password))
}

//// Sanitize ...
//func (c *Customer) Sanitize() {
//	c.ID = primitive.NewObjectID()
//	c.Password = ""
//	c.RepeatPassword = ""
//	c.ReleaseDate = time.Now()
//	c.RegistrationDate = time.Now()
//}

/*
func (c *Config) parseJWT(tokenString string, key []byte) (*Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if token.Valid {
		if claims, ok := token.Claims.(Claims); ok {
			return &claims, nil
		}
		return nil, err
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, err
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
*/
