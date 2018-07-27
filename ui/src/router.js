import Vue from 'vue'
import Router from 'vue-router'
import Login from './components/Login.vue'
import MessageBoard from './components/MessageBoard.vue'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'MessageBoard',
            component: MessageBoard
        },
        {
            path: '/login',
            name: 'Login',
            component: Login
        },
    ]
})
