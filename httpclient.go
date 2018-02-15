package httpclient

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	logger "github.com/paulofilipem/iup-simple-logger"
)

var (
	//Request Params
	Method       string
	Url          string
	Body         string
	HeadersValue *[]string
	CookieValue  *[]string
	CaFileName   string
	CertFileName string
	KeyFileName  string

	EnableCookie = true //Storage received cookie on memory

	//Response Params
	ResponseCode    int
	ResponseBody    string
	ResponseHeaders *[]string
)

func Test() {

}

func Clean() {
	Method = "GET"
	Url = ""
	Body = ""
	HeadersValue = &[]string{}
	CookieValue = &[]string{}
	CaFileName = ""
	CertFileName = ""
	KeyFileName = ""
}

func Send() string {

	if HeadersValue == nil {
		HeadersValue = &[]string{}
	}

	if CookieValue == nil {
		CookieValue = &[]string{}
	}

	return HTTPRequest(Method, Url, Body, HeadersValue, CookieValue, CaFileName, CertFileName, KeyFileName)
}

func HTTPRequest(method string, url string, body string, headersValue *[]string, cookieValue *[]string, caFileName string, certFileName string, keyFileName string) string {

	var debugHTTP = "HTTP Request, to=" + url + ", method=" + method

	/**
	 * HTTP(s) request
	 */
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	checkError(err)

	for _, value := range *headersValue {
		debugHTTP += ",Header=" + value
		result := strings.Split(value, ":")
		req.Header.Add(result[0], strings.Trim(result[1], " "))
	}

	if *cookieValue != nil {
		for _, value := range *cookieValue {
			debugHTTP += ",Cookie=" + value
			req.Header.Add("Cookie", value)
		}
	}

	/**
	 * Simple HTTP Adapter Client
	 */
	client := &http.Client{}

	/**
	 * SSL/TLS
	 */
	if caFileName != "" {
		//client := &http.Client{Transport: transport}
		debugHTTP += ",CAFile=" + caFileName
	}

	/**
	 * SSL/TLS CLIENT
	 */
	if certFileName != "" && keyFileName != "" {
		var (
			caFile   = flag.String("CA", caFileName, "")
			certFile = flag.String("cert", certFileName, "")
			keyFile  = flag.String("key", keyFileName, "")
		)

		flag.Parse()

		// Load client cert
		cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
		if err != nil {
			log.Fatal(err)
		}

		// Load CA cert
		caCert, err := ioutil.ReadFile(*caFile)
		if err != nil {
			log.Fatal(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		// Setup HTTPS client
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		}
		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		client = &http.Client{Transport: transport}
	}

	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	//If u enabled "EnableCookie"
	if EnableCookie {
		if resp.Header["Set-Cookie"] != nil {
			*cookieValue = resp.Header["Set-Cookie"]
		}
	}

	ResponseHeaders = &[]string{}
	for HeaderName, HeaderValue := range resp.Header {
		var headerContent = HeaderName + ": " + strings.Join(HeaderValue, ",")
		*ResponseHeaders = append(*ResponseHeaders, headerContent)
	}

	ResponseCode = resp.StatusCode
	ResponseBody = string(data)

	debugHTTP += "\nrequestBody: \n" + body
	debugHTTP += "\nresponseBody: \n" + ResponseBody

	logger.Debug(debugHTTP)

	return ResponseBody
}

func HttpURLEncode(value string) string {
	return url.QueryEscape(value)
}

func checkError(err error) {
	if err != nil {
		logError(err)
		os.Exit(1)
	}
}

func logError(err error) {
	logger.Error(err.Error())
}
