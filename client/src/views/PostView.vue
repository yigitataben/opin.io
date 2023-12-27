<template>
  <v-card
      max-width="550"
      class="mx-auto"
      variant="flat"
      rounded="xl"
  >
    <v-sheet class="pa-4 text-center">
      <v-autocomplete
          chips
          ref="category"
          placeholder="Please select a category..."
          v-model="category"
          :rules="[() => !!category]"
          :items="categories"
          required
          variant="outlined"
          rounded="xl"
      ></v-autocomplete>

      <v-textarea
          variant="outlined"
          placeholder="Please fill this area..."
          auto-grow
          full-width
          rows="1"
          hide-details
          rounded="xl"
          ref="post"
          v-model="post"
          :rules="[() => !!post, postCheck]"
          required
      ></v-textarea>
    </v-sheet>

    <div class="d-flex justify-space-between pa-4 pb-0">
      <v-btn-toggle
          variant="outlined"
          divided="true"
          rounded="xl"
      >
        <v-btn>
          <v-icon icon="mdi-image"></v-icon>
        </v-btn>

        <v-btn>
          <v-icon icon="mdi-gif"></v-icon>
        </v-btn>

        <v-btn>
          <v-icon icon="mdi-sticker-emoji"></v-icon>
        </v-btn>
      </v-btn-toggle>

      <v-btn-toggle
          multiple="true"
          variant="outlined"
          divided="true"
          rounded="xl"
          justify="center"
      >
        <v-btn>
          <v-icon icon="mdi-format-italic"></v-icon>
        </v-btn>

        <v-btn>
          <v-icon icon="mdi-format-bold"></v-icon>
        </v-btn>

        <v-btn>
          <v-icon icon="mdi-format-underline"></v-icon>
        </v-btn>
      </v-btn-toggle>

      <v-btn-toggle
          variant="outlined"
          rounded="xl"
          justify="center"
          color="success"
          @click="share"
      >
        <v-btn>
          <v-icon icon="mdi-share"></v-icon>
        </v-btn>
      </v-btn-toggle>
    </div>
  </v-card>
</template>
<script>

import axios from "axios";

export default {
  data() {
    return {
      categories: [],
      errorMessages: '',
      post: null,
      category: null,
      formHasErrors: false,
    };
  },

  computed: {
    form() {
      return {
        post: this.post,
        category: this.category,
      };
    },
  },

  async created() {
    await this.fetchCategories();
  },

  methods: {
    async fetchCategories() {
      try {
        const response = await axios.get('http://127.0.0.1:8080/categories');
        this.categories = response.data.map(category => category.category_name);
      } catch (error) {
        console.error('An error occurred while fetching the categories.', error);
      }
    },

    postCheck() {
      this.errorMessages = this.post
          ? (this.category ? '' : 'Hey! Select a category...')
          : 'Hey! Post is required...';

      return this.post && this.category;
    },

    resetForm() {
      this.errorMessages = '';
      this.formHasErrors = false;

      if (this.$refs.post) {
        this.$refs.post.reset();
      }
      if (this.$refs.category) {
        this.$refs.category.reset();
        this.category = null;
      }
    },

    async share() {
      const formData = {
        content: this.post,
        category_name: this.category,
        user_name: this.user_name,
        created_at: this.created_at,
      };

      try {
        const response = await axios.post('http://127.0.0.1:8080/post', formData);
        console.log('Data sent successfully.', response.data);

        this.resetForm();
        window.location.reload();
      } catch (error) {
        console.error('Data sending error.', error);
        this.formHasErrors = true;
      }
    },
  },
}
</script>