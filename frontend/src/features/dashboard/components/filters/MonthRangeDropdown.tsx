import { Select } from "antd";
import dayjs, { Dayjs } from "dayjs";

export type MonthRangeOption = "1y" | "2y" | "5y" | "all-time";

interface MonthRangeDropdownProps {
  value: [Dayjs, Dayjs];
  onChange: (range: [Dayjs, Dayjs]) => void;
}

const MonthRangeDropdown = ({ value, onChange }: MonthRangeDropdownProps) => {
  const getCurrentOption = (): MonthRangeOption => {
    const today = dayjs();
    const start = value[0];

    if (start.isSame(today.clone().subtract(1, "year"), "month")) return "1y";
    if (start.isSame(today.clone().subtract(2, "year"), "month")) return "2y";
    if (start.isSame(today.clone().subtract(5, "year"), "month")) return "5y";
    if (start.isSame(today.clone().subtract(10, "year"), "month"))
      return "all-time";
    return "1y";
  };

  const handleSelect = (option: MonthRangeOption) => {
    const end = dayjs().endOf("day");
    let start = dayjs();

    switch (option) {
      case "1y":
        start = end.subtract(1, "year").startOf("day");
        break;
      case "2y":
        start = end.subtract(2, "year").startOf("day");
        break;
      case "5y":
        start = end.subtract(5, "year").startOf("day");
        break;
      case "all-time":
        start = end.subtract(10, "year").startOf("day");
        break;
    }

    onChange([start, end]);
  };

  return (
    <Select
      value={getCurrentOption()}
      style={{ width: 150 }}
      onChange={(val) => handleSelect(val as MonthRangeOption)}
    >
      <Select.Option value="1y">Last Year</Select.Option>
      <Select.Option value="2y">Last 2 Years</Select.Option>
      <Select.Option value="5y">Last 5 Years</Select.Option>
      <Select.Option value="all-time">All Time</Select.Option>
    </Select>
  );
};

export default MonthRangeDropdown;
