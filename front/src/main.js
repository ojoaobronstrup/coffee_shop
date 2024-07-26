import App from './App.vue'
import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import routes from './routes'
import store from './store'

const app = createApp(App)

app.use(routes)
app.use(ElementPlus)
app.use(store)

app.mount('#app')
