// Code generated by 'yaegi extract github.com/Autumn-27/ScopeSentry-Scan/internal/types'. DO NOT EDIT.

package symbols

import (
	"github.com/Autumn-27/ScopeSentry-Scan/internal/types"
	"reflect"
)

func init() {
	Symbols["github.com/Autumn-27/ScopeSentry-Scan/internal/types/types"] = map[string]reflect.Value{
		// type definitions
		"AssetChangeLog":       reflect.ValueOf((*types.AssetChangeLog)(nil)),
		"AssetHttp":            reflect.ValueOf((*types.AssetHttp)(nil)),
		"AssetOther":           reflect.ValueOf((*types.AssetOther)(nil)),
		"ChangeLog":            reflect.ValueOf((*types.ChangeLog)(nil)),
		"CrawlerResult":        reflect.ValueOf((*types.CrawlerResult)(nil)),
		"CrawlerTask":          reflect.ValueOf((*types.CrawlerTask)(nil)),
		"DirResult":            reflect.ValueOf((*types.DirResult)(nil)),
		"DomainResolve":        reflect.ValueOf((*types.DomainResolve)(nil)),
		"DomainSkip":           reflect.ValueOf((*types.DomainSkip)(nil)),
		"HttpResponse":         reflect.ValueOf((*types.HttpResponse)(nil)),
		"HttpSample":           reflect.ValueOf((*types.HttpSample)(nil)),
		"KatanaResult":         reflect.ValueOf((*types.KatanaResult)(nil)),
		"NotificationApi":      reflect.ValueOf((*types.NotificationApi)(nil)),
		"NotificationConfig":   reflect.ValueOf((*types.NotificationConfig)(nil)),
		"PageMonit":            reflect.ValueOf((*types.PageMonit)(nil)),
		"PageMonitBody":        reflect.ValueOf((*types.PageMonitBody)(nil)),
		"PocData":              reflect.ValueOf((*types.PocData)(nil)),
		"PortAlive":            reflect.ValueOf((*types.PortAlive)(nil)),
		"PortDict":             reflect.ValueOf((*types.PortDict)(nil)),
		"Project":              reflect.ValueOf((*types.Project)(nil)),
		"SecretResults":        reflect.ValueOf((*types.SecretResults)(nil)),
		"SensitiveResult":      reflect.ValueOf((*types.SensitiveResult)(nil)),
		"SensitiveRule":        reflect.ValueOf((*types.SensitiveRule)(nil)),
		"SubTakeResult":        reflect.ValueOf((*types.SubTakeResult)(nil)),
		"SubdomainResult":      reflect.ValueOf((*types.SubdomainResult)(nil)),
		"SubdomainTakerFinger": reflect.ValueOf((*types.SubdomainTakerFinger)(nil)),
		"TmpPageMonitResult":   reflect.ValueOf((*types.TmpPageMonitResult)(nil)),
		"UrlResult":            reflect.ValueOf((*types.UrlResult)(nil)),
		"VulnResult":           reflect.ValueOf((*types.VulnResult)(nil)),
		"WebFinger":            reflect.ValueOf((*types.WebFinger)(nil)),
	}
}