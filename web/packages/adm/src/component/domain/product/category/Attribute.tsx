import { Button, Space, Table } from "antd";

import { api, model } from "@/service";
import { useEffect, useState } from "react";

interface AttributeProps {
    category: model.Category;
}

export function Attribute(props: AttributeProps) {
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
    const [category, setCategory] = useState(new model.Category());
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
        let res = await api.Product.findAttribute(data);
        if (res.code == 1000) {
            attribute.list = res.data.list;
        } else {
            attribute.list = [];
        }
        setAttribute({ ...attribute });
    }

    useEffect(() => {
        Object.assign(category, props.category);
        setCategory({ ...category });
        getAttribute();
    }, [props.category]);


    return <>
        <div className="box-head">
            <div>
                <label>类目：</label>
                <span>{category.name}</span>
            </div>
            <Space>
                <Button type="primary">新增属性</Button>
            </Space>
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