package pkg

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func QuickRender(jsxPath, outputPath string) error {
	log.Println("Starting render...")

	// Create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, renderToPDF(&buf)); err != nil {
		return fmt.Errorf("render failed: %v", err)
	}

	if err := os.WriteFile(outputPath, buf, 0644); err != nil {
		return fmt.Errorf("failed to write PDF: %v", err)
	}

	return nil
}

func renderToPDF(res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`data:text/html,
            <!DOCTYPE html>
            <html>
            <head>
                <script src="https://unpkg.com/react@17/umd/react.production.min.js"></script>
                <script src="https://unpkg.com/react-dom@17/umd/react-dom.production.min.js"></script>
                <script src="https://unpkg.com/babel-standalone@6/babel.min.js"></script>
            </head>
            <body>
                <div id="root">Loading...</div>
                <script type="text/babel">
                    function Simple() {
                        return (
                            <div style={{ padding: '20px', border: '2px solid black' }}>
                                <h1>Hello PDFgo-kit!</h1>
                                <p>If you see this in a PDF, it works! 🎉</p>
                            </div>
                        );
                    }
                    
                    ReactDOM.render(
                        <Simple />,
                        document.getElementById('root')
                    );
                </script>
            </body>
            </html>
        `),
		chromedp.Sleep(2000), // Give React time to render
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
