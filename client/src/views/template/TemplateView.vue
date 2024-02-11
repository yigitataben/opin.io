<template>
  <v-app id="inspire">
    <v-app-bar
        class="px-3"
        color="light-blue lighten-5"
        density="compact"
        flat="true"
    >
      <v-btn
          fixed
          icon
          size="32"
          @click="openPage(0)"
      >
        <v-avatar size="32">
          <v-img
              src="https://freepngimg.com/save/130245-letter-o-png-free-photo/1195x1195"
          ></v-img>
        </v-avatar>
      </v-btn>

      <v-spacer></v-spacer>

      <v-responsive
          centered
          max-width="300"
      >
        <v-text-field
            append-inner-icon="mdi-card-search"
            density="compact"
            flat="true"
            hide-details
            label="Search..."
            rounded="lg"
            single-line
            variant="solo-filled"
        ></v-text-field>
      </v-responsive>

      <v-spacer></v-spacer>

      <v-btn
          fixed
          icon
          size="32"
          @click="openPage(4)"
      >
        <v-avatar size="32">
          <v-img
              src="https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=2080&q=80"
          ></v-img>
        </v-avatar>
      </v-btn>
    </v-app-bar>

    <v-main class="bg-grey-lighten-3">
      <v-container>
        <v-row>
          <v-col
              cols="12"
              md="2"
          >
            <v-sheet
                min-height="268"
                rounded="lg"
            >
              <v-list dense rounded="lg">
                <v-list-item
                    v-for="(item, index) in items"
                    :key="index"
                    v-model="selectedItem"
                    @click="openPage(index)"
                >
                  <template v-slot:prepend>
                    <v-icon>{{ item.icon }}</v-icon>
                  </template>

                  <v-list-item-title>{{ item.text }}</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col
              cols="12"
              md="8"
          >
            <v-sheet
                min-height="70vh"
                rounded="lg"
            >
              <router-view/>
            </v-sheet>
          </v-col>

          <v-col
              cols="12"
              md="2"
          >
            <v-sheet
                min-height="268"
                rounded="lg"
            >
              <router-view>
                <BookmarksView></BookmarksView>
              </router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import router from "@/router";
import BookmarksView from "@/views/BookmarksView.vue";

export default {
  components: {BookmarksView},
  data() {
    return {
      selectedItem: null,
      items: [
        {text: 'Home', icon: 'mdi-home', url: 'content'},
        {text: 'Explore', icon: 'mdi-pound-box', url: 'explore'},
        {text: 'Bookmarks', icon: 'mdi-bookmark', url: 'bookmarks'},
        {text: 'Comments', icon: 'mdi-comment', url: 'comments'},
        {text: 'Profile', icon: 'mdi-account', url: 'profile'},
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
