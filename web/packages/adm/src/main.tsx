import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import { ConfigProvider } from 'antd';
import zhCN from 'antd/locale/zh_CN';

import App from '@/app/App';
import "@radix-ui/themes/styles.css";
import { Theme } from "@radix-ui/themes";
import '@/assets/style/index.less';
import { UserProvider } from '@/context';


const rootDom = document.getElementById('root') as HTMLElement;
ReactDOM.createRoot(rootDom).render(
  <BrowserRouter future={{v7_relativeSplatPath: true, v7_startTransition: true}}>
    <Theme accentColor="teal">
    <UserProvider>
    <ConfigProvider locale={zhCN}>
      <App />
    </ConfigProvider>
    </UserProvider>
    </Theme>
  </BrowserRouter>
);
