import { Select } from "antd";
import dayjs, { Dayjs } from "dayjs";

export type DateRangeOption = "7d" | "30d" | "90d" | "1y";

interface DateRangeDropdownProps {
  value: [Dayjs, Dayjs];
  onChange: (range: [Dayjs, Dayjs]) => void;
}

const DateRangeDropdown = ({ value, onChange }: DateRangeDropdownProps) => {
  const getCurrentOption = (): DateRangeOption => {
    const today = dayjs();
    const start = value[0];

    if (start.isSame(today.clone().subtract(7, "day"), "day")) return "7d";
    if (start.isSame(today.clone().subtract(30, "day"), "day")) return "30d";
    if (start.isSame(today.clone().subtract(90, "day"), "day")) return "90d";
    if (start.isSame(today.clone().subtract(1, "year"), "day")) return "1y";
    return "7d";
  };

  const handleSelect = (option: DateRangeOption) => {
    const end = dayjs();
    let start = dayjs();

    switch (option) {
      case "7d":
        start = end.subtract(7, "day");
        break;
      case "30d":
        start = end.subtract(30, "day");
        break;
      case "90d":
        start = end.subtract(90, "day");
        break;
      case "1y":
        start = end.subtract(1, "year");
        break;
    }

    onChange([start, end]);
  };

  return (
    <Select
      value={getCurrentOption()}
      style={{ width: 150 }}
      onChange={(val) => handleSelect(val as DateRangeOption)}
    >
      <Select.Option value="7d">Last 7 Days</Select.Option>
      <Select.Option value="30d">Last 30 Days</Select.Option>
      <Select.Option value="90d">Last 90 Days</Select.Option>
      <Select.Option value="1y">Last 1 Year</Select.Option>
    </Select>
  );
};

export default DateRangeDropdown;
