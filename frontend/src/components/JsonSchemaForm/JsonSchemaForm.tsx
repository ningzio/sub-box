import React from 'react';
import { JsonSchemaFormProps } from './types';
import { Form } from 'antd';
import FormItem from './FormItem';
import { resolveSchemaRef } from './utils';
import { JSONSchema7 } from 'json-schema';

const JsonSchemaForm: React.FC<JsonSchemaFormProps> = ({
  schema,
  value,
  onChange,
}) => {
  // Extract definitions from the schema
  const definitions = (schema.definitions || {}) as Record<string, JSONSchema7>;
  
  // Resolve the root schema reference
  const resolvedSchema = schema.$ref ? resolveSchemaRef(schema, definitions) : schema;

  const handleChange = (newValue: any) => {
    onChange?.(newValue);
  };

  return (
    <Form layout="vertical">
      <FormItem
        schema={resolvedSchema}
        value={value}
        onChange={handleChange}
        path={[]}
      />
    </Form>
  );
};

export default JsonSchemaForm;
