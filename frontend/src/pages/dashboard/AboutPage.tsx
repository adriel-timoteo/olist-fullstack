import { Col, Divider, Row, Typography } from "antd";
import DashboardLayout from "../../components/DashboardLayout";

const { Title, Paragraph, Text } = Typography;

const AboutPage = () => {
  return (
    <DashboardLayout>
      <Row justify="center" gutter={[24, 24]}>
        <Col xs={24} md={18} lg={14}>
          <Typography>
            <Title level={2}>About Olist Dashboard</Title>
            <Paragraph>
              The <strong>Olist Dashboard</strong> is designed to be a central
              hub for analyzing key metrics and tracking performance from an
              e-commerce. The data used in this project is derived from the
              Brazillian ecommerce Olist.
              (https://www.kaggle.com/datasets/olistbr/brazilian-ecommerce)
            </Paragraph>

            <Divider />

            <Title level={4}>Built With</Title>
            <ul className="list-disc pl-6 space-y-2">
              <li>React & TypeScript</li>
              <li>Golang & Gin for REST API</li>
              <li>Ant Design</li>
              <li>PostgreSQL + AWS RDS</li>
              <li>AWS S3 for raw data storage</li>
              <li>Apache Airflow for data pipeline orchestration</li>
            </ul>

            <Divider />

            <Paragraph type="secondary" className="text-center">
              Designed and developed by Adriel Timoteo
            </Paragraph>
          </Typography>
        </Col>
      </Row>
    </DashboardLayout>
  );
};

export default AboutPage;
