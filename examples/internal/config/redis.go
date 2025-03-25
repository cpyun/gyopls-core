package config

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/redis/go-redis/v9"
)

var (
	RedisConfig = new(Redis)
	_redis      *redis.Client
)

type Redis struct {
	Dsn string `yaml:"dsn" json:"dsn"`
	Tls *Tls   `yaml:"tls" json:"tls"`
}

type Tls struct {
	Cert string `yaml:"cert" json:"cert"`
	Key  string `yaml:"key" json:"key"`
	Ca   string `yaml:"ca" json:"ca"`
}

// GetRedisClient 获取redis客户端
func GetRedisClient() *redis.Client {
	return _redis
}

// SetRedisClient 设置redis客户端
func SetRedisClient(c *redis.Client) {
	if _redis != nil && _redis != c {
		_redis.Shutdown(context.TODO())
	}
	_redis = c
}

func (e *Redis) GetRedisOptions() (*redis.Options, error) {
	ro, err := redis.ParseURL(e.Dsn)
	if err != nil {
		return nil, err
	}

	if ro.TLSConfig != nil {
		ro.TLSConfig, err = e.getTLS(e.Tls)
	}
	return ro, err
}

func (e *Redis) getTLS(c *Tls) (*tls.Config, error) {
	if c != nil && c.Cert != "" {
		// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
		cert, err := tls.LoadX509KeyPair(c.Cert, c.Key)
		if err != nil {
			fmt.Printf("tls.LoadX509KeyPair err: %v\n", err)
			return nil, err
		}
		// 创建一个新的、空的 CertPool，并尝试解析 PEM 编码的证书，解析成功会将其加到 CertPool 中
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(c.Ca)
		if err != nil {
			fmt.Printf("ioutil.ReadFile err: %v\n", err)
			return nil, err
		}

		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			fmt.Println("certPool.AppendCertsFromPEM err")
			return nil, err
		}
		return &tls.Config{
			// 设置证书链，允许包含一个或多个
			Certificates: []tls.Certificate{cert},
			// 要求必须校验客户端的证书
			ClientAuth: tls.RequireAndVerifyClientCert,
			// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
			ClientCAs: certPool,
		}, nil
	}
	return nil, nil
}
