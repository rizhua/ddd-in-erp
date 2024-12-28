import styled from "styled-components";
import { Form, Input, Button, Checkbox, message } from 'antd';
import { Link } from "react-router-dom";
import md5 from 'md5';

import { api } from "@/service";
import { useState } from "react";

const Container = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    background-color: #f0f0f0;

    .wrap {
        max-width: 500px;
        background-color: #fff;
        padding: 50px 64px;
        border-radius: 4px;
    }

    .link {
    text-align: center;
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

export function SignUp() {
    document.title = `注册`;
    const [userForm] = Form.useForm();
    const [captchaText, setCaptcha] = useState('获取验证码');

    const getCaptcha = async () => {
        let mobile = userForm.getFieldValue('mobile');
        if (!mobile) {
            return;
        }

        let data = {
            mobile,
        };
        let res = await api.System.sendSms(data);
        if (res.code == 1000) {
            let num = 60;
            let intv = window.setInterval(() => {
                setCaptcha(`${num}秒重新发送`);
                if (0 >= num) {
                    window.clearInterval(intv);
                    setCaptcha('获取验证码');
                }
                num--;
            }, 1000);
        }
    }

    userForm.setFieldValue('agree', false);

    const submit = async () => {
        let valid = await userForm.validateFields();
        if (!valid) {
            return;
        }
        if (!userForm.getFieldValue('agree')) {
            message.error('请阅读用户协议及隐私政策并勾选！');
            return;
        }
        let data = {
            email: userForm.getFieldValue('email'),
            mobile: userForm.getFieldValue('mobile'),
            password: md5(userForm.getFieldValue('password')),
            captcha: userForm.getFieldValue('captcha'),
        };
        let res = await api.User.signUp(data);
        if (res.code == 1000) {
            message.success('注册成功，请前往登录');
        } else {
            message.error(res.desc);
        }
    }

    return <Container>
        <div className="wrap">
            <h2>欢迎注册</h2>
            <Form form={userForm} size="large">
                <Form.Item name="mobile" rules={[{required: true, message: '请输入手机号'}]}>
                    <Input placeholder="手机号码" />
                </Form.Item>
                <Form.Item name="captcha" rules={[{required: true, message: '请输入验证码'}]}>
                    <Input suffix={<button className="btn btn-link" onClick={() => getCaptcha()} type="button">{captchaText}</button>} placeholder="手机验证码" />
                </Form.Item>
                {/* <Form.Item name="email" rules={[{required: true, message: '请输入邮箱'}]}>
                    <Input placeholder="邮箱，作为登录账号" />
                </Form.Item> */}
                <Form.Item name="password" rules={[{ required: true, message: '请输入密码' }]}>
                    <Input.Password placeholder="密码，8-16位，必须含大小写字母+数字" />
                </Form.Item>
                <Form.Item name="orgName">
                    <Input placeholder="企业名称" />
                </Form.Item>
                <Form.Item name="agree" valuePropName="checked">
                    <Checkbox>
                        &nbsp;我已阅读并同意
                        <Link to="">《服务用户协议》</Link>和<Link to="">《隐私权政策》</Link>
                    </Checkbox>
                </Form.Item>
                <Form.Item>
                    <Button type="primary" onClick={() => submit()} block>同意协议并提交</Button>
                </Form.Item>
            </Form>
            <div className="link">
                <Link to="/auth/signin">登录已有账号</Link>
            </div>
        </div>
    </Container>
}