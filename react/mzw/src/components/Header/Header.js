import React, { useState } from 'react';
import { Row, Menu, Button, Col } from 'antd';
import 'antd/dist/antd.css'
import './Header.css'

const ButtonGroup = Button.Group;
function Header() {
    const [meun, setMeun] = useState("top");
    return (
        <header>
            <Row>
                <Col span={20}>
                    <Menu mode="horizontal" selectedKeys={[meun]} onClick={(e) => { setMeun(e.key) }}>
                        <Menu.Item key="top">
                            PPT模板
                    </Menu.Item>
                        <Menu.Item key="shehui">
                            设计素材
                    </Menu.Item>
                        <Menu.Item key="guonei">
                            平面模板
                    </Menu.Item>
                        <Menu.Item key="guoji">
                            图片
                    </Menu.Item>
                        <Menu.Item key="yule">
                            视频
                     </Menu.Item>
                        <Menu.Item key="tiyu">
                            音乐配乐
                    </Menu.Item>
                        <Menu.Item key="keji">
                            更多
                    </Menu.Item>
                        <Menu.Item key="shishang">
                            时尚
                   </Menu.Item>
                </Menu>
                </Col>
                <Col span={4} className="btnitem">
                    <ButtonGroup>
                        <Button type="primary" ghost>
                            登录
                        </Button>
                        <Button type="primary" ghost>
                            注册
                        </Button>
                    </ButtonGroup>
                </Col>
            </Row>
        </header>
    );
}

export default Header
