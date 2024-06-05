package myokxapi

import (
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"sync"
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

	//ok模拟盘额外参数
	if NowNetType == TEST_NET {
		req.Header.Set("x-simulated-trading", "1")
	}

	client := &http.Client{}
	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}

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
	OKX_API_HTTP          = "www.okx.com"
	OKX_API_WEBSOCKET     = "ws.okx.com:8443"
	OKX_API_HTTP_AWS      = "aws.okx.com"
	OKX_API_WEBSOCKET_AWS = "wsaws.okx.com:8443"
	IS_GZIP               = true
)

type ServerType int

const (
	BASE ServerType = iota
	AWS
)

var SERVER_TYPE = BASE

func SetServerType(serverType ServerType) {
	SERVER_TYPE = serverType
}

type NetType int

const (
	MAIN_NET NetType = iota
	TEST_NET
)

var NowNetType = MAIN_NET

func SetNetType(netType NetType) {
	NowNetType = netType
}

type APIType int

const (
	REST APIType = iota
	WS_PUBLIC
	WS_PRIVATE
	WS_BUSINESS
)

// 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION ： 期权 ANY： 全部
type InstType string

const (
	SPOT    InstType = "SPOT"
	MARGIN  InstType = "MARGIN"
	SWAP    InstType = "SWAP"
	FUTURES InstType = "FUTURES"
	OPTION  InstType = "OPTION"
	ANY     InstType = "ANY"
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

	//log.Warn(hmacSha256Data)
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
	var argBuffer bytes.Buffer

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		argName := t.Field(i).Tag.Get("json")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			argBuffer.WriteString(argName + "=" + v.Field(i).Elem().String() + "&")
		case reflect.Int, reflect.Int64:
			argBuffer.WriteString(argName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
		case reflect.Float32, reflect.Float64:
			argBuffer.WriteString(argName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
		case reflect.Bool:
			argBuffer.WriteString(argName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			ToStringMethod := sv.MethodByName("String")
			args := make([]reflect.Value, 0)
			result := ToStringMethod.Call(args)
			argBuffer.WriteString(argName + "=" + result[0].String() + "&")
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			argBuffer.WriteString(argName + "=" + url.QueryEscape(string(d)) + "&")
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
		}
	}
	return strings.TrimRight(argBuffer.String(), "&")
}

func OkxGetRestHostByAPIType(apiType APIType) string {

	switch apiType {
	case REST:
		switch SERVER_TYPE {
		case BASE:
			return OKX_API_HTTP
		case AWS:
			return OKX_API_HTTP_AWS
		}
	default:
	}
	return ""
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

type MySyncMap[K any, V any] struct {
	smap sync.Map
}

func NewMySyncMap[K any, V any]() MySyncMap[K, V] {
	return MySyncMap[K, V]{
		smap: sync.Map{},
	}
}
func (m *MySyncMap[K, V]) Load(k K) (V, bool) {
	v, ok := m.smap.Load(k)

	if ok {
		return v.(V), true
	}
	var resv V
	return resv, false
}
func (m *MySyncMap[K, V]) Store(k K, v V) {
	m.smap.Store(k, v)
}

func (m *MySyncMap[K, V]) Delete(k K) {
	m.smap.Delete(k)
}
func (m *MySyncMap[K, V]) Range(f func(k K, v V) bool) {
	m.smap.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}

func (m *MySyncMap[K, V]) Length() int {
	length := 0
	m.Range(func(k K, v V) bool {
		length += 1
		return true
	})
	return length
}

func (m *MySyncMap[K, V]) MapValues(f func(k K, v V) V) *MySyncMap[K, V] {
	var res = NewMySyncMap[K, V]()
	m.Range(func(k K, v V) bool {
		res.Store(k, f(k, v))
		return true
	})
	return &res
}
