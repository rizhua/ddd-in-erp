import { useEffect, useState } from "react";
import ReactDOM from "react-dom";
import styled from "styled-components";
import { message } from "antd";

import { api, model } from "@/service";

const Container= styled.div`
    display: grid;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    left: 0;
    z-index: 100;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.3);

    .block {
        padding: 0 1rem;

        a {
            cursor: pointer;
        }
    }

    .block-search {
        display: flex;
        position: relative;
        height: 34px;

        input {
            width: 100%;
            border: 1px solid #ccc;
            padding: 4px 8px;
            height: 34px;
            box-sizing: border-box;
            border-radius: 4px;
        }

        .suffix {
            position: absolute;
            right: 8px;
            height: 34px;
            display: flex;
            align-items: center;
        }
    }

    .box-list {
        display: grid;
        gap: 8px;
        margin: 1rem 0;

        a {
            color: #666;
        }
    }

    .box-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 8px;
    }

    .block-head {
        display: flex;
        justify-content: space-between;
        align-items: center;
        height: 34px;
    }
`

const Box = styled.div`
    background-color: #fff;
    border-radius: 4px;
    min-width: 500px;

    .box-head {
        display: flex;
        justify-content: space-between;
        padding: 0 1rem;
        line-height: 50px;

        a {
            cursor: pointer;
            color: #666;
        }
    }

    .box-body {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        min-height: 30vh;
        max-height: 80vh;
        overflow-y: auto;
    }

    .box-foot {
        display: flex;
        justify-content: flex-end;
        gap: 1rem;
        padding: 1rem;
    }

    .btn {
        cursor: pointer;
    }
`

type PropsType = {
    open: boolean;
    role: model.Role;
    onClose: () => void;
}

export function EmpTransfer(props: PropsType) {
    const [empModel, setEmpModel] = useState({
        list: new Array<model.Emp>(),
        selected: new Array<model.Emp>(),
        keyword: "",
    });

    const getEmp = async (e?: any) => {
        let data = {
            current: 1,
            pageSize: 10,
            queryBy: [
                {
                    field: "keyword",
                    value: empModel.keyword,
                },
            ],
        }
        let res = await api.Structure.findEmp(data);
        if (res.code == 1000) {
            empModel.list = res.data.list;
        } else {
            empModel.list = [];
        }
        setEmpModel({ ...empModel });
    }

    const optEmp = (e: boolean, emp: model.Emp) => {
        let s = new Set(empModel.selected);
        if (e) {
            s.add(emp);
        } else {
            s.delete(emp);
        }
        empModel.selected = Array.from(s);
        setEmpModel({ ...empModel });
    }

    const clear = (id?: number) => {
        if (!id) {
            empModel.selected = [];
        } else {
            empModel.selected = empModel.selected.filter(e => e.id !== id);
        }
        setEmpModel({ ...empModel });
    }

    const addEmp = async () => {
        if (empModel.selected.length == 0) {
            message.error("请选择员工");
            return;
        }
        let userArr = new Array<number>();
        empModel.selected.forEach((v) => {
            userArr.push(v.userId);
        });
        let data = {
            roleId: props.role.id,
            userId: userArr,
        };
        let res = await api.Role.addUser(data);
        if (res.code == 1000) {
            message.success("添加成功");
            clear();
            props.onClose();
        } else {
            message.error(res.desc);
        }
    }

    const close = () => {
        clear();
        props.onClose();
    }

    useEffect(() => {
        if (props.open) {
            getEmp();
        }
    }, [props.open]);

    return props.open ? ReactDOM.createPortal(<Container>
        <Box>
            <div className="box-head">
                <div className="text">添加成员</div>
                <div className="close" onClick={close}>
                    <a className="iconfont">&#xe6b3;</a>
                </div>
            </div>
            <div className="box-body">
                <form className="block" onSubmit={(e) => {e.preventDefault(); getEmp()}}>
                    <div className="block-search">
                        <input value={empModel.keyword} type="text" onChange={(e) => setEmpModel({ ...empModel, keyword: e.target.value })} />
                        <a className="suffix iconfont" onClick={() => getEmp()}>&#xe63f;</a>
                    </div>
                    <ul className="box-list">
                        {empModel.list.map((v, k) => (
                            <li className="box-item" key={k}>
                                <div>
                                    <input checked={empModel.selected.includes(v)} type="checkbox" onChange={(e) => optEmp(e.target.checked, v)} />
                                    <span>{v.name}</span>
                                </div>
                            </li>
                        ))}
                    </ul>
                </form>
                <div className="block">
                    {empModel.selected.length > 0 && <div className="block-head">
                        <label htmlFor="">已选：{empModel.selected.length}名员工</label>
                        <a onClick={() => clear()}>清空</a>
                    </div>}
                    <ul className="box-list">
                        {empModel.selected.map((v, k) => {
                            return <li className="box-item" key={k}>
                                <div>
                                    {v.name}
                                </div>
                                <a className="iconfont" onClick={() => clear(v.id)}>&#xe6b3;</a>
                            </li>
                        })}
                    </ul>
                </div>
            </div>
            <div className="box-foot">
                <button className="btn btn-default" onClick={() => close()}>取消</button>
                <button className="btn btn-primary" onClick={() => addEmp()}>确定</button>
            </div>
        </Box>
    </Container>, document.body) : null;
}