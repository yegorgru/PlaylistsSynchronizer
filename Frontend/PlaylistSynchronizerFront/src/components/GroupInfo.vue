<script setup>
import { useRoute } from "vue-router";
import { ref, onMounted } from "vue";
import axios from "axios";

const route = useRoute();
const groupInfo = ref([]);
const playlistInfo = ref([]);
const userMeInfo = ref([]);
const userInfo = ref([]);

onMounted(async () => {
  try {
    const accessToken = localStorage.getItem("access_token");
    const headers = {
      "Content-Type": "application/json",
      Authorization: "Bearer " + accessToken,
    };

    const response = await axios.get(
      "http://localhost:8080/api/groups/" + route.params.group_id,
      { headers }
    );
    groupInfo.value = response.data;
    console.log(groupInfo.value);

    const response2 = await axios.get(
      "http://localhost:8080/api/playlists/" + groupInfo.value.playListID,
      { headers }
    );
    playlistInfo.value = response2.data;
    console.log(playlistInfo.value);

    const response3 = await axios.get(
      "http://localhost:8080/api/users/me",
      { headers }
    );
    userMeInfo.value = response3.data;
    console.log(userMeInfo.value);

    const response4 = await axios.get(
      "http://localhost:8080/api/groups/" + route.params.group_id + "/users/" + userMeInfo.value.id,
      { headers }
    );
    userInfo.value = response4.data;
    console.log(userInfo.value);
  } catch (error) {
    console.error("Error fetching groups:", error);
  }
});
</script>

<template>
  <div class="group-info container py-4">
    <h2>Group Information</h2>

    <!-- Group Name -->
    <div class="info-section">
      <h3>Group Name</h3>
      <p>{{ groupInfo.name }}</p>
    </div>

    <!-- Group Description -->
    <div class="info-section">
      <h3>Group Description</h3>
      <p>{{ groupInfo.description }}</p>
    </div>

    <!-- Playlist Name -->
    <div class="info-section">
      <h3>Playlist Name</h3>
      <p>{{ playlistInfo.name }}</p>
    </div>

    <!-- Playlist Description -->
    <div class="info-section">
      <h3>Playlist Description</h3>
      <p>{{ playlistInfo.description }}</p>
    </div>

    <!-- User's Role in Group -->
    <div class="info-section">
      <h3>User's Role in Group</h3>
      <p>{{ userRole }}</p>
    </div>

    <!-- List of Songs -->
    <div class="info-section">
      <h3>List of Songs</h3>
      <ul class="list-group">
        <li
          class="list-group-item"
          v-for="(song, index) in groupInfo.tracks"
          :key="index"
        >
          <div class="row">
            <div class="col-md-12">{{ song.name }}</div>
          </div>
        </li>
      </ul>
      <div v-if="!groupInfo.tracks">
        There are no tracks
      </div>
    </div>

    <!-- List of Users -->
    <div class="info-section">
      <h3>List of Users</h3>
      <ul class="list-group">
        <li class="list-group-item" v-for="(user, index) in groupInfo.users" :key="index">
          <div class="row">
            <div class="col-md-4">{{ user.username }}</div>
            <div class="col-md-4">{{ user.platform }}</div>
            <div class="col-md-4">{{ user.roleName }}</div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped></style>
