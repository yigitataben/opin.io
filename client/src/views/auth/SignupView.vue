<template>
  <div>
    <v-img class="mx-auto my-6" max-width="228"></v-img>

    <v-card class="mx-auto pa-12 pb-8" elevation="8" max-width="448" rounded="lg">
      <div class="text-subtitle-1 text-medium-emphasis">Email Address</div>

      <v-text-field
          v-model="email"
          :rules="[rules.required, rules.email]"
          density="compact"
          placeholder="Email Address"
          prepend-inner-icon="mdi-email-outline"
          variant="outlined"
      ></v-text-field>

      <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
        Password
      </div>

      <v-text-field
          v-model="password"
          :append-inner-icon="visible ? 'mdi-eye-off' : 'mdi-eye'"
          :rules="[rules.password, rules.length(6)]"
          :type="visible ? 'text' : 'password'"
          density="compact"
          placeholder="Enter your password."
          prepend-inner-icon="mdi-lock-outline"
          variant="outlined"
          @click:append-inner="visible = !visible"
      ></v-text-field>

      <v-card class="mb-12" color="surface-variant" variant="tonal"></v-card>

      <v-btn block class="mb-8" color="blue" size="large" variant="tonal" @click="signIn">
        Sign Up
      </v-btn>

      <v-card-text class="text-center">
        <v-btn
            color="blue"
            variant="plain"
            @click="redirectToLogIn"
        >
          Already have an account?
        </v-btn>
      </v-card-text>
    </v-card>
  </div>
</template>
<script>

import axios from "axios";

export default {
  data: () => ({
    visible: false,
    email: '',
    rules: {
      required: value => !!value || 'Required.',
      email: value => {
        const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        return pattern.test(value) || 'Invalid e-mail.'
      },
      length: len => v => (v || '').length >= len || `Invalid character length, required ${len}`,
      password: v =>
          !!(v || '').match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*(_|[^\w])).+$/) ||
          'Password must contain an upper case letter, a numeric character, and a special character',
    },
    password: '',
  }),
  methods: {
    redirectToLogIn() {
      this.$router.push('/login');
    },
    async signIn() {
      const userData = {
        email: this.email,
        password: this.password,
      };

      try {
        const response = await axios.post('http://127.0.0.1:3000/signup', userData);
        const data = response.data;

        console.log("Response from server:", data);

        if (data) {
          console.log("Successful registration!");
          location.reload();
        } else {
          console.error("Registration failed:", "Unexpected response from server");
        }
      } catch (error) {
        console.error('Data sending error.', error.message);
      }
    }
  },
};
</script>
