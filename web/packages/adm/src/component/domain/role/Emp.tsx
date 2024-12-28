import { useEffect, useState } from "react";
import { Button, Form, Input, message, Space, Select, Table } from "antd";
import { useForm } from "antd/lib/form/Form";

import { model, api } from "@/service";
import './role.less';
import { EmpTransfer } from "@/component/common/EmpTransfer";


interface EmpProps {
    role: model.Role,
    parent: model.Role,
}

export function Emp(props: EmpProps) {
    const columns = [
        {
            title: '姓名',
            dataIndex: 'name',
        }, {
            title: '工号',
            dataIndex: 'number',
        }, {
            title: '职位',
            dataIndex: 'position',
        }, {
            title: '手机',
            dataIndex: 'mobile',
        }, {
            title: '邮箱',
            dataIndex: 'email',
        }, {
            title: '入职日期',
            dataIndex: 'joinTime',
            render: (text: string) => {
                return text || '-';
            }
        }, 
    ];
    let [req, setReq] = useState({
        ...new model.Request({ current: 1, pageSize: 10 }),
        queryBy: new Array(),
    });
    const [queryForm] = useForm();
    const [emp, setEmp] = useState({
        list: new Array<model.Emp>(),
        total: 0,
        loading: false,
        open: false,
        selected: new Array<model.Emp>(),
    });

    const getEmp = async () => {
        emp.loading = true;
        setEmp({ ...emp });
        let data = {
            ...req,
            roleID: props.role.id || 1
        }
        if (!!queryForm.getFieldValue('value')) {
            data.queryBy.push({ field: queryForm.getFieldValue('field'), value: queryForm.getFieldValue('value') });
        }
        let res = await api.Role.findUser(data);
        if (res.code == 1000) {
            emp.list = res.data.list;
            emp.total = res.data.total;
        } else {
            emp.list = [];
            emp.total = 0;
        }
        emp.loading = false;
        setEmp({ ...emp });
    }

    useEffect(() => {
        getEmp();
    }, [props.role.id]);

    const delEmp = async () => {
        let data = {
            roleId: props.role.id,
            userId: new Array<number>(),
        };
        if (emp.selected.length == 0) {
            message.error("请选择要移除的数据");
            return;
        }
        emp.selected.forEach(item => {
            data.userId.push(item.userId);
        });
        let res = await api.Role.removeUser(data);
        if (res.code == 1000) {
            message.success("移除成功");
            getEmp();
        } else {
            message.error(res.desc);
        }
    }

    const close = () => {
        setEmp({...emp, open: false});
    }

    const rowSelection = {
        onChange: (selectedRowKeys: React.Key[], selectedRows: model.Emp[]) => {
            emp.selected = selectedRows.filter(item => selectedRowKeys.some(i => i === item.id));
            setEmp({...emp});
        },
    };

    useEffect(() => {
        if (!emp.open) {
            getEmp();
        }
    }, [emp.open]);

    return <>
        <div className="box-head">
            <div className="text">
                <span><label>分组:</label>{props.parent.name}</span>
                <span><label>角色:</label>{props.role.name}({emp.total}人)</span>
            </div>
            <Form layout="inline" form={queryForm} initialValues={{ field: 'number' }}>
                <Form.Item name="field">
                    <Select options={[{ label: '工号', value: 'number' }, { label: '姓名', value: 'name' }]} />
                </Form.Item>
                <Form.Item name="value">
                    <Input allowClear />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" onClick={getEmp} loading={emp.loading}>搜索</Button>
                </Form.Item>
                <Space size={16}>
                    <Button type="primary" onClick={() => setEmp({...emp, open: true})}>添加成员</Button>
                    <Button type="primary" onClick={() => delEmp()}>移除成员</Button>
                </Space>
            </Form>
            <EmpTransfer open={emp.open} role={props.role} onClose={() => close()} />
        </div>
        <div className="box-body">
            <Table
                rowKey={(record) => record.id}
                columns={columns}
                dataSource={emp.list}
                pagination={{ current: req.current, pageSize: req.pageSize, total: emp.total }}
                onChange={(e) => setReq({...req, ...e})}
                rowSelection={{...rowSelection}}
            />
        </div>
    </>
}