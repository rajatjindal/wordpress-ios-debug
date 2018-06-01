package blog

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func ServeBlog() {
	server := &http.Server{Addr: ":443"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))

		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()

		makeRequest(r)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server.TLSConfig = &tls.Config{
		NextProtos: []string{"h2"},
	}
	fmt.Println("starting blog on 443")
	fmt.Println(server.ListenAndServeTLS("tls.crt", "tls.key"))
	fmt.Println("done blog on 443")
}

func makeRequest(r *http.Request) {
	request, err := http.NewRequest(http.MethodPost, "https://blog.rajatjindal.com/xmlrpc.php", r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	request.Header = r.Header
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("response code %d", resp.StatusCode)
	fmt.Println("response body %s", resp.Body)
}
