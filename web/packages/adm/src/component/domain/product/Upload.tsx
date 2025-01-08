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
                <a onClick={e => onDelete(e, m)}>
                <svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M5.5 1C5.22386 1 5 1.22386 5 1.5C5 1.77614 5.22386 2 5.5 2H9.5C9.77614 2 10 1.77614 10 1.5C10 1.22386 9.77614 1 9.5 1H5.5ZM3 3.5C3 3.22386 3.22386 3 3.5 3H5H10H11.5C11.7761 3 12 3.22386 12 3.5C12 3.77614 11.7761 4 11.5 4H11V12C11 12.5523 10.5523 13 10 13H5C4.44772 13 4 12.5523 4 12V4L3.5 4C3.22386 4 3 3.77614 3 3.5ZM5 4H10V12H5V4Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
                </svg>
                </a>
            </div>
        </Card>)}
        {fileList.length < maxCount && <Card onClick={onUpload}>
            <input type="file" ref={inputRef} hidden onChange={onFileChange} accept={accept} multiple={multiple} />
            <div className="upload-box">
                <svg width="15" height="15" viewBox="0 0 15 15" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M8 2.75C8 2.47386 7.77614 2.25 7.5 2.25C7.22386 2.25 7 2.47386 7 2.75V7H2.75C2.47386 7 2.25 7.22386 2.25 7.5C2.25 7.77614 2.47386 8 2.75 8H7V12.25C7 12.5261 7.22386 12.75 7.5 12.75C7.77614 12.75 8 12.5261 8 12.25V8H12.25C12.5261 8 12.75 7.77614 12.75 7.5C12.75 7.22386 12.5261 7 12.25 7H8V2.75Z" fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"></path>
                </svg>
                <p>上传图片</p>
            </div>
        </Card>}
    </Container>
}