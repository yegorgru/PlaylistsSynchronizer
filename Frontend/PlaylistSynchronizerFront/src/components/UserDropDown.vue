<script setup>
import axios from "axios";
</script>

<template>
    <div>
      <select v-model="selectedUserRole">
        <option value="USER">USER</option>
        <option value="ADMIN">ADMIN</option>
        <option value="SUPER ADMIN">SUPER ADMIN</option>
      </select>
      <button class="btn btn-primary col-lg-6 mb-3" @click="updateUserRole">Update</button>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      userId: {
        type: Number,
        required: true,
      },
      myId: {
        type: Number,
        required: true,
      },
      groupId: {
        type: Number,
        required: true,
      },
    },
    data() {
      return {
        selectedUserRole: 'USER',
      };
    },
    methods: {
      updateUserRole() {
        if(this.userId != this.myId) {
            const accessToken = localStorage.getItem("access_token");
            const roleData = {
                role: this.selectedUserRole,
            };
            const headers = {
                "Content-Type": "application/json",
                Authorization: "Bearer " + accessToken,
            };
            axios.put('http://localhost:8080/api/groups/' + this.groupId + "/users/" + this.userId, roleData, { headers })
                .then(response => {
                    console.log(response.data);
                    window.location.href = '/group_info/' + this.groupId;
                })
                .catch(error => {
                if(error.data.error.includes("api error: Request had invalid authentication")) {
                    localStorage.removeItem("access_token");
                    window.location.href = '/login';
                }
                console.error('Error updating user role:', error);
            });
        }
      },
    },
  };
  </script>