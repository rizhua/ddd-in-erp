import { Link } from "react-router-dom";
import { Button, Space } from "antd";
import styled from "styled-components";


const Box = styled.div`

    .title {
        border-bottom: 1px solid #ddd;
    }

    .card {
        flex: 1;
        min-height: 500px;
        text-align: center;
        background-color: #f0f0f0;

        &:last-child {
            margin-left: 24px;
        }

        h3 {
            line-height: 72px;
            background-color: #ddd;
        }
    }
`;

export function Invite() {

    return <Box className="box">
        <div className="box-head">
            <Link to="../">
                <i className="iconfont">&#xe613;</i>邀请成员
            </Link>
        </div>
        <div className="box-body">
            <div className="box-head">
                <div className="text">可通过以下任意方式邀请成员加入，成员申请后需管理员审核。</div>
                <Space className="tool" size={36}>
                    <a><i className="iconfont">&#xe622;</i>刷新链接</a>
                    <a><i className="iconfont">&#xe6d2;</i>停止邀请</a>
                </Space>
            </div>
            <div className="row">
                <div className="card">
                    <h3>方式1：分享二维码邀请加入</h3>
                    <Button>下载二维码图片</Button>
                </div>
                <div className="card">
                    <h3>方式2：成员访问链接申请加入</h3>
                    <Button>复制邀请链接</Button>
                </div>
            </div>
        </div>
    </Box>
}