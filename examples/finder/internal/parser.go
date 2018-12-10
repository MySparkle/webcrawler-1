package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"webcrawler/module"
)

func genResponseParsers() []module.ParseResponse {
	parseLink := func(httpResp *http.Response, respDepth uint32) ([]module.Data, []error) {
		dataList := make([]module.Data, 0)
		//检查响应
		if httpResp == nil {
			return nil, []error{fmt.Errorf("nil HTTP response")}
		}
		httpReq := httpResp.Request
		if httpReq == nil {
			return nil, []error{fmt.Errorf("nil HTTP resquest")}
		}
		reqURL := httpReq.URL
		if httpResp.StatusCode != 200 {
			err := fmt.Errorf("unsupported status code %d (requestURL: %s)", httpResp.StatusCode, reqURL)
			return nil, []error{err}
		}
		body := httpResp.Body
		if body == nil {
			err := fmt.Errorf("nil HTTP response body (requestURL: %s)", reqURL)
			return nil, []error{err}
		}
		//检查HTTP响应头中的内容类型
		var matchedContentType bool
		if httpResp.Header != nil {
			contentTypes := httpResp.Header["Content-Type"]
			for _, ct := range contentTypes {
				if strings.HasPrefix(ct, "text/html") {
					matchedContentType = true
					break

				}
			}
		}
		if !matchedContentType {
			return dataList, nil
		}
		//解析HTTP响应体
		doc, err := goquery.NewDocumentFormReader(body)
		if err != nil {
			return dataList, []error{err}
		}
		errs := make([]error, 0)
		//查找a标签并提取链接地址
		doc.Find("a").Each(func(index int, sel *goquery.Selection) {
			href, exists := sel.Attr("href")
			//前期过滤
			if !exists || href == "" || href == "#" || href == "/" {
				return
			}
			href = strings.TrimSpace(href)
			lowerHref := strings.ToLower(href)
			if href == "" || strings.HasPrefix(lowerHref, "javascript") {
				return
			}
			aURL, err := url.Parse(href)
			if err != nil {
				logger.Warnf("An error occurs when parsing attribute %q in tag %q : %s (href: %s)", err, "href", "a", href)
				return
			}
			if aURL.IsAbs() {
				aURL = reqURL.ResolveReference(aURL)
			}
			httpReq, err := http.NewRequest("GET", aURL.String(), nil)
			if err != nil {
				errs = append(errs, err)
			} else {
				req := module.NewRequest(httpReq, respDepth)
				dataList = append(dataList, req)
			}
		})
		//查找img标签并提取地址
		doc.Find("img").Each(func(index int, sel *goquery.Selection) {
			imgSrc, exists := sel.Attr("src")
			if !exists || imgSrc == "" || imgSrc == "#" || imgSrc == "/" {
				return
			}
			imgSrc = strings.TrimSpace(imgSrc)
			imgURL, err := url.Parse(imgSrc)
			if err != nil {
				errs = append(errs, err)
				return
			}
			if !imgURL.IsAbs() {
				imgURL = reqURL.ResolveReference(imgURL)
			}
			httpReq, err := http.NewRequest("GET", imgURL.String(), nil)
			if err != nil {
				errs = append(errs, err)
			} else {
				req := module.NewRequest(httpReq, respDepth)
				dataList = append(dataList, req)
			}
		})
		return dataList, errs
	}
	parseImg := func(httpResp *http.Response, resqDepth uint32) ([]module.Data, []error{
		return nil,[]error{fmt.Errorf("XXX")}
	})

}
