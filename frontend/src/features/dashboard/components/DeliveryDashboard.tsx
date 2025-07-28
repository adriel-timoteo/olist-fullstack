import { Col, Row } from "antd";
import { deliveredTrendApi } from "../api/trend";
import ChartCard from "./ChartCard";
import LineChart from "./charts/LineChart";
import NumberDisplay from "./charts/NumberDisplay";
import DateRangeDropdown from "./filters/DateRangeDropdown";
import MonthRangeDropdown from "./filters/MonthRangeDropdown";
import dayjs from "dayjs";

const DeliveryDashboard = () => {
  return (
    <Row gutter={[16, 16]}>
      {/* KPI */}
      <Col xs={24} md={12}>
        <NumberDisplay title="Total Delivered Products" value={200} />
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

export default DeliveryDashboard;
