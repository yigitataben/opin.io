import { createRouter, createWebHashHistory } from 'vue-router'

import BookmarksView from "@/views/BookmarksView.vue";
import ExploreView from "@/views/ExploreView.vue";
import LoginView from "@/views/LoginView.vue";
import PostView from "@/views/PostView.vue";
import ProfileView from "@/views/ProfileView.vue";
import TimelineView from "@/views/TimelineView.vue";

const routes = [
  {
    path: '/',
    name: 'home',
    component: TimelineView
  },
  {
    path: '/bookmarks',
    name: 'bookmarks',
    component: BookmarksView
  },
  {
    path: '/explore',
    name: 'explore',
    component: ExploreView
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView
  },
  {
    path: '/post',
    name: 'post',
    component: PostView
  },
  {
    path: '/profile',
    name: 'profile',
    component: ProfileView
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router

/*
component: () => import(/!* webpackChunkName: "about" *!/ '../views/BookmarksView.vue')
*/
