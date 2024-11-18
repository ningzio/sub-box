import React from 'react';
import { Form, InputNumber, Tooltip } from 'antd';
import { InfoCircleOutlined } from '@ant-design/icons';
import { NumberFieldProps } from '../types';

const NumberField: React.FC<NumberFieldProps> = ({
  schema,
  value,
  onChange,
  path,
  minimum,
  maximum,
}) => {
  const fieldName = path[path.length - 1];

  const label = (
    <span>
      {schema.title || fieldName}
      {schema.description && (
        <Tooltip title={schema.description}>
          <InfoCircleOutlined style={{ marginLeft: 4, color: '#999', fontSize: 14 }} />
        </Tooltip>
      )}
    </span>
  );

  return (
    <Form.Item
      label={label}
      className="form-item-horizontal"
    >
      <InputNumber
        value={value}
        onChange={onChange}
        min={minimum}
        max={maximum}
        placeholder={schema.examples ? String(schema.examples) : undefined}
        style={{ width: '100%' }}
      />
    </Form.Item>
  );
};

export default NumberField;
