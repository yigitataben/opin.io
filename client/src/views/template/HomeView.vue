<template>
  <v-app id="inspire">
    <v-app-bar color="light-blue lighten-5" flat="true">
      <v-container class="mx-auto d-flex align-center justify-center">
        <v-btn
            fixed
            icon
            size="52"
            @click="openPage(3)"
        >
          <v-avatar size="52">
            <v-img
                src="https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2080&q=80"
            ></v-img>
          </v-avatar>
        </v-btn>

        <v-spacer></v-spacer>

        <v-responsive max-width="160">
          <v-text-field
              density="compact"
              flat="true"
              hide-details
              label="Search..."
              rounded="lg"
              single-line
              variant="solo-filled"
              append-inner-icon="mdi-card-search"
          ></v-text-field>
        </v-responsive>
      </v-container>
    </v-app-bar>

    <v-main class="bg-grey-lighten-3">
      <v-container>
        <v-row>
          <v-col cols="2">
            <v-sheet rounded="lg">
              <v-list dense rounded="lg">
                <v-list-item
                    v-model="selectedItem"
                    v-for="(item, index) in items"
                    :key="index"
                    @click="openPage(index)"
                >
                  <template v-slot:prepend>
                    <v-icon>{{item.icon}}</v-icon>
                  </template>

                  <v-list-item-title>{{item.text}}</v-list-item-title>
                </v-list-item>

                <v-divider class="my-2"></v-divider>

                <v-list-item
                    color="grey-lighten-4"
                    link
                    title="Refresh Page"
                ></v-list-item>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col>
            <v-sheet
                min-height="70vh"
                rounded="lg"
            >
              <router-view/>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>

import router from "@/router";

export default {
  data () {
    return {
      selectedItem: null,
      items: [
        {text: 'Timeline', icon: 'mdi-timeline', url: '/'},
        {text: 'Explore', icon: 'mdi-pound-box', url: '/explore'},
        {text: 'Bookmarks', icon: 'mdi-bookmark', url: '/bookmarks'},
        {text: 'Profile', icon: 'mdi-account', url: '/profile'},
        {text: 'New Post', icon: 'mdi-new-box', url: '/post'},
      ],
    }
  },

  methods: {
    openPage(index) {
      console.log(index)
      router.push(this.items[index].url)
    }
  }
}
</script>