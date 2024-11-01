<template>
<div class="home-wrapper">
    <navbar-home />
    <img src="./img/dprd.png" alt="" style="width: 100%; display: block; margin-top: -80px; filter: brightness(50%);">
    <div class="wrapper-title text-center">
      <h1 class="title">Selamat Datang</h1>
      <h1>Website Transparansi oleh POKIR/JASMAS DPRD Jawa Timur</h1>
    </div>
</div>
<v-container class="" style="padding:16px !important;">
  <div class="section-about">
    <v-row>
  <v-col cols="12" md="4" >
    <h1 class="text-left">Latar Belakang</h1>
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
  <div class=" section-intro tutorial">
    <v-container>
      <div class="langkah">
        <v-row>
          <v-col class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/icon/account.png" alt="" style="width: 128px; height: auto;">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
          <v-col class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/icon/pencil.png" alt="" style="width: 128px; height: auto;">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
          <v-col class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/icon/list-box.png" alt="" style="width: 128px; height: auto;">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
          <v-col class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/icon/send.png" alt="" style="width: 128px; height: auto;">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
        </v-row>
      </div>
    </v-container>
  </div>
  <div class="" style="margin-bottom: 128px;">
    <v-container>
      <v-row>
          <v-col  v-for="(item,index) in program.slice(0,4)" :key="item.id" cols="3">
          <div v-if="index <3 " class="border-lg card-program pa-4">
            <div class="image-container">
              <img :src="`${getImageUrl(item.foto_before)}`" alt="" class="program-image">
            </div>
            <h2 class="mt-5">{{ item.nama_program }}</h2>
            <p class="description-text mt-3 mb-6">{{ item.deskripsi }}</p>
            <a @click="showMore(item.id)" style="color: #BF3232; text-decoration: none;">Lihat Selengkapnya <v-icon>mdi-chevron-right</v-icon></a>
          </div>
          <div v-else class="border-lg pa-4 card-program d-flex justify-center align-center">
            <a href="/program">
              <h2 style="color: #BF3232;">Lihat <br>Lebih <br>Banyak</h2>
            </a>
          </div>
        </v-col>
      </v-row>
    </v-container>
  </div>


<footerHomeVue></footerHomeVue>

</template>

<script>
import footerHomeVue from "@/components/footer-home.vue";
import NavbarHome from "@/components/navbar-home.vue";
import axios from "axios";
import { BASE_URL, getImageUrl } from "@/config/foto.js";
export default{
  components:{
    footerHomeVue,
    NavbarHome
  },
  data() {
    return {
      getImageUrl,
      program:[]
    }
  },
  mounted(){
    this.getprogram()
  },
  methods: {
    getprogram(){
      axios.get("/program")
      .then(res=>{
        this.program = res.data.data
      })
    },
    showMore(id){
      this.$router.push(`/program/detail/${id}`)
    }
  },
}
</script>
<style scooped lang="scss">
.section-about{
  margin-top: 64px;
}
.section-intro{
  background-color: #E8E8E8;
  margin-top: 128px;
  margin-bottom: 128px;
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

  .langkah{
    margin-top: 128px;
    margin-bottom: 128px;
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
  position: relative;
  .wrapper-title{
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.8);
  color: white;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    color: white;
    .title{
      font-size: 96px;
      font-weight: bold;
    }
  }
}



.card-program{
  height: 100%;
  .description-text {
  display: -webkit-box;
  -webkit-line-clamp: 6; /* Jumlah baris yang ingin ditampilkan */
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
  .image-container{
    width: 100%;
  height: 240px; /* Sesuaikan tinggi yang diinginkan */
  overflow: hidden;
  position: relative;
  .program-image{
    width: 100%;
  height: 100%;
  object-fit: cover; /* Ini akan membuat gambar mengisi container dan terpotong jika berlebih */
  object-position: center; /* Mengatur posisi gambar di tengah */
  }

}
}

</style>
