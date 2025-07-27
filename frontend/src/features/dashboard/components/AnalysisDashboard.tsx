import { Col, Row } from "antd";
import BarChart from "../components/charts/BarChart";
import ColumnChart from "../components/charts/ColumnChart";
import LineChart from "../components/charts/LineChart";
import PieChart from "../components/charts/PieChart";
import NumberDisplay from "./charts/NumberDisplay";

const AnalysisDashboard = () => {
  const dailyTrendData = [
    { date: "2025-07-01", value: 120 },
    { date: "2025-07-02", value: 200 },
    { date: "2025-07-03", value: 150 },
    { date: "2025-07-04", value: 80 },
    { date: "2025-07-05", value: 170 },
  ];

  const monthlyTrendData = [
    { month: "Jan", value: 300 },
    { month: "Feb", value: 450 },
    { month: "Mar", value: 320 },
    { month: "Apr", value: 510 },
  ];

  const productStatusData = [
    { type: "Delivered", value: 27 },
    { type: "In Transit", value: 25 },
    { type: "Pending", value: 18 },
    { type: "Cancelled", value: 10 },
  ];

  const topCitiesData = [
    { city: "Jakarta", customers: 200 },
    { city: "Surabaya", customers: 150 },
    { city: "Bandung", customers: 100 },
    { city: "Medan", customers: 80 },
  ];

  return (
    <Row gutter={[16, 16]}>
      <Col xs={24} md={12}>
        <NumberDisplay title="Test" value={200} />
      </Col>
      <Col xs={24} md={12}>
        <LineChart
          title="Daily Trend of Delivered Products"
          data={dailyTrendData}
          xField="date"
          yField="value"
          height={300}
        />
      </Col>
      <Col xs={24} md={12}>
        <ColumnChart
          title="Monthly Trend of Delivered Products"
          data={monthlyTrendData}
          xField="month"
          yField="value"
          height={300}
        />
      </Col>
      <Col xs={24} md={12}>
        <PieChart
          title="Daily Product Purchase Status"
          data={productStatusData}
          angleField="value"
          colorField="type"
          height={300}
        />
      </Col>
      <Col xs={24} md={12}>
        <BarChart
          title="Top N Customer Cities"
          data={topCitiesData}
          xField="customers"
          yField="city"
          seriesField="city"
          height={300}
        />
      </Col>
    </Row>
  );
};

export default AnalysisDashboard;
