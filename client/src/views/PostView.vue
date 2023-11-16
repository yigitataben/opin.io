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
          v-model="category"
          :rules="[() => !!category]"
          :items="categories"
          required
          variant="outlined"
          rounded="xl"
      ></v-autocomplete>

      <v-textarea
          variant="outlined"
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
  data: () => ({
    categories: [ 'General Discussion and Helping',
      'Artificial Intelligence Ethics and Law',
      'Business and Technology',
      'Cloud Technologies',
      'Data Science and Artificial Intelligence',
      'Data-Driven Applications',
      'Database Management',
      'DevOps and Continuous Integration/Continuous Deployment (CI/CD)',
      'Education and Learning',
      'Hardware and Electronics',
      'Internet and Network Technologies',
      'Internet of Things (IoT)',
      'Programming Languages',
      'Robotics and Automation',
      'Security and Hacking',
      'Software Development',
      'Technology News and Trends',
      'VR/AR and Augmented Reality (AR/VR/XR)',
      'Web Development',
      'Mobile Application Development' ],
    errorMessages: '',
    post: null,
    category: null,
    formHasErrors: false,
  }),

  computed: {
    form () {
      return {
        post: this.post,
        category: this.category,
      }
    },
  },
  methods: {
    postCheck () {
      this.errorMessages = this.post
          ? `Hey! I'm required...`
          : ''

      return true
    },
    resetForm () {
      this.errorMessages = []
      this.formHasErrors = false

      Object.keys(this.form).forEach(f => {
        this.$refs[f].reset()
      })
    },
    async share() {
      const formData = {
        PostBody: this.post,
        PostCategory: this.category
      };

      try {
        const response = await axios.post('http://127.0.0.1:8080/post', formData);
        console.log('Data sent successfully.', response.data);

        this.resetForm();
      } catch (error) {
        console.error('Data sending error.', error);
        this.formHasErrors = true;
      }
    }
  }
}
</script>