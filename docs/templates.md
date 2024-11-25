# Templates

PDFgo-kit allows you to use React components as PDF templates, giving you access to the entire React ecosystem.

## Quick Start

1. Install the package:
   npm install @pdfgo/react

2. Create your template (any .jsx/.tsx file):

```jsx
import { Table, Typography } from "antd";
import "antd/dist/antd.css";
import MyChart from "./components/Chart";
import logo from "./assets/logo.png";
export default function InvoiceTemplate({ data }) {
  return (
    <div className="pdf-container">
      <img src={logo} alt="Company Logo" />
      <Typography.Title>Invoice #{data.id}</Typography.Title>
      {/ Use any React components /}
      <Table dataSource={data.items} pagination={false} />
      {/ Custom components work too /}
      <MyChart data={data.analytics} />
    </div>
  );
}
```

3. Generate PDF:

```javascript
import { renderToPDF } from "@pdfgo/react";
const pdf = await renderToPDF(InvoiceTemplate, {
  data: {
    id: "INV-123",
    items: [/ your data /],
    analytics: [/ chart data /],
  },
  config: {
    size: "A4", // 'A4', 'Letter', or custom dimensions
    dpi: 300, // Output quality
    margin: 20, // Page margins in mm
    timeout: 5000, // Max render wait time in ms
    orientation: "portrait",
  },
});
// Save to file
await pdf.save("invoice.pdf");
// Or get buffer
const buffer = await pdf.toBuffer();
```

## Features

- ✨ Use ANY React component library (AntD, MUI, etc.)
- 📊 Charts and graphics support
- 🎨 Full CSS support
- 🖼️ Image handling
- 📝 Custom fonts
- 📱 Responsive layouts
- 🔄 Dynamic content

## Page Breaks

Control page breaks using CSS:

```css
.page-break {
  page-break-after: always;
}
.keep-together {
  page-break-inside: avoid;
}
```

## Best Practices

1. Container Sizing

   - Use fixed widths or percentages
   - A4 dimensions: 210mm × 297mm
   - Letter dimensions: 215.9mm × 279.4mm

2. Images

   - Use absolute paths or import images
   - Supported formats: PNG, JPEG, SVG
   - Recommended resolution: 300 DPI

3. Fonts

   - Web fonts are supported
   - Local fonts must be properly loaded
   - Use font-display: block for consistency

4. Performance
   - Avoid animations
   - Disable interactive elements
   - Set reasonable timeout values
   - Consider breaking large documents into sections

## Examples

### Business Invoice

```jsx
import { Layout, Table, Typography } from "antd";
import "./invoice.css";
export default function Invoice({ data }) {
  return (
    <Layout className="invoice">
      <header>
        <img src={data.logo} />
        <Typography.Title>Invoice</Typography.Title>
      </header>
      <Table
        dataSource={data.items}
        columns={[
          { title: "Item", dataIndex: "name" },
          { title: "Price", dataIndex: "price" },
        ]}
      />
      <div className="page-break" />
      <footer>
        <p>Thank you for your business!</p>
      </footer>
    </Layout>
  );
}
```

### Crazy Startup Pitch Deck

```jsx
import { Card, Space, Typography } from "antd";
import { Chart } from "react-chartjs-2";
import { GoogleMap, Marker } from "@react-google-maps/api";
import { calculateROI, predictGrowth } from "./financial-magic";
import { SplineChart } from "./components/3d-charts";
import { ModelViewer } from "@google/model-viewer";
import companyLogo from "./assets/unicorn.svg";
import coolProduct from "./assets/product.glb"; // 3D model
import confetti from "canvas-confetti";
// Custom component that renders beautiful gradients
import { GradientBg } from "./components/fancy-bg";

export default function CrazyPitch({ data }) {
  // Do some serious mathematics
  const roi = calculateROI(data.investment, data.returns);
  const growth = predictGrowth(data.users, data.timeline);
  // All our office locations
  const offices = data.locations.map((loc) => ({
    lat: loc.latitude,
    lng: loc.longitude,
  }));
  return (
    <Space direction="vertical" size="large">
      <GradientBg>
        <img src={companyLogo} alt="Next Unicorn Ltd." />
        <Typography.Title>
          We're Disrupting {data.industry}
          With {data.buzzwords.join(" & ")}
        </Typography.Title>
      </GradientBg>
      {/ 3D product showcase /}
      <Card title="Revolutionary Product">
        <ModelViewer
          src={coolProduct}
          auto-rotate
          camera-controls
          shadow-intensity="1"
        />
      </Card>
      {/ Fancy 3D chart showing "growth" /}
      <div className="keep-together">
        <Card title="🚀 Exponential Growth">
          <SplineChart data={growth} theme="cyberpunk" animation="smooth" />
        </Card>
      </div>
      {/ Real-time market data /}
      <Card title="Market Opportunity">
        <Chart
          type="radar"
          data={data.marketSize}
          options={{
            plugins: {
              glitter: true, // why not
            },
          }}
        />
        <Typography.Text>ROI: {roi.toFixed(2)}% (trust us)</Typography.Text>
      </Card>
      {/ Global presence /}
      <div className="page-break" />
      <Card title="Global Empire">
        <GoogleMap
          zoom={2}
          center={{ lat: 0, lng: 0 }}
          mapContainerStyle={{
            height: "400px",
            width: "100%",
          }}
        >
          {offices.map((office, i) => (
            <Marker
              key={i}
              position={office}
              icon={{
                url: companyLogo,
                scaledSize: { width: 30, height: 30 },
              }}
            />
          ))}
        </GoogleMap>
      </Card>
      {/ QR code for their "app" /}
      <Card title="Download Our App">
        <QRCode value={data.appUrl} style={{ margin: "0 auto" }} />
      </Card>
      {/ Weather widget because why not /}
      <WeatherWidget latitude={offices[0].lat} longitude={offices[0].lng} />
      <footer style={{ textAlign: "center" }}>
        <Typography.Text type="secondary">
          All numbers generated using advanced AI
          <br />
          AI may or may not be a random number generator
        </Typography.Text>
      </footer>
    </Space>
  );
}
```

[Back to Documentation](README.md)
