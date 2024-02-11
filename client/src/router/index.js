import {createRouter, createWebHashHistory} from 'vue-router';
import BookmarksView from "@/views/BookmarksView.vue";
import ExploreView from "@/views/ExploreView.vue";
import PostView from "@/views/PostView.vue";
import ProfileView from "@/views/ProfileView.vue";
import ContentView from "@/views/ContentView.vue";
import CommentsView from "@/views/CommentsView.vue";
import TemplateView from "@/views/template/TemplateView.vue";
import LoginView from "@/views/authentication/LogInView.vue";
import SignupView from "@/views/authentication/SignInView.vue";

const routes = [
  // Auth Routes:
  {
    path: '/',
    redirect: '/log-in',
  },
  {
    path: '/log-in',
    name: 'LogIn',
    component: LoginView,
  },
  {
    path: '/sign-in',
    name: 'SignIn',
    component: SignupView,
  },

  // Dashboard Routes:
  {
    path: '/homepage',
    name: 'Homepage',
    component: TemplateView,
    meta: {requiresAuth: true},
    children: [
      {
        path: 'content',
        name: 'Content',
        component: ContentView,
      },
      {
        path: 'bookmarks',
        name: 'Bookmarks',
        component: BookmarksView,
      },
      {
        path: 'explore',
        name: 'Explore',
        component: ExploreView,
      },
      {
        path: 'post',
        name: 'Post',
        component: PostView,
      },
      {
        path: 'profile',
        name: 'Profile',
        component: ProfileView,
      },
      {
        path: 'comments',
        name: 'Comments',
        component: CommentsView,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
