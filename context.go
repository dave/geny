package geny

import (
	"context"
	"fmt"
	"strings"
)

type genyCtxKeyType int

const genyCtxKey genyCtxKeyType = 0

type genyCtx struct {
	Imports map[string]string
	Package string
	Name    string
}

func (g *genyCtx) RegisterAnonymousPackage(path string) {
	if _, ok := g.Imports[path]; !ok {
		g.Imports[path] = "_"
	}
}

func (g *genyCtx) RegisterPackage(path string) (alias string) {

	// first we should normalize the package path e.g. remove trailing slashes
	path = strings.TrimRight(path, "/")

	if path == g.Package {
		panic("Should not register the local package")
	}

	// if the path has already been imported, use the existing alias, but
	// not if it's anonymous (an underscore)
	if alias, ok := g.Imports[path]; ok {
		if alias != "_" {
			return alias
		}
	}

	// get the preferred alias for this package
	p := guessPackageName(path)

	found := func(alias string) bool {
		for _, a := range g.Imports {
			if a == alias {
				return true
			}
		}
		return false
	}

	alias = p
	count := 0
	for {
		if !found(alias) {
			break
		}
		count++
		alias = fmt.Sprintf("%s%d", p, count)
	}

	g.Imports[path] = alias
	return alias

}

func Context(ctx context.Context, pkg string, name string) context.Context {
	gc := &genyCtx{
		Package: pkg,
		Name:    name,
		Imports: map[string]string{},
	}
	return context.WithValue(ctx, genyCtxKey, gc)
}

func FromContext(ctx context.Context) *genyCtx {
	return ctx.Value(genyCtxKey).(*genyCtx)
}

func Scope(ctx context.Context) context.Context {
	// TODO
	return ctx
}

func guessPackageName(path string) string {
	preferred := path
	if strings.Contains(preferred, "/") {
		// if the path contains a "/", use the last part
		preferred = preferred[strings.LastIndex(preferred, "/")+1:]
	}
	if strings.Contains(preferred, "-") {
		// the name usually follows a hyphen - e.g. github.com/foo/go-bar
		// if the package name contains a "-", use the last part
		preferred = preferred[strings.LastIndex(preferred, "-")+1:]
	}
	if strings.Contains(preferred, ".") {
		// dot is commonly usually used as a version - e.g. github.com/foo/bar.v1
		// if the package name contains a ".", use the first part
		preferred = preferred[:strings.Index(preferred, ".")]
	}
	return preferred
}
