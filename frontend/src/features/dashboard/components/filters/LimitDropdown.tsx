import { Select } from "antd";

export type LimitOption = 5 | 10 | 20 | 50;

interface LimitDropdownProps {
  value: number;
  onChange: (limit: LimitOption) => void;
}

const LimitDropdown = ({ value, onChange }: LimitDropdownProps) => {
  const options: LimitOption[] = [5, 10, 20, 50];

  const isValid = options.includes(value as LimitOption);
  const fallback = 10;

  return (
    <Select
      value={isValid ? value : fallback}
      style={{ width: 150 }}
      onChange={(val) => onChange(val as LimitOption)}
    >
      {options.map((limit) => (
        <Select.Option key={limit} value={limit}>
          Top {limit}
        </Select.Option>
      ))}
    </Select>
  );
};

export default LimitDropdown;
