import { message } from "antd";
import {  useState } from "react";
import styled from "styled-components";

import { api, model } from "@/service";

const Container = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100vw;
    height: 100vh;
    position: absolute;
    top: 0;
    left: 0;
    z-index: 999;

    .mask {
        position: absolute;
        width: 100vw;
        height: 100vh;
        background-color: rgba(0, 0, 0, .5);
        z-index: -1;
    }
`
const Box = styled.div`
    border-radius: 8px;
    background-color: #fff;
    min-width: 420px;

    .box-head {
        display: flex;
        justify-content: space-between;
        padding: 8px 16px;
        font-weight: 600;

        a{
            cursor: pointer;
        }
    }

    .box-body {
        padding: 16px;
        max-height: 50vh;
        overflow-y: auto;
    }

    .form-item {
        display: grid;
        grid-template-columns: 80px 1fr;
        gap: 16px;
        margin-bottom: 16px;

        label {
            text-align: right;
        }

        input {
            padding: 2px 8px;
            min-width: 200px;
            min-height: 34px;
            border-radius: 4px;
            border: 1px solid #ccc;
            box-sizing: border-box;

            &[type="checkbox"] {
                min-width: 16px;
                min-height: 16px;
            }
        }

        .required {
            color: #f00;
        }
    }

    .form-control {

        .error {
            color: #f00;
            font-size: 12px;
            visibility: hidden;
        }

        .tip {
            font-size: 12px;
            color: #999;
        }
    }

    .input-group {
        display: flex;
        align-items: center;
        gap: 4px;
    }

    .radio-group {
        display: flex;
        align-items: center;
        gap: 8px;

        label {
            cursor: pointer;
            display: flex;
            align-items: center;
            gap: 4px;
        }

        input {
            min-width: 16px;
            min-height: 16px;
        }
    }

    .form-foot {
        display: flex;
        justify-content: center;
        gap: 16px;
        padding: 16px 0;
    }

    .box-list {
    display: grid;
    gap: 16px;
    padding: 16px;
    background-color: #f0f0f0;
}

.box-item {
    display: grid;
    grid-template-columns: 1fr 20px;
    align-items: center;
    gap: 8px;

    &:hover {
        a {
            display: block;
        }
    }

    a {
        cursor: pointer;
        display: none;
    }

    button {
        cursor: pointer;
        border: 1px solid #ccc;
        min-height: 34px;
        border-radius: 4px;
    }
}
`

interface DialogProps {
    title?: string;
    attribute: model.CategoryAttribute;
    mask?: boolean;
    onClose?: () => void;
    onOk?: () => void;
}

export function Dialog(props: DialogProps) {
    const [attribute, setAttribute] = useState(props.attribute);
    const [error, setError] = useState({
        label: {
            valid: false,
            message: '请输入属性名称',
        }
    });

    const validate = (label: string): boolean => {
        let flag = true;
        if (label == '') {
            error.label.valid = true;
            flag = false;
        } else {
            error.label.valid = false;
        }
        setError({...error});

        return flag;
    }

    const onAttribute = async () => {
        if (!validate(attribute.label)) {
            return;
        }
        let data = {
            ...attribute,
        }
        let res: model.Response;
        if (data.id > 0) {
            res = await api.Category.updateAttribute(data);
        } else {
            res = await api.Category.createAttribute(data);
        }
        if (res.code == 1000) {
            props.onOk && props.onOk();
        } else {
            message.error(res.desc);
        }
    }

    const addValue = () => {
        attribute.value.push('');
        setAttribute({ ...attribute });
    }

    const delValue = () => {
        attribute.value.pop();
        setAttribute({ ...attribute });
    }

    const onChange = (name: string, value: string | boolean) => {
        Object.assign(attribute, { [name]: value });
        setAttribute({ ...attribute });
    }

    const onInput = (e: any) => {
        let value = e.target.value;
        let name = Number(e.target.name);
        attribute.value[name] = value;
        setAttribute({ ...attribute });
    }

    const onClose = () => {
        props.onClose && props.onClose();
    }

    return <Container>
        <div className="mask" onClick={() => props.mask && onClose()}></div>
        <Box>
            <div className="box-head">
                <h2>{props.title}</h2>
                <a className="iconfont" onClick={() => onClose()}>&#xe6b3;</a>
            </div>
            <form className="box-body">
                <div className="form-item">
                    <label><i className="required">*</i>属性名称:</label>
                    <div className="form-control">
                        <input name="label" defaultValue={attribute.label} onChange={e => onChange(e.target.name, e.target.value)} />
                        <div className="error" style={{ visibility: error.label.valid ? 'visible' : 'hidden' }}>{error.label.message}</div>
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="multi">属性类型:</label>
                    <div className="form-control">
                        <div className="radio-group">
                            <label>
                                <input type="radio" name="type" defaultChecked={attribute.type == 'SELECT'} onChange={() => onChange('type', 'SELECT')} />下拉框
                            </label>
                            <label>
                                <input type="radio" name="type" defaultChecked={attribute.type == 'INPUT'} onChange={() => onChange('type', 'INPUT')} />输入框
                            </label>
                            <label>
                                <input type="radio" name="type" defaultChecked={attribute.type == 'UPLOAD'} onChange={() => onChange('type', 'UPLOAD')} />上传框
                            </label>
                        </div>
                    </div>
                </div>
                <div className="form-item">
                    <label>是否必选:</label>
                    <div className="form-control">
                        <div className="input-group">
                            <input type="checkbox" name="required" defaultChecked={attribute.required} onChange={(e) => onChange('required', e.target.checked)} />
                            <span>启用</span>
                        </div>
                        <div className="tip">启用后，商品发布时必须选择该属性值</div>
                    </div>
                </div>
                {attribute.type == 'SELECT' && <div className="form-item">
                    <label>属性值:</label>
                    <div className="box-list">
                        {attribute.value.map((v: string, k: number) => <div className="box-item" key={k}>
                            <input name={k.toString()} defaultValue={v} onChange={e => onInput(e)} />
                            <a className="iconfont" onClick={() => delValue()}>&#xe618;</a>
                        </div>)}
                        <div className="box-item">
                            <button type="button" onClick={() => addValue()}>
                                <i className="iconfont">&#xe678;</i>
                            </button>
                        </div>
                    </div>
                </div>}
                <div className="form-foot">
                    <button className="btn btn-default" onClick={() => onClose()}>取消</button>
                    <button className="btn btn-primary" type="button" onClick={() => onAttribute()}>确定</button>
                </div>
            </form>
        </Box>
    </Container>;
}