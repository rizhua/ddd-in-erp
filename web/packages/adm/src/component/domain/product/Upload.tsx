import { api } from "@/service";
import { message } from "antd";
import { useRef, useState } from "react";
import styled from "styled-components";

const MAX_SIZE = 1024 * 1024 * 5; // 5MB

const Container = styled.div`
    display: flex;
    gap: 16px;
`;

const Card = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 4px;
    width: 100px;
    height: 100px;
    border: 1px dashed #e0e0e0;
    border-radius: 4px;
    cursor: pointer;
    overflow: hidden;
    position: relative;

    img {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }

    .btn-box {
        display: flex;
        justify-content: center;
        align-items: center;
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-color: rgba(0, 0, 0, 0.3);
        color: #fff;
        opacity: 0;
        transition: opacity 0.3s;
        cursor: pointer;
        z-index: 10;
    }

    &:hover {
        .btn-box {
            opacity: 1;
        }
    }

    .upload-box {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        cursor: pointer;

        i {
            font-size: 24px;
        }

        p {
            margin: 4px 0 0;
            font-size: 12px;
        }
    }
`;

export interface UploadProps {
    onChange: (fileList: string[]) => void;
    maxCount?: number;
    accept?: string;
    multiple?: boolean;
}

export function Upload(props: UploadProps) {
    const { maxCount = 2, accept = 'image/*', multiple = false } = props;
    const inputRef = useRef<HTMLInputElement>(null);

    const onUpload = (e: React.MouseEvent) => {
        e.stopPropagation();
        if (inputRef.current) {
            inputRef.current.click();
        }
    }

    const [fileList, setFileList] = useState(new Array<string>());

    const onFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
        const files = e.target.files;
        if (!files || files?.length == 0) {
            return;
        }
        
        for (let file of files) {
            await doFile(file);
        }
    }

    // 处理单个文件
    const doFile = async (file: File) => {
        if (file.size > MAX_SIZE) {
            message.error('文件大小不能超过5MB');
            return;
        }
        let formData = new FormData();
        formData.append("file", file);
        let res = await api.File.upload(formData);
        if (res.code === 1000) {
            let s = new Set(fileList);
            s.add(res.data);
            let tmp = Array.from(s);
            setFileList([...tmp]);
            props.onChange(tmp);
        } else {
            message.error(res.desc);
        }
    }

    const onDelete = (e: React.MouseEvent, url: string) => {
        e.stopPropagation();
        let s = new Set(fileList);
        s.delete(url);
        setFileList([...Array.from(s)]);
    }

    return <Container>
        {fileList.map(m => <Card key={m}>
            <img src={m} alt="" />
            <div className="btn-box">
                <a className="iconfont" onClick={e => onDelete(e, m)}>&#xe618;</a>
            </div>
        </Card>)}
        {fileList.length < maxCount && <Card onClick={onUpload}>
            <input type="file" ref={inputRef} hidden onChange={onFileChange} accept={accept} multiple={multiple} />
            <div className="upload-box">
                <i className="iconfont">&#xe678;</i>
                <p>上传图片</p>
            </div>
        </Card>}
    </Container>
}