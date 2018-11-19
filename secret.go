package musicspider

import (
	"fmt"
	"strings"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
)

func aesEncrypt(text, seckey, vi string) string {
	pad := 16 - len(text) % 16
	text = text + strings.Repeat(fmt.Sprintf("%c", rune(pad)), pad)
	block, _ := aes.NewCipher([]byte(seckey))
	blockMode := cipher.NewCBCEncrypter(block, []byte(vi))
	encryptor := make([]byte, len(text))
	blockMode.CryptBlocks(encryptor, []byte(text))
	cipherText := base64.StdEncoding.EncodeToString(encryptor)
	return cipherText
}


func NeteaseAESCBC(text string) string {
	const modulus =	"157794750267131502212476817800345498121872783333389747424011531025366277535262539913701806290766479189477533597854989606803194253978660329941980786072432806427833685472618792592200595694346872951301770580765135349259590167490536138082469680638514416594216629258349130257685001248172188325316586707301643237607"
	const pubkey  = "65537"
	const nonce  = "0CoJUm6Qyw8W8jud"
	const vi = "0102030405060708"
	const skey = "B3v3kH4vRPWRJFfH"
	//const randomHex  = "dbd747a1f907d04b"

	encText := aesEncrypt(aesEncrypt(text, nonce, vi), skey, vi)
	encSecKey := "85302b818aea19b68db899c25dac229412d9bba9b3fcfe4f714dc016bc1686fc446a08844b1f8327fd9cb623cc189be00c5a365ac835e93d4858ee66f43fdc59e32aaed3ef24f0675d70172ef688d376a4807228c55583fe5bac647d10ecef15220feef61477c28cae8406f6f9896ed329d6db9f88757e31848a6c2ce2f94308"
	return fmt.Sprintf(`{"params": %s, "encSecKey": %s}`, encText, encSecKey)
}































