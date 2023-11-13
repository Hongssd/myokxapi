package myokxapi

import (
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"time"

	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_64 = 64
	BIT_SIZE_32 = 32
)

type RequestType string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

var NIL_REQBODY = []byte{}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

func GetPointer[T any](v T) *T {
	return &v
}

func HmacSha256(secret, data string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return h.Sum(nil)
}

// Request 发送请求
func Request(url string, reqBody []byte, method string, isGzip bool) ([]byte, error) {
	return RequestWithHeader(url, reqBody, method, map[string]string{}, isGzip)
}

func RequestWithHeader(url string, reqBody []byte, method string, headerMap map[string]string, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}
	req.Close = true
	req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	log.Debug("reqURL: ", req.URL.String())
	if reqBody != nil && len(reqBody) > 0 {
		log.Debug("reqBody: ", string(reqBody))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	data, err := io.ReadAll(body)
	// log.Debug(string(data))
	return data, err
}

type MyOkx struct {
}

const (
	OKX_API_HTTP      = "www.okx.com"
	OKX_API_WEBSOCKET = "wss://ws.okx.com:8443/ws/v5/public"
	IS_GZIP           = true
)

type APIType int

const (
	REST APIType = iota
)

type Client struct {
	APIKey     string
	SecretKey  string
	Passphrase string
}

type RestClient struct {
	c *Client
}

type PublicRestClient RestClient

type PrivateRestClient RestClient

func (*MyOkx) NewRestClient(APIKey, SecretKey, Passphrase string) *RestClient {
	client := &RestClient{
		c: &Client{
			APIKey:     APIKey,
			SecretKey:  SecretKey,
			Passphrase: Passphrase,
		},
	}
	return client
}

func (c *RestClient) PublicRestClient() *PublicRestClient {
	return &PublicRestClient{
		c: c.c,
	}
}

func (c *RestClient) PrivateRestClient() *PrivateRestClient {
	return &PrivateRestClient{
		c: c.c,
	}
}

// 通用接口调用
func okxCallAPI[T any](client *Client, url url.URL, reqBody []byte, method string) (*OkxRestRes[T], error) {
	body, err := Request(url.String(), reqBody, method, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

// 通用鉴权接口调用
func okxCallAPIWithSecret[T any](client *Client, url url.URL, reqBody []byte, method string) (*OkxRestRes[T], error) {

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	requestPath := url.Path
	query := url.RawQuery

	hmacSha256Data := timestamp + method + requestPath
	if query != "" {
		hmacSha256Data += "?" + query
	}
	if len(reqBody) != 0 {
		hmacSha256Data += string(reqBody)
	}
	sign := base64.StdEncoding.EncodeToString(HmacSha256(client.SecretKey, hmacSha256Data))

	// log.Warn("timestamp: ", timestamp)
	// log.Warn("method: ", method)
	// log.Warn("requestPath: ", requestPath)
	// log.Warn("query: ", query)
	// log.Warn("reqBody: ", string(reqBody))
	// log.Warn("hmacSha256Data: ", hmacSha256Data)
	// log.Warn("sign: ", sign)

	body, err := RequestWithHeader(url.String(), reqBody, method,
		map[string]string{
			"OK-ACCESS-KEY":        client.APIKey,
			"OK-ACCESS-SIGN":       sign,
			"OK-ACCESS-TIMESTAMP":  timestamp,
			"OK-ACCESS-PASSPHRASE": client.Passphrase,
		}, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

// URL标准封装 带路径参数
func okxHandlerRequestAPIWithPathQueryParam[T any](apiType APIType, request *T, name string) url.URL {
	query := okxHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     OkxGetRestHostByAPIType(apiType),
		Path:     name,
		RawQuery: query,
	}
	return u
}

// URL标准封装 不带路径参数
func okxHandlerRequestAPIWithoutPathQueryParam(apiType APIType, name string) url.URL {
	// query := okxHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     OkxGetRestHostByAPIType(apiType),
		Path:     name,
		RawQuery: "",
	}
	return u
}

func okxHandlerReq[T any](req *T) string {
	var paramBuffer bytes.Buffer

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		paramName := t.Field(i).Tag.Get("json")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			paramBuffer.WriteString(paramName + "=" + v.Field(i).Elem().String() + "&")
		case reflect.Int, reflect.Int64:
			paramBuffer.WriteString(paramName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
		case reflect.Float32, reflect.Float64:
			paramBuffer.WriteString(paramName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
		case reflect.Bool:
			paramBuffer.WriteString(paramName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			ToStringMethod := sv.MethodByName("String")
			params := make([]reflect.Value, 0)
			result := ToStringMethod.Call(params)
			paramBuffer.WriteString(paramName + "=" + result[0].String() + "&")
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			paramBuffer.WriteString(paramName + "=" + url.QueryEscape(string(d)) + "&")
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", paramName, v.Field(i).Elem().Kind())
		}
	}
	return strings.TrimRight(paramBuffer.String(), "&")
}

func OkxGetRestHostByAPIType(apiType APIType) string {
	switch apiType {
	case REST:
		return OKX_API_HTTP
	default:
		return ""
	}
}

func interfaceStringToFloat64(inter interface{}) float64 {
	return stringToFloat64(inter.(string))
}

func interfaceStringToInt64(inter interface{}) int64 {
	return int64(inter.(float64))
}

func stringToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, BIT_SIZE_64)
	return f
}
