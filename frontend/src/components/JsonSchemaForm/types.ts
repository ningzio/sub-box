import { JSONSchema7, JSONSchema7Definition } from 'json-schema';

export interface JsonSchemaFormProps {
  schema: JSONSchema7;
  value?: any;
  onChange?: (value: any) => void;
}

export interface FormItemProps {
  schema: JSONSchema7;
  value: any;
  onChange: (value: any) => void;
  path: string[];
}

export interface ObjectFieldProps extends FormItemProps {
  required?: string[];
}

export interface ArrayFieldProps extends FormItemProps {
  itemSchema: JSONSchema7Definition;
}

export interface StringFieldProps extends FormItemProps {
  enum?: string[];
}

export interface NumberFieldProps extends FormItemProps {
  minimum?: number;
  maximum?: number;
}

export interface BooleanFieldProps extends FormItemProps {}
