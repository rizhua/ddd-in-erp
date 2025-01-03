import { Button, Form, Input, message, Modal, Popconfirm, Space, Table } from "antd";

import { api, model } from "@/service";
import { useEffect, useState } from "react";
import { Dialog } from "./Dialog";

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
            render: (text: undefined, record: model.CategoryAttribute) => {
                return <Space size={24}>
                    <a className="iconfont" title="编辑" onClick={() => disAttribute(record)}>&#xe640;</a>
                    <Popconfirm 
                        title="确定要删除吗？"
                        onConfirm={() => delAttribute(record)}
                    >
                        <a className="iconfont" title="删除">&#xe618;</a>
                    </Popconfirm>
                </Space>
            }
        },
    ];
    const [attribute, setAttribute] = useState({
        info: new model.CategoryAttribute(),
        list: new Array<model.CategoryAttribute>(),
        dialog: {
            open: false,
            title: '',
        }
    });

    // 获取属性
    const getAttribute = async () => {
        let data: model.Request = {
            current: 1,
            pageSize: 1000,
            queryBy: [
                {
                    field: 'categoryId',
                    value: props.category.id,
                }
            ]
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

    // 新增、编辑属性
    const disAttribute = async (item?: model.CategoryAttribute) => {
        attribute.dialog.open = !attribute.dialog.open;
        Object.assign(attribute.info, new model.CategoryAttribute());
        if (!!item) {
            attribute.dialog.title = '编辑属性';
            Object.assign(attribute.info, item);
        } else {
            attribute.dialog.title = '新增属性';
        }
        attribute.info.categoryId = props.category.id;
        setAttribute({ ...attribute });
    }

    const onOk = () => {
        getAttribute();
        attribute.dialog.open = false;
        setAttribute({ ...attribute });
    }

    // 删除属性
    const delAttribute = async (item: model.CategoryAttribute) => {
        let data = {
            categoryId: props.category.id,
            id: [item.id],
        }
        let res = await api.Category.deleteAttribute(data);
        if (res.code == 1000) {
            getAttribute();
        } else {
            message.error(res.desc);
        }
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
            {attribute.dialog.open && <Dialog
                title={attribute.dialog.title}
                attribute={attribute.info}
                onOk={onOk}
                onClose={() => disAttribute()}
            />}
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