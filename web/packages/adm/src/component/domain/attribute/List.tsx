import { Button, message, Popconfirm, Space, Table } from "antd";
import { useEffect, useState } from "react";

import { api, model } from "@/service";
import { Dialog } from "./Dialog";
import topbar from "topbar";

export function List() {
    const columns = [
        {
            title: '属性项',
            dataIndex: 'label',
        }, {
            title: '属性值',
            dataIndex: 'value',
            render: (text: string[]) => text.join(','),
        }, {
            title: '操作',
            dataIndex: 'action',
            width: 100,
            render: (text: any, record: model.Attribute) => {
                return <Space size={24}>
                    <a className="iconfont" onClick={() => disAttribute(record)} title="编辑">&#xe640;</a>
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

    const [req, setReq] = useState(new model.Request({ current: 1, pageSize: 10 }));
    const [attribute, setAttribute] = useState({
        info: new model.Attribute(),
        list: new Array<model.Attribute>(),
        total: 0,
        dialog: {
            title: '',
            open: false,
        },
        selectedRowKeys: new Array<number>(),
    });

    const getAttribute = async () => {
        topbar.show();
        let data = {
            ...req,
        }
        let res = await api.Product.findAttribute(data);
        if (res.code == 1000) {
            attribute.list = res.data.list;
            attribute.total = res.data.total;
        } else {
            attribute.list = [];
            attribute.total = 0;
        }
        setAttribute({ ...attribute });
        topbar.hide();
    };

    useEffect(() => {
        getAttribute();
    }, [req.current]);

    const disAttribute = (item?: model.Attribute) => {
        attribute.dialog.open = !attribute.dialog.open;
        Object.assign(attribute.info, new model.Attribute());
        if (!!item) {
            attribute.dialog.title = '编辑属性';
            Object.assign(attribute.info, item);
        } else {
            attribute.dialog.title = '新增属性';
        }
        setAttribute({ ...attribute });
        if (!attribute.dialog.open) {
            getAttribute();
        }
    }

    const delAttribute = async (item: model.Attribute) => {
        let data = {
            id: [item.id],
        }
        let res = await api.Product.deleteAttribute(data);
        if (res.code == 1000) {
            getAttribute();
        } else {
            message.error(res.desc);
        }
    }

    return <>
        <div className="box-head">
            <h1>商品属性</h1>
            <Space>
                <Button type="primary" onClick={() => disAttribute()}>新增属性</Button>
            </Space>
            {attribute.dialog.open && <Dialog attribute={attribute.info} mask onClose={() => disAttribute()}></Dialog>}
        </div>
        <div className="box-body">
            <Table
                columns={columns}
                dataSource={attribute.list}
                rowKey="id"
                pagination={{ current: req.current, pageSize: req.pageSize, total: attribute.total }}
                onChange={(p) => setReq({ ...p })}
            />
        </div>
    </>
}