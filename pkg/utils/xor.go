package utils

import (
	"encoding/base64"
	"encoding/json"
)

// XorKey OAuth state 加解密 key（对齐源 ChannelConstants::XOR_KEY）
const XorKey = "FGY454DSWGLLNNqtrxxpoeCUUEE788611VEfdEFASA"

// XorEncrypt 与源 xorEncrypt 一致：按 key 长度分块 XOR 后 base64
func XorEncrypt(str, key string) string {
	if key == "" {
		key = XorKey
	}
	sb := []byte(str)
	kb := []byte(key)
	kLen := len(kb)
	cipher := make([]byte, 0, len(sb))
	for i := 0; i < len(sb); i += kLen {
		end := i + kLen
		if end > len(sb) {
			end = len(sb)
		}
		chunk := sb[i:end]
		for j := range chunk {
			cipher = append(cipher, chunk[j]^kb[j%kLen])
		}
	}
	return base64.StdEncoding.EncodeToString(cipher)
}

// XorDecrypt 与源 xorDecrypt 一致：base64 解码后按 key 长度分块 XOR
func XorDecrypt(str, key string) ([]byte, error) {
	if key == "" {
		key = XorKey
	}
	raw, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	kb := []byte(key)
	kLen := len(kb)
	plain := make([]byte, 0, len(raw))
	for i := 0; i < len(raw); i += kLen {
		end := i + kLen
		if end > len(raw) {
			end = len(raw)
		}
		chunk := raw[i:end]
		for j := range chunk {
			plain = append(plain, chunk[j]^kb[j%kLen])
		}
	}
	return plain, nil
}

// XorDecryptJSON 解密并解析为 map[string]interface{}
func XorDecryptJSON(str, key string) (map[string]interface{}, error) {
	plain, err := XorDecrypt(str, key)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(plain, &m); err != nil {
		return nil, err
	}
	return m, nil
}
