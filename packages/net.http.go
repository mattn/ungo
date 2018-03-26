// +build go1.6

package packages

import (
	"net/http"
	"reflect"
)

func init() {
	if _, ok := Packages["net/http"]; !ok {
		Packages["net/http"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["net/http"]; !ok {
		PackageTypes["net/http"] = make(map[string]interface{})
	}
	var client http.Client
	var closenotifier http.CloseNotifier
	var connstate http.ConnState
	var cookie http.Cookie
	var cookiejar http.CookieJar
	var dir http.Dir
	var file http.File
	var filesystem http.FileSystem
	var flusher http.Flusher
	var handler http.Handler
	var handlerfunc http.HandlerFunc
	var header http.Header
	var hijacker http.Hijacker
	var protocolerror http.ProtocolError
	var request http.Request
	var response http.Response
	var responsewriter http.ResponseWriter
	var roundtripper http.RoundTripper
	var servemux http.ServeMux
	var server http.Server
	var transport http.Transport
	Packages["net/http"]["CanonicalHeaderKey"] = http.CanonicalHeaderKey
	Packages["net/http"]["DefaultClient"] = http.DefaultClient
	Packages["net/http"]["DefaultMaxHeaderBytes"] = http.DefaultMaxHeaderBytes
	Packages["net/http"]["DefaultMaxIdleConnsPerHost"] = http.DefaultMaxIdleConnsPerHost
	Packages["net/http"]["DefaultServeMux"] = http.DefaultServeMux
	Packages["net/http"]["DefaultTransport"] = http.DefaultTransport
	Packages["net/http"]["DetectContentType"] = http.DetectContentType
	Packages["net/http"]["ErrBodyNotAllowed"] = http.ErrBodyNotAllowed
	Packages["net/http"]["ErrBodyReadAfterClose"] = http.ErrBodyReadAfterClose
	Packages["net/http"]["ErrContentLength"] = http.ErrContentLength
	Packages["net/http"]["ErrHandlerTimeout"] = http.ErrHandlerTimeout
	Packages["net/http"]["ErrHeaderTooLong"] = http.ErrHeaderTooLong
	Packages["net/http"]["ErrHijacked"] = http.ErrHijacked
	Packages["net/http"]["ErrLineTooLong"] = http.ErrLineTooLong
	Packages["net/http"]["ErrMissingBoundary"] = http.ErrMissingBoundary
	Packages["net/http"]["ErrMissingContentLength"] = http.ErrMissingContentLength
	Packages["net/http"]["ErrMissingFile"] = http.ErrMissingFile
	Packages["net/http"]["ErrNoCookie"] = http.ErrNoCookie
	Packages["net/http"]["ErrNoLocation"] = http.ErrNoLocation
	Packages["net/http"]["ErrNotMultipart"] = http.ErrNotMultipart
	Packages["net/http"]["ErrNotSupported"] = http.ErrNotSupported
	Packages["net/http"]["ErrShortBody"] = http.ErrShortBody
	Packages["net/http"]["ErrSkipAltProtocol"] = http.ErrSkipAltProtocol
	Packages["net/http"]["ErrUnexpectedTrailer"] = http.ErrUnexpectedTrailer
	Packages["net/http"]["ErrWriteAfterFlush"] = http.ErrWriteAfterFlush
	Packages["net/http"]["Error"] = http.Error
	Packages["net/http"]["FileServer"] = http.FileServer
	Packages["net/http"]["Get"] = http.Get
	Packages["net/http"]["Handle"] = http.Handle
	Packages["net/http"]["HandleFunc"] = http.HandleFunc
	Packages["net/http"]["Head"] = http.Head
	Packages["net/http"]["ListenAndServe"] = http.ListenAndServe
	Packages["net/http"]["ListenAndServeTLS"] = http.ListenAndServeTLS
	Packages["net/http"]["MaxBytesReader"] = http.MaxBytesReader
	Packages["net/http"]["MethodConnect"] = http.MethodConnect
	Packages["net/http"]["MethodDelete"] = http.MethodDelete
	Packages["net/http"]["MethodGet"] = http.MethodGet
	Packages["net/http"]["MethodHead"] = http.MethodHead
	Packages["net/http"]["MethodOptions"] = http.MethodOptions
	Packages["net/http"]["MethodPatch"] = http.MethodPatch
	Packages["net/http"]["MethodPost"] = http.MethodPost
	Packages["net/http"]["MethodPut"] = http.MethodPut
	Packages["net/http"]["MethodTrace"] = http.MethodTrace
	Packages["net/http"]["NewFileTransport"] = http.NewFileTransport
	Packages["net/http"]["NewRequest"] = http.NewRequest
	Packages["net/http"]["NewServeMux"] = http.NewServeMux
	Packages["net/http"]["NotFound"] = http.NotFound
	Packages["net/http"]["NotFoundHandler"] = http.NotFoundHandler
	Packages["net/http"]["ParseHTTPVersion"] = http.ParseHTTPVersion
	Packages["net/http"]["ParseTime"] = http.ParseTime
	Packages["net/http"]["Post"] = http.Post
	Packages["net/http"]["PostForm"] = http.PostForm
	Packages["net/http"]["ProxyFromEnvironment"] = http.ProxyFromEnvironment
	Packages["net/http"]["ProxyURL"] = http.ProxyURL
	Packages["net/http"]["ReadRequest"] = http.ReadRequest
	Packages["net/http"]["ReadResponse"] = http.ReadResponse
	Packages["net/http"]["Redirect"] = http.Redirect
	Packages["net/http"]["RedirectHandler"] = http.RedirectHandler
	Packages["net/http"]["Serve"] = http.Serve
	Packages["net/http"]["ServeContent"] = http.ServeContent
	Packages["net/http"]["ServeFile"] = http.ServeFile
	Packages["net/http"]["SetCookie"] = http.SetCookie
	Packages["net/http"]["StateActive"] = http.StateActive
	Packages["net/http"]["StateClosed"] = http.StateClosed
	Packages["net/http"]["StateHijacked"] = http.StateHijacked
	Packages["net/http"]["StateIdle"] = http.StateIdle
	Packages["net/http"]["StateNew"] = http.StateNew
	Packages["net/http"]["StatusAccepted"] = http.StatusAccepted
	Packages["net/http"]["StatusBadGateway"] = http.StatusBadGateway
	Packages["net/http"]["StatusBadRequest"] = http.StatusBadRequest
	Packages["net/http"]["StatusConflict"] = http.StatusConflict
	Packages["net/http"]["StatusContinue"] = http.StatusContinue
	Packages["net/http"]["StatusCreated"] = http.StatusCreated
	Packages["net/http"]["StatusExpectationFailed"] = http.StatusExpectationFailed
	Packages["net/http"]["StatusForbidden"] = http.StatusForbidden
	Packages["net/http"]["StatusFound"] = http.StatusFound
	Packages["net/http"]["StatusGatewayTimeout"] = http.StatusGatewayTimeout
	Packages["net/http"]["StatusGone"] = http.StatusGone
	Packages["net/http"]["StatusHTTPVersionNotSupported"] = http.StatusHTTPVersionNotSupported
	Packages["net/http"]["StatusInternalServerError"] = http.StatusInternalServerError
	Packages["net/http"]["StatusLengthRequired"] = http.StatusLengthRequired
	Packages["net/http"]["StatusMethodNotAllowed"] = http.StatusMethodNotAllowed
	Packages["net/http"]["StatusMovedPermanently"] = http.StatusMovedPermanently
	Packages["net/http"]["StatusMultipleChoices"] = http.StatusMultipleChoices
	Packages["net/http"]["StatusNetworkAuthenticationRequired"] = http.StatusNetworkAuthenticationRequired
	Packages["net/http"]["StatusNoContent"] = http.StatusNoContent
	Packages["net/http"]["StatusNonAuthoritativeInfo"] = http.StatusNonAuthoritativeInfo
	Packages["net/http"]["StatusNotAcceptable"] = http.StatusNotAcceptable
	Packages["net/http"]["StatusNotFound"] = http.StatusNotFound
	Packages["net/http"]["StatusNotImplemented"] = http.StatusNotImplemented
	Packages["net/http"]["StatusNotModified"] = http.StatusNotModified
	Packages["net/http"]["StatusOK"] = http.StatusOK
	Packages["net/http"]["StatusPartialContent"] = http.StatusPartialContent
	Packages["net/http"]["StatusPaymentRequired"] = http.StatusPaymentRequired
	Packages["net/http"]["StatusPreconditionFailed"] = http.StatusPreconditionFailed
	Packages["net/http"]["StatusPreconditionRequired"] = http.StatusPreconditionRequired
	Packages["net/http"]["StatusProxyAuthRequired"] = http.StatusProxyAuthRequired
	Packages["net/http"]["StatusRequestEntityTooLarge"] = http.StatusRequestEntityTooLarge
	Packages["net/http"]["StatusRequestHeaderFieldsTooLarge"] = http.StatusRequestHeaderFieldsTooLarge
	Packages["net/http"]["StatusRequestTimeout"] = http.StatusRequestTimeout
	Packages["net/http"]["StatusRequestURITooLong"] = http.StatusRequestURITooLong
	Packages["net/http"]["StatusRequestedRangeNotSatisfiable"] = http.StatusRequestedRangeNotSatisfiable
	Packages["net/http"]["StatusResetContent"] = http.StatusResetContent
	Packages["net/http"]["StatusSeeOther"] = http.StatusSeeOther
	Packages["net/http"]["StatusServiceUnavailable"] = http.StatusServiceUnavailable
	Packages["net/http"]["StatusSwitchingProtocols"] = http.StatusSwitchingProtocols
	Packages["net/http"]["StatusTeapot"] = http.StatusTeapot
	Packages["net/http"]["StatusTemporaryRedirect"] = http.StatusTemporaryRedirect
	Packages["net/http"]["StatusText"] = http.StatusText
	Packages["net/http"]["StatusTooManyRequests"] = http.StatusTooManyRequests
	Packages["net/http"]["StatusUnauthorized"] = http.StatusUnauthorized
	Packages["net/http"]["StatusUnavailableForLegalReasons"] = http.StatusUnavailableForLegalReasons
	Packages["net/http"]["StatusUnsupportedMediaType"] = http.StatusUnsupportedMediaType
	Packages["net/http"]["StatusUseProxy"] = http.StatusUseProxy
	Packages["net/http"]["StripPrefix"] = http.StripPrefix
	Packages["net/http"]["TimeFormat"] = http.TimeFormat
	Packages["net/http"]["TimeoutHandler"] = http.TimeoutHandler
	PackageTypes["net/http"]["Client"] = reflect.TypeOf(&client).Elem()
	PackageTypes["net/http"]["CloseNotifier"] = reflect.TypeOf(&closenotifier).Elem()
	PackageTypes["net/http"]["ConnState"] = reflect.TypeOf(&connstate).Elem()
	PackageTypes["net/http"]["Cookie"] = reflect.TypeOf(&cookie).Elem()
	PackageTypes["net/http"]["CookieJar"] = reflect.TypeOf(&cookiejar).Elem()
	PackageTypes["net/http"]["Dir"] = reflect.TypeOf(&dir).Elem()
	PackageTypes["net/http"]["File"] = reflect.TypeOf(&file).Elem()
	PackageTypes["net/http"]["FileSystem"] = reflect.TypeOf(&filesystem).Elem()
	PackageTypes["net/http"]["Flusher"] = reflect.TypeOf(&flusher).Elem()
	PackageTypes["net/http"]["Handler"] = reflect.TypeOf(&handler).Elem()
	PackageTypes["net/http"]["HandlerFunc"] = reflect.TypeOf(&handlerfunc).Elem()
	PackageTypes["net/http"]["Header"] = reflect.TypeOf(&header).Elem()
	PackageTypes["net/http"]["Hijacker"] = reflect.TypeOf(&hijacker).Elem()
	PackageTypes["net/http"]["ProtocolError"] = reflect.TypeOf(&protocolerror).Elem()
	PackageTypes["net/http"]["Request"] = reflect.TypeOf(&request).Elem()
	PackageTypes["net/http"]["Response"] = reflect.TypeOf(&response).Elem()
	PackageTypes["net/http"]["ResponseWriter"] = reflect.TypeOf(&responsewriter).Elem()
	PackageTypes["net/http"]["RoundTripper"] = reflect.TypeOf(&roundtripper).Elem()
	PackageTypes["net/http"]["ServeMux"] = reflect.TypeOf(&servemux).Elem()
	PackageTypes["net/http"]["Server"] = reflect.TypeOf(&server).Elem()
	PackageTypes["net/http"]["Transport"] = reflect.TypeOf(&transport).Elem()
}
