import { Collapse } from 'antd';
import { CaretRightOutlined } from '@ant-design/icons';
import { ObjectFieldTemplateProps } from '@rjsf/utils';

const ComplexField = ({
  title,
  description,
  properties,
  idSchema,
}: ObjectFieldTemplateProps<any>) => {
  const isRoot = !idSchema.$id.includes('_');
  const isEmpty = properties.length === 0;
  const hasDescription = typeof description === 'string' && description !== '';

  if (isRoot) {
    return (
      <div className="root-section">
        {properties.map((prop) => prop.content)}
      </div>
    );
  }

  if (isEmpty) {
    return null;
  }

  return (
    <div className="field-container">
      <div className="simple-field-row">
        <div className="field-label">
          {title}
          {hasDescription && (
            <span className="field-description">{description}</span>
          )}
        </div>
        <div className="field-input">
          <Collapse
            className="field-collapse"
            defaultActiveKey={[]}
            ghost={true}
            expandIcon={({ isActive }) => (
              <CaretRightOutlined rotate={isActive ? 90 : 0} style={{ fontSize: '12px' }} />
            )}
          >
            <Collapse.Panel
              key={idSchema.$id}
              header={null}
              showArrow={true}
            >
              {properties.map((prop) => (
                <div key={prop.name}>
                  {prop.content}
                </div>
              ))}
            </Collapse.Panel>
          </Collapse>
        </div>
      </div>
    </div>
  );
};

export default ComplexField;
