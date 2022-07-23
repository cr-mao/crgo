package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"hash"
	"io"
	"math/rand"
	"strings"
	"time"
	"unsafe"

	"github.com/golang/protobuf/proto"

	"crgo/infra/bizerror"
)

// MustHash make sure hash function works on io.Reader, otherwise panic.
//
// Example:
//     MustHash(md5.New(), strings.NewReader("this is string"))
//     MustHash(sha256.New(), bytes.NewReader([]byte("this is bytes")))
//     MustHash(hmac.New(sha256.New, []byte("bytes")), bytes.NewReader([]byte("bytes too")))
//
// Or use crypto package
//
// Example:
//     MustHash(crypto.MD5.New(), strings.NewReader("this is string"))
//
// Note os.File is also io.Reader
//
// Example:
//     f, _ := os.Open("/path/to/file")
//     MustHash(crypto.SHA512.New(), f)
//
// See https://pkg.go.dev/crypto#Hash for all available hash functions
func MustHash(h hash.Hash, r io.Reader) string {
	_, err := io.Copy(h, r)
	if err != nil {
		panic(bizerror.Wrap(1001, "encoding failed.", err))
	}

	return hex.EncodeToString(h.Sum(nil))
}

// MD5
func MD5(p []byte) string {
	return MustHash(md5.New(), bytes.NewReader(p))
}

// SHA1
func SHA1(p []byte) string {
	return MustHash(sha1.New(), bytes.NewReader(p))
}

// SHA256
func SHA256(p []byte) string {
	return MustHash(sha256.New(), bytes.NewReader(p))
}

// HMAC-SHA256
// p content
// k key
func HMACSHA256(p []byte, k []byte) string {
	return MustHash(hmac.New(sha256.New, k), bytes.NewReader(p))
}

// MD5 file
func MD5File(src io.Reader) string {
	return MustHash(md5.New(), src)
}

// MustMarshal make sure marshal(interface{}) works, otherwise panic.
//
// Example:
//     MustMarshal(json.Marshal, []string{"hello", "world"})
//     MustMarshal(xml.Marshal, []string{"hello", "world"})
//     MustMarshal(MarshalWrapper(proto.Marshal), msg) // msg := proto.Message
func MustMarshal(marshal func(interface{}) ([]byte, error), v interface{}) []byte {
	bs, err := marshal(v)
	if err != nil {
		panic(bizerror.NewError(1001, fmt.Sprintf("%T marshal failed", marshal), err.Error()))
	}
	return bs
}

func MarshalWrapper(marshal func(proto.Message) ([]byte, error)) func(interface{}) ([]byte, error) {
	return func(m interface{}) ([]byte, error) {
		return marshal(m.(proto.Message))
	}
}

// MustUnmarshal make sure unmarshal([]byte, interface{}) works, otherwise panic.
//
// Example:
//     MustMarshal(json.Unmarshal, []byte(``), v)
//     MustMarshal(xml.Unmarshal, []byte(``), v)
//     MustMarshal(UnmarshalWrapper(proto.Unmarshal), []byte(``), v)
func MustUnmarshal(unmarshal func([]byte, interface{}) error, bs []byte, v interface{}) {
	err := unmarshal(bs, v)
	if err != nil {
		panic(bizerror.NewError(1001, fmt.Sprintf("%T unmarshal failed", unmarshal), err.Error()))
	}
}

func UnmarshalWrapper(unmarshal func([]byte, proto.Message) error) func([]byte, interface{}) error {
	return func(bs []byte, v interface{}) error {
		return unmarshal(bs, v.(proto.Message))
	}
}

// json marshal
func JSONMarshal(v interface{}) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		panic(bizerror.NewError(1001, "json marshal failed", err.Error()))
	}
	return bs
}

// json unmarshal
func JSONUnmarshal(bs []byte, v interface{}) error {
	err := json.Unmarshal(bs, &v)
	return err
}

// xml marshal
func XMLMarshal(v interface{}) []byte {
	bs, err := xml.Marshal(v)
	if err != nil {
		panic(bizerror.NewError(1001, "解析 xml 失败", err.Error()))
	}
	return bs
}

// xml unmarshal
func StructToMap(v interface{}) map[string]interface{} {
	bm := make(map[string]interface{})

	bs := JSONMarshal(v)
	d := json.NewDecoder(strings.NewReader(string(bs)))
	d.UseNumber()
	err := d.Decode(&bm)
	if err != nil {
		panic(bizerror.NewError(1001, "struct to map failed", err.Error()))
	}
	return bm
}

// 获取随机字符串
//    length：字符串长度
func GetRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	var (
		result []byte
		b      []byte
		r      *rand.Rand
	)
	b = []byte(str)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
