import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'

// 引入TDesign样式
import 'tdesign-vue-next/es/style/index.css'
import './styles/theme.css'
import './styles/form.css'
import dayjs from 'dayjs';
// 引入自定义组件
import FlexRow from './components/FlexRow.vue'

const app = createApp(App)
import isoWeek from 'dayjs/plugin/isoWeek';
dayjs.extend(isoWeek);
app.use(createPinia())
app.use(router)

// 全局注册组件
app.component('FlexRow', FlexRow)

app.mount('#app')
