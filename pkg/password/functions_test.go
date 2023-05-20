package password

import (
	"reflect"
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	type args struct {
		saltSize int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateSalt(tt.args.saltSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateSalt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSalt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbkdf2Encode(t *testing.T) {
	type args struct {
		rawPassword string
		salt        []byte
		iterations  int
		keyLength   int
		fn          HashFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pbkdf2Encode(tt.args.rawPassword, tt.args.salt, tt.args.iterations, tt.args.keyLength, tt.args.fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pbkdf2Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pbkdf2Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPbkdf2Decode(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		want1   *int
		want2   *[]byte
		want3   *[]byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, err := Pbkdf2Decode(tt.args.encodedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pbkdf2Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pbkdf2Decode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Pbkdf2Decode() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Pbkdf2Decode() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("Pbkdf2Decode() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

func TestScryptEncode(t *testing.T) {
	type args struct {
		rawPassword string
		salt        []byte
		N           int
		r           int
		p           int
		keyLen      int
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ScryptEncode(tt.args.rawPassword, tt.args.salt, tt.args.N, tt.args.r, tt.args.p, tt.args.keyLen)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScryptEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScryptEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScryptDecode(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		want1   *int
		want2   *int
		want3   *int
		want4   *[]byte
		want5   *[]byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, got4, got5, err := ScryptDecode(tt.args.encodedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScryptDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScryptDecode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ScryptDecode() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ScryptDecode() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("ScryptDecode() got3 = %v, want %v", got3, tt.want3)
			}
			if !reflect.DeepEqual(got4, tt.want4) {
				t.Errorf("ScryptDecode() got4 = %v, want %v", got4, tt.want4)
			}
			if !reflect.DeepEqual(got5, tt.want5) {
				t.Errorf("ScryptDecode() got5 = %v, want %v", got5, tt.want5)
			}
		})
	}
}

func TestArgon2Encode(t *testing.T) {
	type args struct {
		rawPassword string
		salt        []byte
		iterations  int
		memory      int
		threads     int
		keyLen      int
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Argon2Encode(tt.args.rawPassword, tt.args.salt, tt.args.iterations, tt.args.memory, tt.args.threads, tt.args.keyLen)
			if (err != nil) != tt.wantErr {
				t.Errorf("Argon2Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Argon2Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgon2Decode(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		args    args
		want    *string
		want1   *int
		want2   *int
		want3   *int
		want4   *int
		want5   *[]byte
		want6   *[]byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, got4, got5, got6, err := Argon2Decode(tt.args.encodedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Argon2Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Argon2Decode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Argon2Decode() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Argon2Decode() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("Argon2Decode() got3 = %v, want %v", got3, tt.want3)
			}
			if !reflect.DeepEqual(got4, tt.want4) {
				t.Errorf("Argon2Decode() got4 = %v, want %v", got4, tt.want4)
			}
			if !reflect.DeepEqual(got5, tt.want5) {
				t.Errorf("Argon2Decode() got5 = %v, want %v", got5, tt.want5)
			}
			if !reflect.DeepEqual(got6, tt.want6) {
				t.Errorf("Argon2Decode() got6 = %v, want %v", got6, tt.want6)
			}
		})
	}
}
