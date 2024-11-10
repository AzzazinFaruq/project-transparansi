<template>
  <v-container>
    <div class="mt-1 pa-4">
        <div class="d-flex align-center justify-start mt-2 mb-3">
          <h2>Daftar Program</h2>
          </div>
          <div class="d-flex justify-space-between mb-3">
            <div class="" style="width: 300px; margin-bottom: -20px;">
              <v-text-field
              class=""
              density="compact"
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              v-model="search"
              label="Cari Program"
              @keyup.enter="searchProgram"
              ></v-text-field>
            </div>
            <div class="d-flex" style="margin-top: 8px;">
              <div class="" style="">
                <p class="mr-2" style="">Filter : </p>
              </div>
           <div class="ml-2" style=" min-width: 128px;">
            <v-btn
                variant="text"
                style="border-color: #BF3232;  text-transform: none; letter-spacing: 0.5px; margin-top: -5px; padding: 0;"
                append-icon="mdi-chevron-down"
              >
                {{ filter }}
                <v-menu activator="parent">
                  <v-list>
                    <v-list-item
                    v-for="item in filterList"
                    :key="item"
                    @click="selectedFilter(item);"
                    >
                      <v-list-item-title >{{ item }}</v-list-item-title>
                    </v-list-item>
                  </v-list>
                </v-menu>
              </v-btn>
           </div>
            </div>
        </div>

        <v-divider class="mx-2"></v-divider>
        <div class="">
          <v-table class="no-divider ">
            <thead style="">
              <tr class="">
                <th style="min-width: 100px;" class=" font-weight-bold">
                  Tanggal
                </th>
                <th style="min-width: 200px;" class=" font-weight-bold">
                  Nama
                </th>
                <th style="min-width: 150px;" class=" font-weight-bold">
                  Status
                </th>
                <th class=" font-weight-bold">
                  Aksi
                </th>
              </tr>
            </thead>
              <tbody>
                <tr
                v-for="(item, index) in paginatedItems" :key="index"
                >
              <td>{{ item.created_at}}</td>
              <td>{{ item.program.nama_program }}</td>
              <td
              style="min-width: 150px;"
              >
                <v-badge v-if="item.status=='Belum Ditanggapi'" dot inline color="#FFE642"></v-badge>
                <v-badge v-else-if="item.status=='Sudah Ditanggapi'" dot inline color="#4A975B"></v-badge>
                {{ item.status }}
              </td>
              <td>
                <v-btn class="rounded-lg" style="background-color:#3884B0;color: white;text-transform: none;" @click="detailAduan(item.id)">
                  Detail
                </v-btn>
              </td>

              </tr>
              </tbody>
          </v-table>
        </div>
        <v-divider class="my-3"></v-divider>
        <div class="d-flex justify-end align-center">
          <p class="mr-3">Data/Halaman :</p>
          <div class="mr-5" style="width: 100px;height: 40px;">
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
    </div>
  </v-container>
</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      filter:'All Status',
      filterList:["All Status","Belum Ditanggapi","Sudah Ditanggapi"],
      search:'',
      headers: [
        { title: "Nama", value: "username" },
        { title: "Jabatan", value: "" },
        { title: "Email", value: "email" },
        { title: "No Hp", value: "no_hp" },
        { title: "Action", value: "actions" },
      ],
      Userlist:[],
      currentPage: 1,
      itemsPerPage: 10,
    }
  },
  mounted() {
    this.user();
  },
  watch: {
    search(newVal) {
      if (!newVal || newVal.trim() === '') {
        this.user();
        this.currentPage = 1; // Reset ke halaman pertama
      }
    }
  },
  methods: {
    selectedFilter(item){
      this.filter = item
      if (this.filter == "All Status") {
        axios.get("/api/index-aduan")
        .then(res=>{
          this.Userlist = res.data.data
        })
      } else {
        axios.get(`/api/aduan/status/${item}`)
        .then(res=>{
          this.Userlist = res.data.data
        })
      }
    },
    searchProgram(){
      axios.get(`/api/aduan/aduan-by-program/${this.search}`)
      .then(res=>{
        this.Userlist = res.data.data
      })
    },
    user(){
      axios.get("/api/index-aduan")
      .then(res=>{
        console.log(res.data.data)
        this.Userlist=res.data.data
      })
    },
    changePage(page) {
      this.currentPage = page;
    },
    detailAduan(id){
      this.$router.push(`/admin/keluhan/${id}`)
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
