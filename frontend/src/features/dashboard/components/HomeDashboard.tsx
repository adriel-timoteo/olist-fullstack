import { Col, Row, Typography } from "antd";
import { repeatPurchaseRateApi, totalUniqueCustApi } from "../api/customer";
import { averageOrderValue, totalRevenue } from "../api/order";
import NumberDisplay from "./charts/NumberDisplay";

const { Title, Paragraph } = Typography;

const HomeDashboard = () => {
  return (
    <div className="p-6 min-h-screen bg-gray-50">
      <Row gutter={[16, 16]}>
        {/* KPIs */}
        <Col xs={12} md={6}>
          <NumberDisplay
            title="Total Revenue"
            fetchData={() => totalRevenue().then((res) => res.data.count)}
            prefix="R$"
          />
        </Col>
        <Col xs={12} md={6}>
          <NumberDisplay
            title="Total Unique Customer"
            fetchData={() => totalUniqueCustApi().then((res) => res.data.count)}
          />
        </Col>
        <Col xs={12} md={6}>
          <NumberDisplay
            title="Repeat Purchase Rate"
            fetchData={() =>
              repeatPurchaseRateApi().then((res) => res.data.rate * 100)
            }
            suffix="%"
            precision={2}
          />
        </Col>
        <Col xs={12} md={6}>
          <NumberDisplay
            title="Average Order Value"
            fetchData={() => averageOrderValue().then((res) => res.data.count)}
            prefix="R$"
            precision={2}
          />
        </Col>
      </Row>

      {/* Additional Content */}
      <Row gutter={[20, 20]} className="mt-6">
        <Col span={24}>
          <Title level={4}>Welcome to Olist Dashboard</Title>
          <Paragraph>
            This dashboard helps you monitor the core performance indicators of
            the Olist marketplace. Here’s what your data is saying:
          </Paragraph>
          <ul className="list-disc pl-5 space-y-1">
            <li>
              <strong>Total Revenue:</strong> Overview of total sales across all
              time.
            </li>
            <li>
              <strong>Total Unique Customers:</strong> Number of distinct
              buyers.
            </li>
            <li>
              <strong>Repeat Purchase Rate:</strong> Indicates customer loyalty
              and satisfaction.
            </li>
            <li>
              <strong>Average Order Value:</strong> Measures the average spend
              per order.
            </li>
          </ul>
          <Paragraph className="mt-4">
            More data can be seen in the Analysis Page.
          </Paragraph>
        </Col>
      </Row>
    </div>
  );
};

export default HomeDashboard;
