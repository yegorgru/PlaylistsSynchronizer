<script setup>
    import { ref } from 'vue';
    import axios from 'axios';

    const groupName = ref('');
    const groupDescription = ref('');
    const playlistName = ref('');
    const playlistDescription = ref('');

    const createGroup = async () => {
        const groupData = {
            groupName: groupName.value,
            groupDescription: groupDescription.value,
            playListName: playlistName.value,
            playListDescription: playlistDescription.value,
        };

        const accessToken = localStorage.getItem('access_token');

        const headers = {
            'Content-Type': 'application/json',
            'Authorization': "Bearer " + accessToken,
        };

        try {
            const response = await axios.post('http://localhost:8080/api/groups', groupData, {headers});
            if (response.status === 200) {
                console.log("Created group " + response.data.id);
                window.location.href = '/home';
            }
            else {
                alert("Error!");
            }
        } catch (error) {
            console.error('Error creating group:', error);
        }
    };
</script>

<template>
    <div class="container mt-5">
      <div class="card">
        <div class="card-header">
          <h3>Create Group</h3>
        </div>
        <div class="card-body">
          <form @submit.prevent="createGroup">
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
            <button type="submit" class="btn btn-primary">Create Group</button>
          </form>
        </div>
      </div>
    </div>
</template>
  
<style scoped>

</style>