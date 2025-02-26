import { useEffect, useState } from "react";

import topbar from "topbar";
import { Button, Drawer, Form, Input, InputNumber, message, Popconfirm, Select, Space, Table } from "antd";

import { model, api } from "@/service";
import { Node } from './Node';
import { License } from "./License";

export function List() {
    const columns = [{
        title: '名称',
        dataIndex: 'name',
    }, {
        title: '期限(月)',
        dataIndex: 'term',
    }, {
        title: '工位',
        dataIndex: 'quota',
    }, {
        title: '价格',
        dataIndex: 'price',
    }, {
        title: '创建时间',
        dataIndex: 'createAt',
    }, {
        title: '操作',
        dataIndex: 'action',
        width: 150,
        render: (v: null, r: model.Bundle) => <Space size={24}>
            <a className="iconfont" onClick={() => disNode(r)} title="权限">&#xe680;</a>
            <a className="iconfont" onClick={() => disLicense(r)} title="许可">&#xe6c2;</a>
            <a className="iconfont" onClick={() => disBundle(r)} title="编辑">&#xe640;</a>
            <Popconfirm
                title="确定要删除吗？"
                description="删除后不可恢复"
                onConfirm={() => delBundle(r)}
            >
                <a className="iconfont" title="删除">&#xe618;</a>
            </Popconfirm>
        </Space>
    }];
    const [req, setReq] = useState(new model.Request({current:1, pageSize: 10}));
    const [bundle, setBundle] = useState({
        info: new model.Bundle(),
        list: new Array<model.Bundle>(),
        total: 0,
        nodeDrawer: false,
        licenseDrawer: false,
        dialog: {
            open: false,
            title: '',
        }
    });
    const [searchForm] = Form.useForm();

    const getBundle = async () => {
        topbar.show();
        let data: model.Request = {
            ...req,
        };
        if (!!searchForm.getFieldValue('value')) {
            data.queryBy = [{
                ...searchForm.getFieldsValue(),
            }];
        }
        let res = await api.Bundle.find(data);
        if (res.code == 1000) {
            bundle.list = res.data.list;
            bundle.total = res.data.total;
        } else {
            bundle.list = [];
            bundle.total = 0;
        }
        setBundle({ ...bundle });
        topbar.hide();
    }

    const [bundleForm] = Form.useForm();

    const disBundle = (item?: model.Bundle) => {
        bundle.dialog.open = !bundle.dialog.open;
        bundleForm.resetFields();
        if (!!item) {
            bundle.dialog.title = '编辑套餐';
            bundleForm.setFieldsValue(item);
        } else {
            bundle.dialog.title = '新增套餐';
        }
        setBundle({ ...bundle });
    }

    const onBundle = async () => {
        let valid = await bundleForm.validateFields();
        if (!valid) {
            return;
        }
        let data = {
            ...bundleForm.getFieldsValue(),
        };
        let res: model.Response;
        if (data.id > 0) {
            res = await api.Bundle.update(data);
        } else {
            res = await api.Bundle.create(data);
        }
        if (res.code == 1000) {
            getBundle();
            disBundle();
        } else {
            message.error(res.desc);
        }
    }

    const delBundle = async (item: model.Bundle) => {
        let data = {
            id: [item.id],
        }
        let res = await api.Bundle.delete(data);
        if (res.code == 1000) {
            message.success('删除成功');
            getBundle();
        } else {
            message.error(res.desc);
        }
    }

    useEffect(() => {
        getBundle();
    }, []);

    const disNode = (item?: model.Bundle) => {
        bundle.nodeDrawer = !bundle.nodeDrawer;
        if (!!item) {
            bundle.info = item;
        }
        if (!bundle.nodeDrawer) {
            getBundle();
        }
        setBundle({ ...bundle });
    }

    const disLicense = (item?: model.Bundle) => {
        bundle.licenseDrawer = !bundle.licenseDrawer;
        if (!!item) {
            bundle.info = item;
        }
        setBundle({ ...bundle });
    }

    return <>
        <div className="box-head">
            <h1>资源套餐</h1>
            <form className="form-inline">
                <div className="form-item">
                    <select>
                        <option value="name">名称</option>
                    </select>
                </div>
                <div className="form-item">
                    <input type="text" />
                </div>
                <div className="form-item">
                    <Button type="primary">搜索</Button>
                </div>
            </form>
            <Drawer
                open={bundle.dialog.open}
                title={bundle.dialog.title}
                onClose={() => disBundle()}
                footer={<Space>
                    <Button onClick={() => disBundle()}>取消</Button>
                    <Button type="primary" onClick={() => onBundle()}>确定</Button>
                </Space>}
            >
                <Form
                    form={bundleForm}
                    layout="vertical"
                >
                    <Form.Item name="name" label="名称" rules={[{ required: true }]}>
                        <Input />
                    </Form.Item>
                    <Form.Item name="term" label="期限(月)">
                        <InputNumber precision={0} style={{ width: "100%" }} />
                    </Form.Item>
                    <Form.Item name="quota" label="配额(工位)">
                        <InputNumber precision={0} style={{ width: "100%" }} />
                    </Form.Item>
                    <Form.Item name="price" label="价格">
                        <InputNumber
                            formatter={(value) => `￥ ${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                            style={{ width: "100%" }}
                        />
                    </Form.Item>
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
                dataSource={bundle.list}
                pagination={{ current: req.current, pageSize: req.pageSize, total: bundle.total }}
                onChange={(e) => setReq({ ...e })}
            />
            <Node open={bundle.nodeDrawer} bundle={bundle.info} onClose={() => disNode()} />
            <License open={bundle.licenseDrawer} bundle={bundle.info} onClose={() => disLicense()} />
        </div>
    </>
}
