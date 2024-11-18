import React from 'react';
import { FormItemProps } from './types';
import StringField from './fields/StringField';
import NumberField from './fields/NumberField';
import BooleanField from './fields/BooleanField';
import ArrayField from './fields/ArrayField';
import ObjectField from './fields/ObjectField';
import { resolveSchemaRef } from './utils';
import { JSONSchema7 } from 'json-schema';

const FormItem: React.FC<FormItemProps> = (props) => {
  const { schema } = props;

  // Add a type guard to ensure schema is a valid JSONSchema7 object
  if (!schema || typeof schema !== 'object') {
    return null;
  }

  // Resolve any schema references
  const resolvedSchema = schema.$ref
    ? resolveSchemaRef(schema, (schema.definitions || {}) as Record<string, JSONSchema7>)
    : schema;

  switch (resolvedSchema.type) {
    case 'string':
      return <StringField {...props} schema={resolvedSchema} enum={resolvedSchema.enum as string[]} />;
    case 'number':
    case 'integer':
      return (
        <NumberField
          {...props}
          schema={resolvedSchema}
          minimum={resolvedSchema.minimum}
          maximum={resolvedSchema.maximum}
        />
      );
    case 'boolean':
      return <BooleanField {...props} schema={resolvedSchema} />;
    case 'array':
      if (resolvedSchema.items) {
        // If items is an array, use the first schema as a template
        const itemSchema = Array.isArray(resolvedSchema.items)
          ? resolvedSchema.items[0]
          : resolvedSchema.items;
        return <ArrayField {...props} schema={resolvedSchema} itemSchema={itemSchema} />;
      }
      return null;
    case 'object':
      return (
        <ObjectField
          {...props}
          schema={resolvedSchema}
          required={resolvedSchema.required}
        />
      );
    default:
      return null;
  }
};

export default FormItem;
