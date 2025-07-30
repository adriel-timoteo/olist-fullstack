import { Col, Row } from "antd";
import dayjs, { Dayjs } from "dayjs";
import { totalUniqueCustApi } from "../api/count";
import { onTimeDeliveryRateApi, repeatPurchaseRateApi } from "../api/rate";
import { deliveredTrendApi } from "../api/trend";
import ChartCard from "./ChartCard";
import LineChart from "./charts/LineChart";
import NumberDisplay from "./charts/NumberDisplay";
import DateRangeDropdown from "./filters/DateRangeDropdown";
import MonthRangeDropdown from "./filters/MonthRangeDropdown";
import LimitDropdown from "./filters/LimitDropdown";
import BarChart from "./charts/BarChart";
import { topCategoriesApi, topCitiesApi } from "../api/categorical";

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
          title="On Time Delivery Rate"
          fetchData={() =>
            onTimeDeliveryRateApi().then((res) => res.data.rate * 100)
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

      {/* Daily Trend */}
      <Col xs={24} md={12}>
        <ChartCard<{ dateRange: [Dayjs, Dayjs] }>
          title="Daily Trend of Delivered Products"
          defaultFilters={{
            dateRange: [
              dayjs().subtract(7, "day").startOf("day"),
              dayjs().endOf("day"),
            ],
          }}
          filters={(filters, setFilters) => (
            <DateRangeDropdown
              value={filters.dateRange}
              onChange={(range) => setFilters({ ...filters, dateRange: range })}
            />
          )}
        >
          {({ filters }) => (
            <LineChart
              key="daily-trend"
              xField="time"
              yField="count"
              fetchData={(start, end) =>
                deliveredTrendApi("day", start, end).then((res) => res.data)
              }
              dateRange={filters.dateRange}
            />
          )}
        </ChartCard>
      </Col>

      {/* Monthly Trend */}
      <Col xs={24} md={12}>
        <ChartCard<{ dateRange: [Dayjs, Dayjs] }>
          title="Monthly Trend of Delivered Products"
          defaultFilters={{
            dateRange: [
              dayjs().subtract(7, "day").startOf("day"),
              dayjs().endOf("day"),
            ],
          }}
          filters={(filters, setFilters) => (
            <MonthRangeDropdown
              value={filters.dateRange}
              onChange={(range) => setFilters({ ...filters, dateRange: range })}
            />
          )}
        >
          {({ filters }) => (
            <LineChart
              key="monthly-trend"
              xField="time"
              yField="count"
              fetchData={(start, end) =>
                deliveredTrendApi("month", start, end).then((res) => res.data)
              }
              dateRange={filters.dateRange}
            />
          )}
        </ChartCard>
      </Col>
      <Col xs={24} md={12}>
        <ChartCard
          title="Top Cities"
          defaultFilters={{ limit: 10 }}
          filters={(filters, setFilters) => (
            <LimitDropdown
              value={filters.limit}
              onChange={(limit) => setFilters({ ...filters, limit })}
            />
          )}
        >
          {({ filters }) => (
            <BarChart
              key="top-cities"
              xField="city"
              yField="count"
              limit={filters.limit}
              fetchData={(limit) =>
                topCitiesApi(limit.toString()).then((res) => res.data)
              }
            />
          )}
        </ChartCard>
      </Col>
      <Col xs={24} md={12}>
        <ChartCard
          title="Top Categories"
          defaultFilters={{ limit: 10 }}
          filters={(filters, setFilters) => (
            <LimitDropdown
              value={filters.limit}
              onChange={(limit) => setFilters({ ...filters, limit })}
            />
          )}
        >
          {({ filters }) => (
            <BarChart
              key="top-categories"
              xField="category"
              yField="count"
              limit={filters.limit}
              fetchData={(limit) =>
                topCategoriesApi(limit.toString()).then((res) => res.data)
              }
            />
          )}
        </ChartCard>
      </Col>
    </Row>
  );
};

export default AnalysisDashboard;
