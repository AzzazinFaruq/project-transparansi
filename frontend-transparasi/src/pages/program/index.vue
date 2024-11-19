<template>
  <div class="">
  <navbar-program style=""></navbar-program>
<div class="" style="margin-top: 100px">
  <div class="container-home mb-5">
    <v-text-field
      density="compact"
      v-model="searchQuery"
      label="Temukan Program yang anda inginkan"
      variant="outlined"
      hide-details="true"
      @keyup.enter="searchPrograms"
      prepend-inner-icon="mdi-magnify"
      @click:append-inner="searchPrograms"
    ></v-text-field>
  </div>
  <div class=""  v-for="(program, index) in paginatedPrograms" :key="program.id">
    <div class="" v-if="index == 2 || index == 3" style="background-color: #F5F5F5; ">
      <div class="container-home list-program margin-32">
        <a @click="tanggapi(program.id)">
        <v-row>
          <v-col cols="12" sm="4">
            <div class="wrapper-image">
              <img v-if="program.foto_before == ''" src="../img/dummy-image.jpg" alt="" width="100%" class="image-program">
              <img v-else :src="`${getImageUrl(program.foto_before)}`" class="image-program" alt="" width="100%">
            </div>
            </v-col>
            <v-col cols="12" sm="8">
            <h2>{{ program.nama_program }}</h2>
            <p class="mt-2 clamp-text-6">{{ program.deskripsi }}</p>
            <p class="mt-4" style="font-weight: lighter;">Diupdate pada {{ program.updated_at }}</p>
          </v-col>
        </v-row>
        </a>
      </div>
    </div>
    <div class="" v-else style=" ">
      <div class="container-home list-program  margin-32">
        <a @click="tanggapi(program.id)">
        <v-row>
          <v-col cols="12" sm="4">
            <div class="wrapper-image">
              <img v-if="program.foto_before == ''" src="../img/dummy-image.jpg" alt="" width="100%" class="image-program">
              <img v-else :src="`${getImageUrl(program.foto_before)}`" class="image-program" alt="">
            </div>
            </v-col>
            <v-col cols="12" sm="8">
            <h2>{{ program.nama_program }}</h2>
            <p class="mt-2 clamp-text-6">{{ program.deskripsi }}</p>
            <p class="mt-4" style="font-weight: lighter;">Diupdate pada {{ program.updated_at }}</p>
          </v-col>
        </v-row>
        </a>
      </div>
    </div>
  </div>
</div>

      <div class="container-home d-flex justify-center mt-5 margin-32">
        <v-pagination
          v-model="currentPage"
          :length="totalPages"
          :total-visible="3"
          rounded="circle"
        ></v-pagination>
      </div>

<div class="container-home d-flex justify-center">
</div>
<div v-if="program.length == 0" class="container-home d-flex justify-center">
  <h3 class="mb-5">Tidak ada program yang ditemukan</h3>
</div>
<footer-home></footer-home>
</div>
</template>
<script>
import axios from 'axios';
import navbarProgram from '@/components/navbar-program.vue';
import footerHome from '@/components/footer-home.vue';
import { getImageUrl } from "@/config/foto.js";
export default {
  components: { navbarProgram, footerHome },
  setup() {
    return {};
  },
  data() {
    return {
      getImageUrl,
      program: [],
      searchQuery: '',
      currentPage: 1,
      itemsPerPage: 6, // Jumlah item per halaman
    }
  },
  mounted() {
    this.getProgram();
  },
  methods: {
    searchPrograms() {
      axios.get(`/api/program/search?nama=${this.searchQuery}`)
      .then(res => {
        this.program = res.data.data;
        this.currentPage = 1; // Reset ke halaman pertama saat search
      })
      .catch(err => {
        console.log(err);
      })
    },
    async getProgram() {
    try {
      axios.get('/api/program/status/Publish')
      .then(res => {
        this.program = res.data.data;
      })
    } catch (err) {
      console.log(err);
    }
  },
    tanggapi(id) {
      this.$router.push(`/program/detail/${id}`);
    }
  },
  watch: {
    searchQuery(newVal) {
      if (!newVal || newVal.trim() === '') {
        this.getProgram();
        this.currentPage = 1; // Reset ke halaman pertama
      }
    }
  },
  computed: {
    // Hitung total halaman
    totalPages() {
      return Math.ceil(this.program.length / this.itemsPerPage)
    },
    // Data yang ditampilkan di halaman saat ini
    paginatedPrograms() {
      const start = (this.currentPage - 1) * this.itemsPerPage
      const end = start + this.itemsPerPage
      return this.program.slice(start, end)
    }
  },
}
</script>
<style lang="scss" scoped>
.list-program{
  padding-top: 32px;
  padding-bottom: 32px;
  .wrapper-image{
    width: 100%;
    height: 200px;
    overflow: hidden;
    .image-program{
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
  }
}
.margin-32 {
  padding-top: 32px;
  padding-bottom: 32px;
}
.clamp-text-6 {
  display: -webkit-box;
  -webkit-line-clamp: 5; /* Jumlah baris yang ingin ditampilkan */
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
