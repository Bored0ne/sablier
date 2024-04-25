package caddy

import (
	"context"
	"fmt"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"io"
	"net/http"
)

func init() {
	caddy.RegisterModule(SablierMiddleware{})
}

type SablierMiddleware struct {
	Config  Config
	client  *http.Client
	request *http.Request
}

func FindReplaceAll(repl *caddy.Replacer, arr []string) (output []string) {
	for _, item := range arr {
		output = append(output, repl.ReplaceAll(item, "ERROR_REPLACEMENT2"))
	}
	return output
}

// CaddyModule returns the Caddy module information.
func (SablierMiddleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.sablier",
		New: func() caddy.Module { return new(SablierMiddleware) },
	}
}

// Provision implements caddy.Provisioner.
func (m *SablierMiddleware) Provision(ctx caddy.Context) error {
	//	req, err := m.Config.BuildRequest()
	//	if err != nil {
	//		return err
	//	}
	//	m.request = req
	//	m.client = &http.Client{}

	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (sm SablierMiddleware) ServeHTTP(rw http.ResponseWriter, req *http.Request, next caddyhttp.Handler) error {
	repl := req.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)
	fmt.Println("repl")

	//	sm.Config.Names = FindReplaceAll(repl, sm.Config.Names)
	req2, err := sm.Config.BuildRequest(repl)
	fmt.Println("req")
	if err != nil {
		fmt.Println("err")
		return err
	}
	fmt.Println("noerr")
	sm.client = &http.Client{}
	fmt.Println("sm.client")
	sablierRequest := req2.Clone(context.TODO())
	fmt.Println("sablierRequest")
	//	sablierRequest.
	resp, err2 := sm.client.Do(sablierRequest)
	fmt.Println("resp")
	if err2 != nil {
		fmt.Println(err2)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return nil
	}
	fmt.Println("noerr2")
	defer resp.Body.Close()

	fmt.Println("good?")
	if resp.Header.Get("X-Sablier-Session-Status") == "ready" {
		fmt.Println("next")
		err3 := next.ServeHTTP(rw, req)
		fmt.Println("err")
		if err3 != nil {
			fmt.Println(err3.Error())
			return err3
		}
	} else {
		fmt.Println("forward")
		forward(resp, rw)
	}
	fmt.Println("nil")
	return nil
}

func forward(resp *http.Response, rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	rw.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
	io.Copy(rw, resp.Body)
}

// Interface guards
var (
	_ caddy.Provisioner           = (*SablierMiddleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*SablierMiddleware)(nil)
)
