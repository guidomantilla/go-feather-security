package password

import (
	"strings"

	"go.uber.org/zap"
)

var SupportedDecoders = map[string]PasswordEncoder{
	Argon2PrefixKey: NewArgon2PasswordEncoder(),
	BcryptPrefixKey: NewBcryptPasswordEncoder(),
	Pbkdf2PrefixKey: NewPbkdf2PasswordEncoder(),
	ScryptPrefixKey: NewScryptPasswordEncoder(),
}

type DelegatingPasswordEncoderOption func(encoder *DelegatingPasswordEncoder)

type DelegatingPasswordEncoder struct {
	encoder  PasswordEncoder
	decoders map[string]PasswordEncoder
}

func NewDelegatingPasswordEncoder(encoder PasswordEncoder, options ...DelegatingPasswordEncoderOption) *DelegatingPasswordEncoder {

	if encoder == nil {
		zap.L().Fatal("starting up - error setting up delegating password encoder: encoder is nil")
	}

	delegator := &DelegatingPasswordEncoder{
		encoder:  encoder,
		decoders: SupportedDecoders,
	}

	for _, opt := range options {
		opt(delegator)
	}

	return delegator
}

func WithSupportedDecoders(decoders map[string]PasswordEncoder) DelegatingPasswordEncoderOption {
	return func(encoder *DelegatingPasswordEncoder) {
		encoder.decoders = decoders
	}
}

func (encoder *DelegatingPasswordEncoder) Encode(rawPassword string) (*string, error) {
	return encoder.encoder.Encode(rawPassword)
}

func (encoder *DelegatingPasswordEncoder) Matches(encodedPassword string, rawPassword string) (*bool, error) {

	for prefix, encoder := range encoder.decoders {
		if strings.HasPrefix(encodedPassword, prefix) {
			return encoder.Matches(encodedPassword, rawPassword)
		}
	}

	return nil, ErrPasswordEncoderNotFound
}

func (encoder *DelegatingPasswordEncoder) UpgradeEncoding(encodedPassword string) (*bool, error) {

	for prefix, encoder := range encoder.decoders {
		if strings.HasPrefix(encodedPassword, prefix) {
			return encoder.UpgradeEncoding(encodedPassword)
		}
	}

	return nil, ErrPasswordEncoderNotFound
}
