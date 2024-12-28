import { useEffect, useState } from "react";
import styled from "styled-components";
import { Table } from "antd";

import { api, model } from "@/service";
import { s } from "vite/dist/node/types.d-aGj9QkWt";

const Contianer = styled.div`
    position: absolute;
    top: 0;
    right: 0;
    width: 50vw;
    height: 100vh;
    background-color: #fff;
    box-shadow: -1px 0px 2px #eee;
    z-index: 100;

    .box-head {
        display: grid;
        grid-template-columns: 50px 1fr;
        align-items: center;

        .title {
            margin: 0;
            padding: 0 1rem;
        }

        .close {
            cursor: pointer;
            display: flex;
            justify-content: center;
        }
    }

    .box-body {
        padding: 1rem;
    }

    .box-item {
        border-bottom: 1px solid #ddd;
        padding-bottom: 1rem;
        margin-bottom: 1rem;
    }
`;

export interface LicenseProps {
    open: boolean;
    onClose: () => void;
    bundle: model.Bundle;
}
export function License(props: LicenseProps) {
    const [license, setLicense] = useState({
        list: new Array<model.License>(),
        total: 0,
    });

    const getLicense = async () => {
        let data = {
            ...new model.Request(),
            queryBy: [{
                field: 'bizId',
                value: props.bundle.id,
            }],
        }
        let res = await api.Bundle.findLicense(data);
        if (res.code == 1000) {
            license.list = res.data.list;
            license.total = res.data.total;
        } else {
            license.list = [];
            license.total = 0;
        }
        setLicense({...license});
    }

    useEffect(() => {
        if (props.open) {
            getLicense();
        }
    }, [props.open]);

    const close = () => {
        license.list = [];
        license.total = 0;
        setLicense({...license});
        props.onClose();
    }

    return props.open && <Contianer>
        <div className="box-head">
            <a className="close iconfont" onClick={() => close()}>&#xe6b3;</a>
        </div>
        <div className="box-body">
            <h1>授权密钥</h1>
            <div className="box-list">
                {license.list.map((v) => <div className="box-item" key={v.id}>
                    <div className="head">{v.createAt}</div>
                    <div className="text">{v.code}</div>
                </div>)}
            </div>
        </div>
    </Contianer>
}