import { createRouter, createWebHistory } from 'vue-router'

const routes = [
	{
		path: '/',
		name: 'Main',
		component: () => import('../views/Main.vue')
	},
	{
		path: '/login',
		name: 'Login',
		component: () => import('../views/Login.vue')
	},
	{
		path: '/registration',
		name: 'Registration',
		component: () => import('../views/Registration.vue')
	},
	{
		path: '/profile',
		name: 'Profile',
		component: () => import('../views/Profile.vue')
	},
	{
		path: '/main',
		name: 'Home',
		component: () => import('../views/Home.vue')
	},



	{
		path: '/:catchAll(.*)',
		component: () => import('../views/NotFound.vue')
	},
]

const router = createRouter({
	history: createWebHistory(process.env.BASE_URL),
	routes
})

export default router
