<script setup>
  import { ref } from "vue";

  const groups = ref([
    { name: "Group 1", description: "Description for Group 1" },
    { name: "Group 2", description: "Description for Group 2" },
    // Add more groups here if needed
  ]);

  function createGroup() {
    // Logic to create a new group
    // You can add new groups by modifying the 'groups' array
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
