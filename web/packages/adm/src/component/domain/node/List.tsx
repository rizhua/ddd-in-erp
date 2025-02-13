import { useEffect, useState } from "react";
import styled from "styled-components";
import { Button, Drawer, Form, Input, message, Popconfirm, Select, Space, Switch, Table, Tooltip, TreeSelect } from "antd";
import topbar from "topbar";

import { model, api } from "@/service";


export function List() {
    document.title = '运营-节点管理';
    const columns = [
        {
            title: '名称',
            dataIndex: 'name',
        }, {
            title: '图标',
            dataIndex: 'icon',
            render: (v: string) => <i className="iconfont" dangerouslySetInnerHTML={{ __html: v }}></i>
        }, {
            title: '元数据',
            dataIndex: 'meta',
        }, {
            title: '类型',
            dataIndex: 'type',
            render: (f: number) => {
                let m = new Map([
                    [0, '应用'],
                    [1, '功能'],
                    [2, '菜单'],
                    [3, '操作'],
                    [4, '接口'],
                ]);
                return m.has(f) ? m.get(f) : '-';
            }
        }, {
            title: '状态',
            dataIndex: 'status',
            render: (v: boolean, r: model.Node) => <Switch checked={v} onChange={() => setStatus(r)} size="small" />
        }, {
            title: '排序',
            dataIndex: 'sort',
            width: '100px',
            render: (v: number, r: model.Node) => {
                let e = <span onClick={() => selCell(r)}>{v}</span>;
                if (editable.some(e => e == r.id)) {
                    e = <Input defaultValue={v} onPressEnter={(event) => setSort(event, r)} onBlur={(event) => setSort(event, r)} size="small" />;
                }
                return <Tooltip title="单击编辑">{e}</Tooltip>
            }
        }, {
            title: '操作',
            dataIndex: 'action',
            width: '110px',
            render: (text: undefined, record: model.Node) => {
                return <Space size={16}>
                    <a className="iconfont" onClick={() => disNode(record, true)} title="新增子节点">&#xe600;</a>
                    <a className="iconfont" onClick={() => disNode(record)} title="编辑">&#xe640;</a>
                    <Popconfirm
                        placement="left"
                        title="确定要删除吗？"
                        onConfirm={() => delNode(record)}
                    >
                        <a className="iconfont" href="#" title="删除">&#xe618;</a>
                    </Popconfirm>
                </Space>
            }
        },
    ];

    const [node, setNode] = useState({
        info: new model.Node(),
        list: new Array<model.Node>(),
    });
    const [queryBy] = Form.useForm();

    const [editable, setEditable] = useState<number[]>([]);
    const selCell = (item: model.Node) => {
        let s = new Set([...editable]);
        s.add(item.id);
        setEditable([...s]);
    }

    const getNode = async () => {
        topbar.show();
        let data: model.Request = {
            current: 1,
            pageSize: 1000,
        };
        if (!!queryBy.getFieldValue('field')) {
            data.queryBy = [];
            data.queryBy.push({ 'field': queryBy.getFieldValue('field'), 'value': queryBy.getFieldValue('field') });
        }
        let res = await api.Node.find(data);
        if (res.code == 1000) {
            node.list = res.data.list;
        } else {
            node.list = [];
        }
        setNode({ ...node });
        topbar.hide();
    }

    const [drawer, setDrawer] = useState({ open: false, title: '' });
    const [nodeForm] = Form.useForm();

    const disNode = (item?: model.Node, isSub?: boolean) => {
        drawer.open = !drawer.open;
        nodeForm.resetFields();
        Object.assign(node.info, new model.Node());
        if (!!item) {
            Object.assign(node.info, item);
            if (!!isSub) {
                drawer.title = '新增节点';
                let node = new model.Node();
                node.parentId = item.id;
                nodeForm.setFieldsValue(node);
            } else {
                drawer.title = '编辑节点';
                nodeForm.setFieldsValue(item);
            }
        } else {
            drawer.title = '新增节点';
        }
        setDrawer({ ...drawer });
    }

    const onNode = async () => {
        let valid = await nodeForm.validateFields().catch(e => console.log(e));
        if (!valid) {
            return;
        }
        let data = {
            ...nodeForm.getFieldsValue(),
        };
        let res: model.Response;
        if (data.id > 0) {
            res = await api.Node.update(data);
        } else {
            res = await api.Node.create(data);
        }
        if (res.code == 1000) {
            disNode();
            getNode();
        } else {
            message.error(res.desc);
        }
    }

    const delNode = async (item: model.Node) => {
        let data = {
            id: [item.id],
        };
        let res = await api.Node.delete(data);
        if (res.code == 1000) {
            message.success('删除成功');
            getNode();
        } else {
            message.error(res.desc);
        }
    }

    const setSort = (event: any, r: model.Node) => {
        let sort = Number(event.target.value);
        if (sort != r.sort) {
            let data = {
                id: r.id,
                sort: sort,
            };
            api.Node.setSort(data).then(res => {
                if (res.code == 1000) {
                    getNode();
                    message.success('修改成功');
                } else {
                    message.error(res.desc);
                }
            });
        }
        let s = new Set([...editable]);
        s.delete(r.id);
        setEditable([...s]);
    }

    const setStatus = async (r: model.Node) => {
        let data = {
            id: r.id,
            status: Number(!r.status),
        };
        let res = await api.Node.setStatus(data);
        if (res.code == 1000) {
            getNode();
        } else {
            message.error(res.desc);
        }
    }

    useEffect(() => {
        getNode();
    }, []);

    return <>
        <div className="box-head">
            <div className="text">
                <h3>节点列表</h3>
            </div>
            <Form layout="inline" form={queryBy} onFinish={getNode} initialValues={{ col: 'name' }}>
                <Form.Item name="col">
                    <Select
                        options={[{ label: '名称', value: 'name' }]}
                        value="name"
                    />
                </Form.Item>
                <Form.Item name="val">
                    <Input />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">搜索</Button>
                </Form.Item>
                <Button type="primary" onClick={() => disNode()}>新增节点</Button>
            </Form>
            {/* 新增、编辑 */}
            <Drawer
                open={drawer.open}
                onClose={() => disNode()}
                title={drawer.title}
                footer={
                    <Space>
                        <Button onClick={() => disNode()}>取消</Button>
                        <Button type="primary" onClick={onNode}>确定</Button>
                    </Space>
                }
            >
                <Form form={nodeForm} layout="vertical">
                    <Form.Item
                        name="name"
                        label="名称"
                        rules={[{ required: true, message: '请输入节点名称!' }]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item name="type" label="类型" rules={[{ required: true, message: '请选择类型!' }]}>
                        <Select options={[
                            { label: '应用', value: 0 },
                            { label: '功能', value: 1 },
                            { label: '菜单', value: 2 },
                            { label: '操作', value: 3 },
                            { label: '接口', value: 4 },
                        ]} />
                    </Form.Item>
                    <Form.Item name="meta" label="元数据" rules={[{ required: true, message: '请输入元数据!' }]}>
                        <Input />
                    </Form.Item>
                    <Form.Item name="icon" label="图标">
                        <Input />
                    </Form.Item>
                    {<Form.Item name="parentId" label="上级节点">
                        <TreeSelect treeData={node.list} fieldNames={{ label: 'name', value: 'id' }} showSearch allowClear />
                    </Form.Item>}
                    <Form.Item name="id" hidden>
                        <Input />
                    </Form.Item>
                </Form>
            </Drawer>
        </div>
        <div className="box-body">
            <Table
                rowKey={(record) => record.id}
                columns={columns}
                dataSource={node.list}
                pagination={false}
            />
        </div>
    </>
}