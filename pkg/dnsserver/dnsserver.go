package dnsserver

import (
	"fmt"
	"log"
	"strconv"

	"github.com/miekg/dns"
)

var records = map[string]string{
	"test.service.":                  "192.168.0.2",
	"blog.rajatjindal.com.":          "192.168.0.103",
	"carenshare.zindagisikhati.com.": "192.168.0.103",
	"myblog.com.":                    "192.168.0.103",
}

//ServeDNS serves fake dns
func ServeDNS() {
	// attach request handler func
	dns.HandleFunc("rajatjindal.com.", handleDnsRequest)
	dns.HandleFunc("zindagisikhati.com.", handleDnsRequest)
	dns.HandleFunc("myblog.com.", handleDnsRequest)

	// start server
	port := 53
	server := &dns.Server{Addr: ":" + strconv.Itoa(port), Net: "udp"}
	log.Printf("Starting at %d\n", port)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		fmt.Println(m.Question)
		switch q.Qtype {
		case dns.TypeA:
			log.Printf("Query for %s\n", q.Name)
			ip := records[q.Name]
			if ip != "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		}
	}
}

func handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}

	w.WriteMsg(m)
}
