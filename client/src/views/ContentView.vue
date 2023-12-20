<template>
  <div>
    <PostView></PostView>
    <v-container fluid>
      <v-sheet v-scroll-y="onScroll" style="max-height: 400px; overflow-y: auto;">
        <v-col v-for="(post, index) in posts" :key="index">
          <v-card class="mx-auto" max-width="auto" variant="elevated">
            <v-card-item>
              <div>
                <div class="text-overline mb-1">In <v-btn outlined rounded elevation="0" size="x-small" color="light-blue lighten-5">@{{ post.user_name }}</v-btn> 's opinion:</div>
                <div class="text-h6 mb-1">{{ post.category_name }}</div>
                <div class="text-caption">{{ post.content }}</div>
                <div class="text-overline mb-1">{{ formatDate(post.created_at) }}</div>
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
      </v-sheet>
    </v-container>
  </div>
</template>

<script setup>
import PostView from "@/views/PostView.vue";
import { ref, onMounted } from "vue";
import axios from "axios";

const posts = ref([]);

const fetchDataFromDatabase = async () => {
  try {
    const response = await axios.get("http://127.0.0.1:8080/post");
    posts.value = response.data;
  } catch (error) {
    console.error("Error fetching data from the database:", error);
  }
};

const formatDate = (dateString) => {
  if (!dateString) return "Invalid Date";

  const date = new Date(dateString);

  if (isNaN(date.getTime())) return "Invalid Date";

  const options = {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "numeric",
    minute: "numeric",
    second: "numeric"
  };

  return date.toLocaleDateString('en-US', options);
};

onMounted(() => {
  fetchDataFromDatabase().then(() => {
    posts.value.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
  });
});

</script>
