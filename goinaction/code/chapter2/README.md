## 代码说明
本代码为第二章 快速开始对应的相关代码  
此处做了一些记录

### 项目目录结构
<pre><code>
├── sample                               <=
│ ├── data                           <= 
│ │  ├── form_static                    <=  统计页面的主文件夹
│ │  │  ├── form_static_detail.vue      <=  下半部分的数据展示子组件。主要包含： map_single 地图缩略图显示子组件（暂时使用div填充即可）， typhoon_detail_chart 可复用的echarts子组件
│ ├── matchers                            <= 
│ ├── search                            <= 业务逻辑都在search包中
│ ├── main.go                            <= 程序入口文件，包名称为main，且拥有main函数
</code></pre>

每个可执行的go程序都有两个特征：
1. 名为main的函数
2. main.go的包名为main  
注意 `main` 函数必须保存在 `main` 的包中；  
go中每一个.go文件都属于一个包；  
可以把`包名称`类比为`命名空间`
注意每个`.go` 的代码文件里的`init`函数都会在main函数执行前调用

### 部分语法知识点的汇总
map是一种特殊的数据结构：一种类似 k-v的 键值对（可以理解为字典）  
可以理解为 `字典`