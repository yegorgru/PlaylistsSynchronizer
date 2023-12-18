<script setup>
    import { ref } from 'vue';
    import axios from 'axios';
    import { useRoute } from "vue-router";

    const route = useRoute();

    const spotifyUri = ref('');
    const ytMusicId = ref('');
    const name = ref('');

    const addTrack = async () => {
        const trackData = {
            spotifyUri: spotifyUri.value,
            youTubeMusicID: ytMusicId.value,
            name: name.value
        };

        const accessToken = localStorage.getItem('access_token');

        const headers = {
            'Content-Type': 'application/json',
            'Authorization': "Bearer " + accessToken,
        };

        try {
            const response = await axios.post("http://localhost:8080/api/playlists/" + route.params.playlist_id + "/tracks", trackData, {headers});
            if (response.status === 200) {
                console.log("Added track " + response.data.id);
                window.location.href = "/group_info/" + route.params.playlist_id;
            }
            else {
                alert("Error! Invalid data provided");
            }
        } catch (error) {
            console.error('Error adding track:', error);
        }
    };
</script>

<template>
    <div class="container mt-5">
      <div class="card">
        <div class="card-header">
          <h3>Add Track</h3>
        </div>
        <div class="card-body">
          <form @submit.prevent="addTrack">
            <div class="mb-3">
              <label for="name" class="form-label">Track Name</label>
              <input type="text" v-model="name" class="form-control" id="name" maxlength="50" required>
            </div>
            <div class="mb-3">
              <label for="spotifyUri" class="form-label">Spotify Uri</label>
              <input v-model="spotifyUri" class="form-control" id="spotifyUri" maxlength="50" required>
            </div>
            <div class="mb-3">
              <label for="ytMusicId" class="form-label">YouTube Music Id</label>
              <input type="text" v-model="ytMusicId" class="form-control" id="ytMusicId" maxlength="50" required>
            </div>
            <button type="submit" class="btn btn-primary">Add Track</button>
          </form>
        </div>
      </div>
    </div>
</template>
  
<style scoped>

</style>