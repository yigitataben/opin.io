import { createRouter, createWebHashHistory } from 'vue-router'

import BookmarksView from "@/views/BookmarksView.vue";
import ExploreView from "@/views/ExploreView.vue";
import LoginView from "@/views/LoginView.vue";
import PostView from "@/views/PostView.vue";
import ProfileView from "@/views/ProfileView.vue";
import ContentView from "@/views/ContentView.vue";
import CommentsView from "@/views/CommentsView.vue";

const routes = [
  {
    path: '/',
    name: 'Content',
    component: ContentView
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
  },
  {
    path: '/comments',
    name: 'Comments',
    component: CommentsView
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
