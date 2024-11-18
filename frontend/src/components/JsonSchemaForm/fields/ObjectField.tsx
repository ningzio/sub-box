import React from 'react';
import { Form, Card } from 'antd';
import { ObjectFieldProps } from '../types';
import FormItem from '../FormItem';
import { JSONSchema7 } from 'json-schema';

const ObjectField: React.FC<ObjectFieldProps> = ({
  schema,
  value = {},
  onChange,
  path,
  required = [],
}) => {
  const fieldName = path[path.length - 1];
  const properties = schema.properties || {};

  const handlePropertyChange = (propertyName: string, propertyValue: any) => {
    onChange({
      ...value,
      [propertyName]: propertyValue,
    });
  };

  return (
    <Form.Item
      label={schema.title || fieldName}
      help={schema.description}
    >
      <Card size="small">
        {Object.entries(properties).map(([propertyName, propertySchema]) => {
          // Create a new schema object for the property
          const propertySchemaObj: JSONSchema7 = {
            ...(propertySchema as JSONSchema7),
            // Only set required array if this property is required
            required: required.includes(propertyName) ? [propertyName] : undefined,
          };

          return (
            <FormItem
              key={propertyName}
              schema={propertySchemaObj}
              value={value?.[propertyName]}
              onChange={(newValue) => handlePropertyChange(propertyName, newValue)}
              path={[...path, propertyName]}
            />
          );
        })}
      </Card>
    </Form.Item>
  );
};

export default ObjectField;
