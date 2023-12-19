<script setup>
    import { useRoute } from "vue-router";
    import { ref, onMounted } from "vue";
    import axios from "axios";

    const route = useRoute();
    const groupName = ref('');
    const groupDescription = ref('');
    const playlistName = ref('');
    const playlistDescription = ref('');

    const amendGroup = async () => {
        const groupData = {
            name: groupName.value,
            description: groupDescription.value
        };

        const playlistData = {
            name: playlistName.value,
            description: playlistDescription.value
        };

        const accessToken = localStorage.getItem('access_token');

        const headers = {
            'Content-Type': 'application/json',
            'Authorization': "Bearer " + accessToken,
        };

        try {
            const response = await axios.put('http://localhost:8080/api/groups/' + route.params.group_id, groupData, {headers});
            if (response.status === 200) {
                console.log("Amended group " + response.data.id);
            }
            else {
                alert("Error!");
            }
        } catch (error) {
            console.error('Error amending group:', error);
        }

        try {
            const response = await axios.put('http://localhost:8080/api/playlists/' + route.params.group_id, playlistData, {headers});
            if (response.status === 200) {
                console.log("Amended playlist " + response.data.id);
                window.location.href = '/home';
            }
            else {
                alert("Error!");
            }
        } catch (error) {
            console.error('Error amending group:', error);
        }
    };
</script>

<template>
    <div class="container mt-5">
      <div class="card">
        <div class="card-header">
          <h3>Amend Group</h3>
        </div>
        <div class="card-body">
          <form @submit.prevent="amendGroup">
            <div class="mb-3">
              <label for="groupName" class="form-label">Group Name</label>
              <input type="text" v-model="groupName" class="form-control" id="groupName" maxlength="20" required>
            </div>
            <div class="mb-3">
              <label for="groupDescription" class="form-label">Group Description</label>
              <textarea v-model="groupDescription" class="form-control" id="groupDescription" maxlength="50" required></textarea>
            </div>
            <div class="mb-3">
              <label for="playlistName" class="form-label">Playlist Name</label>
              <input type="text" v-model="playlistName" class="form-control" id="playlistName" maxlength="20" required>
            </div>
            <div class="mb-3">
              <label for="playlistDescription" class="form-label">Playlist Description</label>
              <textarea v-model="playlistDescription" class="form-control" id="playlistDescription" maxlength="50" required></textarea>
            </div>
            <button type="submit" class="btn btn-primary">Amend Group</button>
          </form>
        </div>
      </div>
    </div>
</template>
  
<style scoped>

</style>