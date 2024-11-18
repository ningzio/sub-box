import { FieldTemplateProps, FormContextType, RJSFSchema } from '@rjsf/utils';
import { Tooltip } from 'antd';
import { InfoCircleOutlined } from '@ant-design/icons';

const SimpleField = ({
  label,
  required,
  description,
  errors,
  rawErrors = [],
  children,
}: FieldTemplateProps<any, RJSFSchema, FormContextType>) => {
  // 检查 description 是否为有效的非空字符串
  const hasDescription = typeof description === 'string' && description !== '';

  return (
    <div className="simple-field-row">
      <div className="field-label">
        {label}
        {required && <span className="required-mark">*</span>}
        {hasDescription && (
          <Tooltip title={description}>
            <InfoCircleOutlined className="field-info-icon" />
          </Tooltip>
        )}
      </div>
      <div className="field-input">
        {children}
        {errors}
        {rawErrors.length > 0 && (
          <div className="error-text">{rawErrors.join(', ')}</div>
        )}
      </div>
    </div>
  );
};

export default SimpleField;
