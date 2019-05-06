package http_ece

import (
	"crypto/elliptic"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDecryptWithAESGCM(t *testing.T) {
	authSecret := d("9HcXsQe3xLMG/w2HsYKrOA==")
	salt := d("mRGYnIzSJGeZnJ19lgQcfw==")
	privateKey := d("yfSYB+/vCEoWklHCG7F99cQ1vRwemFYn87jZc8PHBwU=")
	senderPublicKey := d("BGJXZ4zDA04RfSgTufdauZXcNYbe3oF/yEri5ETSuZLDx70gYi7w2ytak8U82H01P1HYnIvr2fEeX7NZpeHdnhM=")
	content := d("vOjpVgZE4IYn/uEJKk3DzZ4X+Qr1dgSSUkuIzQE=")
	plaintext, err := Decrypt(content,
		WithEncoding(AESGCM),
		WithSalt(salt),
		WithAuthSecret(authSecret),
		WithPrivate(privateKey),
		WithDh(senderPublicKey),
	)

	if assert.Nil(t, err) {
		assert.Equal(t, "hello world", string(plaintext))
	}
}

func TestDecryptWithAESGCM_2record(t *testing.T) {
	authSecret := d("9HcXsQe3xLMG/w2HsYKrOA==")
	salt := d("mRGYnIzSJGeZnJ19lgQcfw==")
	privateKey := d("yfSYB+/vCEoWklHCG7F99cQ1vRwemFYn87jZc8PHBwU=")
	senderPublicKey := d("BGJXZ4zDA04RfSgTufdauZXcNYbe3oF/yEri5ETSuZLDx70gYi7w2ytak8U82H01P1HYnIvr2fEeX7NZpeHdnhM=")
	content := d("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")
	plaintext, err := Decrypt(content,
		WithEncoding(AESGCM),
		WithSalt(salt),
		WithAuthSecret(authSecret),
		WithDh(senderPublicKey),
		WithPrivate(privateKey),
	)

	if assert.Nil(t, err) {
		assert.Equal(t, strings.Repeat("a", 4095), string(plaintext))
	}
}

func TestDecryptWithAESGCM_WrongKey(t *testing.T) {
	authSecret := d("9HcXsQe3xLMG/w2HsYKrOA==")
	salt := d("mRGYnIzSJGeZnJ19lgQcfw==")
	privateKey, _, _ := randomKey(elliptic.P256())
	senderPublicKey := d("BGJXZ4zDA04RfSgTufdauZXcNYbe3oF/yEri5ETSuZLDx70gYi7w2ytak8U82H01P1HYnIvr2fEeX7NZpeHdnhM=")
	content := d("vOjpVgZE4IYn/uEJKk3DzZ4X+Qr1dgSSUkuIzQE=")
	plaintext, err := Decrypt(content,
		WithEncoding(AESGCM),
		WithSalt(salt),
		WithAuthSecret(authSecret),
		WithPrivate(privateKey),
		WithDh(senderPublicKey),
	)

	assert.EqualError(t, err, "cipher: message authentication failed")
	assert.Nil(t, plaintext)
}

func TestDecryptWithAESGCM_NoAuthSecret(t *testing.T) {
	salt := d("mRGYnIzSJGeZnJ19lgQcfw==")
	privateKey := d("yfSYB+/vCEoWklHCG7F99cQ1vRwemFYn87jZc8PHBwU=")
	senderPublicKey := d("BGJXZ4zDA04RfSgTufdauZXcNYbe3oF/yEri5ETSuZLDx70gYi7w2ytak8U82H01P1HYnIvr2fEeX7NZpeHdnhM=")
	content := d("oK6DiDHgs7kq5E+kzbbgJ14m5j9NhR0zYa8YY/c=")
	plaintext, err := Decrypt(content,
		WithEncoding(AESGCM),
		WithSalt(salt),
		WithPrivate(privateKey),
		WithDh(senderPublicKey),
	)

	if assert.Nil(t, err) {
		assert.Equal(t, "hello world", string(plaintext))
	}
}

func TestDecryptWithAES128GCM(t *testing.T) {
	authSecret := d("9HcXsQe3xLMG/w2HsYKrOA==")
	privateKey := d("yfSYB+/vCEoWklHCG7F99cQ1vRwemFYn87jZc8PHBwU=")
	senderPublicKey := d("BGJXZ4zDA04RfSgTufdauZXcNYbe3oF/yEri5ETSuZLDx70gYi7w2ytak8U82H01P1HYnIvr2fEeX7NZpeHdnhM=")
	content := d("mRGYnIzSJGeZnJ19lgQcfwAAEABBBGJXZ4zDA04RfSgTufdauZXcNYbe3oF/yEri5ETSuZLDx70gYi7w2ytak8U82H01P1HYnIvr2fEeX7NZpeHdnhPF+Ah4eiBMGQcXDvtjM/2s1KUn64dsYvM2ljQ1")
	plaintext, err := Decrypt(content,
		WithEncoding(AES128GCM),
		WithAuthSecret(authSecret),
		WithDh(senderPublicKey),
		WithPrivate(privateKey),
	)

	if assert.Nil(t, err) {
		assert.Equalf(t, "hello world", string(plaintext), "")
	}
}

func TestDecryptWithAES128GCM_WrongKey(t *testing.T) {
	authSecret := d("9HcXsQe3xLMG/w2HsYKrOA==")
	privateKey, _, _ := randomKey(elliptic.P256())
	senderPublicKey := d("BGJXZ4zDA04RfSgTufdauZXcNYbe3oF/yEri5ETSuZLDx70gYi7w2ytak8U82H01P1HYnIvr2fEeX7NZpeHdnhM=")
	content := d("mRGYnIzSJGeZnJ19lgQcfwAAEABBBGJXZ4zDA04RfSgTufdauZXcNYbe3oF/yEri5ETSuZLDx70gYi7w2ytak8U82H01P1HYnIvr2fEeX7NZpeHdnhPF+Ah4eiBMGQcXDvtjM/2s1KUn64dsYvM2ljQ1")
	plaintext, err := Decrypt(content,
		WithEncoding(AES128GCM),
		WithAuthSecret(authSecret),
		WithDh(senderPublicKey),
		WithPrivate(privateKey),
	)

	assert.EqualError(t, err, "cipher: message authentication failed")
	assert.Nil(t, plaintext)
}

func TestDecryptWithAES128GCM_NoAuthSecret(t *testing.T) {
	privateKey := d("yfSYB+/vCEoWklHCG7F99cQ1vRwemFYn87jZc8PHBwU=")
	senderPublicKey := d("BGJXZ4zDA04RfSgTufdauZXcNYbe3oF/yEri5ETSuZLDx70gYi7w2ytak8U82H01P1HYnIvr2fEeX7NZpeHdnhM=")
	content := d("mRGYnIzSJGeZnJ19lgQcfwAAEABBBGJXZ4zDA04RfSgTufdauZXcNYbe3oF/yEri5ETSuZLDx70gYi7w2ytak8U82H01P1HYnIvr2fEeX7NZpeHdnhPF+Ah4eiBMGQcXDvtjM/2s1KUn64dsYvM2ljQ1")
	plaintext, err := Decrypt(content,
		WithEncoding(AES128GCM),
		WithDh(senderPublicKey),
		WithPrivate(privateKey),
	)

	assert.EqualError(t, err, "no authentication secret for webpush")
	assert.Nil(t, plaintext)
}