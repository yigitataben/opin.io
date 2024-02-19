<template>
  <div>
    <v-img
        class="mx-auto my-6"
        max-width="228"
    ></v-img>

    <v-card
        class="mx-auto pa-12 pb-8"
        elevation="8"
        max-width="448"
        rounded="lg"
    >
      <div class="text-subtitle-1 text-medium-emphasis">Account</div>

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

        <a
            class="text-caption text-decoration-none text-blue"
            href="#"
            rel="noopener noreferrer"
            target="_blank"
        >
          Forgot login password?
        </a>
      </div>

      <v-text-field
          v-model="user_password"
          :append-inner-icon="visible ? 'mdi-eye-off' : 'mdi-eye'"
          :rules="[rules.user_password, rules.length(6)]"
          :type="visible ? 'text' : 'password'"
          density="compact"
          placeholder="Enter your password."
          prepend-inner-icon="mdi-lock-outline"
          variant="outlined"
          @click:append-inner="visible = !visible"
      ></v-text-field>

      <v-card
          class="mb-12"
          color="surface-variant"
          variant="tonal"
      >
        <v-card-text class="text-medium-emphasis text-caption">
          Warning: After 3 consecutive failed login attempts, you account will be temporarily locked for three hours. If
          you must login now, you can also click "Forgot login password?" below to reset the login password.
        </v-card-text>
      </v-card>

      <v-btn
          block
          class="mb-8"
          color="blue"
          size="large"
          variant="tonal"
          @click="logIn"
      >
        Log In
      </v-btn>

      <v-card-text class="text-center">
        <v-btn
            color="blue"
            variant="plain"
            @click="redirectToSignIn"
        >
          Sign in now!
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
    email_address: '',
    rules: {
      required: value => !!value || 'Required.',
      email_address: value => {
        const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        return pattern.test(value) || 'Invalid e-mail.'
      },
      length: len => v => (v || '').length >= len || `Invalid character length, required ${len}`,
      user_password: v =>
          !!(v || '').match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*(_|\W)).+$/) ||
          'Password must contain an upper case letter, a numeric character, and a special character',
    },
    user_password: '',
  }),
  methods: {
    redirectToSignIn() {
      this.$router.push('/sign-in');
    },
    redirectToHomePage() {
      this.$router.push('/homepage/content');
    },
    async logIn() {
      const userData = {
        email_address: this.email_address,
        user_password: this.user_password,
      };

      try {
        const response = await axios.post('http://127.0.0.1:8080/login', userData);
        const data = response.data;

        if (data.success) {
          console.log("Successful login!");
          this.redirectToHomePage();
        } else {
          console.error("Login failed:", data.message);
        }
      } catch (error) {
        console.error('Data sending error.', error.message);
      }
    },
  },
};
</script>
