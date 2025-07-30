import { Col, Row } from "antd";
import dayjs from "dayjs";
import { totalUniqueCustApi } from "../api/count";
import { repeatPurchaseRateApi } from "../api/rate";
import { deliveredTrendApi } from "../api/trend";
import ChartCard from "./ChartCard";
import LineChart from "./charts/LineChart";
import NumberDisplay from "./charts/NumberDisplay";
import DateRangeDropdown from "./filters/DateRangeDropdown";
import MonthRangeDropdown from "./filters/MonthRangeDropdown";

const AnalysisDashboard = () => {
  return (
    <Row gutter={[16, 16]}>
      {/* KPI */}
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
          title="Total Unique Customer"
          fetchData={() => totalUniqueCustApi().then((res) => res.data.count)}
        />
      </Col>
      <Col xs={12} md={6}>
        <NumberDisplay
          title="Total Unique Customer"
          fetchData={() => totalUniqueCustApi().then((res) => res.data.count)}
        />
      </Col>

      {/* Daily Trend */}
      <Col xs={24} md={12}>
        <ChartCard
          title="Daily Trend of Delivered Products"
          filters={(range, setRange) => (
            <DateRangeDropdown value={range} onChange={setRange} />
          )}
        >
          {({ dateRange }) => (
            <LineChart
              key="daily-trend"
              xField="time"
              yField="count"
              fetchData={(start, end) =>
                deliveredTrendApi("day", start, end).then((res) => res.data)
              }
              dateRange={dateRange}
            />
          )}
        </ChartCard>
      </Col>

      {/* Monthly Trend */}
      <Col xs={24} md={12}>
        <ChartCard
          title="Monthly Trend of Delivered Products"
          defaultRange={[dayjs().subtract(1, "year"), dayjs()]}
          filters={(range, setRange) => (
            <MonthRangeDropdown value={range} onChange={setRange} />
          )}
        >
          {({ dateRange }) => (
            <LineChart
              key="monthly-trend"
              xField="time"
              yField="count"
              fetchData={(start, end) =>
                deliveredTrendApi("month", start, end).then((res) => res.data)
              }
              dateRange={dateRange}
            />
          )}
        </ChartCard>
      </Col>
    </Row>
  );
};

export default AnalysisDashboard;
