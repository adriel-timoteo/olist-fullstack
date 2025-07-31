import { Col, Divider, Row, Typography } from "antd";

const { Title, Paragraph, Text } = Typography;

const AboutPage = () => {
  return (
    <Row justify="center" gutter={[24, 24]}>
      <Col xs={24} md={18} lg={14}>
        <Typography>
          <Title level={2}>About Olist Dashboard</Title>
          <Paragraph>
            The <strong>Olist Dashboard</strong> is your centralized platform
            for analyzing e-commerce data, visualizing key metrics, and enabling
            strategic decisions based on the Olist Brazilian dataset. It's
            designed to be fast, insightful, and easy to use.
          </Paragraph>

          <Divider />

          <Title level={4}>Key Features</Title>
          <ul className="list-disc pl-6 space-y-2">
            <li>
              <Text strong>Real-Time KPIs:</Text> Instantly track revenue,
              customer trends, and repeat purchases.
            </li>
            <li>
              <Text strong>Interactive Visuals:</Text> Drill down into charts
              for time-series, category, and geographic insights.
            </li>
            <li>
              <Text strong>Modern Stack:</Text> Powered by React, PostgreSQL,
              and Airflow for end-to-end visibility.
            </li>
            <li>
              <Text strong>Responsive Design:</Text> Accessible from desktop or
              mobile.
            </li>
          </ul>

          <Divider />

          <Title level={4}>How to Run Locally</Title>
          <Paragraph>
            <ol className="list-decimal pl-6 space-y-2">
              <li>
                Clone the repository from GitHub:
                <pre className="bg-gray-100 p-2 rounded mt-1">
                  git clone https://github.com/your-username/olist-dashboard.git
                </pre>
              </li>
              <li>
                Start the Airflow pipeline using Docker Compose:
                <pre className="bg-gray-100 p-2 rounded mt-1">
                  docker compose up airflow
                </pre>
              </li>
              <li>
                Start the backend server with live reload (Go + Air):
                <pre className="bg-gray-100 p-2 rounded mt-1">air</pre>
              </li>
              <li>
                Start the frontend development server:
                <pre className="bg-gray-100 p-2 rounded mt-1">
                  npm install
                  <br />
                  npm run dev
                </pre>
              </li>
              <li>
                Ensure PostgreSQL is running and pre-loaded with:
                <pre className="bg-gray-100 p-2 rounded mt-1">
                  DB: olist_db
                  <br />
                  Use <code>init.sql</code> to initialize schema & data
                </pre>
              </li>
            </ol>
          </Paragraph>

          <Divider />

          <Title level={4}>Alternatively: Hosted on AWS</Title>
          <Paragraph>
            This application is also hosted on AWS with the following services:
            <ul className="list-disc pl-6 space-y-2 mt-2">
              <li>
                <Text strong>Backend:</Text> Deployed via EC2 with Go backend
                service
              </li>
              <li>
                <Text strong>Frontend:</Text> Hosted using AWS Amplify or S3 +
                CloudFront
              </li>
              <li>
                <Text strong>Database:</Text> PostgreSQL on AWS RDS using
                <code>olist_db</code>
              </li>
              <li>
                <Text strong>ETL Pipeline:</Text> Managed using Apache Airflow
                running in Docker on EC2
              </li>
            </ul>
          </Paragraph>

          <Divider />

          <Title level={4}>Built With</Title>
          <ul className="list-disc pl-6 space-y-2">
            <li>React & TypeScript</li>
            <li>Ant Design & TailwindCSS</li>
            <li>PostgreSQL + AWS RDS</li>
            <li>AWS S3 for raw data storage</li>
            <li>Apache Airflow for ETL orchestration</li>
            <li>Go (Gin) for backend REST API</li>
          </ul>

          <Divider />

          <Paragraph type="secondary" className="text-center">
            Designed & developed by Adriel Timoteo as a final assignment
          </Paragraph>
        </Typography>
      </Col>
    </Row>
  );
};

export default AboutPage;
