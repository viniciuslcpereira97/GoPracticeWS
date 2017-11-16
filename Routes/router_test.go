package routes

import (
    "fmt"
    "net/http"
    "testing"
)

func GenericHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Print("test")
}

func TestRouterIsInitialized(t *testing.T) {
    InitRouter()
    
    if Router == nil {
        t.Fatalf("Router not initialized")
    }
}

func TestIsRegisteringGetRoutes(t *testing.T) {
    Get("test.get", "/test/route", GenericHandler)

    if len(routes) != 1 {
        t.Fatalf("GET Routes - expected length 1 but received %d", len(routes))
    }
}

func TestIsRegisteringPostRoutes(t *testing.T) {
    Post("test.post", "/test/route", GenericHandler)

    if len(routes) != 2 {
        t.Fatalf("POST Routes - expected length 2 but received %d", len(routes))
    }
}
