<template>
<div class="home-wrapper">
    <navbar-home />
    <img src="./img/dprd.png" alt="" style="width: 100%; display: block; margin-top: -80px;">
</div>
<v-container class="" style="padding:16px !important;">
  <div class="section-about">
    <v-row>
  <v-col cols="12" md="4" >
    <h1 class="text-left">Profil Singkat</h1>
  </v-col>
  <v-col cols="12" md="4" >
    <p class="text-left ">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </v-col>
  <v-col cols="12" md="4" >
    <p class="text-left ">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
  </v-col>
</v-row>
  </div>
</v-container>
<div class="section-intro">
  <v-container class="content">
    <v-row>
      <v-col cols="12" md="6">
        <h1 class="text-left">Latar Belakang</h1>
        <p class="text-left caption">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
        <p class="text-left caption">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>
      </v-col>
      <v-col cols="12" md="6">
        <div class="d-flex align-center justify-center">
          <img src="./img/logo-dprd-1.png" class="align-self-center" alt="" width="70%">
        </div>
      </v-col>
    </v-row>
  </v-container>
  </div>
  <div class="tutorial">
    <v-container>
      <div class="title">
        <h1 class="text-left">Langkah - Langkah <br>Mengajukan Keluhan</h1>
      </div>
      <div class="langkah">
        <v-row>
          <v-col class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/dummy-icon.png" alt="" style="width: 128px; height: auto;">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
          <v-col class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/dummy-icon.png" alt="" style="width: 128px; height: auto;">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
          <v-col class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/dummy-icon.png" alt="" style="width: 128px; height: auto;">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
          <v-col class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/dummy-icon.png" alt="" style="width: 128px; height: auto;">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
        </v-row>
      </div>
    </v-container>
  </div>

<footerHomeVue></footerHomeVue>

</template>

<script>
import footerHomeVue from "@/components/footer-home.vue";
import NavbarHome from "@/components/navbar-home.vue";
export default{
  components:{
    footerHomeVue,
    NavbarHome
  }
}
</script>
<style scooped lang="scss">
.section-about{
  margin-top: 64px;
}
.section-intro{
  background-color: #E8E8E8;
  margin-top: 64px;
  padding: 16px;
  .content{
    margin-top: 64px;
    margin-bottom: 64px;
    padding: 16px !important;
    .caption{
      margin-top: 32px;
    }
  }
}
.tutorial{
  margin-top: 128px;
  margin-bottom: 128px;
  .langkah{
    margin-top: 64px;
    .image{
      width: 240px;
      height: 240px;
      background-color: #BF3232;
      border-radius: 50%;
    }
  }
}

.position-relative {
  position: relative;
}

.navbar-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 999;
}

.home-wrapper {
  margin: 0;
  padding: 0;
  line-height: 0;
}
</style>
