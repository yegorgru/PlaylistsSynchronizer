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

    function logout() {
        const accessToken = localStorage.getItem("access_token");
        const headers = {
            "Content-Type": "application/json",
            Authorization: "Bearer " + accessToken,
        };
        axios.post('http://localhost:8080/auth/logout/', {}, { headers })
            .then(response => {
            console.log(response.data);
            localStorage.removeItem("access_token");
            window.location.reload();
            })
            .catch(error => {
            if(error.data.error.includes("api error: Request had invalid authentication")) {
                localStorage.removeItem("access_token");
                window.location.href = '/login';
            }
            console.error('Error joining group:', error);
            });
    }
</script>

<template>
  <div id="app" class="container py-4">
    <div class="d-flex justify-content-end mb-3">
      <button class="btn btn-primary mb-3 bg-danger" @click="logout">Logout</button>
    </div>
    <div class="d-flex justify-content-center mb-3">
      <button class="btn btn-primary mb-3" @click="createGroup">Create Group</button>
    </div>
    <div class="list-group">
        <router-link
            v-for="(group, index) in groups"
            :key="index"
            class="list-group-item list-group-item-action"
            :to="{ path: '/group_info/' + group.id }"
            >
            <h5 class="mb-1">{{ group.name }}</h5>
            <p class="mb-1">{{ group.description }}</p>
        </router-link>
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