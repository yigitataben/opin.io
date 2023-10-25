<template>
  <v-card elevation="0" ref="form">
    <v-card-text>
      <v-autocomplete
          outlined
          ref="category"
          v-model="category"
          :rules="[() => !!category || 'This field is required']"
          :items="categories"
          label="Category"
          placeholder="Select..."
          required
      ></v-autocomplete>
      <v-text-field
          outlined
          ref="title"
          v-model="title"
          :rules="[
              () => !!title || 'This field is required',
              () => !!title && title.length <= 25 || 'Title must be less than 25 characters',
              titleCheck
            ]"
          label="Title"
          counter="25"
          required
      ></v-text-field>
      <v-textarea
          auto-grow
          outlined
          ref="post"
          v-model="post"
          :rules="[() => !!post || 'This field is required', titleCheck]"
          label="Post"
          required
      ></v-textarea>
    </v-card-text>
    <v-card-actions>
      <v-btn>
        Cancel
      </v-btn>
      <v-spacer></v-spacer>
      <v-slide-x-reverse-transition>
        <v-tooltip
            v-if="formHasErrors"
            left
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
                icon
                class="my-0"
                v-bind="attrs"
                @click="resetForm"
                v-on="on"
            >
              <v-icon>mdi-refresh</v-icon>
            </v-btn>
          </template>
          <span>Refresh form</span>
        </v-tooltip>
      </v-slide-x-reverse-transition>
      <v-btn
          color="primary"
          @click="submit"
      >
        Submit
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
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
    title: null,
    post: null,
    category: null,
    formHasErrors: false,
  }),

  computed: {
    form () {
      return {
        title: this.title,
        post: this.post,
        category: this.category,
      }
    },
  },

  methods: {
    titleCheck () {
      this.errorMessages = this.title
          ? `Hey! I'm required`
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
    submit () {
      this.formHasErrors = false

      Object.keys(this.form).forEach(f => {
        if (!this.form[f]) this.formHasErrors = true

        this.$refs[f].validate(true)
      })
    },
  },
}
</script>