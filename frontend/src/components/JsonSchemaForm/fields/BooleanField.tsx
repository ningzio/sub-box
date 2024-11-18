import React from 'react';
import { Form, Switch, Tooltip } from 'antd';
import { InfoCircleOutlined } from '@ant-design/icons';
import { BooleanFieldProps } from '../types';

const BooleanField: React.FC<BooleanFieldProps> = ({
  schema,
  value,
  onChange,
  path,
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
      valuePropName="checked"
      className="form-item-horizontal"
    >
      <Switch
        checked={value}
        onChange={onChange}
      />
    </Form.Item>
  );
};

export default BooleanField;
