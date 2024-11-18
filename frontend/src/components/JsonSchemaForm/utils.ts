import { JSONSchema7 } from 'json-schema';

export function resolveSchemaRef(schema: JSONSchema7, definitions: Record<string, JSONSchema7>): JSONSchema7 {
  if (!schema) return schema;

  // Handle $ref
  if (typeof schema.$ref === 'string') {
    // Remove the #/definitions/ prefix if it exists
    const refPath = schema.$ref.replace('#/definitions/', '');
    const resolvedSchema = definitions[refPath];
    if (!resolvedSchema) {
      console.warn(`Could not resolve schema reference: ${schema.$ref}`);
      return schema;
    }
    // Recursively resolve any nested refs
    return resolveSchemaRef(resolvedSchema, definitions);
  }

  // Handle nested objects
  if (schema.properties) {
    const resolvedProperties: Record<string, JSONSchema7> = {};
    Object.entries(schema.properties).forEach(([key, value]) => {
      resolvedProperties[key] = resolveSchemaRef(value as JSONSchema7, definitions);
    });
    return { ...schema, properties: resolvedProperties };
  }

  // Handle arrays
  if (schema.items) {
    if (Array.isArray(schema.items)) {
      return {
        ...schema,
        items: schema.items.map(item => resolveSchemaRef(item as JSONSchema7, definitions))
      };
    } else {
      return {
        ...schema,
        items: resolveSchemaRef(schema.items as JSONSchema7, definitions)
      };
    }
  }

  return schema;
}
