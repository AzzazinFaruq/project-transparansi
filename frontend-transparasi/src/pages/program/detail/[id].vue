<template>
  <div class="" >
  <navbar-program></navbar-program>
  <div class="container-home" style="margin-top: 64px; margin-bottom: 64px;">
    <v-btn style="margin-left: -15px; margin-bottom: 20px;" variant="text" prepend-icon="mdi-chevron-left" href="/program">Kembali</v-btn>
    <div class="">
      <h2>{{ detail.nama_program }}</h2>
      <p class="mt-3">{{ detail.deskripsi }}</p>
      <p class="mt-5" style="text-transform: uppercase; "><v-icon>mdi-map-marker</v-icon> {{ detail.dusun }}, {{ desa }}, {{ kecamatan }}, {{ kabupaten }}</p>
    </div>
    <div class="mt-5" style="width:100%;max-width:500px" >
      
      <label for="" class="label-form">Foto Before</label>
      <img :src="`${getImageUrl(detail.foto_before)}`" alt="" width="100%">
      <label for="" class="label-form">Foto Progress</label>
      <img :src="`${getImageUrl(detail.foto_progress)}`" alt="" width="100%">
      <label for="" class="label-form">Foto After</label>
      <img :src="`${getImageUrl(detail.foto_after)}`" alt="" width="100%">
    </div>
    <v-form class="mt-5">
      <v-textarea variant="outlined" label="Keluhan" v-model="aduan.keluhan"></v-textarea>
      <div class="d-flex justify-end">
        <v-btn variant="tonal" color="white" style="background-color: #387144; text-transform: none;" append-icon="mdi-send" @click="simpanAduan">Kirim Tanggapan</v-btn>
      </div>
    </v-form>
  </div>
  <div class="">
    <div class="container-home" v-for="item in aduanList" :key="item.id">
      <div class="d-flex">
        <div class="mr-3" style="width: 50px;">
          <img v-if="item.user.foto_profil == ''" src="@/assets/profile.png" alt="" width="100%">
          <img v-else :src="`${getImageUrl(item.user.foto_profil)}`" alt="" width="100%">
        </div>
        <div class="ml-2" style="width: 100%;">
        <h3 class="mb-1">{{ item.user.username }}</h3>
        <p class="mb-1">{{ item.keluhan }}</p>
        <p class="mb-1">{{ item.created_at }}</p>
        <v-row class="mt-2 mb-5 " v-if="item.status == 'Sudah Ditanggapi'">
          <v-col  cols="1">
          </v-col>
          <v-col cols="11">
            <div class="d-flex">
              <div class="mr-3" style="width: 50px;">
                <img v-if="item.user_tanggapan.foto_profil == ''" src="@/assets/profile.png" alt="" width="100%">
                <img v-else :src="`${getImageUrl(item.user_tanggapan.foto_profil)}`" alt="" width="100%">
              </div>
              <div class="ml-2" style="width: 100%;">
              <h3 v-if="item.user_tanggapan.username == ''" class="mb-1">Anonim (DPRD)</h3> 
              <h3 v-else class="mb-1">{{ item.user_tanggapan.username }} (DPRD)</h3>
              <p class="mb-1">{{ item.tanggapan }}</p>
              <p class="mb-1">{{ item.updated_at }}</p>
              </div>
            </div>
          </v-col>
        </v-row>
        </div>
      </div>
    </div>
  </div>
  <footer-home></footer-home>
</div>
</template>
<script>
import axios from 'axios';
import navbarProgram from '@/components/navbar-program.vue';
import footerHome from '@/components/footer-home.vue';
import { getImageUrl } from "@/config/foto.js";
export default{
  components:{
    navbarProgram,
    footerHome
  },
  data() {
    return {
      getImageUrl,
      desa:'',
      kecamatan:'',
      kabupaten:'',
      detail: [],
      aduan:{
        program_id:0,
        user_id:0,
        keluhan:''
      },
      aduanList:[]
    }
  },
  mounted() {
    this.getDetail();
    this.getUser();
    this.getAduan();
  },
  methods: {
    getAduan(){
      axios.get(`/api/aduan/aduan-by-program/${this.$route.params.id}`)
      .then(res=>{
        this.aduanList = res.data.data;
      })
    },
    getUser(){
  axios.get("api/user")
  .then(res=>{
    this.aduan.user_id = parseInt(res.data.data.Id); // atau +res.data.data.Id
    this.aduan.program_id = parseInt(this.$route.params.id); // atau +this.$route.params.id
  })
},
    getDetail(){
      axios.get(`/api/program/${this.$route.params.id}`)
      .then(res=>{
        this.detail = res.data.data;
        this.desa = res.data.data.Desa.nama_desa;
        this.kecamatan = res.data.data.Kecamatan.nama_kecamatan;
        this.kabupaten = res.data.data.Kabupaten.nama_kabupaten
      })
    },
    simpanAduan(){
      axios.post("/api/create-aduan",this.aduan)
      .then(res=>{
        this.$router.go(0)
      })
    }
  },

}
</script>
