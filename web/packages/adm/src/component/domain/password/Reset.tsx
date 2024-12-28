import styled from "styled-components";
import { Form, Input, Button, message } from 'antd';
import md5 from 'md5';

import { api } from "@/service";
import { useState } from "react";

const Container = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: calc(100vh - 152px);

    .wrap {
        width: 310px;
        max-width: 500px;
    }
`;

export function Reset(props: any) {
    let [email, setEmail] = useState('');
    let [token, setToken] = useState('');
    let uri = new URLSearchParams(location.search.split('?')[1]);
    if (uri.hasOwnProperty('email')) {
        setEmail(uri.get('email')!.toString);
    }
    if (uri.hasOwnProperty('token')) {
        setToken(uri.get('token')!.toString);
    }
    const [userForm] = Form.useForm();

    const submit = async () => {
        let data = {
            email,
            token,
            password: md5(userForm.getFieldValue('password'))
        };
        let res = await api.User.reset(data);
        if (res.code == 1000) {
            message.success("密码修改成功！");
            setTimeout(() => {
                props.history('/auth/signin');
            }, 5000);
        }
    }

    return <Container>
        <div className="wrap">
            <h2>Hi，{email}</h2>
            <p>请输入您的新密码进行重置操作。</p>
            <Form form={userForm} onFinish={submit}>
                <Form.Item name="password">
                    <Input placeholder="密码，8-16位，必须含大小写字母+数字" />
                </Form.Item>
                <Form.Item name="mobile">
                    <Input placeholder="输入手机号码" />
                </Form.Item>
                <Form.Item>
                    <Button type="primary" htmlType="submit" block>重置密码</Button>
                </Form.Item>
            </Form>
        </div>
    </Container>
}