import styled from "styled-components";
import { ReactNode, useEffect, useState } from "react";
import { Button, Text } from "@radix-ui/themes";

import { api, model } from "@/service";
import { Upload } from "./Upload";

const Container = styled.form`

    h2 {
        margin: 0;
    }

    .box-head {
        background-color: #f0f0f0;
        padding: 0 16px;
        border-radius: 4px;
    }

    ul.box-body {
        display: grid;
        grid-template-columns: repeat(4, minmax(100px, 240px));
        gap: 16px;

        .list-item {
            border: 1px solid #e0e0e0;
            border-radius: 4px;
            padding: 16px;
            text-align: center;
        }
    }

    .form-item {
        display: flex;
        gap: 8px;
        margin: 24px 0;
    }

    .form-label {
        min-width: 100px;
        text-align: right;

        &::after {
            content: ':';
            margin-right: 0.5rem;
        }
    }

    .form-control {
        display: grid;
        gap: 8px;
        flex: 1;

        input {
            max-width: 360px;
            border-radius: 4px;
            padding: 0 8px;
            border: 1px solid #e0e0e0;
        }

        input[type="text"] {
            width: 100%;
            height: 30px;
        }

        input[type="number"] {
            width: 80px;
            height: 24px;
        }
    }

    .checkbox-group {
        display: flex;
        gap: 16px;
    }

    .checkbox-item {
        display: flex;
        align-items: center;
        gap: 4px;
    }

    .radio-group {
        display: flex;
        flex-direction: column;
        gap: 8px;

        .row {
            display: flex;
            align-items: center;
            gap: 4px;
        }
    }

    .radio-item {
        display: flex;
        align-items: center;
        gap: 4px;
    }

    .box-foot {
        display: flex;
        justify-content: center;
        gap: 16px;
        padding: 16px 0;
        border-top: 1px solid #e0e0e0;
    }
`;
const Sell = styled.div`
    border: 1px solid #e0e0e0;
    padding: 8px;

    .sell-item {
        position: relative;

        &:hover > .sell-item-delete {
            display: flex;
        }

        .sell-item-delete {
            cursor: pointer;
            position: absolute;
            right: 10px;
            top: 10px;
            display: none;
            align-items: center;
            justify-content: center;
            width: 20px;
            height: 20px;
            border-radius: 50%;
            background-color: rgba(0, 0, 0, 0.3);
            color: #fff;
            font-size: 14px;
            font-weight: bold;
        }
    }

    .sell-item-name {
        display: grid;
        grid-template-columns: 50px 1fr 40px;
        gap: 8px;
        background-color: #f5f5f5;
        padding: 8px;

        input[type="text"] {
            width: 100px;
        }

        .column {
            display: flex;
            flex-direction: column;
            gap: 8px;
        }
    }

    .sell-item-value {
        display: grid;
        grid-template-columns: 50px 1fr;
        gap: 8px;
        padding: 8px;

        .input-box {
            position: relative;

            &:hover > a {
                display: flex;
            }

            input {
                width: 100px;
            }

            a {
                cursor: pointer;
                position: absolute;
                right: -10px;
                top: -10px;
                display: none;
                align-items: center;
                justify-content: center;
                width: 20px;
                height: 20px;
                border-radius: 50%;
                background-color: rgba(0, 0, 0, 0.3);
                color: #fff;
                font-size: 12px;
                font-weight: bold;
            }
        }

        .row {
            display: flex;
            align-items: center;
            gap: 8px;
        }
    }

    .sell-foot {
        display: flex;
        padding: 8px;
        background-color: #f5f5f5;

        .btn-default {
            cursor: pointer;
            padding: 4px 8px;
            border-radius: 4px;
            background-color: #f0f0f0;
            color: #333;
        }
    }
`;

export function Publish() {
    const [spu, setSpu] = useState({
        info: {
            images: new Array<string>(),
            freight: {
                deliveryMode: 0,
                shippingCost: {
                    type: 0,
                    value: 0,
                },
            },
            video: '',
        }
    });
    const [attribute, setAttribute] = useState(new Array<model.Attribute>());

    const addAttribute = () => {
        setAttribute([...attribute, new model.Attribute()]);
    };

    const removeAttribute = (index: number) => {
        setAttribute(attribute.filter((_, i) => i !== index));
    };

    const addAttributeItem = (index: number) => {
        const item = attribute[index];
        item.value.push('');
        setAttribute([...attribute]);
    };

    const removeAttributeItem = (index: number, itemIndex: number) => {
        const item = attribute[index].value;
        if (item && item.length > 0) {
            item.splice(itemIndex, 1);
        }
        setAttribute([...attribute]);
    };

    const getAttribute = async () => {
        let data = {
            ...new model.Request({ current: 1, pageSize: 1000 }),
            queryBy: [
                {
                    field: 'categoryId',
                    value: 1
                }
            ]
        }
        let res = await api.Product.findAttribute(data);
        if (res.code == 1000) {
            setAttribute(res.data.list);
        } else {
            // message.error(res.desc);
        }
    };

    // 获取运费模板
    const getFreight = async () => {
        let data = {
            ...new model.Request({ current: 1, pageSize: 1000 }),
            queryBy: [
                {
                    field: 'categoryId',
                    value: 1
                }
            ]
        }
        let res = await api.Product.findFreight(data);
        if (res.code == 1000) {
            setSpu({ ...spu, info: { ...spu.info, freight: res.data.list[0] } });
        } else {
            // message.error(res.desc);
        }
    };

    useEffect(() => {
        getAttribute();
        getFreight();
    }, []);

    let [tr, setTr] = useState<ReactNode[]>([]);
    useEffect(() => {
        const data = combineData(attribute).map((m, i) => <tr key={i}>
            {m.map((n, j) => <td key={j}>{n}</td>)}
            <td>
                <input type="text" />
            </td>
            <td>
                <input type="text" />
            </td>
            <td>
                <input type="text" />
            </td>
        </tr>);
        setTr(data);
    }, [attribute]);

    function combineData(data: model.Attribute[]): string[][] {
        const result: string[][] = [];
        // if (data.length < 2) {
        //     return result;
        // }
        // 用于存储所有属性值的数组
        const allValuesArrays: string[][]= [];
        data.forEach((item) => {
            if (item.multi) {
                allValuesArrays.push([item.value.join(',')]);
            } else {
                allValuesArrays.push(item.value);
            }
        });

        // 生成所有可能的组合
        function generateCombinations(index: number, currentCombination: string[]): void {
            if (index === allValuesArrays.length) {
                result.push([...currentCombination]);
                return;
            }
            const currentValues = allValuesArrays[index];
            for (let i = 0; i < currentValues.length; i++) {
                currentCombination.push(currentValues[i]);
                generateCombinations(index + 1, currentCombination);
                currentCombination.pop();
            }
        }
        generateCombinations(0, []);
        return result;
    }


    return <Container>
        <div className="box-head">
            <h2>商品类型</h2>
        </div>
        <ul className="box-body">
            <li className="list-item">
                <label>实物商品</label>
                <div>(物流发货)</div>
            </li>
            <li className="list-item">
                <label>虚拟商品</label>
                <div>(无需物流)</div>
            </li>
            <li className="list-item">
                <label>服务商品</label>
                <div>(无需物流)</div>
            </li>
            <li className="list-item">
                <label>电子卡券</label>
                <div>(无需物流)</div>
            </li>
        </ul>
        <div className="box-head">
            <h2>基础信息</h2>
        </div>
        <div className="box-body">
            <div className="form-item">
                <label className="form-label">商品名称</label>
                <div className="form-control">
                    <input type="text" />
                </div>
            </div>
            <div className="form-item">
                <label className="form-label">商品编码</label>
                <div className="form-control">
                    <input type="text" />
                </div>
            </div>
            <div className="form-item">
                <label className="form-label">商品图片</label>
                <div className="form-control">
                    <Upload onChange={(v) => setSpu({ ...spu, info: { ...spu.info, images: v } })} maxCount={3} />
                </div>
            </div>
            <div className="form-item">
                <label className="form-label">讲解视频</label>
                <div className="form-control">
                    <Upload onChange={(v) => setSpu({ ...spu, info: { ...spu.info, video: v[0] } })} maxCount={1} />
                </div>
            </div>
            <div className="form-item">
                <label className="form-label">商品类目</label>
                <button className="btn btn-default">选择商品类目</button>
            </div>
        </div>
        <div className="box-head">
            <h2>销售信息</h2>
        </div>
        <div className="box-body">
            <div className="form-item">
                <label className="form-label">销售属性</label>
                <Sell className="form-control">
                    {attribute.map((m, i) => <div className="sell-item" key={i}>
                        <div className="sell-item-name">
                            <label>规格名:</label>
                            <div className="column">
                                <input type="text" value={m.label} onChange={(e) => {
                                    const item = attribute[i];
                                    item.label = e.target.value;
                                    setAttribute([...attribute]);
                                }} />
                                {i == 0 && <label htmlFor="">
                                    <input type="checkbox" />添加规格图片
                                </label>}
                            </div>
                        </div>
                        <div className="sell-item-value">
                            <label>规格值:</label>
                            <div className="row">
                                {m.value.map((x, j) => <div className="input-box" key={j}>
                                    <input type="text" value={x} onChange={(e) => {
                                        const item = attribute[i];
                                        item.value[j] = e.target.value;
                                        setAttribute([...attribute]);
                                    }} />
                                    <a onClick={() => removeAttributeItem(i, j)}>x</a>
                                </div>)}
                                <a onClick={() => addAttributeItem(i)}>添加规格值</a>
                            </div>
                        </div>
                        <a className="sell-item-delete" onClick={() => removeAttribute(i)}>x</a>
                    </div>)}
                    <div className="sell-foot">
                        <a className="btn-default" onClick={addAttribute}>添加规格</a>
                    </div>
                </Sell>
            </div>
            <div className="form-item">
                <label className="form-label">销售规格</label>
                <div className="form-control">
                    <table className="table">
                        <thead>
                            {attribute.map((m, i) => <th key={i}>{m.label}</th>)}
                            <th>价格(元)</th>
                            <th>库存</th>
                            <th>规格编码</th>
                        </thead>
                        <tbody>
                            {tr.map(m => m)}
                        </tbody>
                    </table>
                </div>
            </div>
            <div className="form-item">
                <label className="form-label">商品价格</label>
                <div className="form-control">
                    <input type="text" />
                </div>
            </div>
            <div className="form-item">
                <label className="form-label">商品库存</label>
                <div className="form-control">
                    <input type="text" />
                    <label className="tips"><input type="checkbox" />商品详情、购物车不显示剩余件数</label>
                </div>
            </div>
            <div className="form-item">
                <label className="form-label">商品条码</label>
                <div className="form-control">
                    <input type="text" />
                </div>
            </div>
        </div>
        <div className="box-head">
            <h2>物流信息</h2>
        </div>
        <div className="box-body">
            <div className="form-item">
                <label className="form-label">配送方式</label>
                <div className="form-control">
                    <div className="checkbox-group">
                        <label className="checkbox-item">
                            <input type="checkbox" />快递发货
                        </label>
                        <label className="checkbox-item">
                            <input type="checkbox" />同城配送
                        </label>
                        <label className="checkbox-item">
                            <input type="checkbox" />到店自提
                        </label>
                    </div>
                </div>
            </div>
            <div className="form-item">
                <label className="form-label">快递运费</label>
                <div className="form-control">
                    <div className="radio-group">
                        <div className="row">
                            <label className="radio-item">
                                <input name="express-fee" type="radio" />统一邮费
                            </label>
                            <input type="number" placeholder="请输入邮费" />
                        </div>
                        <div className="row">
                            <label className="radio-item">
                                <input name="express-fee" type="radio" />运费模板
                            </label>
                            <select name="express-fee-template">
                                <option value="0">请选择运费模板</option>
                            </select>
                            <a href="#">新建运费模板</a>
                        </div>
                        <div className="col">
                            <label className="radio-item">
                                <input name="express-fee" type="radio" />运费到付
                            </label>
                            <Text className="tips">设置运费到付后，需要买家在收到商品后自行支付运费，运费最终以物流公司计算为准。</Text>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div className="box-head">
            <h2>售后服务</h2>
        </div>
        <div className="box-body">
            <div className="form-item">
                <label className="form-label">上架时间</label>
                <div className="form-control">
                    <div className="radio-group">
                        <div className="row">
                            <label className="radio-item">
                                <input name="publish-time" type="radio" />立即上架
                            </label>
                        </div>
                        <div className="row">
                            <label className="radio-item">
                                <input name="publish-time" type="radio" />定时上架
                            </label>
                            <input type="datetime-local" placeholder="请选择上架时间" />
                        </div>
                        <div className="row">
                            <label className="radio-item">
                                <input name="publish-time" type="radio" />放入仓库
                            </label>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div className="box-foot">
            {/* <button className="btn btn-primary">保存</button> */}
            <Button>保存</Button>
            <button className="btn btn-primary">预览</button>
        </div>
    </Container>;
}