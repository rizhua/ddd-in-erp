import { useEffect, useState } from "react";
import { Button, Form, Input, Pagination, Select, Space, Table } from "antd";
import styled from "styled-components";

import { model, api } from "@/service";

const Container = styled.div`
    .box-head > .text {
        display: flex;
        gap: 0;
        
        label {
            color: #666;
        }
    }

    .box-body > .cell {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 16px 0;
    }
`

export interface EmpProps {
    dept: model.Dept;
}

export function Emp(props: EmpProps) {
    const columns = [
        {
            title: '姓名',
            dataIndex: 'name',
        }, {
            title: '工号',
            dataIndex: 'number',
            render: (v: string) => {
                return v == '' ? '-' : v;
            }
        }, {
            title: '职位',
            dataIndex: 'position',
            render: (v: string) => {
                return v == '' ? '-' : v;
            }
        }, {
            title: '电话',
            dataIndex: 'tel',
            render: (v: string) => {
                return v == '' ? '-' : v;
            }
        }, {
            title: '邮箱',
            dataIndex: 'email',
            render: (v: string) => {
                return v == '' ? '-' : v;
            }
        }, {
            title: '入职日期',
            dataIndex: 'joinTime',
            render: (v: string) => {
                return v == '' ? '-' : v;
            }
        }, {
            title: '状态',
            dataIndex: 'status',
            render: (v: number) => {
                return model.EmpStatus[v];
            }
        },
    ];

    const [req, setReq] = useState(new model.Request({ current: 1, pageSize: 10 }));
    const [emp, setEmp] = useState({
        info: new model.Emp(),
        list: new Array<model.Emp>(),
        total: 0,
    });
    const [queryBy] = Form.useForm();

    const getEmp = async () => {
        let data = {
            ...req,
            queryBy: new Array<model.QueryBy>(),
        };
        if (!!queryBy.getFieldValue('value')) {
            data.queryBy.push({ 'field': queryBy.getFieldValue('field'), 'value': queryBy.getFieldValue('value') });
        }
        if (props.dept.id > 0) {
            data.queryBy.push({ 'field': 'deptId', 'value': props.dept.id });
        }
        if (data.queryBy.length == 0) {
            data.queryBy = undefined!;
        }
        let res = await api.Structure.findEmp(data);
        if (res.code == 1000) {
            emp.list = res.data.list;
            emp.total = res.data.total;
        } else {
            emp.list = [];
            emp.total = 0;
        }
        setEmp({ ...emp });
    }

    const rowSelection = {

    }

    useEffect(() => {
        getEmp();
    }, [props.dept]);

    return <Container>
        <div className="box-head">
        {props.dept.id > 0 ? <div className="text"><label>部门:</label><span>{props.dept.name}</span></div> : <div></div>}
            <Form layout="inline" form={queryBy} onFinish={getEmp} initialValues={{ field: 'number' }}>
                <Form.Item name="field">
                    <Select options={[
                        { label: '工号', value: 'number' },
                        { label: '手机', value: 'mobile' },
                    ]} />
                </Form.Item>
                <Form.Item name="value">
                    <Input allowClear />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">搜索</Button>
                </Form.Item>
                <Space size={16}>
                    <Button type="default">调整部门</Button>
                </Space>
            </Form>
        </div>
        <div className="box-body">
            <Table
                rowKey={(record) => record.id}
                columns={columns}
                dataSource={emp.list}
                pagination={false}
                rowSelection={rowSelection}
            />
            <div className="cell">
                <Space size={16}>
                    <Button type="default">调整部门</Button>
                </Space>
                <Pagination current={req.current} pageSize={req.pageSize} total={emp.total} onChange={(e) => setReq({ ...req, current: e })} />
            </div>
        </div>
    </Container>
}