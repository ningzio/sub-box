import React from 'react';
import { Form, Button, Space, Card } from 'antd';
import { DeleteOutlined, PlusOutlined } from '@ant-design/icons';
import { ArrayFieldProps } from '../types';
import FormItem from '../FormItem';
import { JSONSchema7 } from 'json-schema';

const ArrayField: React.FC<ArrayFieldProps> = ({
  schema,
  value = [],
  onChange,
  path,
  itemSchema,
}) => {
  const fieldName = path[path.length - 1];

  const handleAdd = () => {
    onChange([...value, undefined]);
  };

  const handleRemove = (index: number) => {
    const newValue = [...value];
    newValue.splice(index, 1);
    onChange(newValue);
  };

  const handleItemChange = (index: number, itemValue: any) => {
    const newValue = [...value];
    newValue[index] = itemValue;
    onChange(newValue);
  };

  // Convert itemSchema to JSONSchema7 object
  let actualItemSchema: JSONSchema7;
  if (Array.isArray(itemSchema)) {
    actualItemSchema = itemSchema[0] as JSONSchema7;
  } else if (typeof itemSchema === 'boolean') {
    // If itemSchema is true, accept any value
    // If itemSchema is false, reject all values
    actualItemSchema = itemSchema ? {} : { not: {} };
  } else {
    actualItemSchema = itemSchema;
  }

  return (
    <Form.Item
      label={schema.title || fieldName}
      help={schema.description}
    >
      <Space direction="vertical" style={{ width: '100%' }}>
        {value.map((item: any, index: number) => (
          <Card
            key={index}
            size="small"
            extra={
              <Button
                type="text"
                icon={<DeleteOutlined />}
                onClick={() => handleRemove(index)}
                danger
              />
            }
          >
            <FormItem
              schema={actualItemSchema}
              value={item}
              onChange={(newValue) => handleItemChange(index, newValue)}
              path={[...path, index.toString()]}
            />
          </Card>
        ))}
        <Button
          type="dashed"
          onClick={handleAdd}
          icon={<PlusOutlined />}
          block
        >
          Add Item
        </Button>
      </Space>
    </Form.Item>
  );
};

export default ArrayField;
