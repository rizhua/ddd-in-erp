import { Button, Form, Input, message } from "antd";
import { Link, useNavigate } from "react-router-dom";
import styled from "styled-components";
import md5 from 'md5';
import { useState } from "react";

import { api } from "@/service";

import bgPic from '@/assets/image/20201203214346381.jpg'

const Container = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    min-width: 100vw;
    min-height: 100vh;
    background-image: url(${bgPic});
    background-size: cover;

    .main {
        width: 360px;
        border-radius: 8px;
        background-color: #fff;

        form {
            margin: 32px;
        }
    }

    .slogan {
        text-align: center;

        h2 {
            margin-top: 32px;
        }
    }

    .link {
        display: flex;
        justify-content: center;
        align-items: center;
        margin-bottom: 30px;
    }

    .justify-between {
        display: flex;
        justify-content: space-between;
    }

    .tab {
        display: flex;
        gap: 24px;
        margin-bottom: 24px;

        a {
            color: #333;
        }
    }

    .tab-item {
        cursor: pointer;

        &.active {
            &:after {
                content: '';
                display: block;
                margin-top: 8px;
                width: 100%;
                height: 2px;
                background-color: #f40;
            }
        }
    }

    .btn {
        display: flex;
        justify-content: center;
        align-items: center;
        outline: none;
    }

    .btn-link {
        background-color: transparent;
        border: none;
        height: 24px;
    }
`;

export function SignIn() {
    document.title = `登录`;
    const uri = new URLSearchParams(location.search.split('?')[1]);
    const navigate = useNavigate();
    const [userModel, setUserModel] = useState({
        tabindex: 1,
        mode: 0,
        sms: {
            text: '获取验证码',
            disabled: false,
        }
    });
    const [userForm] = Form.useForm();

    const onTab = (index: number) => {
        if (index == 2) {
            userModel.mode = 0;
        }
        setUserModel({
            ...userModel,
            tabindex: index
        });
    }

    const getCaptcha = async (phone: string) => {
        let data = {
            mobile: phone
        }
        let res = await api.System.sendSms(data);
        if (res.code === 1000) {
            let second = 60;
            let intv = setInterval(() => {
                second--;
                if (second <= 0) {
                    clearInterval(intv);
                    userModel.sms.disabled = true;
                } else {
                    userModel.sms.text = second + "秒后重发";
                    userModel.sms.disabled = true;
                }
                setUserModel({ ...userModel });
            }, 1000);
        } else {
            message.error(res.desc)
        }
    }

    const submit = async (values: any) => {
        let valid = await userForm.validateFields();
        if (!valid) {
            return;
        }
        let data = {
            account: values.account,
            password: md5(values.password),
        };
        let res = await api.User.signIn(data);
        if (res.code == 1000) {
            if (uri.has('redirect_url')) {
                location.href = `${uri.get('redirect_url')}?token=${res.data.token}`;
            } else {
                localStorage.setItem('token', res.data);
                navigate('/');
            }
        } else {
            message.error(res.desc);
        }
    }

    return <Container>
        <div className="main">
            <div className="slogan">
                <h2>美好生活，日抓开启</h2>
                <p>欢迎来到日抓，请登录</p>
            </div>
            <Form form={userForm} layout="vertical" size="large" onFinish={submit}>
                <div className="tab">
                    <a className={"tab-item" + (userModel.tabindex == 1 && " active")} onClick={() => onTab(1)}>手机号登录</a>
                    <a className={"tab-item" + (userModel.tabindex == 2 && " active")} onClick={() => onTab(2)}>邮箱登录</a>
                </div>
                <Form.Item name="account"  rules={[{ required: true, message: userModel.tabindex == 1 ? '请输入手机号' : '请输入邮箱' }]}>
                    <Input placeholder={userModel.tabindex == 1 ? '手机号' : '邮箱'} />
                </Form.Item>
                {userModel.mode == 0 && <Form.Item name="password" rules={[{ required: true, message: '请输入密码' }]}>
                    <Input.Password placeholder="密码" />
                </Form.Item>}
                {userModel.mode == 1 && <Form.Item name="captcha" rules={[{ required: true, message: '请输入验证码' }]}>
                    <Input
                        suffix={<button
                            className="btn btn-link"
                            onClick={() => getCaptcha(userForm.getFieldValue('account'))}
                            type="button"
                            disabled={userModel.sms.disabled}
                        >{userModel.sms.text}</button>} placeholder="验证码" />
                </Form.Item>}
                <Form.Item>
                    <div className="justify-between">
                        {userModel.tabindex == 1 ? (userModel.mode == 0 ? <a onClick={() => setUserModel({ ...userModel, mode: 1 })}>短息登录</a> : <a onClick={() => setUserModel({ ...userModel, mode: 0 })}>密码登录</a>) : <a></a>}
                        <Link to="/password/forget">忘记密码</Link>
                    </div>
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit" block>登录</Button>
                </Form.Item>
            </Form>
            <div className="link">
                还没有帐号?<Link to="/auth/signup">免费注册</Link>
            </div>
        </div>
    </Container>
}