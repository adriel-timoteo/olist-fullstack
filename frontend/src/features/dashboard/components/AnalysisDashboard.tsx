import DashboardCard from "../components/DashboardCard";
import DailyTrendChart from "../components/charts/DailyTrendChart";
import MonthlyTrendChart from "../components/charts/MonthlyTrendChart";
import ProductStatusChart from "../components/charts/ProductStatusChart";
import TopCustomerCitiesChart from "../components/charts/TopCustomerCitiesChart";

const AnalysisDashboard = () => {
  return (
    <div
      style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: "16px" }}
    >
      <DashboardCard title="Daily Trend of Delivered Products">
        <DailyTrendChart />
      </DashboardCard>

      <DashboardCard title="Monthly Trend of Delivered Products">
        <MonthlyTrendChart />
      </DashboardCard>

      <DashboardCard title="Daily Product Purchase Status">
        <ProductStatusChart />
      </DashboardCard>

      <DashboardCard title="Top N Customer Cities">
        <TopCustomerCitiesChart />
      </DashboardCard>
    </div>
  );
};

export default AnalysisDashboard;
