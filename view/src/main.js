// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.

import 'normalize.css'
import 'spectre.css/dist/spectre.css'
import 'spectre.css/dist/spectre-exp.css'
import 'spectre.css/dist/spectre-icons.css'

import 'luxbar/build/luxbar.min.css'
import 'font-awesome/css/font-awesome.css'
//import 'element-ui/lib/theme-chalk/index.css'

import Vue from 'vue'
import App from './App'
import axios from 'axios'
import router from './router'

Vue.prototype.$app = {
  name: "all-wing",
  context: "/allwing/service"
}

Vue.prototype.$http = axios
Vue.prototype.$axios = axios
Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {App}
})
