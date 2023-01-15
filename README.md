# Tiktok-backend

## Milestones for the project

- Basic features
  - Video feed stream
  - Vide upload
  - Homepage
- Likes list
- Comments
- Relationship list
- Messages

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

[抖音项目方案说明](https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#)

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

### 功能说明

接口功能不完善，仅作为示例

- 用户登录数据保存在内存中，单次运行过程中有效
- 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试

test 目录下为不同场景的功能测试 case，可用于验证功能实现正确性

其中 common.go 中的 _serverAddr_ 为服务部署的地址，默认为本机地址，可以根据实际情况修改

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

## Android emulator

[极简抖音 App 使用说明及下载](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)

**Windows**

[Android Studio](https://developer.android.com/studio)

After installation, click virtual deveice manager to install virtual device.

Drag apk from windows to virtual device can install the app.

To connect backend, change baseUrl to "HTTP://10.0.2.2:8080"
