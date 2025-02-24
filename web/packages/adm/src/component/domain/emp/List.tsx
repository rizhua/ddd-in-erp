import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import topbar from "topbar";
import { Flex, Table, Button, Select, Text, TextField } from '@radix-ui/themes';
import { toast } from 'react-toastify';
import { useForm } from "react-hook-form";

import { model, api } from "@/service";
import { Dialog, Field } from '@/component/common';
import dayjs, { Dayjs } from "dayjs";


export function List() {
    document.title = '员工列表';
    const [req, setReq] = useState(new model.Request());
    // const [queryForm] = Form.useForm();
    const [emp, setEmp] = useState({
        list: new Array<model.Emp>(),
        total: 0,
        selectedRowKeys: new Array<number>(),
        dialog: {
            open: false,
            title: '',
        }
    });
    const empForm = useForm({mode: 'onChange'});

    const getEmp = async () => {
        topbar.show();
        let data = {
            ...req,
        };
        // if (!!queryForm.getFieldValue('val')) {
        //     data.queryBy = [];
        //     data.queryBy.push({ field: queryForm.getFieldValue('col'), value: queryForm.getFieldValue('val') });
        // }
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

    const disEmp = (item?: model.Emp) => {
        emp.dialog.open = !emp.dialog.open;
        empForm.reset();
        if (!!item) {
            emp.dialog.title = '编辑成员';
            empForm.reset(item);
        } else {
            emp.dialog.title = '新增成员';
        }
        setEmp({ ...emp });
    }

    const onEmp = async (data: any) => {
        data.joinTime = dayjs(data.joinTime);
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
            toast.error(res.desc);
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
            <Flex gap="16px">
                <Button>办理入职</Button>
                <Button>转正申请</Button>
                <Button>办理离职</Button>
                <Button type="button">
                    <Link to="invite">邀请加入</Link>
                </Button>
                <Button type="button" onClick={() => disEmp()}>添加成员</Button>
                <form className="form-inline">
                    <div className="form-item">
                        <Select.Root defaultValue="name">
                            <Select.Trigger />
                            <Select.Content>
                                <Select.Item value="name">姓名</Select.Item>
                            </Select.Content>
                        </Select.Root>
                    </div>
                    <div className="form-item">
                        <TextField.Root></TextField.Root>
                    </div>
                    <div className="form-item">
                        <Button type="submit">搜索</Button>
                    </div>
                </form>
            </Flex>
        </div>
        <Dialog
            open={emp.dialog.open}
            title={emp.dialog.title}
            placement="right"
            onOk={() => empForm.handleSubmit(onEmp)()}
            onClose={() => setEmp({ ...emp, dialog: { ...emp.dialog, open: false } })}
        >
            <form className="form-vertical">
                <input type="number" name="id" hidden />
                <div className="form-item">
                    <label htmlFor="">姓名</label>
                    <div className="form-control">
                        <TextField.Root {...empForm.register('name', { required: true })} />
                        {empForm.formState.errors.name && <div className="error">请输入姓名</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">工号</label>
                    <div className="form-control">
                        <TextField.Root {...empForm.register('number', {required: true})} />
                        {empForm.formState.errors.number && <div className="error">请输入工号</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">职位</label>
                    <div className="form-control">
                        <TextField.Root {...empForm.register('position', {required: true})} />
                        {empForm.formState.errors.position && <div className="error">请输入职位</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">职级</label>
                    <div className="form-control">
                        <TextField.Root {...empForm.register('grade', {required: true})} />
                        {empForm.formState.errors.grade && <div className="error">请输入职级</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">入职时间</label>
                    <div className="form-control">
                        <Field {...empForm.register('joinTime', {required: true})} type="date" />
                        {empForm.formState.errors.joinTime && <div className="error">请输入入职时间</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">工作电话</label>
                    <div className="form-control">
                        <TextField.Root {...empForm.register('tel')} />
                        {empForm.formState.errors.tel && <div className="error">请输入工作电话</div>} 
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">工作邮箱</label>
                    <div className="form-control">
                        <TextField.Root {...empForm.register('email')} />
                        {empForm.formState.errors.email && <div className="error">请输入工作邮箱</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">办公地点</label>
                    <div className="form-control">
                        <TextField.Root {...empForm.register('address')} />
                        {empForm.formState.errors.address && <div className="error">请输入办公地点</div>}
                    </div>
                </div>
            </form>
        </Dialog>

        <div className="box-body">
            <Table.Root>
                <Table.Header>
                    <Table.Row>
                        <Table.ColumnHeaderCell>姓名</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>性别</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>生日</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>职位</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>职级</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>入职时间</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>联系方式</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>状态</Table.ColumnHeaderCell>
                        <Table.ColumnHeaderCell>操作</Table.ColumnHeaderCell>
                    </Table.Row>
                </Table.Header>
                {emp.list.map((m) => <Table.Row>
                    <Table.Cell>{m.name}</Table.Cell>
                    <Table.Cell>{m.gender}</Table.Cell>
                    <Table.Cell>{m.birthday}</Table.Cell>
                    <Table.Cell>{m.position}</Table.Cell>
                    <Table.Cell>{m.grade}</Table.Cell>
                    <Table.Cell>{m.joinTime}</Table.Cell>
                    <Table.Cell>{m.mobile}</Table.Cell>
                    <Table.Cell>{m.status}</Table.Cell>
                    <Table.Cell>
                        <Flex gap="16px">
                            <a className="iconfont" title="编辑" onClick={() => disEmp(m)}>&#xe640;</a>
                        </Flex>
                    </Table.Cell>
                </Table.Row>)}
            </Table.Root>
        </div>
    </>
}