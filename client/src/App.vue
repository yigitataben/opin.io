<template>
  <v-app>
    <v-main>
      <TemplateView />
    </v-main>
  </v-app>
</template>

<script>
import TemplateView from "@/views/template/TemplateView.vue";

export default {
  components: {
    TemplateView,
  },
};
</script>
