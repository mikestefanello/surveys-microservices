import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import SurveyForm from '../forms/Survey.vue'
import Survey from '../views/Survey.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/new',
    name: 'SurveyForm',
    component: SurveyForm
  },
  {
    path: '/survey/:id',
    name: 'Survey',
    component: Survey
  },
  {
    path: '*',
    redirect: '/'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
