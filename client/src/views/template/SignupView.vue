<template>
  <div>
    <v-img class="mx-auto my-6" max-width="228"></v-img>

    <v-card class="mx-auto pa-12 pb-8" elevation="8" max-width="448" rounded="lg">
      <div class="text-subtitle-1 text-medium-emphasis">Personal Informations</div>

      <v-text-field
          v-model="first_name"
          :rules="[rules.required, rules.user_name]"
          density="compact"
          placeholder="First Name"
          prepend-inner-icon="mdi-card-account-details-outline"
          variant="outlined"
      ></v-text-field>

      <v-text-field
          v-model="last_name"
          :rules="[rules.required, rules.user_name]"
          density="compact"
          placeholder="Last Name"
          prepend-inner-icon="mdi-card-account-details-outline"
          variant="outlined"
      ></v-text-field>

      <v-text-field
          v-model="user_name"
          :rules="[rules.required, rules.user_name]"
          density="compact"
          placeholder="Username"
          prepend-inner-icon="mdi-at"
          variant="outlined"
      ></v-text-field>

      <v-text-field
          v-model="email_address"
          :rules="[rules.required, rules.email_address]"
          density="compact"
          placeholder="E-mail Address"
          prepend-inner-icon="mdi-email-outline"
          variant="outlined"
      ></v-text-field>

      <div class="text-subtitle-1 text-medium-emphasis d-flex align-center justify-space-between">
        Password
      </div>

      <v-text-field
          v-model="user_password"
          :rules="[rules.user_password, rules.length(6)]"
          :append-inner-icon="visible ? 'mdi-eye-off' : 'mdi-eye'"
          :type="visible ? 'text' : 'password'"
          density="compact"
          placeholder="Enter your password."
          prepend-inner-icon="mdi-lock-outline"
          variant="outlined"
          @click:append-inner="visible = !visible"
      ></v-text-field>

      <v-card class="mb-12" color="surface-variant" variant="tonal"></v-card>

      <v-btn block class="mb-8" color="blue" size="large" variant="tonal" @click="signup">
        Sign Up
      </v-btn>
    </v-card>
  </div>
</template>

<script>

import axios from "axios";

export default {
  data: () => ({
    visible: false,
    email_address: '',
    rules: {
      required: value => !!value || 'Required.',
      user_name: value => /^[a-zA-Z\s]+$/.test(value) || 'Invalid character.',
      email_address: value => {
        const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        return pattern.test(value) || 'Invalid e-mail.'
      },
      length: len => v => (v || '').length >= len || `Invalid character length, required ${len}`,
      user_password: v =>
          !!(v || '').match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*(_|[^\w])).+$/) ||
          'Password must contain an upper case letter, a numeric character, and a special character',
    },
    first_name: '',
    last_name: '',
    user_name: '',
    user_password: '',
  }),
  methods: {
    async signup() {
      const userData = {
        user_name: this.user_name,
        first_name: this.first_name,
        last_name: this.last_name,
        email_address: this.email_address,
        user_password: this.user_password,
      };

      try {
        const response = await axios.post('http://127.0.0.1:8080/user', userData);
        const data = response.data;

        if (data.success) {
          console.log("Successful registration!");
          window.location.reload();
        } else {
          console.error("Registration failed:", data.message);
        }
      } catch (error) {
        console.error('Data sending error.', error.message);
      }
    }
    },
  };
</script>
