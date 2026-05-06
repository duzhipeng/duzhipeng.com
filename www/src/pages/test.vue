<script setup>
import { ref } from 'vue'
import { jsPDF } from 'jspdf'
import autoTable from 'jspdf-autotable'

// 👉 你要注入的参数
const pdfData = ref({
  title: '测试报告',
  name: '杜志鹏',
  age: 26,
  department: '技术部',
  tableData: [
    { id: 1, project: '项目A', hours: 100 },
    { id: 2, project: '项目B', hours: 200 },
  ]
})

// 生成 PDF
const createPDF = () => {
  const doc = new jsPDF()

  // 写入内容（模板 + 注入参数）
  doc.text(`标题：${pdfData.value.title}`, 20, 20)
  doc.text(`姓名：${pdfData.value.name}`, 20, 30)
  doc.text(`部门：${pdfData.value.department}`, 20, 40)

  // 表格
  autoTable(doc, {
    head: [['ID', '项目', '工时']],
    body: pdfData.value.tableData.map(i => [i.id, i.project, i.hours]),
    startY: 50,
  })

  // 下载
  doc.save('报告.pdf')
}
</script>

<template>
  <button @click="createPDF">生成 PDF</button>
</template>