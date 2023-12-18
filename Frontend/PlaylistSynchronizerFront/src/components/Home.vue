<script setup>
    import { ref, onMounted } from "vue";
    import axios from "axios";

    const groups = ref([]);

    onMounted(async () => {
    try {
        const accessToken = localStorage.getItem('access_token');
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': "Bearer " + accessToken,
        };

        const response = await axios.get("http://localhost:8080/api/groups", {headers});
        groups.value = response.data;
    } catch (error) {
        console.error("Error fetching groups:", error);
    }
    });

    function createGroup() {
    window.location.href = '/create_group';
    }
</script>

<template>
  <div id="app" class="container py-4">
    <div>
      <button class="btn btn-primary mb-3" @click="createGroup">Create Group</button>
    </div>
    <div class="list-group">
      <a
        v-for="(group, index) in groups"
        :key="index"
        class="list-group-item list-group-item-action"
      >
        <h5 class="mb-1">{{ group.name }}</h5>
        <p class="mb-1">{{ group.description }}</p>
      </a>
      <div v-if="groups.length === 0" class="text-muted">No groups available</div>
    </div>
  </div>
</template>

<style>
#app {
  min-width: 320px;
}
body {
  background-color: #fff78a !important;
}
</style>