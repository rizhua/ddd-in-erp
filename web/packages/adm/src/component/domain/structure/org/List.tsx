import { useEffect, useState } from "react";

import { Badge, Button, Form, Input, Select, Space, Table } from "antd";

import { model, api } from "@/service";
import topbar from "topbar";

export function List() {
    const columns = [{
        title: '名称',
        dataIndex: 'fullName',
    }, {
        title: '联系人',
        dataIndex: 'contact',
    }, {
        title: '联系电话',
        dataIndex: 'tel',
    }, {
        title: '通讯地址',
        dataIndex: 'address',
    }, {
        title: '创建时间',
        dataIndex: 'createAt',
    }, {
        title: '状态',
        dataIndex: 'status',
        render: (text: number) => {
            let badge;
            switch (text) {
                case 1:
                    badge = <Badge color="green" text="正常" />
                    break;
                default:
                    badge = <Badge color="gold" text="冻结" />
            }
            return badge;
        }
    }, {
        title: '操作',
        dataIndex: 'action',
        width: '80px',
        render: (v: null, r: model.Org) => <Space size={16}>
            <a className="iconfont" title="免密登录">&#xe61f;</a>
        </Space>
    }];
    const [req, setReq] = useState(new model.Request());
    const [orgModel, setOrgModel] = useState({
        list: new Array<model.Org>(),
        total: 0,
    });
    const [searchForm] = Form.useForm();

    const getOrg = async () => {
        topbar.show();
        let data: model.Request = {
            ...req,
        };
        if (!!searchForm.getFieldValue('val')) {
            data.queryBy = [{
                ...searchForm.getFieldsValue()
            }];
        }
        let res = await api.Structure.list(data);
        if (res.code == 1000) {
            orgModel.list = res.data.list;
            orgModel.total = res.data.total;
        } else {
            orgModel.list = [];
            orgModel.total = 0;
        }
        setOrgModel({ ...orgModel });
        topbar.hide();
    }

    useEffect(() => {
        getOrg();
    }, [req.current]);

    return <>
        <div className="box-head">
            <h1>组织列表</h1>
            <Form
                form={searchForm}
                onFinish={getOrg}
                initialValues={{ col: 'name' }}
                layout="inline"
            >
                <Form.Item name="col">
                    <Select options={[{ label: '名称', value: 'name' }]} />
                </Form.Item>
                <Form.Item name="val">
                    <Input allowClear />
                </Form.Item>
                <Button type="primary" htmlType="submit">搜索</Button>
            </Form>
        </div>
        <div className="box-body">
            <Table
                rowKey={(record) => record.id}
                columns={columns}
                dataSource={orgModel.list}
                pagination={{ current: req.current, pageSize: req.pageSize, total: orgModel.total }}
                onChange={(e) => setReq(e)}
            />
        </div>
    </>
}