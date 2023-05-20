package password

import "go.uber.org/zap"

type DefaultPasswordManager struct {
	passwordEncoder   PasswordEncoder
	passwordGenerator PasswordGenerator
}

func NewDefaultPasswordManager(passwordEncoder PasswordEncoder, passwordGenerator PasswordGenerator) *DefaultPasswordManager {

	if passwordEncoder == nil {
		zap.L().Fatal("starting up - error setting up default password manager: passwordEncoder is nil")
	}

	if passwordGenerator == nil {
		zap.L().Fatal("starting up - error setting up default password manager: passwordGenerator is nil")
	}

	return &DefaultPasswordManager{
		passwordEncoder:   passwordEncoder,
		passwordGenerator: passwordGenerator,
	}
}

func (manager *DefaultPasswordManager) Encode(rawPassword string) (*string, error) {
	return manager.passwordEncoder.Encode(rawPassword)
}

func (manager *DefaultPasswordManager) Matches(encodedPassword string, rawPassword string) (*bool, error) {
	return manager.passwordEncoder.Matches(encodedPassword, rawPassword)
}

func (manager *DefaultPasswordManager) UpgradeEncoding(encodedPassword string) (*bool, error) {
	return manager.passwordEncoder.UpgradeEncoding(encodedPassword)
}

func (manager *DefaultPasswordManager) Generate() string {
	return manager.passwordGenerator.Generate()
}

func (manager *DefaultPasswordManager) Validate(rawPassword string) error {
	return manager.passwordGenerator.Validate(rawPassword)
}
