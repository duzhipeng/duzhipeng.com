// ==UserScript==
// @name         让阿杜来
// @namespace    https://www.duzhipeng.com/
// @version      0.1.0
// @description  这事让你干我怎么好意思呢？让我来！
// @author       DU ZHIPENG
// @match        https://*.shinwell.cn/*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=greasyfork.org
// @grant        none
// ==/UserScript==

(function () {
  "use strict";

  // Your code here...
  //debugger;
  console.log("ADU-程序初始化完成。");

  // 新增维保单自动填表
  function autoFillAddForm() {
    const label_ok = document.querySelector('label[title="是否抢修"]');
    if (label_ok) {
      const label = label_ok.parentNode.nextSibling;
      const formItem = label.querySelector('input[type="radio"]');
      // 要选择的单选按钮的选择器;
      // 检查单选按钮是否已经选中;
      if (!formItem.checked) {
        // 模拟点击以触发 Vue 的事件
        formItem.click();
        // 手动设置 checked 属性
        formItem.checked = true;
      }
    }
  }

  // 创建一个 MutationObserver 实例
  const observer = new MutationObserver((mutationsList) => {
    for (const mutation of mutationsList) {
      if (mutation.type === "childList") {
        // 遍历新增的节点
        mutation.addedNodes.forEach((node) => {
          document.addEventListener("click", function () {
            // 获取节点的文本内容
            const nodeText = node.textContent || node.innerText;
            // 新增维保单界面
            const key_word = "新增维保单";
            if (nodeText.includes(key_word)) {
              // 为整个文档绑定点击事件
              autoFillAddForm(); // 是否抢修：是
              // 停止监听
              observer.disconnect();
            }
          });
        });
      }
    }
  });

  // 配置观察选项
  const config = { childList: true, subtree: true };

  // 开始观察页面的变化
  observer.observe(document.body, config);
})();
