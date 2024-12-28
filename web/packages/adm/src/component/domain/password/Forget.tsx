import { useState } from "react";
import { Link } from "react-router-dom";
import { Form, Input, Button, message } from "antd";
import styled from "styled-components";

import { api } from "@/service";
import { Captcha } from "@/component/common";

const Container = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;

    .wrap {
        min-width: 360px;
    }

    .link {
        display: flex;
        justify-content: space-between;
        margin-bottom: .5rem;
    }

    .main {
        padding: 60px 1rem;
        text-align: center;
        border: 1px solid #ddd;

        .btn {
            margin-top: 30px;
        }
    }

    .send {
        margin-top: 2rem;
        padding: 0.5rem 1rem;
        a {
            cursor: pointer;
            color: #f60;
        }
    }
`;

export function Forget() {
    const [userForm] = Form.useForm();
    const [step, setStep] = useState(1);
    const [mode, setMode] = useState({ value: true, text: '手机号找回' });
    const website = new Map([
        ['qq.com', 'https://mail.qq.com'],
        ['126.com', 'https://www.126.com'],
        ['163.com', 'https://mail.163.com'],
        ['yeah.net', 'https://www.yeah.net'],
        ['outlook.com', 'https://outlook.live.com'],
        ['live.com', 'https://outlook.live.com']
    ]);
    const [mobile, setMobile] = useState('');

    const toEmail = () => {
        let domain = userForm.getFieldValue('email').split("@")[1];
        if (website.has(domain)) {
            window.open(website.get(domain));
        }
    }

    const submit = async () => {
        let data = {
            ...userForm.getFieldsValue(),
        };
        let res = await api.User.forget(data);
        if (res.code == 1000) {
            setStep(2);
        } else {
            message.error(res.desc);
        }
    }

    return <Container>
        <div className="wrap">
            <Form form={userForm} onFinish={submit} size="large" style={{ display: step == 1 ? 'block' : 'none' }}>
                {mode.value ? <Form.Item name="email" rules={[{ required: true }]}>
                    <Input placeholder="注册邮箱" />
                </Form.Item> : <>
                    <Form.Item name="mobile" rules={[{ required: true }]}>
                        <Input addonBefore="+86" onChange={(e) => setMobile(e.target.value)} placeholder="手机号" />
                    </Form.Item>
                    <Form.Item name="captcha" rules={[{ required: true }]}>
                        <Captcha phoneNumber={mobile} />
                    </Form.Item>
                </>
                }
                <div className="link">
                    <a onClick={() => setMode({ value: !mode.value, text: mode.value ? '邮箱找回' : '手机号找回' })}>{mode.text}</a>
                    <Link to="/auth/signin">返回登录</Link>
                </div>
                <Button type="primary" htmlType="submit" block>找回密码</Button>
            </Form>
            {step == 2 && <>
                <div className="main">
                    <h3>Hi，{userForm.getFieldValue('email')}</h3>
                    <p>验证邮件发送成功。</p>
                    <p>请至邮箱查收验证邮件，进行邮箱确认操作。</p>
                    <Button className="btn" onClick={toEmail} size="large" type="primary">立刻登录邮箱完成验证</Button>
                </div>
                <div className="send">
                    <span>若您没有收到邮件：检查您的垃圾邮件中，是否包含验证邮件；或者</span>
                    <a onClick={submit}>重发验证邮件</a>
                </div>
            </>}
        </div>
    </Container>
}