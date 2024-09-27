// fingerprintx-------------------------------------
// @file      : fingerprintx.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2024/9/26 21:20
// -------------------------------------------

package fingerprintx

import (
	"errors"
	"fmt"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/interfaces"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/types"
	"github.com/Autumn-27/ScopeSentry-Scan/pkg/logger"
	"github.com/praetorian-inc/fingerprintx/pkg/plugins"
	"github.com/praetorian-inc/fingerprintx/pkg/scan"
	"net/netip"
	"strconv"
	"strings"
	"time"
)

type Plugin struct {
	Name      string
	Module    string
	Parameter string
	Result    chan interface{}
}

func NewPlugin() *Plugin {
	return &Plugin{
		Name:   "fingerprintx",
		Module: "PortFingerprint",
	}
}

func (p *Plugin) SetResult(ch chan interface{}) {
	p.Result = ch
}

func (p *Plugin) SetName(name string) {
	p.Name = name
}

func (p *Plugin) GetName() string {
	return p.Name
}

func (p *Plugin) SetModule(module string) {
	p.Module = module
}

func (p *Plugin) GetModule() string {
	return p.Module
}

func (p *Plugin) Install() error {
	return nil
}

func (p *Plugin) Check() error {
	return nil
}

func (p *Plugin) SetParameter(args string) {
	p.Parameter = args
}

func (p *Plugin) GetParameter() string {
	return p.Parameter
}

func (p *Plugin) Execute(input interface{}) (interface{}, error) {
	asset, ok := input.(*types.Asset)
	if !ok {
		logger.SlogError(fmt.Sprintf("%v error: %v input is not a string\n", p.Name, input))
		return nil, errors.New("input is not a string")
	}
	fxConfig := scan.Config{
		DefaultTimeout: time.Duration(3) * time.Second,
		FastMode:       false,
		Verbose:        false,
		UDP:            false,
	}
	ip, _ := netip.ParseAddr(asset.IP)
	portUint64, err := strconv.ParseUint(asset.Port, 10, 16)
	if err != nil {
		fmt.Println("转换错误:", err)
		logger.SlogError(fmt.Sprintf("%v 端口转换错误: %v ", p.GetName(), err))
		return nil, err
	}
	target := plugins.Target{
		Address: netip.AddrPortFrom(ip, uint16(portUint64)),
		Host:    asset.Host,
	}
	fingerResults, err := scan.ScanTargets([]plugins.Target{target}, fxConfig)
	for _, fingerResult := range fingerResults {
		if strings.Contains(fingerResult.Protocol, "http") {
			asset.Type = "http"
		} else {
			asset.Type = "other"
		}
		asset.Protocol = fingerResult.Protocol
		asset.TLS = fingerResult.TLS
		asset.Transport = fingerResult.Transport
		asset.Version = fingerResult.Version
		asset.Raw = fingerResult.Raw
		return nil, nil
	}
	return nil, nil
}

func (p *Plugin) Clone() interfaces.Plugin {
	return &Plugin{
		Name:   p.Name,
		Module: p.Module,
	}
}