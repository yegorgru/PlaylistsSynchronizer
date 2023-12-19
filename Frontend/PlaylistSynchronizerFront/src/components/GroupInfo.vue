<script setup>
import { useRoute } from "vue-router";
import { ref, onMounted } from "vue";
import axios from "axios";
import UserDropdown from "./UserDropDown.vue";

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

    const response3 = await axios.get("http://localhost:8080/api/users/me", { headers });
    userMeInfo.value = response3.data;
    console.log(userMeInfo.value);

    try {
      const response4 = await axios.get(
        "http://localhost:8080/api/groups/" +
          route.params.group_id +
          "/users/" +
          userMeInfo.value.id,
        { headers }
      );
      userInfo.value = response4.data;
      console.log(userInfo.value);
    } catch (err) {
      userInfo.value = { roleName: "" };
    }
  } catch (error) {
    console.error("Error fetching groups:", error);
  }
});

function addTrack() {
  window.location.href = "/add_track/" + groupInfo.value.playListID;
}

function amendGroup() {
  window.location.href = "/amend_group/" + route.params.group_id;
}

function joinGroup() {
  const accessToken = localStorage.getItem("access_token");
  const headers = {
    "Content-Type": "application/json",
    Authorization: "Bearer " + accessToken,
  };
  axios
    .post(
      "http://localhost:8080/api/groups/" + route.params.group_id + "/users",
      {},
      { headers }
    )
    .then((response) => {
      console.log(response.data);
      window.location.reload();
    })
    .catch((error) => {
      if (error.data.error.includes("api error: Request had invalid authentication")) {
        localStorage.removeItem("access_token");
        window.location.href = "/login";
      }
      console.error("Error joining group:", error);
    });
}

function kickUser(userId) {
  const accessToken = localStorage.getItem("access_token");
  const headers = {
    "Content-Type": "application/json",
    Authorization: "Bearer " + accessToken,
  };
  axios
    .delete(
      "http://localhost:8080/api/groups/" + route.params.group_id + "/users/" + userId,
      { headers }
    )
    .then((response) => {
      console.log(response.data);
      window.location.reload();
    })
    .catch((error) => {
      if (error.data.error.includes("api error: Request had invalid authentication")) {
        localStorage.removeItem("access_token");
        window.location.href = "/login";
      }
      console.error("Error kicking user:", error);
    });
}

function deleteTrack(songid) {
  const accessToken = localStorage.getItem("access_token");
  const headers = {
    "Content-Type": "application/json",
    Authorization: "Bearer " + accessToken,
  };
  axios
    .delete(
      "http://localhost:8080/api/playlists/" +
        route.params.group_id +
        "/tracks/" +
        songid,
      { headers }
    )
    .then((response) => {
      console.log(response.data);
      window.location.reload();
    })
    .catch((error) => {
      if (error.data.error.includes("api error: Request had invalid authentication")) {
        localStorage.removeItem("access_token");
        window.location.href = "/login";
      }
      console.error("Error deleting track:", error);
    });
}

function leaveGroup() {
  const accessToken = localStorage.getItem("access_token");
  const headers = {
    "Content-Type": "application/json",
    Authorization: "Bearer " + accessToken,
  };
  axios
    .post(
      "http://localhost:8080/api/groups/" + route.params.group_id + "/leave",
      {},
      { headers }
    )
    .then((response) => {
      console.log(response.data);
      window.location.reload();
    })
    .catch((error) => {
      if (error.data.error.includes("api error: Request had invalid authentication")) {
        localStorage.removeItem("access_token");
        window.location.href = "/login";
      }
      console.error("Error leaving group:", error);
    });
}

function deleteGroup() {
  const accessToken = localStorage.getItem("access_token");
  const headers = {
    "Content-Type": "application/json",
    Authorization: "Bearer " + accessToken,
  };
  axios
    .delete("http://localhost:8080/api/groups/" + route.params.group_id, { headers })
    .then((response) => {
      console.log(response.data);
      window.location.href = "/home";
    })
    .catch((error) => {
      if (error.data.error.includes("api error: Request had invalid authentication")) {
        localStorage.removeItem("access_token");
        window.location.href = "/login";
      }
      console.error("Error deleting group:", error);
    });
}
</script>

<template>
  <div class="group-info container py-4">
    <div class="group-info container py-4 info">
      <h2>Group Information</h2>

      <div class="card">
        <div class="card-body">
          <div class="info-section">
            <h3>Group Name</h3>
            <p>{{ groupInfo.name }}</p>
          </div>
          <div class="info-section">
            <h3>Group Description</h3>
            <p>{{ groupInfo.description }}</p>
          </div>
          <div class="info-section">
            <h3>Playlist Name</h3>
            <p>{{ playlistInfo.name }}</p>
          </div>
          <div class="info-section">
            <h3>Playlist Description</h3>
            <p>{{ playlistInfo.description }}</p>
          </div>
          <div class="info-section">
            <h3>User's Role in Group</h3>
            <p>{{ userInfo.roleName }}</p>
          </div>
        </div>
      </div>
    </div>

    <div v-if="!userInfo.roleName" class="d-flex justify-content-center p-3">
      <button class="btn btn-primary mb-3" @click="joinGroup">Join Group</button>
    </div>
    <div v-else-if="userInfo.roleName === 'SUPER ADMIN'" class="d-flex justify-content-center p-3">
      <button class="btn btn-primary m-2 mt-2 bg-danger" @click="deleteGroup">Delete Group</button>
      <button class="btn btn-primary m-2 mt-2" @click="amendGroup">Amend Group</button>
    </div>
    <div v-else-if="userInfo.roleName === 'ADMIN'" class="d-flex justify-content-center p-3">
      <button class="btn btn-primary mb-3" @click="amendGroup">Amend Group</button>
    </div>
    <div v-else class="d-flex justify-content-center p-3">
      <button class="btn btn-primary mb-3" @click="leaveGroup">Leave Group</button>
    </div>

    <!-- List of Songs -->
    <div class="info-section">
      <h3>List of Songs</h3>
      <div v-if="userInfo.roleName">
        <button class="btn btn-primary mb-3" @click="addTrack">Add Track</button>
      </div>
      <ul class="list-group">
        <li
          class="list-group-item"
          v-for="(song, index) in groupInfo.tracks"
          :key="index"
        >
          <div class="d-flex justify-content-between">
            <div class="col-6">{{ song.name }}</div>
            <div v-if="userInfo.roleName === 'ADMIN'">
              <button
                class="btn btn-primary col-6 mb-3 w-100 bg-danger"
                @click="deleteTrack(song.id)"
              >
                Delete Track
              </button>
            </div>
            <div v-if="userInfo.roleName === 'SUPER ADMIN'">
              <button
                class="btn btn-primary col-6 mb-3 w-100 bg-danger"
                @click="deleteTrack(song.id)"
              >
                Delete Track
              </button>
            </div>
          </div>
        </li>
      </ul>
      <div v-if="!groupInfo.tracks">There are no tracks</div>
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
          <div v-if="userInfo.roleName === 'SUPER ADMIN'" class="d-flex justify-content-center">
            <div v-if="user.roleName === 'USER'" class="w-25">
              <button class="btn btn-primary col-lg-6 mb-3 bg-danger" @click="kickUser(user.id)">
                Kick User
              </button>
            </div>
            <div v-if="user.id !== userMeInfo.id" class="w-75">
              <UserDropdown
                :userId="user.id"
                :myId="userMeInfo.id"
                :groupId="Number(route.params.group_id)"
              ></UserDropdown>
            </div>
          </div>
          <div v-if="userInfo.roleName === 'ADMIN'">
            <div v-if="user.roleName === 'USER'" class="w-25">
              <button class="btn btn-primary col-lg-6 mb-3" @click="kickUser(user.id)">
                Kick User
              </button>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  components: {
    UserDropdown,
  },
};
</script>

<style scoped>
    .info {
        background-color: #FFE382;
        border-radius: 10px;
        color: #9A3B3B;
    }
</style>
