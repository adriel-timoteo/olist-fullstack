import { Col, Row } from "antd";
import dayjs, { Dayjs } from "dayjs";
import {
  repeatPurchaseRateApi,
  topCitiesApi,
  totalUniqueCustApi,
} from "../api/customer";
import {
  averageOrderValue,
  deliveredStatusApi,
  deliveredTrendApi,
  onTimeDeliveryRateApi,
  totalRevenue,
} from "../api/order";
import { topCategoriesApi } from "../api/product";
import ChartCard from "./ChartCard";
import BarChart from "./charts/BarChart";
import LineChart from "./charts/LineChart";
import NumberDisplay from "./charts/NumberDisplay";
import DateRangeDropdown from "./filters/DateRangeDropdown";
import LimitDropdown from "./filters/LimitDropdown";
import MonthRangeDropdown from "./filters/MonthRangeDropdown";
import StackedColumnChart from "./charts/StackedColumnChart";

const AnalysisDashboard = () => {
  return (
    <Row gutter={[16, 16]}>
      {/* KPI */}
      <Col xs={12} sm={6} md={4}>
        <NumberDisplay
          title="Total Unique Customer"
          fetchData={() => totalUniqueCustApi().then((res) => res.data.count)}
        />
      </Col>
      <Col xs={12} sm={6} md={4}>
        <NumberDisplay
          title="Repeat Purchase Rate"
          fetchData={() =>
            repeatPurchaseRateApi().then((res) => res.data.rate * 100)
          }
          suffix="%"
          precision={2}
        />
      </Col>
      <Col xs={12} sm={6} md={4}>
        <NumberDisplay
          title="On Time Delivery Rate"
          fetchData={() =>
            onTimeDeliveryRateApi().then((res) => res.data.rate * 100)
          }
          suffix="%"
          precision={2}
        />
      </Col>
      <Col xs={12} sm={6} md={4}>
        <NumberDisplay
          title="Total Revenue"
          fetchData={() => totalRevenue().then((res) => res.data.count)}
          prefix="R$"
        />
      </Col>
      <Col xs={12} sm={6} md={4}>
        <NumberDisplay
          title="Average Order Value"
          fetchData={() => averageOrderValue().then((res) => res.data.count)}
          prefix="R$"
          precision={2}
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
      <Col xs={24} md={12}>
        <ChartCard<{ dateRange: [Dayjs, Dayjs] }>
          title="Daily Status of Purchased Products"
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
            <StackedColumnChart
              key="purchased-status"
              xField="time"
              yField="count"
              fetchData={(start, end) =>
                deliveredStatusApi(start, end).then((res) => res.data)
              }
              dateRange={filters.dateRange}
              colorField={"status"}
            />
          )}
        </ChartCard>
      </Col>
    </Row>
  );
};

export default AnalysisDashboard;
