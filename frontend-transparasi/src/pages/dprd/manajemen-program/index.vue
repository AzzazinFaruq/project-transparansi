<template>
  <div class="mt-3 pa-5">
    <v-card>
      <div class="d-flex justify-space-between pa-3">
        <h2 class="mr-4">Daftar Program</h2>
        <v-text-field
        class="mr-4"
        variant="outlined"
        density="compact"
          name="name"
          label="Cari Program"
          v-model="searchQuery"
          prepend-inner-icon="mdi-magnify"
          @keyup.enter="searchPrograms"
        ></v-text-field>
        <div class="d-flex jsutify-center mr-2 mt-1" style="">
          <p class="mt-1">Filter :</p>
          <v-btn variant="text" style="text-transform: none;" append-icon="mdi-chevron-down">{{ filter }}</v-btn>
          <v-menu activator="parent">
            <v-list>
              <v-list-item
              v-for="item in filterList"
              :key="item"
              @click="selectedFilter(item)"
              >
                {{ item }}
              </v-list-item>
            </v-list>
          </v-menu>
        </div>
        <v-btn
        class="mt-1"
        style="text-transform: none; width: 130px;"
        color="#BF3232"
        prepend-icon="mdi-plus"
        href="/dprd/manajemen-program/tambah"
        >
        Program
      </v-btn>
      </div>
      <v-divider class="mx-2"></v-divider>
        <div class="pa-2">
          <v-table class="no-divider ">
            <thead style="">
              <tr class="">
                <th style="min-width: 120px;" class=" font-weight-bold">
                  Tanggal
                </th>
                <th style="min-width: 160px;" class=" font-weight-bold">
                  Nama
                </th>
                <th style="min-width: 120px;" class=" font-weight-bold">
                  Status
                </th>
                <th style="min-width: 120px;" class=" font-weight-bold">
                  Action
                </th>
              </tr>
            </thead>
              <tbody>
                <tr
                v-for="(item, index) in paginatedItems" :key="index"
                >
              <td>{{ item.created_at}}</td>
              <td>{{ item.nama_program }}</td>
              <td>
                <v-badge v-if="item.status=='Draft'" dot inline color="#4BB0EB"></v-badge>
                <v-badge v-else-if="item.status=='Publish'" dot inline color="#4A975B"></v-badge>
                {{ item.status }}
              </td>
              <td>
                <v-btn class="rounded-lg" style="background-color:#3884B0;color: white;text-transform: none;" @click="detail(item.id)">
                  Lihat Detail
                </v-btn>
              </td>

              </tr>
              </tbody>
          </v-table>
        </div>
        <v-divider class="my-3 "></v-divider>
        <div class="d-flex justify-end align-center mb-2">
          <p class="mr-3 d-none d-sm-block">Data/Halaman :</p>
          <div class="mr-5 d-none d-sm-block" style="width: 100px;height: 40px;">
            <v-select
            density="compact"
            class="custom-outline"
            v-model="itemsPerPage"
            :items="[5, 10, 25, 50, 100]"
            variant="outlined"
          ></v-select>
          </div>
          <p class="mr-5">{{ currentPage }} dari {{ totalPages }} Halaman</p>
          <v-pagination
            class="custom-pagination"
            :total-visible="0"
            border
            v-model="currentPage"
            :length="totalPages"
            @input="changePage"
          >
          </v-pagination>
        </div>

    </v-card>
    </div>
</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      searchQuery: '',
      filter:'All Status',
      filterList:["All Status","Draft","Publish"],
      detailProgram:[],
      Userlist:[],
      currentPage: 1,
      itemsPerPage: 10,
    }
  },
  mounted() {
    this.user();
  },
  methods: {
    searchPrograms() {
      axios.get(`/api/program/search?nama=${this.searchQuery}`)
      .then(res => {
        this.Userlist = res.data.data;
        this.currentPage = 1; // Reset ke halaman pertama saat search
      })
      .catch(err => {
        console.log(err);
      })
    },
    selectedFilter(item){
      this.filter = item
      if (this.filter == "All Status") {
        axios.get("/api/index-program")
        .then(res=>{
          this.Userlist = res.data.data
          this.currentPage = 1; // Reset ke halaman pertama saat filter
        })
      } else {
        axios.get(`/api/program/status/${item}`)
        .then(res=>{
          this.Userlist = res.data.data
          this.currentPage = 1; // Reset ke halaman pertama saat filter
        })
      }
    },
    user(){
     axios.get("/api/index-program")
     .then(res=>{
      this.Userlist = res.data.data
     })
    },
    changePage(page) {
      this.currentPage = page;
    },
    detail(id){
      this.$router.push(`/dprd/manajemen-program/${id}`)
    },
    disetujui(id){
      console.log(id)
      axios.get(`/api/program/accept/${id}`)
      .then(res=>{
        console.log(res.data.data)
        this.user();
      })
    },
    ditolak(id){
      console.log(id)
      axios.get(`/api/program/reject/${id}`)
      .then(res=>{
        console.log(res.data.data)
        this.user();
      })
    }
  },
  watch: {
    searchQuery(newVal) {
      if (!newVal || newVal.trim() === '') {
        this.user();
        this.currentPage = 1; // Reset ke halaman pertama
      }
    }
  },
  computed: {
    totalPages() {
      return Math.ceil(this.Userlist.length / this.itemsPerPage);
    },
    paginatedItems() {
      const start = (this.currentPage - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.Userlist.slice(start, end);
    },
  },
}
</script>


