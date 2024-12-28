import { Alert, Button, Col, Form, Input, Row, Select } from "antd";
import { useContext } from "react";
import styled from "styled-components";

import { UserContext } from "@/context";

const Container = styled.div`
    margin: 1rem;

    h3 {
        margin-left: 5rem;
        font-weight: 500;
    }

    .box {
        margin-bottom: 1rem;
        padding: 1rem;
    }

    .avatar {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .item {
        display: flex;
        line-height: 50px;

        label {
            min-width: 100px;
        }

        a {
            margin-left: 1rem;
        }
    }

    .ant-alert {
        margin-bottom: 2rem;
    }
`;

const Avatar = styled.a`

    img {
        width: 100px;
        border-radius: 50%;
    }
`;

export function Profile() {
    document.title = '基本信息';
    let userContext = useContext(UserContext);

    return <Container>
        <div className="box-head">
            <div className="text">
                <div className="avatar">
                    <img src="" alt="" />
                </div>
                <div>
                    <h1>{userContext.state.nickname}</h1>
                    <div>
                        <div>
                            <label htmlFor="">关注</label>
                            <span>1</span>
                        </div>
                        <div>
                            <label htmlFor="">粉丝</label>
                            <span>1</span>
                        </div>
                        <div>
                            <label htmlFor="">获赞</label>
                            <span>1</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div className="box">
            <Alert message="基本资料以实名信息为准，以下信息仅供参考，填写以下信息方便我们更好为您服务。" type="info" />
            <Form labelCol={{ span: 8 }} style={{ maxWidth: 600 }}>
                <div className="box-head">
                    <h3>核心信息</h3>
                </div>
                <div className="box-body">
                    <Form.Item label="会员身份">
                        <span>个人</span>
                    </Form.Item>
                    <Form.Item label="真实姓名">
                        <span>**三</span>
                    </Form.Item>
                </div>
                <div className="box-head">
                    <h3>业务信息</h3>
                </div>
                <div className="box-body">
                    <Form.Item label="应用行业">
                        <Select options={[]} />
                    </Form.Item>
                    <Form.Item label="产品名称">
                        <Input />
                    </Form.Item>
                    <Form.Item label="企业网址">
                        <Input />
                    </Form.Item>
                </div>
                <div className="box-head">
                    <h3>联系信息</h3>
                </div>
                <div className="box-body">
                    <Form.Item label="国家/地区">
                        <span>中国</span>
                    </Form.Item>
                    <Form.Item label="所在地区">
                        <Input />
                    </Form.Item>
                    <Form.Item label="街道地址">
                        <Input />
                    </Form.Item>
                    <Form.Item label="联系电话">
                        <Input />
                    </Form.Item>
                </div>
                <Row>
                    <Col span={8}></Col>
                    <Button type="primary">保存</Button>
                </Row>
            </Form>
        </div>
    </Container>
}