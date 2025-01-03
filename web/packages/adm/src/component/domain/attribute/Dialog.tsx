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
        max-height: 60vh;
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
        gap: 8px;
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
    attribute: model.Attribute;
    mask?: boolean;
    onClose?: () => void;
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
            res = await api.Product.updateAttribute(data);
        } else {
            res = await api.Product.createAttribute(data);
        }
        if (res.code == 1000) {
            onClose();
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

    const inputValue = (e: any) => {
        let name = e.target.name;
        let value = e.target.value;
        Object.assign(attribute, { [name]: value });
        if (name == 'label') {
            validate(value);
        }
        setAttribute({ ...attribute });
    }

    const checkValue = (e: any) => {
        let name = e.target.name;
        let checked = e.target.checked;
        Object.assign(attribute, { [name]: checked });
        setAttribute({ ...attribute });
    }

    const onChange = (e: any) => {
        let value = e.target.value;
        let name = Number(e.target.name);
        attribute.value[name] = value;
        setAttribute({ ...attribute });
    }

    const onClose = () => {
        props.onClose && props.onClose();
    }

    return <Container>
        {props.mask && <div className="mask" onClick={() => onClose()}></div>}
        <Box>
            <div className="box-head">
                <h2>新增属性</h2>
                <a className="iconfont" onClick={() => onClose()}>&#xe6b3;</a>
            </div>
            <form className="box-body">
                <div className="form-item">
                    <label><i className="required">*</i>属性名称:</label>
                    <div className="form-control">
                        <input name="label" defaultValue={attribute.label} onChange={inputValue} />
                        <div className="error" style={{ visibility: error.label.valid ? 'visible' : 'hidden' }}>{error.label.message}</div>
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="multi">属性值多选:</label>
                    <div className="form-control">
                        <div className="input-group">
                            <input type="checkbox" name="multi" defaultChecked={attribute.multi} onChange={checkValue} />启用
                        </div>
                        <div className="tip">启用后，支持选择多个属性值，比如一杯奶茶添加布丁、珍珠等多种配料</div>
                    </div>
                </div>
                <div className="form-item">
                    <label htmlFor="required">属性值必选:</label>
                    <div className="form-control">
                        <div className="input-group">
                            <input type="checkbox" name="required" defaultChecked={attribute.required} onChange={checkValue} />启用
                        </div>
                        <div className="tip">启用后，商品发布时，必须填写该属性</div>
                    </div>
                </div>
                <div className="form-item">
                    <label>属性值:</label>
                    <div className="box-list">
                        {attribute.value.map((v: string, k: number) => <div className="box-item" key={k}>
                            <input name={k.toString()} defaultValue={v} onChange={onChange} />
                            <a className="iconfont" onClick={() => delValue()}>&#xe618;</a>
                        </div>)}
                        <div className="box-item">
                            <button type="button" onClick={() => addValue()}>
                                <i className="iconfont">&#xe678;</i>
                            </button>
                        </div>
                    </div>
                </div>
                <div className="form-foot">
                    <button className="btn btn-default" onClick={() => onClose()}>取消</button>
                    <button className="btn btn-primary" type="button" onClick={() => onAttribute()}>确定</button>
                </div>
            </form>
        </Box>
    </Container>;
}