import { useEffect, useState } from "react";
import { Button, DatePicker, Drawer, Form, Input, message, Select, Space, Table } from "antd";
import topbar from "topbar";

import { model, api } from "@/service";
import { Link } from "react-router-dom";


export function List() {
    document.title = '员工列表';
    const columns = [
        {
            title: '姓名',
            dataIndex: 'name'
        }, {
            title: '性别',
            dataIndex: 'gender',
            render: (v: number) => {
                let m = new Map([
                    [1, '男'],
                    [2, '女'],
                ]);
                return m.has(v) ? m.get(v) : '-';
            }
        }, {
            title: '生日',
            dataIndex: 'birthday'
        }, {
            title: '职位',
            dataIndex: 'position'
        }, {
            title: '职级',
            dataIndex: 'grade'
        }, {
            title: '入职时间',
            dataIndex: 'joinTime'
        }, {
            title: '状态',
            dataIndex: 'status',
            render: (v: number) => {
                let m = new Map([
                    [0, '待入职'],
                    [1, '试用期'],
                    [2, '已转正'],
                    [3, '已离职'],
                ]);
                return m.has(v) ? m.get(v) : '-';
            }
        }, {
            title: '联系方式',
            dataIndex: 'contact',
            render: (v: null, r: model.Emp) => {
                return <>
                    <div>{r.tel}</div>
                    {r.address}
                </>
            }
        }, {
            title: '操作',
            dataIndex: 'action',
            render: (v: null, r: model.Emp) => {
                return <Space size={16}>
                    <a className="iconfont" title="编辑" onClick={() => disEmp(r)}>&#xe640;</a>
                </Space>
            }
        }
    ];
    const [req, setReq] = useState(new model.Request());
    const [queryForm] = Form.useForm();
    const [emp, setEmp] = useState({
        list: new Array<model.Emp>(),
        total: 0,
        selectedRowKeys: new Array<number>(),
        dialog: {
            open: false,
            title: '',
        }
    });

    const getEmp = async () => {
        topbar.show();
        let data = {
            ...req,
        };
        if (!!queryForm.getFieldValue('val')) {
            data.queryBy = [];
            data.queryBy.push({ field: queryForm.getFieldValue('col'), value: queryForm.getFieldValue('val') });
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
        topbar.hide();
    }

    const [empForm] = Form.useForm();

    const disEmp = (item?: model.Emp) => {
        emp.dialog.open = !emp.dialog.open;
        empForm.resetFields();
        if (!!item) {
            emp.dialog.title = '编辑成员';
            empForm.setFieldsValue(item);
        } else {
            emp.dialog.title = '新增成员';
        }
        setEmp({ ...emp });
    }

    const onEmp = async () => {
        let valid = await empForm.validateFields();
        if (!valid) {
            return;
        }
        let data = {
            ...empForm.getFieldsValue(),
        };
        let res: model.Response;
        if (data.id > 0) {
            res = await api.Structure.updateEmp(data);
        } else {
            res = await api.Structure.createEmp(data);
        }
        if (res.code == 1000) {
            disEmp();
            getEmp();
        } else {
            message.error(res.desc);
        }
    }

    const selection = {
        selectedRowKeys: emp.selectedRowKeys,
        onChange: (selectedRowKeys: React.Key[]) => {
            emp.selectedRowKeys = selectedRowKeys as number[];
            setEmp({ ...emp });
        }
    };

    useEffect(() => {
        getEmp();
    }, [req.current]);

    return <>
        <div className="box-head">
            <h1>员工列表</h1>
            <Form form={queryForm} onFinish={getEmp} layout="inline" initialValues={{ col: 'ename' }}>
                <Form.Item name="col">
                    <Select options={[{ label: '姓名', value: 'ename' }]} />
                </Form.Item>
                <Form.Item name="val">
                    <Input allowClear />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit">搜索</Button>
                </Form.Item>
                <Form.Item>
                    <Space size={16}>
                        <Button>办理入职</Button>
                        <Button>转正申请</Button>
                        <Button>办理离职</Button>
                        <Button type="default"><Link to="invite">邀请加入</Link></Button>
                        <Button type="default" onClick={() => disEmp()}>添加成员</Button>
                    </Space>
                </Form.Item>
            </Form>
            <Drawer
                title={emp.dialog.title}
                open={emp.dialog.open}
                footer={<Space>
                    <Button onClick={() => disEmp()}>取消</Button>
                    <Button type="primary" onClick={onEmp}>确定</Button>
                </Space>}
                onClose={() => disEmp()}
            >
                <Form form={empForm} layout="vertical">
                    <Form.Item name="name" label="姓名" rules={[{ required: true, message: '请填写姓名!' }]}>
                        <Input />
                    </Form.Item>
                    {empForm.getFieldValue('id') == 0 && <Form.Item name="mobile" label="手机" rules={[{ required: true, message: '请填写手机号!' }]}>
                        <Input />
                    </Form.Item>}
                    <Form.Item name="number" label="工号">
                        <Input />
                    </Form.Item>
                    <Form.Item name="position" label="职位">
                        <Input />
                    </Form.Item>
                    <Form.Item name="grade" label="职级">
                        <Input />
                    </Form.Item>
                    {/* <Form.Item name="joinTime" label="入职时间">
                        <DatePicker />
                    </Form.Item> */}
                    <Form.Item name="tel" label="工作电话">
                        <Input />
                    </Form.Item>
                    <Form.Item name="email" label="工作邮箱">
                        <Input />
                    </Form.Item>
                    <Form.Item name="id" hidden>
                        <Input />
                    </Form.Item>
                    <Form.Item name="address" label="办公地点">
                        <Input />
                    </Form.Item>
                </Form>
            </Drawer>
        </div>
        <div className="box-body">
            <Table
                rowKey={(record) => record.id}
                columns={columns}
                dataSource={emp.list}
                pagination={{ current: req.current, pageSize: req.pageSize, total: emp.total }}
                onChange={(e) => setReq({ ...e })}
                rowSelection={selection}
            />
        </div>
    </>
}