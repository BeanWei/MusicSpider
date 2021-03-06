package musicspider

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// 创建指定长度的key
func createSecKey(size int, keys string) string {
	rs := ""
	for i := 0; i < size; i++ {
		pos := rand.Intn(len(keys))
		rs += keys[pos : pos+1]
	}
	return rs
}

// 通过 CBC 模式的AES加密，用sKey加密sSrc
func aesEncrypt(sSrc, sKey, vi string) string {
	block, _ := aes.NewCipher([]byte(sKey))
	padding := block.BlockSize() - len([]byte(sSrc))%block.BlockSize()
	src := append([]byte(sSrc), bytes.Repeat([]byte{byte(padding)}, padding)...)
	model := cipher.NewCBCEncrypter(block, []byte(vi))
	cipherText := make([]byte, len(src))
	model.CryptBlocks(cipherText, src)
	return base64.StdEncoding.EncodeToString(cipherText)
}

// 字符串补零
func addPadding(encText, modulus string) string {
	ml := len(modulus)
	for i := 0; ml > 0 && modulus[i:i+1] == "0"; i++ {
		ml--
	}
	num := ml - len(encText)
	prefix := ""
	for i := 0; i < num; i++ {
		prefix += ""
	}
	return prefix + encText
}

// 解析加密过的请求响应
func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}

// 对 Key 加密
func rsaEncrypt(key, pubKey, modulus string) string {
	key = reverseStr(key)
	// 将 key 转 ascii 编码 然后转成 16 进制字符串
	hexKey := ""
	for _, char := range []rune(key) {
		hexKey += fmt.Sprintf("%x", int(char))
	}
	// 将 16进制 的 三个参数 转为10进制的 bigint
	bigRKey, _ := big.NewInt(0).SetString(hexKey, 16)
	bigPubKey, _ := big.NewInt(0).SetString(pubKey, 16)
	bigModulus, _ := big.NewInt(0).SetString(modulus, 16)
	// 执行幂乘取模运算得到最终的bigint结果
	bigRs := bigRKey.Exp(bigRKey, bigPubKey, bigModulus)
	// 将结果转为 16进制字符串
	hexRs := fmt.Sprintf("%x", bigRs)
	// 可能存在不满256位的情况，要在前面补0补满256位
	return addPadding(hexRs, modulus)
}

func NeteaseAESCBC(text string) *strings.Reader {
	const modulus = "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	const pubkey = "010001"
	const nonce = "0CoJUm6Qyw8W8jud"
	const vi = "0102030405060708"
	const keys = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"

	secKey := createSecKey(16, keys)
	aes1 := aesEncrypt(text, nonce, vi)
	aes2 := aesEncrypt(aes1, secKey, vi)
	encSecKey := rsaEncrypt(secKey, pubkey, modulus)
	form := url.Values{}
	form.Set("params", aes2)
	form.Set("encSecKey", encSecKey)
	return strings.NewReader(form.Encode())
}

func NeteaseEncryptId(id string) (idDecoded string) {
	//magic := []string{"3", "g", "o", "8", "&", "$", "8", "*", "3", "*", "3", "h", "0", "k", "(", "2", ")", "2"}
	magic := "3go8&$8*3*3h0k(2)2"
	var s string
	for i := 0; i < len(id); i++ {
		s += fmt.Sprintf("%c", rune(id[i]+magic[i%len(magic)]))
	}
	idDecoded = md5Encrpyt16(s)
	idDecoded = strings.Replace(idDecoded, "/", "_", -1)
	idDecoded = strings.Replace(idDecoded, "+", "-", -1)
	return
}

func BaiduAESCBC(text string) string {
	const key = "DBEECF8C50FD160E"
	const vi = "1231021386755796"
	data := fmt.Sprintf(`songid="%s"&ts="%v"`, text, time.Now().Unix())
	aesText := aesEncrypt(data, key, vi)
	return aesText

}

// 凯撒矩阵解密, 还原真实文本
func Caesar(text string) string {
	// 凯撒行数
	rows, _ := strconv.Atoi(fmt.Sprintf("%c", text[0]))
	// 凯撒矩阵最后一列的长度
	remainder := (len(text) - 1) % rows
	// 凯撒矩阵最后一行的长度
	cutLength := (len(text) - 1 - remainder) / rows
	paddingStr := ""
	for i := 1; i <= remainder; i++ {
		paddingStr += fmt.Sprintf("%c", text[i*cutLength+i])
	}
	matrixStr := ""
	for indexA := 1; indexA <= cutLength; indexA++ {
		for indexB := 0; indexB < rows; indexB++ {
			currentIndex := indexA + indexB*cutLength
			if indexB <= remainder {
				matrixStr += fmt.Sprintf("%c", text[currentIndex+indexB])
			} else {
				matrixStr += fmt.Sprintf("%c", text[currentIndex+remainder])
			}
		}
	}
	matrixStr += paddingStr
	return matrixStr
}

/* TODO 虾米站点的参数解密
func XiamiSign()  {

}
*/
