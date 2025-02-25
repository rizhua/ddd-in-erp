import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import topbar from "topbar";
import { Flex, Badge, Button, Select, Text, TextField } from '@radix-ui/themes';
import { toast } from 'react-toastify';
import { useForm } from "react-hook-form";

import { model, api } from "@/service";
import { Dialog } from '@/component/common';
import dayjs from "dayjs";
import { EmpStatus, BadgeColor } from "@/constant";


export function List() {
    document.title = '员工列表';
    const [req, setReq] = useState(new model.Request());
    const queryForm = useForm();
    const [emp, setEmp] = useState({
        list: new Array<model.Emp>(),
        total: 0,
        selectedRowKeys: new Array<number>(),
        dialog: {
            open: false,
            title: '',
        }
    });
    const empForm = useForm({ mode: 'onChange' });

    const getEmp = async () => {
        topbar.show();
        const {col, val} = queryForm.getValues();
        let data = {
            ...req,
        };
        if (!!col && !!val) {
            data.queryBy = [];
            data.queryBy.push({ field: col, value: val });
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

    const disEmp = (item?: model.Emp) => {
        if (!!item) {
            emp.dialog.title = '编辑成员';
            empForm.reset(item);
        } else {
            empForm.reset();
            emp.dialog.title = '新增成员';
        }
        emp.dialog.open = true;
        setEmp({ ...emp });
    }

    const onEmp = async () => {
        let data:any = empForm.getValues();
        data.joinTime = dayjs(data.joinTime);
        let res: model.Response;
        if (data.id > 0) {
            res = await api.Structure.updateEmp(data);
        } else {
            res = await api.Structure.createEmp(data);
        }
        if (res.code == 1000) {
            emp.dialog.open = false;
            setEmp({...emp});
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

    // 定义颜色枚举
    enum BadgeColor {
        Red = 'red',
        Green = 'green',
        Blue = 'blue',
        Yellow = 'yellow',
        Purple = 'purple'
    }

    // 生成随机颜色的函数
    function getBadgeColor(i: number): BadgeColor {
        const colorValues = Object.values(BadgeColor) as BadgeColor[];
        return colorValues[i];
    }

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
                <form className="form-inline" onSubmit={queryForm.handleSubmit(getEmp)}>
                    <div className="form-item">
                        <select {...queryForm.register('col')}>
                            <option value="name">姓名</option>
                            <option value="number">工号</option>
                        </select>
                    </div>
                    <div className="form-item">
                        <input type="text" {...queryForm.register('val')} />
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
            description="Make changes to your profile."
            placement="right"
            // maskClosable={true}
            onOk={empForm.handleSubmit(onEmp)}
            onClose={() => setEmp({...emp, dialog: {...emp.dialog, open: false}})}
        >
            <form className="form-vertical">
                <input type="number" name="id" hidden />
                <div className="form-item">
                    <label htmlFor="">姓名</label>
                    <div className="form-control">
                        <input type="text" {...empForm.register('name', { required: true })} />
                        {empForm.formState.errors.name && <div className="error">请输入姓名</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">工号</label>
                    <div className="form-control">
                        <input type="text" {...empForm.register('number', { required: true })} />
                        {empForm.formState.errors.number && <div className="error">请输入工号</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">职位</label>
                    <div className="form-control">
                        <input type="text" {...empForm.register('position', { required: true })} />
                        {empForm.formState.errors.position && <div className="error">请输入职位</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">职级</label>
                    <div className="form-control">
                        <input type="text" {...empForm.register('grade', { required: true })} />
                        {empForm.formState.errors.grade && <div className="error">请输入职级</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">入职时间</label>
                    <div className="form-control">
                        <input {...empForm.register('joinTime', { required: true })} type="date" />
                        {empForm.formState.errors.joinTime && <div className="error">请输入入职时间</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">工作电话</label>
                    <div className="form-control">
                        <input type="text" {...empForm.register('tel')} />
                        {empForm.formState.errors.tel && <div className="error">请输入工作电话</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">工作邮箱</label>
                    <div className="form-control">
                        <input type="text" {...empForm.register('email')} />
                        {empForm.formState.errors.email && <div className="error">请输入工作邮箱</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">办公地点</label>
                    <div className="form-control">
                        <input type="text" {...empForm.register('address')} />
                        {empForm.formState.errors.address && <div className="error">请输入办公地点</div>}
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="">办公地点</label>
                    <div className="form-control">
                        <input type="text" {...empForm.register('address')} />
                        {empForm.formState.errors.address && <div className="error">请输入办公地点</div>}
                    </div>
                </div>
            </form>
        </Dialog>
        <div className="box-body">
            <table className="table">
                <thead>
                    <tr>
                        <th>姓名</th>
                        <th>性别</th>
                        <th>生日</th>
                        <th>职位</th>
                        <th>职级</th>
                        <th>入职时间</th>
                        <th>联系方式</th>
                        <th>状态</th>
                        <th align="center">操作</th>
                    </tr>
                </thead>
                <tbody>
                    {emp.list.map((m) => <tr key={m.id}>
                        <td>{m.name}</td>
                        <td>{m.gender}</td>
                        <td>{m.birthday}</td>
                        <td>{m.position}</td>
                        <td>{m.grade}</td>
                        <td>{m.joinTime}</td>
                        <td>{m.mobile}</td>
                        <td>
                            <Badge color={getBadgeColor(m.status)}>{EmpStatus[m.status]}</Badge>
                        </td>
                        <td align="center">
                            <Flex gap="16px">
                                <a className="iconfont" title="编辑" onClick={() => disEmp(m)}>&#xe640;</a>
                            </Flex>
                        </td>
                    </tr>)}
                </tbody>
            </table>
        </div>
    </>
}