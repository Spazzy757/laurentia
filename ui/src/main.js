import '@babel/polyfill'
import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import axios from 'axios';
import VueAxios from 'vue-axios';
import VueCookie from 'vue-cookie';
import Moment from 'vue-moment';
// Tell Vue to use the plugin
Vue.use(VueAxios, axios);
Vue.use(VueCookie)
Vue.use(Moment);
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
