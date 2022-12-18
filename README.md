### 项目启动

#### 前端启动

环境：node.js v16.18.1  npm版本9.1.2	

进入frontend目录，执行npm run serve启动前端，若出现module找不到错误，删除node_modules，package-lock.json 后重新执行npm install导包，再执行npm run serve运行

#### 后端启动

直接运行backend.exe文件，或是在go环境下，在backend文件下执行go build生成可执行文件，运行可执行文件，后端启动端口默认为8083，可在main.go中，路由器组件run方法更改

#### 数据库配置

数据库文件已导出为sys.sql，可在本地MySQL数据库中执行创建数据库，数据库配置信息在后端dao包下mysql.go中dsn变量配置