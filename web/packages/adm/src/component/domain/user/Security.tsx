import { Progress } from "antd";
import styled from "styled-components";


const Container = styled.div`
    margin: 1rem;
    padding: 1rem;
    min-height: calc(100vh - 84px);
`

const Rate = styled.div`
    display: flex;
    justify-content: flex-start;
    padding: 1rem 0;
    border-bottom: 1px dotted #ddd;

    label {
        margin-right: .5rem;
    }

    span {
        margin-left: 24px;
    }
`;

export function Security() {
    document.title = `${document.title} - 安全设置`;

    return <Container className="box">
        <Rate>
            <label>安全等级:</label>
            <Progress percent={30} showInfo={false} style={{ maxWidth: 200 }} />
            <span>建议您启动全部安全设置，以保障帐号及资金安全。 </span>
        </Rate>
        <table className="table table-bordered">
            <tbody>
                <tr>
                    <td>登录密码</td>
                    <td>互联网账号存在被盗风险，建议您定期更改密码以保护账户安全。</td>
                    <td>修改</td>
                </tr>
                <tr>
                    <td>邮箱验证</td>
                    <td>您验证的邮箱：<span>zh*****xi@126.com</span></td>
                    <td>修改</td>
                </tr>
                <tr>
                    <td>手机绑定</td>
                    <td>您已绑定了手机<span>135****6629</span></td>
                    <td>修改</td>
                </tr>
                <tr>
                    <td>支付密码</td>
                    <td>建议您定期更换新的支付密码，提高安全性。</td>
                    <td>修改</td>
                </tr>
            </tbody>
        </table>
    </Container>
}