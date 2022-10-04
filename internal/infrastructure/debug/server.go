package debug

import (
	"log"
	"net/http"
	"net/http/pprof"
)

func StartServer(addr string) error {
	mux := http.NewServeMux()
	mux.Handle("/", http.RedirectHandler("/debug/pprof/", http.StatusFound))
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	return http.ListenAndServe(addr, mux) //nolint:gosec // timeouts not needed for internal debug API
}

func StartServerAsync(addr string) {
	go func() {
		log.Println("Debug server listening on " + addr)
		if err := StartServer(addr); err != nil {
			log.Println("Debug server error:", err)
		}
	}()
}
