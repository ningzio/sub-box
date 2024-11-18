import React from 'react';
import { Form, Input, Select, Tooltip } from 'antd';
import { InfoCircleOutlined } from '@ant-design/icons';
import { StringFieldProps } from '../types';

const StringField: React.FC<StringFieldProps> = ({
  schema,
  value,
  onChange,
  path,
  enum: enumOptions,
}) => {
  const fieldName = path[path.length - 1];
  const isLongField = !enumOptions && (!schema.maxLength || schema.maxLength > 50);
  // Get the first example as string
  const placeholder = Array.isArray(schema.examples) && schema.examples.length > 0
    ? String(schema.examples[0])
    : undefined;

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

  if (enumOptions) {
    return (
      <Form.Item
        label={label}
        className="form-item-horizontal"
      >
        <Select
          value={value}
          onChange={onChange}
          options={enumOptions.map(option => ({
            label: option,
            value: option,
          }))}
        />
      </Form.Item>
    );
  }

  return (
    <Form.Item
      label={label}
      className={isLongField ? undefined : 'form-item-horizontal'}
    >
      <Input
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder={placeholder}
      />
    </Form.Item>
  );
};

export default StringField;
