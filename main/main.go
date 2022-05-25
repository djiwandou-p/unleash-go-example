package main

import (
    "io"
    "fmt"
    "log"
    "net/http"

    "github.com/Unleash/unleash-client-go/v3"
)

type metricsInterface struct {
}

func init() {
    unleash.Initialize(
        unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("new-feature-satu"),
		unleash.WithUrl("https://app.unleash-hosted.com/demo/api/"),
		unleash.WithCustomHeaders(http.Header{"Authorization": {"56907a2fa53c1d16101d509a10b78e36190b0f918d9f122d"}}),
    )
}

func helloServer(w http.ResponseWriter, req *http.Request) {
    if unleash.IsEnabled("new-feature-satu") {
        io.WriteString(w, "Feature satu enabled\n")
        fmt.Println("Feature satu enabled")
    } else {
        io.WriteString(w, "hello, world! feature satu disable\n")
        fmt.Println("hello, world! feature satu disable")
    }
}

func main() {
    http.HandleFunc("/", helloServer)
    log.Fatal(http.ListenAndServe(":8080", nil))
}