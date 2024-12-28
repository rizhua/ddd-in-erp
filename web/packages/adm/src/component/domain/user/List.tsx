import { useEffect, useState } from "react";
import { Badge, Button, Form, Input, Select, Table } from "antd";
import styled from "styled-components";

import { model, api } from "@/service";
import topbar from "topbar";


const columns = [{
    title: 'ID',
    dataIndex: 'id',
},{
    title: '昵称',
    dataIndex: 'nickname',
}, {
    title: '手机号',
    dataIndex: 'mobile',
}, {
    title: '邮箱',
    dataIndex: 'email',
    render: (v: string) => {
        return v || '-';
    }
}, {
    title: '生日',
    dataIndex: 'birthday',
    render: (v: string) => {
        return v == '0001-01-01' ? '-' : v;
    }
}, {
    title: '性别',
    dataIndex: 'gender',
    render: (v: number) => {
        let m = new Map([
            [1, '男'],
            [2, '女']
        ]);
        return m.has(v) ? m.get(v) : '-';
    }
}, {
    title: '状态',
    dataIndex: 'status',
    render: (v: number) => {
        let badge;
        switch (v) {
            case 1:
                badge = <Badge color="green" text="正常" />
                break;
            default:
                badge = <Badge color="gold" text="冻结" />
        }
        return badge;
    }
}, {
    title: '最后登录时间',
    dataIndex: 'lastTime',
}, {
    title: '注册时间',
    dataIndex: 'createAt',
}];

export function List() {
    const [req, setReq] = useState(new model.Request({current:1, pageSize: 10}));
    const [user, setUser] = useState({
        list: new Array<model.User>(),
        total: 0,
    });
    const [searchForm] = Form.useForm();

    const getUser = async () => {
        topbar.show();
        let data: model.Request = {
            ...req,
        };
        if (!!searchForm.getFieldValue('value')) {
            data.queryBy = [{
                ...searchForm.getFieldsValue()
            }];
        }
        let res = await api.User.find(data);
        if (res.code == 1000) {
            user.list = res.data.list;
            user.total = res.data.total;
        } else {
            user.list = [];
            user.total = 0;
        }
        setUser({...user});
        topbar.hide();
    }

    useEffect(() => {
        getUser();
    }, [req.current]);

    return <>
        <div className="box-head">
            <h1>用户列表</h1>
            <Form form={searchForm} onFinish={getUser} layout="inline" initialValues={{ field: 'mobile' }}>
                <Form.Item name="field">
                    <Select options={[{ label: '手机号', value: 'mobile' }, { label: '昵称', value: 'nickname' }]} />
                </Form.Item>
                <Form.Item name="value">
                    <Input allowClear />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">搜索</Button>
                </Form.Item>
            </Form>
        </div>
        <div className="box-body">
            <Table
                rowKey={(record) => record.id}
                columns={columns}
                dataSource={user.list}
                pagination={{ current: req.current, pageSize: req.pageSize, total: user.total }}
                onChange={(e) => setReq({...e})}
            />
        </div>
    </>
}