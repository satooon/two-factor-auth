//go:generate gen
package yaml

import (
	"net/url"
)

type _proxy struct {
	ProxyConfSlice ProxyConfSlice `yaml:"proxy"`
}

// +gen * slice:"Where,All,Any,First"
type ProxyConf struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	StripPath   string `yaml:"strip_path"`
	Dest        string `yaml:"dest"`
}

var proxy = &_proxy{}

func LoadProxy(filename string) error {
	return Load(filename, proxy)
}

func Proxy() *_proxy {
	return proxy
}

func (y *_proxy) GetProxyConfSlice() ProxyConfSlice {
	return y.ProxyConfSlice
}

func (p *ProxyConf) DestUrl() (*url.URL, error) {
	return url.Parse(p.Dest)
}
