package main

import (
    "log"
    "time"
    "github.com/irgipaulius/pdfgo-kit/pkg"
)

func main() {
    start := time.Now()

    err := pkg.QuickRender(
        "test/fixtures/simple.jsx",
        "test/output.pdf",
    )
    
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Total time: %v", time.Since(start))
}