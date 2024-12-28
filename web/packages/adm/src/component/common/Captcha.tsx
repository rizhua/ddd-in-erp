import { api } from "@/service";
import { Input, message } from "antd";
import { useEffect, useState } from "react";
import styled from "styled-components";

const Button = styled.button`
    display: flex;
    justify-content: center;
    align-items: center;
    outline: none;
    background-color: transparent;
    border: none;
    height: 24px;
`

type CaptchaProps = {
    phoneNumber: string
}
export function Captcha(props: CaptchaProps) {
    const [smsModel, setSmsModel] = useState({ text: "获取验证码", disabled: false });

    let intv: NodeJS.Timeout;
    const getCaptcha = async () => {
        if (!props.phoneNumber) {
            return;
        }
        let data = {
            mobile: props.phoneNumber,
        }
        let res = await api.System.sendSms(data);
        if (res.code === 1000) {
            let second = 10;
            intv = setInterval(() => {
                second--;
                if (second <= 0) {
                    clearInterval(intv);
                    smsModel.text = "获取验证码";
                    smsModel.disabled = false;
                } else {
                    smsModel.text = second + "秒后重发";
                    smsModel.disabled = true;
                }
                setSmsModel({ ...smsModel });
            }, 1000);
        } else {
            message.error(res.desc)
        }
    }

    useEffect(() => {
        return () => {
            clearInterval(intv);
        }
    }, []);


    return  <Input
        suffix={<Button
            onClick={() => getCaptcha()}
            type="button"
            disabled={smsModel.disabled}
        >{smsModel.text}</Button>} placeholder="验证码" />
}