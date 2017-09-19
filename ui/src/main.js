// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

import iView from 'iview'
import 'iview/dist/styles/iview.css' // 使用 CSS

Vue.use(iView)

router.beforeEach((to, from, next) => {
  iView.LoadingBar.start()
  next()
})
router.afterEach((to, from, next) => {
  iView.LoadingBar.finish()
})

import axios from 'axios'

var axio = axios.create({
  baseURL: 'http://localhost:9000'
})

axio.interceptors.response.use(resp => {
  if (resp.data.code === 200) {
    return resp.data.data
  } else {
    return Promise.reject(resp.data)
  }
}, error => {
  return Promise.reject(error)
})

Vue.prototype.$http = axio

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
