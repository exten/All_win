import Vue from 'vue'
import Router from 'vue-router'

import Login from '@/components/Login'
import Order from '@/components/Order'
import Register from '@/components/Register'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login
    },
    {
      path: '/Order',
      name: 'Order',
      component: Order
    },
    {
      path: '/Register',
      name: 'Register',
      component: Register
    }
  ]
})
