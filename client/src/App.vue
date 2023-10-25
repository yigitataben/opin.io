<template>
  <v-app>
    <v-main>
      <HomeView />
    </v-main>
  </v-app>
</template>

<script>
import HomeView from 'C:/Users/yigit/Playground/opin.io/client/src/views/template/HomeView.vue';

export default {
  components: {
    HomeView,
  },
};
</script>

<!--<template>
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </nav>
  <router-view/>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>-->
