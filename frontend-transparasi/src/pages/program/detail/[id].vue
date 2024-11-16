<template>
  <div class="">
  <navbar-program></navbar-program>
  <div class="container-custom">
  <div class="container-home" style="margin-top: 64px; margin-bottom: 64px;">
    <v-btn style="margin-left: -15px; margin-bottom: 20px;" variant="text" prepend-icon="mdi-chevron-left" href="/program">Kembali</v-btn>
    <div class="">
      <h2>{{ detail.nama_program }}</h2>
      <p class="mt-3">{{ detail.deskripsi }}</p>
      <p class="mt-5" style="text-transform: uppercase; "><v-icon>mdi-map-marker</v-icon> {{ detail.dusun }}, {{ desa }}, {{ kecamatan }}, {{ kabupaten }}</p>
    </div>

    <div class="mt-5">
      <div v-if="!hasCoordinates" class="text-center pa-4" style="background: #f5f5f5; border-radius: 4px;">
        <p>Koordinat lokasi tidak tersedia</p>
      </div>
      <div v-else id="map" style="height: 400px; width: 100%; margin-bottom: 20px; border: 1px solid #ccc; border-radius: 4px;"></div>
    </div>

    <div class="mt-5" style="width:100%;max-width:500px" >

      <label for="" class="label-form">Foto Before</label>
      <img v-if="detail.foto_before == ''" src="../detail/dummy-image.jpg" alt="" width="100%">
      <img v-else :src="`${getImageUrl(detail.foto_before)}`" alt="" width="100%">
      <label for="" class="label-form">Foto Progress</label>
      <img v-if="detail.foto_progress == ''" src="../detail/dummy-image.jpg" alt="" width="100%">
      <img v-else :src="`${getImageUrl(detail.foto_progress)}`" alt="" width="100%">
      <label for="" class="label-form">Foto After</label>
      <img v-if="detail.foto_after == ''" src="../detail/dummy-image.jpg" alt="" width="100%">
      <img v-else :src="`${getImageUrl(detail.foto_after)}`" alt="" width="100%">
    </div>
    <div class="mt-5">
      <v-btn
        variant="tonal"
        color="white"
        style="background-color: #387144; text-transform: none;"
        @click="dialog = true"
      >
        Ajukan Keluhan
      </v-btn>
    </div>

    <!-- Dialog Form Keluhan -->
    <v-dialog
      v-model="dialog"
      width="500"
      @click:outside="closeDialog"
    >
      <v-card>
        <v-card-title class="text-h5 pa-4">
          Form Keluhan
        </v-card-title>

        <v-card-text>
          <v-form @submit.prevent="simpanAduan">
            <v-text-field
              v-model="aduan.alamat"
              label="Alamat"
              variant="outlined"
              class="mb-3"
              required
            ></v-text-field>

            <v-text-field
              v-model="aduan.no_hp"
              label="Nomor HP"
              variant="outlined"
              type="number"
              class="mb-3"
              required
            ></v-text-field>

            <v-textarea
              v-model="aduan.keluhan"
              label="Isi Keluhan"
              variant="outlined"
              required
            ></v-textarea>
          </v-form>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="grey-darken-1"
            variant="text"
            @click="closeDialog"
          >
            Batal
          </v-btn>
          <v-btn
            color="white"
            style="background-color: #387144"
            variant="tonal"
            @click="simpanAduan"
          >
            Kirim
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
  <div class="">
    <div class="container-home" v-for="item in aduanList" :key="item.id">
      <div class="d-flex">
        <div class="profile-container mt-1 mr-2"><a href="/profile">
          <img v-if="item.user.foto_profil == ''" src="../detail/profile.png"  alt="" class="profile-pictures" >
          <img v-else :src="`${getImageUrl(item.user.foto_profil)}`"  alt="" class="profile-pictures" >
        </a>
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
              <div class="profile-container mt-1 mr-2"><a href="/profile">
                <img v-if="item.user_tanggapan.foto_profil == ''" src="../detail/profile.png"  alt="" class="profile-pictures" >
                <img v-else :src="`${getImageUrl(item.user_tanggapan.foto_profil)}`"  alt="" class="profile-pictures" >
              </a>
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
  </div>
  <footer-home></footer-home>
</div>
</template>
<script>
import axios from 'axios';
import navbarProgram from '@/components/navbar-program.vue';
import footerHome from '@/components/footer-home.vue';
import { getImageUrl } from "@/config/foto.js";
import Swal from 'sweetalert2';
import 'leaflet/dist/leaflet.css';

// Ganti cara import Leaflet
import * as L from 'leaflet';

export default {
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
      latitude: null,
      longitude: null,
      map: null,
      aduan:{
        program_id:0,
        user_id:0,
        alamat: '',
        no_hp: '',
        keluhan:''
      },
      aduanList:[],
      dialog: false,
      hasCoordinates: false,
    }
  },
  async mounted() {
    await this.getDetail();
    this.getUser();
    this.getAduan();
    this.$nextTick(() => {
      this.initMap();
    });
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
    async getDetail(){
      try {
        const res = await axios.get(`/api/program/${this.$route.params.id}`);
        this.detail = res.data.data;
        this.desa = res.data.data.Desa.nama_desa;
        this.kecamatan = res.data.data.Kecamatan.nama_kecamatan;
        this.kabupaten = res.data.data.Kabupaten.nama_kabupaten;

        if (res.data.data.latitude && res.data.data.longitude &&
            !isNaN(res.data.data.latitude) && !isNaN(res.data.data.longitude)) {
          this.latitude = parseFloat(res.data.data.latitude);
          this.longitude = parseFloat(res.data.data.longitude);
          this.hasCoordinates = true;
        } else {
          this.hasCoordinates = false;
          console.log("Koordinat tidak tersedia atau tidak valid");
        }

      } catch (error) {
        console.error("Error fetching detail:", error);
        this.hasCoordinates = false;
      }
    },
    simpanAduan() {
      // Validasi form
      if (!this.aduan.alamat || !this.aduan.no_hp || !this.aduan.keluhan) {
        Swal.fire({
          title: 'Error!',
          text: 'Semua field harus diisi',
          icon: 'error',
          confirmButtonColor: '#387144'
        });
        return;
      }

      // Loading state
      Swal.fire({
        title: 'Mohon Tunggu',
        text: 'Sedang memproses...',
        allowOutsideClick: false,
        showConfirmButton: false,
        didOpen: () => {
          Swal.showLoading()
        }
      });

      axios.post("/api/create-aduan", this.aduan)
        .then(res => {
          this.dialog = false;
          Swal.fire({
            title: 'Berhasil!',
            text: 'Keluhan Anda telah terkirim',
            icon: 'success',
            confirmButtonColor: '#387144'
          }).then(() => {
            this.$router.go(0);
          });
        })
        .catch(err => {
          Swal.fire({
            title: 'Error!',
            text: 'Terjadi kesalahan saat mengirim keluhan',
            icon: 'error',
            confirmButtonColor: '#387144'
          });
          console.error(err);
        });
    },

    // Method untuk reset form
    resetForm() {
      this.aduan = {
        program_id: this.aduan.program_id,
        user_id: this.aduan.user_id,
        alamat: '',
        no_hp: '',
        keluhan: ''
      };
    },

    // Modifikasi ketika dialog ditutup
    closeDialog() {
      this.dialog = false;
      this.resetForm();
    },
    initMap() {
      if (typeof window === 'undefined' || !L) return;
      if (!this.hasCoordinates) return;

      try {
        if (this.map) {
          this.map.remove();
        }

        this.map = L.map('map').setView([this.latitude, this.longitude], 13);

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
          attribution: '© OpenStreetMap contributors'
        }).addTo(this.map);

        if (this.latitude && this.longitude) {
          const marker = L.marker([this.latitude, this.longitude]).addTo(this.map);
          if (this.detail.nama_program) {
            marker.bindPopup(this.detail.nama_program);
          }
        }

        setTimeout(() => {
          this.map.invalidateSize();
        }, 100);

      } catch (error) {
        console.error("Error initializing map:", error);
      }
    }
  },
  beforeUnmount() {
    if (this.map) {
      this.map.remove();
    }
  }
}
</script>

<style>
.leaflet-control-container .leaflet-control {
  z-index: 1000;
}

.leaflet-control-zoom {
  border: 2px solid rgba(0,0,0,0.2);
  background-clip: padding-box;
}

.leaflet-default-icon-path {
  background-image: url("https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon.png");
}

.leaflet-default-shadow-path {
  background-image: url("https://unpkg.com/leaflet@1.7.1/dist/images/marker-shadow.png");
}
</style>
