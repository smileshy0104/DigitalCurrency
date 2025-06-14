package tools

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Post 发起一个POST请求并返回响应体。
// 参数:
//
//	url: 请求的URL。
//	params: 请求的参数，会被序列化为JSON格式。
//
// 返回值:
//
//	响应体的字节切片和可能发生的错误。
func Post(url string, params any) ([]byte, error) {
	// 创建一个带有超时的context，以防止请求无限期地等待。
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()

	// 将参数序列化为JSON格式。
	marshal, _ := json.Marshal(params)
	s := string(marshal)
	reqBody := strings.NewReader(s)

	// 创建POST请求。
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reqBody)
	if err != nil {
		return nil, err
	}

	// 设置请求头为JSON格式。
	httpReq.Header.Add("Content-Type", "application/json")

	// 发起请求并获取响应。
	httpRsp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	// 读取响应体。
	rspBody, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	return rspBody, nil
}

// GetWithHeader 发起一个GET请求，并携带指定的请求头和代理设置。
// 参数:
//
//	path: 请求的URL路径。
//	m: 包含请求头的映射，其中键是头的名称，值是头的值。
//	proxy: 代理服务器的URL，如果不需要代理则为空。
//
// 返回值:
//
//	[]byte: 响应体的内容。
//	error: 如果在请求过程中发生错误，则返回该错误。
func GetWithHeader(path string, m map[string]string, proxy string) ([]byte, error) {
	// 创建一个带有超时的context，以防止请求无限期地等待。
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()

	// 创建GET请求。
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头。
	if m != nil {
		for k, v := range m {
			httpReq.Header.Set(k, v)
		}
	}
	// 设置请求头为JSON格式。
	httpReq.Header.Add("Content-Type", "application/json")

	// 根据是否提供代理，选择性地创建新的HTTP客户端。
	client := http.DefaultClient
	// 如果提供了代理，则创建新的HTTP客户端，并设置代理服务器。
	if proxy != "" {
		// 解析代理服务器的URL。
		proxyAddress, _ := url.Parse(proxy)
		client = &http.Client{
			Transport: &http.Transport{
				// 设置代理服务器。
				Proxy: http.ProxyURL(proxyAddress),
			},
		}
	}

	httpRsp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	rspBody, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	return rspBody, nil
}

// PostWithHeader 发起一个带有自定义头部的POST请求并返回响应体。
// 参数:
//
//	path: 请求的URL。
//	params: 请求的参数，会被序列化为JSON格式。
//	m: 请求头的键值对。
//	proxy: 代理服务器的URL。
//
// 返回值:
//
//	响应体的字节切片和可能发生的错误。
func PostWithHeader(path string, params any, m map[string]string, proxy string) ([]byte, error) {
	// 创建一个带有超时的上下文，以确保请求不会无限期地等待。
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	// 确保在函数返回前取消上下文，以释放资源。
	defer cancelFunc()

	// 将请求参数转换为JSON格式。
	marshal, _ := json.Marshal(params)
	// 将JSON数据转换为字符串。
	s := string(marshal)
	// 创建一个读取请求体的Reader。
	reqBody := strings.NewReader(s)

	// 创建一个HTTP POST请求。
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, path, reqBody)
	if err != nil {
		return nil, err
	}

	// 设置自定义请求头。
	if m != nil {
		for k, v := range m {
			httpReq.Header.Set(k, v)
		}
	}

	// 根据是否提供了代理地址来决定使用默认的HTTP客户端还是自定义的。
	client := http.DefaultClient
	if proxy != "" {
		// 解析代理地址。
		proxyAddress, _ := url.Parse(proxy)
		// 创建一个带有代理的HTTP客户端。
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyAddress),
			},
		}
	}

	// 发起HTTP请求。
	httpRsp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	// 确保在函数返回前关闭响应体，以释放资源。
	defer httpRsp.Body.Close()

	// 读取服务器的响应体。
	rspBody, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	// 返回服务器的响应体。
	return rspBody, nil
}
