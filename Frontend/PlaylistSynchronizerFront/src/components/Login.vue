<script setup>
</script>

<template>
<div class="container mt-5">
  <div class="row">
    <div class="col-md-8 mx-auto">
      <div class="card my-5 login">
        <div class="card-body">
          <h2 class="card-title text-center">Login</h2>
          <p class="card-text text-center">Choose a platform:</p>

          <div class="row">
            <div class="col-sm-6 mb-2">
              <button class="btn btn-success w-100" @click="loginWithPlatform('spotify')">
                <img src="@/assets/icons/spotify-icon.png" alt="Spotify Icon" class="icon" />
                Spotify
              </button>
            </div>

            <div class="col-sm-6 mb-2">
              <button class="btn btn-danger w-100" @click="loginWithPlatform('youtube-music')">
                <img src="@/assets/icons/yt-music-icon.png" alt="YouTube Music Icon" class="icon" />
                YouTube Music
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  methods: {
    loginWithPlatform(platform) {
      console.log(`Logging in with ${platform}`);
    }
  }
};
</script>

<style scoped>
    .icon {
        max-width: 40px;
        height: auto;
    }
    .login {
        background-color: #FFE382;
        border-radius: 10px;
        color: #9A3B3B;
    }
</style>
