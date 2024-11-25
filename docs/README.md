PDFgo-kit Documentation

Transform any data into beautiful PDFs using the full power of React ecosystem.

Quick Links

- [API Reference](api.md) - Complete list of methods and interfaces
- [Templates](templates.md) - Create PDFs using ANY React component
- [Examples](examples.md) - From simple documents to 3D startup pitches

Use Cases

- [SRT to PDF Conversion](guides/srt2pdf.md) - Convert subtitles to readable documents
- [PDF to Video Generation](guides/pdf2mp4.md) - Transform PDFs into video presentations
- [Business Documents](guides/invoices.md) - Generate professional invoices and reports

Integration

- [React Integration](integration/react.md) - Use your favorite React components
- [Command Line Interface](integration/cli.md) - Batch process files
- [Go Library API](integration/api.md) - Core Go functionality

Technical

- [Performance & Benchmarks](performance.md) - Processing speeds and optimization
- [Contributing](contributing.md) - Help us make PDF generation fun again

Installation

Backend:
go get github.com/irgipaulius/pdfgo-kit

Frontend:
npm install @pdfgo/react

Quick Example

React Template:

```javascript
import { Card, Table } from "antd";
import logo from "./assets/logo.png";

export default function SimpleDoc({ data }) {
  return (
    <Card title="Simple Document">
      <img src={logo} />
      <Table dataSource={data.items} pagination={false} />
    </Card>
  );
}
```

Generate PDF:

```javascript
import { renderToPDF } from "@pdfgo/react";

const pdf = await renderToPDF(SimpleDoc, {
  data: {
    items: [
      { name: "Item 1", price: 100 },
      { name: "Item 2", price: 200 },
    ],
  },
});

// Save or use the buffer
await pdf.save("output.pdf");
```

Features

- 🎨 Use ANY React component (AntD, MUI, Charts, Maps, 3D...)
- 📄 Convert SRT files to beautiful documents
- 🎥 Transform PDFs into videos
- 💼 Generate professional business documents
- 🚀 Browser-side processing with WASM
- 🖥️ CLI for batch processing
- 📊 Full CSS and layout support
- 🌈 Custom fonts and styling
- 🔧 Extensive configuration options
- 🎯 Zero server costs for processing

License
MIT License - see [LICENSE](../LICENSE) for details
