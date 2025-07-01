---
title: 配置服务器的操作的备忘（v2023.2)
sections: [ 
  {id: "system", name: "适用系统"},
  {id: "environment", name: "更新环境"},
  {id: "disk", name: "挂载数据盘"},
  {id: "adduser", name: "增加用户"},
  {id: "install", name: "安装必要软件"},
  {id: "nginx", name: "配置 Nginx",sub: [
    {id: "single", name: "仅前端配置"},
    {id: "multiple", name: "前后端配置"},        
    {id: "enable", name: "启用配置"},      
  ]},
  {id: "supervisor", name: "配置 Supervisor",sub: [
    {id: "extension", name: "文件后缀"},      
    {id: "common", name: "常用命令"},      
  ]},
]
---

<script setup>
import {onMounted} from "vue";

const emits = defineEmits(["syncMeta"]);

onMounted(()=>{
    emits("syncMeta", frontmatter);
    
})
</script>

Markdown 目录：
[TOC]

Markdown 标题：
# 这是 H1
## 这是 H2
### 这是 H3

Markdown 列表：
- 列表项目
1. 列表项目

*斜体*或_斜体_
**粗体**
***加粗斜体***
~~删除线~~

Markdown 插入链接：
[链接文字](链接网址 "标题")

Markdown 插入图片：

[//]: # (![shop_qrcode.png]&#40;..%2F..%2Fassets%2Fshop_qrcode.png&#41;)

Markdown 插入代码块：

```javascript {.line-numbers .match-braces .rainbow-braces data-line=4}
$(document).ready(function () {
    alert('RUNOOB');
});
$(document).ready(function () {
    alert('RUNOOB');
});
$(document).ready(function () {
    alert('RUNOOB');
});
$(document).ready(function () {
    alert('RUNOOB');
});
$(document).ready(function () {
    alert('RUNOOB');
});
$(document).ready(function () {
    alert('RUNOOB');
});
```

Markdown 引用：
> 引用内容

Markdown 分割线：
---

Markdown 换行：
<br>

Markdown 段首缩进：
&ensp; or &#8194; 表示一个半角的空格
&emsp; or &#8195;  表示一个全角的空格
&emsp;&emsp; 两个全角的空格（用的比较多）
&nbsp; or &#160; 不断行的空白格

\begin{equation}
  a^2+b^2=c^2
\end{equation}

\begin{equation}
  a^2+b^2=c^2
\end{equation}

\begin{equation}
  \begin{pmatrix}
    A & B \\ B & C
  \end{pmatrix} 
\end{equation}

Euler\'s identity $e^{i\\pi}+1=0$ is a beautiful formula in $\\RR^2$.