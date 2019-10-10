# rem_layout

## 概念

- 固定布局
- 自适应布局
  - 流式布局
    - 百分比
  - 弹性布局
    - 弹性布局
- 响应式布局
  - 媒体查询


## 实现方式

- viewport

- em   

  - https://www.w3cplus.com/css/px-to-em    

  - 当你想要当前元素的 padding，margin，line-height 等值，与当前字体大小成比例的时候，使用 em 单位。

  - 浏览器的默认字体大小是16px

  - 如果元素自身没有设置字体大小

    那么元素自身上的所有属性值如“border、width、height、padding、margin、line-height”等值，按照如下公式计算:

    ```
    1 / 父元素的font-size * 设计尺寸 = em值
    ```

  - 如果元素自身字体设置了大小

    那么元素自身上的所有属性值如“border、width、height、padding、margin、line-height”等值，按照如下公式计算:

    ```
    1 / 元素自身的font-size * 设计尺寸 = em值
    ```

    ​

- rem

  - http://www.alloyteam.com/2016/03/mobile-web-adaptation-tool-rem/
  - 根元素html的font-size
  - meta标签 
    - 否则获取到的是980
  - 获取浏览器宽度
    - window.innerwidth || document.documentElement.clientWidth || document.body.clientWidth
  - width/3.75 (设计图以6 为基准)
    - font-size不可以小于12
  - resize

- 媒体查询

- flex

- 图片自适应









font-size和vertical-align https://www.w3cplus.com/css/css-font-metrics-line-height-and-vertical-align.html  



常用meta https://segmentfault.com/a/1190000004987068    