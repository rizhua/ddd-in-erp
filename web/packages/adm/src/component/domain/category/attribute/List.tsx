import { Button, Form, message, Modal, Space, Table } from "antd";

import { api, model } from "@/service";
import { useEffect, useState } from "react";

interface AttributeProps {
    category: model.Category;
}

export function List(props: AttributeProps) {
    const columns = [
        {
            title: '属性项',
            dataIndex: 'label',
        }, {
            title: '属性值',
            dataIndex: 'value',
            render: (f: string[]) => f.join(',')
        }, {
            title: '属性类型',
            dataIndex: 'type',
            render: () => <>

            </>
        }, {
            title: '操作',
            dataIndex: 'action',
            width: '110px',
            render: (text: undefined, record: model.Attribute) => {
                return <Space size={24}>
                    <a className="iconfont" title="编辑">&#xe640;</a>
                    <a className="iconfont" title="删除">&#xe640;</a>
                </Space>
            }
        },
    ];
    const [attribute, setAttribute] = useState({
        list: new Array<model.Attribute>(),
        dialog: {
            open: false,
            title: '',
        }
    });

    const getAttribute = async () => {
        let data: model.Request = {
            current: 1,
            pageSize: 1000,
        };
        let res = await api.Category.findAttribute(data);
        if (res.code == 1000) {
            attribute.list = res.data.list;
        } else {
            attribute.list = [];
        }
        setAttribute({ ...attribute });
    }

    useEffect(() => {
        getAttribute();
    }, [props.category]);

    const [attributeForm] = Form.useForm();

    const disAttribute = async (item?: model.Attribute) => {
        attribute.dialog.open = !attribute.dialog.open;
        attributeForm.resetFields();
        if (!!item) {
            attribute.dialog.title = '编辑属性';
            attributeForm.setFieldsValue(item);
        } else {
            attribute.dialog.title = '新增属性';
        }
        setAttribute({ ...attribute });
    }

    const onAttribute = async () => {
        let data = attributeForm.getFieldsValue();
        let res: model.Response;
        if (data.id > 0) {
            res = await api.Category.updateAttribute(data);
        } else {
            res = await api.Category.createAttribute(data);
        }
        if (res.code == 1000) {
            attribute.dialog.open = false;
            getAttribute();
        } else {
            message.error(res.desc);
        }
        setAttribute({ ...attribute });
    }


    return <>
        <div className="box-head">
            <div>
                <label>类目：</label>
                <span>{props.category.name}</span>
            </div>
            <Space>
                <Button type="primary" onClick={() => disAttribute()}>新增属性</Button>
            </Space>
            <Modal
                title={attribute.dialog.title}
                open={attribute.dialog.open}
                onCancel={() => attribute.dialog.open = false}
                onOk={onAttribute}
            >
                <Form
                    labelCol={{ span: 4 }}
                    wrapperCol={{ span: 20 }}
                    form={attributeForm}
                >
                    <Form.Item label="属性名称" name="label">
                        <input />
                    </Form.Item>
                    <Form.Item label="属性值" name="value">
                        <input />
                    </Form.Item>
                </Form>
            </Modal>
        </div>
        <div className="box-body">
            <Table
                columns={columns}
                dataSource={attribute.list}
                rowKey="id"
                pagination={false}
            />
        </div>
    </>
}