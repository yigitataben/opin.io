<template>
  <PostView></PostView>
    <v-container>
      <v-row align="center" justify="center">
        <v-col cols="12">
        <!--
        v-for="(variant, i) in variants"
        :key="i"
        -->
      <v-card
          class="mx-auto"
          max-width="auto"
          variant="elevated"
      >
        <v-card-item>
          <div>
            <div class="text-overline mb-1">@yigitataben</div>
            <div class="text-h6 mb-1">General Discussion & Helping</div>
            <div class="text-caption">Lorem ipsum dolor sit amet, consectetur adipiscing elit.
              Quisque eu lacus non sem auctor tristique eget eget orci. In et eros tellus.
              Integer pharetra dolor ac velit ornare lacinia sed non lectus.
              Maecenas molestie eget nisl eget viverra. Donec eu elit bibendum, lobortis quam sed, lobortis leo.
              Nulla at augue ultrices, pretium magna vel, aliquam enim. Praesent dui nisl, ultricies ac justo sed, congue condimentum dui.
              Suspendisse nibh eros, cursus vitae nunc ac, blandit egestas nibh. Aliquam et blandit metus.
              Sed mollis tellus id nisl lobortis porta. Nunc ipsum magna, molestie ut tortor quis, aliquet rhoncus libero.
              Curabitur fringilla at erat id dapibus. In bibendum velit ac tortor porttitor gravida sed in eros.
              Nulla facilisi. Integer condimentum ante ligula, at dignissim turpis iaculis sit amet. In non felis et mauris ultrices mattis.</div>
            <div class="text-overline mb-1">27/10/2023 18:00:50</div>
          </div>
        </v-card-item>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-col cols="auto">
            <v-btn density="default" icon="mdi-comment" size="32" color="light-blue lighten-5"></v-btn>
          </v-col>
          <v-col cols="auto">
            <v-btn density="default" icon="mdi-bookmark" size="32" color="light-blue lighten-5"></v-btn>
          </v-col>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</v-container>
</template>
<script setup>
import PostView from "@/views/PostView.vue";
</script>
