<template>
  <div class="">
<div class="home-wrapper">
    <navbar-home />
    <div class="image-dprd-container">
    <img src="./img/dprd.png" alt=""  class="image-dprd" >
  </div>
    <div class="wrapper-title text-center ">
      <p class="title ">Selamat Datang</p>
      <p class="caption ">Website Transparansi oleh POKIR/JASMAS DPRD Jawa Timur</p>
    </div>
</div>
<v-container class="" style="padding:16px !important;">
  <div class="section-about">
    <v-row>
  <v-col cols="12" md="4" >
    <h1 class="title">Latar Belakang</h1>
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
          <v-col cols="6" sm="3" class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/icon/account.png" class="icon-langkah" alt="">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
          <v-col cols="6" sm="3" class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/icon/pencil.png" class="icon-langkah" alt="">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
          <v-col cols="6" sm="3" class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
                <img src="./img/icon/list-box.png" class="icon-langkah" alt="">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
          <v-col cols="6" sm="3" class="d-flex flex-column align-center">
            <div class="image d-flex justify-center align-center">
              <img src="./img/icon/send.png" class="icon-langkah" alt="">
            </div>
            <p class="text-center mt-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor</p>
          </v-col>
        </v-row>
      </div>
    </v-container>
  </div>
  <div class="container-custom" style="margin-bottom: 128px;">
      <v-row>
          <v-col  v-for="(item,index) in program.slice(0,4)" :key="item.id" cols="12" sm="6" md="3">
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
  </div>
<footerHomeVue></footerHomeVue>
</div>
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
<style scoped lang="scss">

.section-about{
  margin-top: 64px;
  .title{
    font-size: 32px !important;
    @media (max-width:500px) {
      font-size: 24px !important;
    }
  }

}
.section-intro{
  background-color: #E8E8E8;
  margin-top: 128px;
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
    @media (max-width:500px) {
      margin-top: 64px;
      margin-bottom: 64px;
      }
    margin-top: 128px;
    margin-bottom: 128px;
    .image{
      width: 240px;
      height: 240px;
      @media (max-width:1024px) {
        width: 140px;
        height: 140px;
      }
      @media (max-width:500px) {
        width: 90px;
        height: 90px;
      }
      background-color: #BF3232;
      border-radius: 50%;
      .icon-langkah{
        width: 128px;
        height: auto;
        @media (max-width:1024px) {
        width: 60px;
        height: auto;
      }
        @media (max-width:500px) {
        width: 40px;
        height: auto;
      }
      }
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

  .image-dprd-container{
    width: 100%;
    height: 800px;
    @media (max-width:500px) {
      height: 650px;
      }
    margin-top: -80px;
    filter: brightness(50%);
    overflow: hidden;
    position: relative;
    .image-dprd{
      width: 100%;
      height: 100%;
      object-fit: cover; /* Ini akan membuat gambar mengisi container dan terpotong jika berlebih */
      object-position: center; /* Mengatur posisi gambar di tengah */
    }
  }

  .wrapper-title{
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.8);
    width: 90%;
    color: white;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    color: white;
    .title{
      font-size: 100px !important;
      font-weight: bold;
      @media (max-width:1024px) {
        font-size: 80px !important;
      }
      @media (max-width:500px) {
        font-size: 54px !important;
      }
    }
    .caption{
      font-size: 40px !important;
      @media (max-width:1024px) {
        font-size: 18px !important;
      }
    }
  }
}

.v-container{
  padding: 0 !important;
}

.container-custom{
  padding: 16px !important;
  margin-bottom: 64px;
  margin-top: 64px;
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
