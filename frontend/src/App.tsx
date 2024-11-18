import { useState, useEffect } from 'react';
import { Card, message, Button, Space } from 'antd';
import { JSONSchema7 } from 'json-schema';
import Form from '@rjsf/antd';
import validator from '@rjsf/validator-ajv8';
import { RJSFSchema } from '@rjsf/utils';
import {
  Panel,
  PanelGroup,
  PanelResizeHandle,
} from 'react-resizable-panels';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
import { vscDarkPlus } from 'react-syntax-highlighter/dist/esm/styles/prism';
import { ExportOutlined, ReloadOutlined } from '@ant-design/icons';
import { exportJson } from './utils/exportJson';
import SimpleField from './components/JsonEditor/SimpleField';
import ComplexField from './components/JsonEditor/ComplexField';
import './styles/JsonEditor.css';

const App = () => {
  const [schema, setSchema] = useState<JSONSchema7>({});
  const [formData, setFormData] = useState<any>({});
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    fetchSchema();
  }, []);

  const fetchSchema = async () => {
    try {
      setLoading(true);
      const response = await fetch('http://localhost:8080/api/schema');
      const data = await response.json();
      setSchema(data);
    } catch (error) {
      message.error('Failed to fetch schema');
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const handleExport = () => {
    exportJson(formData, 'config.json');
  };

  return (
    <div className="json-editor-container">
      <PanelGroup direction="horizontal" style={{ height: '100%' }}>
        <Panel defaultSize={50} minSize={30}>
          <Card
            title="Sing-box Configuration Editor"
            className="editor-panel"
            bordered={false}
            extra={
              <Space>
                <Button
                  icon={<ReloadOutlined />}
                  onClick={fetchSchema}
                  loading={loading}
                >
                  Reload Schema
                </Button>
                <Button
                  icon={<ExportOutlined />}
                  onClick={handleExport}
                  disabled={!formData || Object.keys(formData).length === 0}
                >
                  Export
                </Button>
              </Space>
            }
          >
            <Form
              schema={schema as RJSFSchema}
              validator={validator}
              onChange={({ formData }) => setFormData(formData)}
              formData={formData}
              templates={{
                ObjectFieldTemplate: ComplexField,
                FieldTemplate: SimpleField,
              }}
            />
          </Card>
        </Panel>

        <PanelResizeHandle className="resize-handle" />

        <Panel defaultSize={50} minSize={30}>
          <Card
            title="JSON Preview"
            className="preview-panel"
            bordered={false}
          >
            <SyntaxHighlighter
              language="json"
              style={vscDarkPlus}
              customStyle={{ margin: 0, borderRadius: 4 }}
            >
              {JSON.stringify(formData, null, 2)}
            </SyntaxHighlighter>
          </Card>
        </Panel>
      </PanelGroup>
    </div>
  );
};

export default App;
