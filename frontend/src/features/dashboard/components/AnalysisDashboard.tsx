import { Row, Col } from "antd";
import DashboardCard from "../components/DashboardCard";
import DailyTrendChart from "../components/charts/DailyTrendChart";
import MonthlyTrendChart from "../components/charts/MonthlyTrendChart";
import ProductStatusChart from "../components/charts/ProductStatusChart";
import TopCustomerCitiesChart from "../components/charts/TopCustomerCitiesChart";

const AnalysisDashboard = () => {
  return (
    <Row gutter={[16, 16]}>
      <Col xs={24} md={12}>
        <DashboardCard title="Daily Trend of Delivered Products" height={350}>
          <DailyTrendChart />
        </DashboardCard>
      </Col>
      <Col xs={24} md={12}>
        <DashboardCard title="Monthly Trend of Delivered Products" height={350}>
          <MonthlyTrendChart />
        </DashboardCard>
      </Col>
      <Col xs={24} md={12}>
        <DashboardCard title="Daily Product Purchase Status" height={350}>
          <ProductStatusChart />
        </DashboardCard>
      </Col>
      <Col xs={24} md={12}>
        <DashboardCard title="Top N Customer Cities" height={350}>
          <TopCustomerCitiesChart />
        </DashboardCard>
      </Col>
    </Row>
  );
};

export default AnalysisDashboard;
